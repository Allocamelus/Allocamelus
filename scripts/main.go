package main

import (
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account"
	apiUser "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user"
	apiUserUpdate "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/update"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	tsGen("web/app/src/models/go_structs_gen.ts", "GEN_",
		post.Post{},
		user.User{},
		user.Session{},
		account.AuthResp{},
		account.AuthRequest{},
		account.AuthA10Token{},
		apiUser.CreateResp{},
		apiUser.CreateRequest{},
		apiUser.CreateA10Token{},
		apiUserUpdate.AvatarResp{},
	)
}

func tsGen(path string, prefix string, structs ...interface{}) {
	converter := typescriptify.New()
	converter.Prefix = prefix
	converter.BackupDir = ""
	for _, s := range structs {
		converter.Add(s)
	}
	if err := converter.ConvertToFile(path); err != nil {
		panic(err.Error())
	}
}
