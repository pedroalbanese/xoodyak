package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pedroalbanese/xoodoo/xoodyak"
)

var (
	key    = flag.String("k", "", "HMAC secret key.")
	target = flag.String("f", "", "Target file. ('-' for STDIN)")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Xoodyak MAC - ALBANESE Lab (c) 2020-2023\n")
		fmt.Println("Usage of", os.Args[0]+":")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var err error
	var file io.Reader
	if *target == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(*target)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
	}
	h := xoodyak.NewXoodyakMac([]byte(*key))
	if _, err = io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
	os.Exit(0)

}
