package main

import (
	"fmt"
	"os"
	"crypto/rsa"
	"crypto/rand"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"encoding/json"
)
func init() {

}

func main() {
	if os.Args == nil || len(os.Args) == 0  {
		fmt.Println("Nor Argument provided")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "publish":
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("-> Searching for publish.json in directory: ",pwd)
		file, fileErr := ioutil.ReadFile(pwd+"/publish.json")
		if fileErr != nil {
			fmt.Println("ERROR: ",err)
			os.Exit(1)
		}

		var publish Meshcraft
		if err := json.Unmarshal(file, &publish); err != nil {
			fmt.Println("ERROR: ",err)
			os.Exit(1)
		}

		os.Exit(0)
	case  "init":
		fmt.Println("-> Generate publishing key. Notice: This can take a moment")
		privKey, err:= rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			fmt.Println("ERROR: Creating publishing key failed!")
			os.Exit(1)
		}

		pemBlock := &pem.Block{
			Type : "RSA PRIVATE KEY",
			Bytes : x509.MarshalPKCS1PrivateKey(privKey),
		}

		data := pem.EncodeToMemory(pemBlock)

		fmt.Println("-> Write publishing key to disk")
		if err := ioutil.WriteFile("./publishing.key", data, 0755); err != nil {
			fmt.Println("ERROR: Writing generated publishing key failed!!")
			os.Exit(1)
		}

		os.Exit(0)
	default:
		fmt.Println("Unknown command")
	}
}