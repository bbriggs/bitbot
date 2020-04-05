package bitbot

import (
	"io"
	"fmt"
	"strings"
	"crypto/cipher"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	// "encoding/json"
	// "net/http"

	"github.com/whyrusleeping/hellabot"
	"github.com/btcsuite/btcutil/base58"
)

const (
	nonceSize 		= 16			// privatebin uses a nonce of 16 bytes by default
	aesKeySize   	= 32			// using aes-256-gcm; for reference only
	gcmTagSize		= 16			// for reference
)

type AuthData struct {			//todo: should be type "json" ?
	// todo: the following fields (till EOF) should be a (json ?) array
	nonce				string			// base64(cipher_iv); getRandomBytes(16) default
	kdfsalt				string			// base64(kdf_salt); getRandomBytes(8) default
	pbkdf_iterations	int				// pbkdf_iterations; default
	pbkdf_keysize		int				// pbkdf_keysize; default
	cipher_tag_size		int				// cipher_tag_size (wtf ?); default
	cipher_algo			string			// cipher_algo; default
	cipher_mode			string			// cipher_mode; default
	compression_type	string			// compression_type; default
	// EOF

	format 				string			// format of the paste
    open_discussion		bool			// open-discussion flag (todo: not sure if bool works)
    burn_after_reading	bool			// burn-after-reading flag (todo: not sure if bool works)

}

type PasteMeta struct {		// todo: should be type "json" ?
	expire string `json:"expire"`
}

type PasteRequest struct {	// todo: should be type "json" ?
	adata AuthData
	meta PasteMeta
	version int `json:"v"`
	ciphertext string `json:"ct"`	// todo: type should be "base64" ?
}

type PasteResponse struct {
	status bool `json:"status"`		// todo: not sure if bool works
	id string  `json:"id"`
	url string `json:"url"`
	deletetoken string `json:"deletetoken"`
}

var PasteTrigger = NamedTrigger {
	ID:		"paste",
	Help:	"returns a pastebin link for a PRIVMSG to bitbot" + 
			"\nUsage: !paste <content>",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Content, "!paste ")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {

		// uncomment below
		// plaintext := []byte(m.Content)		// todo: only fetch the paste content
		// key, nonce, ciphertext, kdfsalt := encrypt(plaintext)
		/*
		todo: 
			1. send request to privatebin
			2. parse the response
			3. check status, fail if not okay
			4. return a valid URL (https://$website?$paste_id#generated_key_in_base58)

		*/
		return false
	},
}

func recvPaste(req *PasteRequest) (PasteResponse) {
	var resp PasteResponse

	return resp
}

func encrypt(plaintext []byte) (string, string, string, string){
	// encrypts the message with a random key, then return it back
	// todo: should generate the key instead of it being a random
	// todo: should return a struct (first array of AuthData)	

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
	keyWithNonce := make([]byte, aesKeySize + nonceSize)
	if _, err := io.ReadFull(rand.Reader, keyWithNonce); err != nil {
		panic(err.Error())
	}

	key, nonce := keyWithNonce[:aesKeySize], keyWithNonce[aesKeySize:]
	block, err := aes.NewCipher(key)	// will auto-pick aes-256 because key size
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		panic(err.Error())
	}

	// encrypt, encode and return
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	encodedKey := base58.Encode(key)
	encodedNonce := base64.StdEncoding.EncodeToString(nonce)
	encodedCipherText := base64.StdEncoding.EncodeToString(ciphertext)
	encodedKDFsalt := "kvDZJC6IahU="	// dummy value for PBKDF as we don't use it

	fmt.Printf("pt: %s\n key: %s\n, ct: %s\n", plaintext, encodedKey, encodedCipherText)
	return encodedKey, encodedNonce, encodedCipherText, encodedKDFsalt
}