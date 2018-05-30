package util

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
)

func LoadBase64CertificateAuthority(path string) (string, error) {
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[CA-debug] ")

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
