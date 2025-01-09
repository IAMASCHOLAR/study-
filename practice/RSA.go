package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// PoW 函数从之前的代码移植过来
func minePow(targetLen int) (int, string, float64) {
	target := ""
	for i := 0; i < targetLen; i++ {
		target += strconv.Itoa(i)
	}

	starttime1 := time.Now()
	nonce := 0
	nickname := "Hz_001"

	for {
		data := fmt.Sprintf("%s%d", nickname, nonce)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))

		if hash[0:targetLen] == target {
			time_spent := time.Since(starttime1).Seconds()
			return nonce, hash, time_spent
		}
		nonce++
	}
}

func main() {
	// 1. 先进行 PoW 计算
	targetLen := 4 // 寻找前4位为"0123"的哈希
	nonce, powHash, timeSpent := minePow(targetLen)
	fmt.Printf("\nPoW 结果：\n")
	fmt.Printf("Nonce: %d\n", nonce)
	fmt.Printf("Hash: %s\n", powHash)
	fmt.Printf("耗时: %.2f 秒\n\n", timeSpent)

	// 2. 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("生成密钥对失败: %v\n", err)
		return
	}
	publicKey := &privateKey.PublicKey

	// 3. 使用 PoW 的结果创建要签名的数据
	nickname := "Hz_001"
	data := fmt.Sprintf("%s%d", nickname, nonce)

	// 4. 计算数据的哈希
	dataHash := sha256.Sum256([]byte(data))

	// 5. 使用私钥对哈希进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, dataHash[:])
	if err != nil {
		fmt.Printf("签名失败: %v\n", err)
		return
	}
	fmt.Printf("RSA签名成功！\n")
	fmt.Printf("签名数据: %s\n", data)
	fmt.Printf("签名结果: %x\n\n", signature)

	// 6. 使用公钥验证签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, dataHash[:], signature)
	if err != nil {
		fmt.Printf("验证签名失败: %v\n", err)
		return
	}
	fmt.Printf("RSA验证签名成功！\n")
}
