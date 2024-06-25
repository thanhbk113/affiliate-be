package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// TimeNow
func TimeNow() time.Time {
	return time.Now()
}

// EncodeToString ...
func EncodeToString(text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return encoded
}

// DecodeToString ...
func DecodeToString(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}

// GeneratePassphrase ...
func GeneratePassphrase(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

// HashWithPassphrase ...
func HashWithPassphrase(data []byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// ASCEncrypt ...
func ASCEncrypt(data []byte, passphrase string) (string, error) {
	block, _ := aes.NewCipher([]byte(GeneratePassphrase(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// ASCDecrypt ...
func ASCDecrypt(msg, passphrase string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(msg)
	key := []byte(GeneratePassphrase(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// generateChecksum ...
func generateChecksum(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// ChecksumIsValid ...
func ChecksumIsValid(stringData, clientChecksum string) bool {
	checksumData := generateChecksum(stringData)
	if checksumData != clientChecksum {
		return false
	}
	return true
}

// GenerateChecksum ...
func GenerateChecksum(stringData string) string {
	return generateChecksum(stringData)
}

// SHA1Encode ...
func SHA1Encode(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// HashAndSalt ...
func HashAndSalt(pwd string) string {
	pwdByte := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords ...
func ComparePasswords(hashedPwd string, pwd string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	log.Println(err)
	return err == nil
}
