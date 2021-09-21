package main

import (
	"bufio"
	"crypto-stuff/crypto"
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const (
	encSuffix = ".enc"
)

var (
	dec = flag.Bool("d", false, "decrypt file")
	enc = flag.Bool("e", false, "encrypt file")

	filepath = flag.String("path", "", "path to file to be encrypted")
	encMode  = flag.String("mode", "", "ctr|gcm")

	op    func([]byte) []byte
)

func init() {
	flag.Parse()
	if !*enc && !*dec {
		panic("you need to set enc|dec flag")
	}
}

func main() {

	f1, err := os.Open(*filepath)
	if err != nil {
		log.Fatalf("cannot open src file: %+v\n", err)
	}
	defer f1.Close()
	var newPath string
	if *enc {
		newPath = *filepath + encSuffix
	} else if *dec {
		if path.Ext(*filepath) == encSuffix {
			newPath = strings.Replace(*filepath, encSuffix, "", -1)
		} else {
			newPath = *filepath + ".txt"
		}
	} else {
		newPath = ""
	}
	f2, err := os.Create(newPath)
	if err != nil {
		log.Fatalf("cannot open dst file: %+v\n", err)
	}
	defer f2.Close()
	//fmt.Print("key: ")
	//key, err := term.ReadPassword(0)
	key := []byte("MatejKarolcik123")
	if err != nil {
		log.Fatalf("cannot read password: %+v\n", err)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cannot use key for cipher: %+v\n", err)
	}

	e := crypto.NewEcryptor(c, crypto.Mode(*encMode))
	if *dec {
		op = e.Decrypt
	} else if *enc {
		op = e.Encrypt
	}

	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		var b []byte
		if *dec {
			b, err = hex.DecodeString(scanner.Text())
			if err != nil {
				log.Fatalf("cannot decode text from hex: %+v\n", err)
			}
			_, err = fmt.Fprintf(f2, "%s%s", op(b), "\n")
			if err != nil {
				log.Fatalf("cannot write decrypted data to file: %+v\n", err)
			}
		} else if *enc {
			b = scanner.Bytes()
			_, err = fmt.Fprintf(f2, "%x%s", op(b), "\n")
			if err != nil {
				log.Fatalf("cannot write encrypted data to file: %+v\n", err)
			}
		}

	}
}
