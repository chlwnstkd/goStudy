// 9-5
package main

import (
	"fmt"
	"crypto/rand"
	"crypto/rsa"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	s := '동해물과 백두산이 마르고 닳도록
	하나님이 보우하사 우리나라 만세.
	무궁화 삼천리 화려강산
	대한 사람, 대한으로 길이 보전하세.'
	
	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s),
	)
	fmt.Printf("%x\n", ciphertext)
	
	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext,
	)
	
	fmt.Println(string(plaintext))
}