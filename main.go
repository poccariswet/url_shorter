package main

import (
	"fmt"
	"log"
	"os"

	"github.com/soeyusuke/url_shorter/storage/mysql"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		log.Println(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

}
