package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path"
)

const algDesc = `The algorithm to use when generating a key pair. Supported algorithms are: ed25519.`

var (
	alg        string
	folderPath string
)

func main() {

	// Parse the command line arguments
	flag.StringVar(&alg, "alg", "", algDesc)

	flag.StringVar(&folderPath, "folder", ".", "The folder to save the keys")

	flag.Parse()

	switch alg {
	case "ed25519":
		// Generate a new key pair
		pub, pri, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated a new key pair for ed25519")

		pub = []byte(base64.StdEncoding.EncodeToString(pub))
		pri = []byte(base64.StdEncoding.EncodeToString(pri))

		if err := os.WriteFile(path.Join(folderPath, "private.key"), pri, 0644); err != nil {
			panic(err)
		}

		if err := os.WriteFile(path.Join(folderPath, "public.key"), pub, 0644); err != nil {
			panic(err)
		}

		fmt.Println("The public key is saved in", path.Join(folderPath, "public.key"))
		fmt.Println("The private key is saved in", path.Join(folderPath, "private.key"))
	default:
		fmt.Println("The algorithm is not supported")
		fmt.Println("Use --help to see the supported algorithms")
	}

}
