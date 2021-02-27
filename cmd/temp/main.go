package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"time"

	"github.com/allocamelus/allocamelus/internal/pkg/backupkey"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/user"
)

func main() {
	privKey, _ := pgp.NewKey("name", "1039@alloc.com")
	fmt.Println(len(privKey.Armored))
	publicKey, _ := privKey.PublicKey()
	fmt.Println(len(publicKey.Armored))

	eText, _ := publicKey.EncryptArmored([]byte("test"))
	text, _ := privKey.DecryptArmored(eText)
	fmt.Println(string(text))

	key, humanCode := backupkey.Create()
	fmt.Println(humanCode)
	decodedKey, _ := backupkey.Decode(humanCode)
	fmt.Println((subtle.ConstantTimeCompare(key, decodedKey) == 1))

	u, _ := user.New("name", "bob", "u@b.c")
	start := time.Now()
	err := u.ValidatePublic()
	end := time.Now()
	fmt.Println(end.Sub(start))
	b, _ := json.Marshal(err)
	fmt.Println(string(b))
}
