package main

import (
	"fmt"
	"github.com/developerdong/review/conf"
	"github.com/developerdong/review/sql"
	"net/url"
	"os"
	"strings"
)

const help = `
usage:
	review <insert|select|delete> [url]
example:
	review insert https://www.google.com
	review select
	review delete https://www.google.com
`

func Fatal(v ...interface{}) {
	fmt.Print(v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func main() {
	// init the storage instance
	var storage sql.Storage
	switch driverName := strings.TrimSpace(conf.GetEnv(conf.DriverName)); driverName {
	case sql.SqliteDriverName:
		storage = new(sql.Sqlite)
	case "":
		Fatalln("the name of sql driver should be set in the environment")
	default:
		Fatalf("the driver name %s is unsupported\n", driverName)
	}
	// execute commands
	if !((len(os.Args) == 2 && os.Args[1] == "select") || (len(os.Args) == 3 && (os.Args[1] == "insert" || os.Args[1] == "delete"))) {
		// the format of input is incorrect
		Fatal(help)
	} else {
		// process different command
		switch os.Args[1] {
		case "insert":
			if u, err := url.Parse(os.Args[2]); err != nil {
				Fatalln(err)
			} else if err := storage.Insert(u); err != nil {
				Fatalln(err)
			}
		case "select":
			if u, r, err := storage.Select(); err != nil {
				Fatalln(err)
			} else {
				fmt.Println(u.String(), r)
			}
		case "delete":
			if u, err := url.Parse(os.Args[2]); err != nil {
				Fatalln(err)
			} else if err := storage.Delete(u); err != nil {
				Fatalln(err)
			}
		}
	}
}
