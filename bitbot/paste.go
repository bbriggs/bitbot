package bitbot

// welp; it turned out the key (the base58 encoded one) is used in a key-derivation function; that is, the key used in AES-GCM is the derived one; consequently, the salt we send is equally important; here're the specs
// key derivation function: pbkdf2
// key: password of paste + our random key
// salt: the one we send
// hash algorithm: sha-256 (used at the end of derivation)
// iterations: 10,000 (can we instruct the server to lower it ?)
// ====================================================
// will be using this: https://pkg.go.dev/golang.org/x/crypto/pbkdf2?tab=doc
import (
	"bytes"
     "golang.org/x/crypto/pbkdf2"
     "crypto/sha256"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	// "net"
	// "time"
	"net/http"
	// "os"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/whyrusleeping/hellabot"
)

const (
	nonceSize  = 16 // privatebin uses a nonce of 16 bytes by default
	aesKeySize = 32 // using aes-256-gcm; for reference only
	gcmTagSize = 16 // for reference
    kdfSaltSize = 8 // for reference
)

// not used directly in the paste request
type Array1 struct { // TODO: more descriptive name
	// TODO: the following fields (till EOF) should be a (json ?) array
	Nonce           []byte // base64(cipher_iv); getRandomBytes(16) default
	Kdfsalt         []byte // base64(kdf_salt); getRandomBytes(8) default
	KdfIterations   int    // pbkdf_iterations; default
	KdfKeySize      int    // pbkdf_keysize; default
	CipherTagSize   int    // cipher_tag_size (wtf ?); default
	CipherAlgo      string // cipher_algo; default
	CipherMode      string // cipher_mode; default
	CompressionType string // compression_type; default
	// EOF
}

type AuthData struct { //TODO: should be type "json" ?
	//
	EncryptionDetails []interface{} // TODO: more descriptive name (type was: Array1)
	Format            string        // format of the paste
	OpenDiscussion    int           // open-discussion flag (TODO: not sure if bool works)
	BurnAfterReading  int           // burn-after-reading flag (TODO: not sure if bool works)
	//
}
type PasteData struct { // !shrug (see https://github.com/PrivateBin/PrivateBin/wiki/Encryption-format#data-passed-in)
	Paste           string        `json:"paste"`    // ciphertext (encrypted zlib'd plaintext)
	Attachment      string        `json:"attachment"`
	AttachementName string        `json:"attachment_name"`
	Children        []interface{} `json:"children"`
}

// ============================================================================================================================================================================================================================

// https://raw.githubusercontent.com/PrivateBin/PrivateBin/master/js/types.jsonld
type PasteMeta struct {
	Expire string `json:"expire"` // ["5min", "10min", "1hour", "1day", "1week", "1month", "1year", "never"]
}

type PasteResponse struct {
	Status      int    `json:"status"`
	Id          string `json:"id"`
	Url         string `json:"url"`
	Deletetoken string `json:"deletetoken"`
}

type PasteRequest struct {
	AuthData   []interface{} `json:"adata"`
	Meta       PasteMeta     `json:"meta"`
	Version    int           `json:"v"`
	CipherText []byte        `json:"ct"`
}



func NewRequest(aData []interface{}, cipherText []byte, expiryDate string) *PasteRequest {

	var (
		req            PasteRequest
	)

	meta := PasteMeta{expiryDate}

	req.AuthData = aData
	req.Meta = meta
	req.Version = 2
	req.CipherText = cipherText
	return &req

}

var PasteTrigger = NamedTrigger{
	ID: "paste",
	Help: "returns a pastebin link for a PRIVMSG to bitbot" +
		"\nUsage: !paste <content>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && // This is a message
			m.Params[0] == irc.Nick && // Private to the bot
			strings.HasPrefix(m.Content, "!paste ") // beginning with !paste
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {

		var (
			pasteResp PasteResponse
			pasteReq  *PasteRequest
			err       error
		)
		plaintext := []byte(m.Content) // TODO: only fetch paste content and options
		key, nonce,  kdfsalt := generateEncryptionParameters()
        adata := generateAuthenticationData(nonce, kdfsalt, "plaintext", 0, 0)
        aesKey := pbkdf2.Key(key, kdfsalt, 100000, aesKeySize, sha256.New)
        ciphertext := encrypt(plaintext, aesKey, nonce,  adata)    // auth tag is appended to ciphertext
		pasteReq = NewRequest(adata, ciphertext, "1week")
		if pasteResp, err = recvPaste(pasteReq); err != nil {
			fmt.Println("could not receive paste")
			return false
		}

		if pasteResp.Status != 0 {
			irc.Reply(m, "This is impossible to see unless the same key AND message were used")
			return false
		}
		url := fmt.Sprintf("https://bin.fraq.io%s#%s", pasteResp.Url, base58.Encode(key))
		deleteURL := fmt.Sprintf("https://bin.fraq.io/?pasteid=%s&deletetoken=%s", pasteResp.Id, pasteResp.Deletetoken)
        fmt.Printf("response: %v\n", pasteResp) // TODO: use logging
        fmt.Printf("url: %s\n delete url: %s\n", url, deleteURL)    // TODO: use logging
        irc.Reply(m, fmt.Sprintf("Link: %s | Delete paste: %s", url, deleteURL))

		return true
	},
}

func generateAuthenticationData(iv []byte, dummyKDFsalt []byte, format string, openDiscussion int, burnAfterReading int) []interface{} {
	// encryptionInfo := Array1{iv, dummyKDFsalt, 10000, 265, 128, "aes", "gcm", "zlib"}
	// encryptionInfo := make([]interface{}, 0)
    var (
		encryptionInfo []interface{}
		aData          []interface{} // or aData := make([]interface{}, 0); then append
    )
    encryptionInfo = append(encryptionInfo, iv, dummyKDFsalt, 100000, 256, 128, "aes", "gcm", "none")       // TODO: rget back zlib support
	aData = append(aData, encryptionInfo, format, openDiscussion, burnAfterReading)
    return aData
}
func generateEncryptionParameters() (key, iv, kdfSalt []byte){


	// since we'll be using a different random key for each paste,
	// a fixed nonce should be OK (but we won't do it anyway)
    totalSize := aesKeySize+nonceSize+kdfSaltSize
	keyWithNonceAndKdfSalt := make([]byte, totalSize) 
	if _, err := io.ReadFull(rand.Reader, keyWithNonceAndKdfSalt); err != nil {
		panic(err.Error())
	}

	key = keyWithNonceAndKdfSalt[:aesKeySize]
    iv = keyWithNonceAndKdfSalt[aesKeySize:aesKeySize+nonceSize]
    kdfSalt = keyWithNonceAndKdfSalt[totalSize-kdfSaltSize:]// dummy value for PBKDF as we don't use it

    return key, iv, kdfSalt

}

func recvPaste(pasteReq *PasteRequest) (resp PasteResponse, err error) {
	var (
		jsonForm []byte
		req      *http.Request
		r        *http.Response
	)

	if jsonForm, err = json.Marshal(pasteReq); err != nil { // Marshal, not NewEncoder
		fmt.Println(err)
	}
	fmt.Printf("marhsalled json: %s\n", jsonForm)
	// ==== cert ==== //
	// TODO: self-signed certificates workaround (https://groups.google.com/d/msg/golang-nuts/v5ShM8R7Tdc/I2wyTy1o118J)
	cert, err := ioutil.ReadFile("/tmp/burp.pem") // TODO: to debug w/ burp, install its cert here
	if err != nil {
		fmt.Println("error in importing cert: ", err)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(cert); ok != true {
		fmt.Println("error in appending cert")
	}
	transportOptions := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            caCertPool,
		},
		Proxy: http.ProxyFromEnvironment,
	}

	httpClient := &http.Client{Transport: transportOptions}
	if req, err = http.NewRequest("POST", "https://bin.fraq.io", bytes.NewReader(jsonForm)); err != nil { // TODO: don't hardcode url
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("X-Requested-With", "JSONHttpRequest") // reason we used http.NewRequest w/ Client.Do()
	// req.Header.Add("Origin", "https://bin.fraq.io")

	// fmt.Println("debugging body: ")
	// if b, err := req.GetBody(); err == nil{
	//     fmt.Println(ioutil.ReadAll(b))
	//     fmt.Println(ioutil.ReadAll(req.Body))
	// }
	// resp, err := http.Post(, "application/json", &jsonForm)
	url, err := http.ProxyFromEnvironment(req)
	fmt.Printf("proxy: %s %s\n", url, err)
	// req.WriteProxy(os.Stdout)

	if r, err = httpClient.Do(req); err != nil {
		fmt.Printf("post to paste site error: %s\n", err) // TODO: use log
		return resp, err
	}
	// fmt.Println(r.Body)
	defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)
	// fmt.Printf("received body: %s, err %s\n", body, err)
	// err = json.Unmarshal(body, &resp)
	if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
		fmt.Printf("json decoding error: %s\n", err) // TODO: use logging
	}

	return resp, err
}

func encrypt(plaintext, key, iv []byte, authenticationData []interface{}) (ciphertext []byte){
	// compresses the message with zlib and encrypts it with a random key

	block, err := aes.NewCipher(key) // will auto-pick aes-256 because key size
	if err != nil {
		panic(err.Error())
	}

    aesgcm, err := cipher.NewGCMWithNonceSize(block, nonceSize) //TODO: should instruct privatebin to use standard nonce size instead
	if err != nil {
		panic(err.Error())
	}

	// compress and encrypt message, then encode key and return
	var (
		compressedCiphertext        bytes.Buffer
        // pasteData       PasteData
		// encodedCompressedPlaintext bytes.Buffer
        cipherJson, authenticatedDataJson []byte
	)
    // pasteData = PasteData{Paste: string(plaintext)}      // TODO: support file attachement and paste linking
    pasteData := struct {
        Paste string `json:"paste"`
    }{
        string(plaintext),
    }
	if cipherJson, err = json.Marshal(pasteData); err != nil { // Marshal, not NewEncoder
		panic(err.Error())
	}
	if authenticatedDataJson, err = json.Marshal(authenticationData); err != nil { // Marshal, not NewEncoder
		panic(err.Error())
	}
    fmt.Printf("marshalled cipher: %s\n", cipherJson)
    fmt.Printf("marshalled adata: %s\n", authenticatedDataJson)

    // TODO: add check for compression support (for now, assuming defaults); get back zlib
	compressedCiphertextWriter := zlib.NewWriter(&compressedCiphertext)
	compressedCiphertextWriter.Write(cipherJson)
	compressedCiphertextWriter.Close()
	// encoder := base64.NewEncoder(base64.StdEncoding, &encodedCompressedPlaintext)
	// encoder.Write(compressedCiphertext.Bytes())
	// encoder.Close()

    // authData is authenticated as well(https://github.com/r4sas/PBinCLI/blob/682b47fbd3e24a8a53c3b484ba896a5dbc85cda2/pbincli/format.py#L122)
    // kudos to filo for hinting about the tag location (https://github.com/golang/go/issues/32742)
    // look for function " decryptOrPromptPassword" in privatebin.js; start debugging there
    // TODO: fully support the API (https://github.com/PrivateBin/PrivateBin/wiki/API)
    ciphertext = aesgcm.Seal(nil, iv, cipherJson, authenticatedDataJson) // TODO: zzzzlib
	// 	encodedNonce := base64.StdEncoding.EncodeToString(nonce)
	// 	encodedCipherText := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("pt: %s\n key: %s\n", plaintext, base58.Encode(key))
	return ciphertext
}
