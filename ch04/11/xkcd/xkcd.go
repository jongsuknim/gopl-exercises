package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Transcript string
}

func main() {
	scaner := bufio.NewScanner(os.Stdin)

	for scaner.Scan() {
		n := scaner.Text()
		f, err := os.Open("../downloader/" + n)
		if err != nil {
			fmt.Printf("%s open error %q", n, err)
			continue
		}

		var d Data
		if err := json.NewDecoder(f).Decode(&d); err != nil {
			f.Close()
			continue
		} else {
			fmt.Println(d.Transcript)
		}
	}
}
