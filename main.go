//go:generate goversioninfo -manifest=testdata/resource/goversioninfo.exe.manifest
package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"os"

	"github.com/pedroalbanese/xoodoo/xoodyak"
)

var (
	dec    = flag.Bool("d", false, "Decrypt instead Encrypt.")
	file   = flag.String("f", "", "Target file. ('-' for STDIN)")
	iter   = flag.Int("i", 1024, "Iterations. (for PBKDF2)")
	key    = flag.String("k", "", "128-bit key to Encrypt/Decrypt.")
	pbkdf  = flag.String("p", "", "PBKDF2.")
	random = flag.Bool("r", false, "Generate random 128-bit cryptographic key.")
	salt   = flag.String("s", "", "Salt. (for PBKDF2)")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Xoodyak Encryption Tool - ALBANESE Research Lab (c) 2020-2022")
		fmt.Fprintln(os.Stderr, "A Lightweight AEAD Cryptographic Scheme written in Pure Go\n")
		fmt.Fprintln(os.Stderr, "Usage of "+os.Args[0]+":")
		fmt.Fprintln(os.Stderr, os.Args[0]+" [-d] -p \"pass\" [-i N] [-s \"salt\"] -f <file.ext>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *random == true {
		var key []byte
		var err error
		key = make([]byte, 16)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.EncodeToString(key))
		os.Exit(0)
	}

	var keyHex string
	var keyRaw []byte
	if *pbkdf != "" {
		keyRaw = pbkdf2.Key([]byte(*pbkdf), []byte(*salt), *iter, 16, xoodyak.NewXoodyakHash)
		keyHex = hex.EncodeToString(keyRaw)
	} else {
		keyHex = *key
	}
	var key []byte
	var err error
	if keyHex == "" {
		key = make([]byte, 16)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "Key=", hex.EncodeToString(key))
	} else {
		key, err = hex.DecodeString(keyHex)
		if err != nil {
			log.Fatal(err)
		}
		if len(key) != 16 {
			log.Fatal(err)
		}
	}

	buf := bytes.NewBuffer(nil)
	var data io.Reader
	if *file == "-" {
		data = os.Stdin
	} else {
		data, _ = os.Open(*file)
	}
	io.Copy(buf, data)
	msg := buf.Bytes()

	aead, err := xoodyak.NewXoodyakAEAD(key)
	if err != nil {
		panic(err)
	}

	if *dec == false {
		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())

		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			log.Fatal(err)
		}

		out := aead.Seal(nonce, nonce, msg, nil)
		fmt.Printf("%s", out)

		os.Exit(0)
	}

	if *dec == true {
		nonce, msg := msg[:aead.NonceSize()], msg[aead.NonceSize():]

		out, err := aead.Open(nil, nonce, msg, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)

		os.Exit(0)
	}
}