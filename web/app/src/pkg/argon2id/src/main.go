//go:build js && wasm

package main

import (
	"encoding/base64"
	"syscall/js"

	"github.com/allocamelus/allocamelus/pkg/argon2id"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/crypto/blake2b"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type exportedHash struct {
	Version int           `json:"version"`
	Cost    argon2id.Cost `json:"cost"`
	// Encoded no key
	Encoded string `json:"encoded"`
	Salt    string `json:"salt"`
	Key     string `json:"key"`
	KeyHash string `json:"keyHash"`
}

func (eh *exportedHash) toMap() map[string]interface{} {
	return map[string]interface{}{
		"version": eh.Version,
		"cost": map[string]interface{}{
			"time":    eh.Cost.Time,
			"memory":  eh.Cost.Memory,
			"threads": eh.Cost.Threads,
			"keyLen":  eh.Cost.KeyLen,
			"saltLen": eh.Cost.SaltLen,
		},
		"encoded": eh.Encoded,
		"salt":    eh.Salt,
		"key":     eh.Key,
		"keyHash": eh.KeyHash,
	}
}

func argon2idHashSalt() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 3 {
			return "Invalid no of arguments passed"
		}

		password, err := base64.StdEncoding.DecodeString(args[0].String())
		if err != nil {
			return err.Error()
		}
		salt, err := base64.StdEncoding.DecodeString(args[1].String())
		if err != nil {
			return err.Error()
		}
		var cost argon2id.Cost
		json.UnmarshalFromString(args[2].String(), &cost)

		// Hash password
		hash := argon2id.HashSalt(string(password), salt, cost)
		return exportHash(hash)
	})
}

func argon2idHash() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}

		password, err := base64.StdEncoding.DecodeString(args[0].String())
		if err != nil {
			return err.Error()
		}
		var cost argon2id.Cost
		json.UnmarshalFromString(args[1].String(), &cost)

		// Hash password
		hash := argon2id.Hash(string(password), cost)
		return exportHash(hash)
	})
}

func exportHash(hash *argon2id.Password) map[string]interface{} {
	exported := exportedHash{
		Version: hash.Version,
		Cost:    hash.Cost,
		Encoded: hash.ToStringNoKey(),
		Salt:    base64.RawStdEncoding.EncodeToString(hash.Salt),
	}

	if len(hash.Key) > 0 {
		// Hash key with blake2b
		keyHash := blake2b.Sum512(hash.Key)
		exported.Key = base64.RawStdEncoding.EncodeToString(hash.Key)
		exported.KeyHash = base64.RawStdEncoding.EncodeToString(keyHash[:])
	}
	return exported.toMap()
}

func argon2idParse() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		hash, _, err := argon2id.Parse(args[0].String())
		if err != nil {
			return err.Error()
		}
		return exportHash(hash)
	})
}

func main() {
	done := make(chan struct{})
	gA2id := js.Global().Get("argon2id")
	if gA2id.IsUndefined() {
		js.Global().Set("argon2id", make(map[string]interface{}))
	}
	gA2id = js.Global().Get("argon2id")
	gA2id.Set("loading", true)
	gA2id.Set("hash", argon2idHash())
	gA2id.Set("hashSalt", argon2idHashSalt())
	gA2id.Set("parse", argon2idParse())
	gA2id.Set("loaded", true)
	gA2id.Set("loading", false)
	<-done
}
