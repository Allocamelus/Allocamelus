package main

import (
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account"
	apiUser "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	tsGen("web/app/src/models/post_gen.ts", "", post.Post{})
	tsGen("web/app/src/models/user_gen.ts", "", user.User{}, user.Session{})
	tsGen("web/app/src/models/api_account_gen.ts", "API_",
		account.AuthResp{},
		account.AuthRequest{},
		account.AuthA10Token{},
	)
	tsGen("web/app/src/models/api_user_gen.ts", "API_",
		apiUser.CreateResp{},
		apiUser.CreateRequest{},
		apiUser.CreateA10Token{},
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
