package bitbot

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	// "encoding/base64"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/whyrusleeping/hellabot"
)

const (
	nonceSize  = 16 // privatebin uses a nonce of 16 bytes by default
	aesKeySize = 32 // using aes-256-gcm; for reference only
	gcmTagSize = 16 // for reference
)

// not used directly in the paste request
type array1 struct { // TODO: more descriptive name
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

type authData struct { //TODO: should be type "json" ?
	//
	EncryptionDetails *array1 // TODO: more descriptive name
	Format            string  // format of the paste
	OpenDiscussion    int     // open-discussion flag (TODO: not sure if bool works)
	BurnAfterReading  int     // burn-after-reading flag (TODO: not sure if bool works)
	//
}

func (req PasteRequest) NewRequest(iv []byte, cipherText []byte, dummyKDFsalt []byte, format string, openDiscussion int, burnAfterReading int, expiryDate string) *PasteRequest {


	encryptionInfo := array1{iv,dummyKDFsalt,10000,265,128,"aes","gcm","zlib"}

	// TODO: this is really messy!
	aData := authData{EncryptionDetails: encryptionInfo, Format: format, OpenDiscussion: openDiscussion, BurnAfterReading: burnAfterReading}
	// aData.EncryptionDetails = encryptionInfo
	// aData.Format = format
	// aData.OpenDiscussion = openDiscussion
	// aData.BurnAfterReading = burnAfterReading
	var aDataInterface [4]interface{} //https://golang.org/doc/faq#convert_slice_of_interface
	// for i, v := range aData {
	//         aDataInterface[i] = v
	// }
	aDataInterface[0] = aData.EncryptionDetails
	aDataInterface[1] = aData.Format
	aDataInterface[2] = aData.OpenDiscussion
	aDataInterface[3] = aData.BurnAfterReading

	meta := PasteMeta{expiryDate}
	
	req.AuthData = aDataInterface
	req.Meta = meta
	req.Version = 2
	req.CipherText = cipherText

	return &req
}

// ============================================================================================================================================================================================================================

type PasteMeta struct { // TODO: should be type "json" ?
	Expire string `json:"expire"`
}

type PasteRequest struct { // TODO: should be type "json" ?
	AuthData   [4]interface{} `json:"adata"`
	Meta       *PasteMeta     `json:"meta"` // TODO: meta is another json
	Version    int            `json:"v"`
	CipherText []byte         `json:"ct"` // TODO: type should be "base64" ?
}

type PasteResponse struct {
	Status      int    `json:"status"` // TODO: not sure if bool works
	Id          string `json:"id"`
	Url         string `json:"url"`
	Deletetoken string `json:"deletetoken"`
}

var PasteTrigger = NamedTrigger{
	ID: "paste",
	Help: "returns a pastebin link for a PRIVMSG to bitbot" +
		"\nUsage: !paste <content>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!paste ")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {

		// uncomment below
		plaintext := []byte(m.Content)		// TODO: only fetch paste content and options
		key, nonce, ciphertext, kdfsalt := encrypt(plaintext)
        pasteReq := PasteRequest.NewRequest(nonce, ciphertext, kdfsalt, "plaintext", 0, 0, "1week")
        if pasteResp, err := recvPaste(pasteReq); err != nil {
            fmt.Println("could not receive paste")
            return false
        }
        fmt.Println("%v %v", key, pasteResp)
		/*
			TODO:
				1. send request to privatebin
				2. parse the response
				3. check status, fail if not okay
				4. return a valid URL (https://$website?$paste_id#generated_key_in_base58)

		*/
		return true
	},
}

func recvPaste(pasteReq *PasteRequest) (resp PasteResponse, err error) {
	var (
		jsonForm []byte
		r, req   *http.Request
	)

	if jsonForm, err := json.Marshal(pasteReq); err != nil { // Marshal, not NewEncoder
		fmt.Println(err)
	}
	httpClient := &http.Client{}
	if req, err := http.NewRequest("POST", "https://bin.fraq.io", bytes.NewReader(jsonForm)); err != nil { // TODO: don't hardcode url
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Requested-With", "JSONHttpRequest")

	// resp, err := http.Post(, "application/json", &jsonForm)
	if r, err := httpClient.Do(req); err != nil {
        fmt.Println(err)    // TODO: use log
	}
	// fmt.Println(r.Body)
	defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	//  print(err)
	// }
	// err = json.Unmarshal(body, &m)
	if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
        fmt.Println(err)    // TODO: use logging
	}

	return resp, err
}

func encrypt(plaintext []byte) (encodedKey string, nonce []byte, ciphertext []byte, dummyKDFsalt []byte) {
	// encrypts the message with a random key, then return it back
	// TODO: should generate the key instead of it being a random
	// TODO: should return a struct (first array of AuthData)

	/*		public function create($pasteid, $paste)
	{
	$pasteid = substr(hash('md5', $paste['data']), 0, 16);

	$paste['data']                      // text
	$paste['meta']['postdate']          // int UNIX timestamp
	$paste['meta']['expire_date']       // int UNIX timestamp
	$paste['meta']['opendiscussion']    // true (if false it is unset)
	$paste['meta']['burnafterreading']  // true (if false it is unset; if true, then opendiscussion is unset)
	$paste['meta']['formatter']         // string
	$paste['meta']['attachment']        // text
	$paste['meta']['attachmentname']    // text
	}
	*/

	// since we'll be using a different random key for each paste,
	// a fixed nonce should be OK (but we won't do it anyway)
	keyWithNonce := make([]byte, aesKeySize+nonceSize)
	if _, err := io.ReadFull(rand.Reader, keyWithNonce); err != nil {
		panic(err.Error())
	}

	key, nonce := keyWithNonce[:aesKeySize], keyWithNonce[aesKeySize:]
	block, err := aes.NewCipher(key) // will auto-pick aes-256 because key size
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		panic(err.Error())
	}

	// encrypt, encode and return
	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	encodedKey = base58.Encode(key)
	// 	encodedNonce := base64.StdEncoding.EncodeToString(nonce)
	// 	encodedCipherText := base64.StdEncoding.EncodeToString(ciphertext)
    dummyKDFsalt = make([]byte, 8)      // TODO: remove magic number
	if _, err := io.ReadFull(rand.Reader, dummyKDFsalt); err != nil {
		fmt.Println("warning, dummy kdfsalt is not initialized properly")
		// encodedKDFsalt := "kvDZJC6IahU=" // dummy value for PBKDF as we don't use it
	}
	fmt.Printf("pt: %s\n key: %s\n, ct: %s\n", plaintext, encodedKey, ciphertext)
	return encodedKey, nonce, ciphertext, dummyKDFsalt
}
