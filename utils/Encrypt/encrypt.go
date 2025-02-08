package Encrypt

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"strconv"
	"strings"
)

func Do(password string) (string, error) {
	// 生成16字节的随机盐
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// 参数设置：N=16384, r=8, p=1, keyLen=32
	costParam := 16384   // CPU/内存开销参数，必须是2的幂
	blockSize := 8       // 块大小参数
	parallelization := 1 // 并行度参数
	keyLength := 32      // 生成的密钥长度

	// 使用scrypt生成密钥
	derivedKey, err := scrypt.Key([]byte(password), salt, costParam, blockSize, parallelization, keyLength)
	if err != nil {
		return "", err
	}

	// 编码为Base64存储
	b64Salt := base64.StdEncoding.EncodeToString(salt)
	b64Key := base64.StdEncoding.EncodeToString(derivedKey)

	// 格式化存储字符串：参数用$分隔
	stored := fmt.Sprintf("%d$%d$%d$%s$%s", costParam, blockSize, parallelization, b64Salt, b64Key)
	return stored, nil
}

func VerifyPassword(storedHash, password string) (bool, error) {
	// 解析存储的哈希值
	parts := strings.Split(storedHash, "$")
	if len(parts) != 5 {
		return false, errors.New("invalid hash format")
	}

	// 提取参数
	costParam, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, err
	}
	blockSize, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, err
	}
	parallelization, err := strconv.Atoi(parts[2])
	if err != nil {
		return false, err
	}
	b64Salt := parts[3]
	b64Key := parts[4]

	// 解码盐和密钥
	salt, err := base64.StdEncoding.DecodeString(b64Salt)
	if err != nil {
		return false, err
	}
	storedKey, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return false, err
	}

	// 重新计算用户输入的密码的哈希
	derivedKey, err := scrypt.Key([]byte(password), salt, costParam, blockSize, parallelization, len(storedKey))
	if err != nil {
		return false, err
	}

	// 安全比较哈希值（防止时序攻击）
	return subtle.ConstantTimeCompare(derivedKey, storedKey) == 1, nil
}
