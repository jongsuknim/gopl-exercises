package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hash = flag.String("h", "SHA256", "hash type")

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	switch *hash {
	case "SHA384":
		c := sha512.Sum384([]byte(text))
		fmt.Printf("%x", c)
	case "SHA512":
		c := sha512.Sum512([]byte(text))
		fmt.Printf("%x", c)
	default:
		c := sha256.Sum256([]byte(text))
		fmt.Printf("%x", c)
	}
}
