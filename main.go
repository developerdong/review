package main

import (
	"fmt"
	"github.com/developerdong/review/conf"
	"github.com/developerdong/review/sql"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const help = `
usage:
	review <insert|select|next|delete> [url]
description:
	insert
		Insert a reading record of the url.
	select
		Select the reading record with lowest retrievability.
	next
		Insert a reading record of the url with lowest retrievability, then select a new lowest one.
	delete
		Delete all reading records of the url.
example:
	review insert https://www.google.com
	review select
	review next
	review delete https://www.google.com
`

var (
	// The back-end implementation of storage.
	storage sql.Storage
)

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

func init() {
	// Init the storage instance.
	switch driverName := strings.TrimSpace(conf.GetEnv(conf.DriverName)); driverName {
	case sql.SqliteDriverName:
		storage = new(sql.Sqlite)
	case "":
		Fatalln("the name of sql driver should be set in the environment")
	default:
		Fatalf("the driver name %s is unsupported\n", driverName)
	}
}

func main() {
	// execute commands
	if !((len(os.Args) == 2 && (os.Args[1] == "select" || os.Args[1] == "next")) || (len(os.Args) == 3 && (os.Args[1] == "insert" || os.Args[1] == "delete"))) {
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
			u, r, err := storage.Select()
			if err != nil {
				Fatalln(err)
			}
			fmt.Println(u, r)
		case "next":
			oldU, oldR, err := storage.Select()
			if err != nil {
				Fatalln(err)
			}
			if err = storage.Insert(oldU); err != nil {
				// If encounter error, at least print the old URL.
				fmt.Println(oldU, oldR, "(old)")
				Fatalln(err)
			}
			newU, newR, err := storage.Select()
			if err != nil {
				// If encounter error, at least print the old URL.
				fmt.Println(oldU, oldR, "(old)")
				Fatalln(err)
			}
			// Print aligned URLs.
			oldStr, newStr := oldU.String(), newU.String()
			alignLen := len(oldStr)
			if len(newStr) > alignLen {
				alignLen = len(newStr)
			}
			alignLenStr := strconv.Itoa(alignLen)
			fmt.Printf("%-"+alignLenStr+"s %f (old)\n", oldStr, oldR)
			fmt.Printf("%-"+alignLenStr+"s %f (new)\n", newStr, newR)
		case "delete":
			if u, err := url.Parse(os.Args[2]); err != nil {
				Fatalln(err)
			} else if err := storage.Delete(u); err != nil {
				Fatalln(err)
			}
		}
	}
}
