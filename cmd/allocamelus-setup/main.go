package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	sqlcdb "github.com/allocamelus/allocamelus/db"
	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/version"
	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

var configPath string

func init() {
	if configPath = getEnvTrim("CONFIG_PATH"); configPath == "" {
		configPath = "./config.json"
	}

	v := flag.Bool("version", false, "Returns version")
	flag.Parse()
	if *v {
		fmt.Println(version.Version)
		os.Exit(0)
	}
}

func main() {
	d := data.New(configPath)
	u, _ := url.Parse(d.Config.DBconnStr())
	db := dbmate.New(u)
	db.FS = sqlcdb.FS
	db.MigrationsDir = []string{"./migrations"}

	fmt.Println("Migrations:")
	migrations, err := db.FindMigrations()
	if err != nil {
		panic(err)
	}
	for _, m := range migrations {
		fmt.Println(m.Version, m.FilePath)
	}

	fmt.Println("\nApplying migrations...")
	err = db.CreateAndMigrate()
	if err != nil {
		panic(err)
	}
}

func getEnvTrim(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
