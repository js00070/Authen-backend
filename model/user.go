package model

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Test 测试
type Test struct {
	gorm.Model
}

// User 用户信息
type User struct {
	// 基础信息
	//ID uint `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	gorm.Model
	LoginName  string `gorm:"unique_index;size:16" json:"login_name"`
	PasswdHash string `gorm:"size:64"`
}

// SetPasswd 设置密码
func (user *User) SetPasswd(passwd string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), 12)
	if err != nil {
		return err
	}
	user.PasswdHash = string(bytes)
	return nil
}

// CheckPasswd 检查密码，若密码正确返回true，否则返回false
func (user *User) CheckPasswd(passwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswdHash), []byte(passwd))
	return err == nil
}

func encryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func decryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}

// GetToken 加密userid得到token
func (user *User) GetToken() string {
	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := fmt.Sprintf("%d,%d,%d", user.ID, time.Now().Unix(), time.Now().UnixNano())

	c := encryptAES([]byte(key), pt)
	return c
}

// GetUIDFromToken 从token获取userid
func GetUIDFromToken(token string) (uint, error) {
	s := decryptAES([]byte("thisis32bitlongpassphraseimusing"), token)
	var id uint
	var unix int64
	var unixNano int64
	_, err := fmt.Sscanf(s, "%d,%d,%d", &id, &unix, &unixNano)
	return id, err
}
