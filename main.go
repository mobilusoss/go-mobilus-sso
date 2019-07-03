package go_mobilus_sso

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"strings"
)

type (
	MobilusSSO struct {
		Secret string
	}
	User struct {
		Name        string `json:"name"`
		PermitLevel int    `json:"permitLevel"`
		Token       string `json:"token"`
		DomainID    string `json:"domainId"`
		PlusID      string `json:"plusId"`
		UserID      string `json:"userId"`
		TenantID    string `json:"tenantId"`
	}
)

func New(secret string) *MobilusSSO {
	return &MobilusSSO{
		Secret: secret,
	}
}

func (this *MobilusSSO) Unmarshal(data string) (*User, error) {
	dataBytes, err := this.decodeBase64(data)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	block, err := aes.NewCipher(this.getHash())
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(dataBytes))
	dec := cipher.NewCBCDecrypter(block, iv)
	dec.CryptBlocks(decrypted, dataBytes)

	s := string(decrypted)
	p := strings.LastIndex(s, "}") // NOTE: \x03とかを回避する

	result := User{}
	err = json.Unmarshal(decrypted[0:p+1], &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (this *MobilusSSO) Marshal(data interface{}) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	dataBytes = this.paddingPKCS5(dataBytes)

	iv := make([]byte, aes.BlockSize)
	block, err := aes.NewCipher(this.getHash())
	if err != nil {
		return "", err
	}

	encrypted := make([]byte, len(dataBytes))
	dec := cipher.NewCBCEncrypter(block, iv)
	dec.CryptBlocks(encrypted, dataBytes)

	result := this.encodeBase64(encrypted)

	return string(result), nil
}

func (this *MobilusSSO) getHash() []byte {
	h := md5.New()
	h.Write([]byte(this.Secret))
	return h.Sum(nil)
}

func (this *MobilusSSO) decodeBase64(data string) ([]byte, error) {
	return base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(data)
}

func (this *MobilusSSO) encodeBase64(data []byte) string {
	r := base64.URLEncoding.EncodeToString(data)
	return strings.Replace(r, "=", "", -1)
}

func (this *MobilusSSO) paddingPKCS5(src []byte) []byte {
	repeat := aes.BlockSize - len(src)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(repeat)}, repeat)
	return append(src, padding...)
}
