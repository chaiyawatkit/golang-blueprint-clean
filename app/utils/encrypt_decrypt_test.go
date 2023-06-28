package utils_test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang-blueprint-clean/app/utils"

	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUtils_EncryptAES(t *testing.T) {
	var (
		strToEncrypt = "string to encrypt"
	)
	t.Run("Happy", func(t *testing.T) {
		bytes := make([]byte, 32)
		rand.Read(bytes)
		key := hex.EncodeToString(bytes)
		encrypted, err := utils.EncryptAES(strToEncrypt, key)
		assert.Nil(t, err)
		assert.NotNil(t, encrypted)
	})

	t.Run("Failure", func(t *testing.T) {
		_, err := utils.EncryptAES("string to encrypt", "")
		assert.NotNil(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		_, err := utils.EncryptAES("string to encrypt", "fb4b22d7e1ea9f29917aafd2b26f59e5fbb68ff93f91c021d8f895895dc346b")
		assert.NotNil(t, err)
	})

	t.Run("Happy", func(t *testing.T) {
		bytes := make([]byte, 32)
		rand.Read(bytes)
		key := hex.EncodeToString(bytes)
		encrypted, _ := utils.EncryptAES(strToEncrypt, key)
		decrypted, err := utils.DecryptAES(fmt.Sprintf("%x", encrypted), key)
		assert.Nil(t, err)
		assert.NotNil(t, decrypted)
		assert.Equal(t, strToEncrypt, *decrypted)
	})

	t.Run("Failure", func(t *testing.T) {
		bytes := make([]byte, 64)
		rand.Read(bytes)
		key := hex.EncodeToString(bytes)
		encrypted, _ := utils.EncryptAES(strToEncrypt, key)
		decrypted, err := utils.DecryptAES(fmt.Sprintf("%x", encrypted), key)
		assert.Nil(t, decrypted)
		log.Printf("debug ====> %v", err.Error())
		assert.NotNil(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		bytes := make([]byte, 32)
		rand.Read(bytes)
		key := hex.EncodeToString(bytes)
		encrypted, _ := utils.EncryptAES(strToEncrypt, key)
		decrypted, err := utils.DecryptAES(fmt.Sprintf("%x", encrypted), "key")
		assert.Nil(t, decrypted)
		log.Printf("debug ====> %v", err.Error())
		assert.NotNil(t, err)
	})
}
