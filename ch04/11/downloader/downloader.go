package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Data struct {
	Transcript string
}

var url = "https://xkcd.com/"

func download(n int) (*Data, error) {
	q := url + strconv.Itoa(n) + "/info.0.json"
	resp, err := http.Get(q)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Data
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	w, err := os.Create(strconv.Itoa(n))
	if err != nil {
		return nil, err
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.Close()
		return nil, err
	}
	w.Close()

	return &result, nil
}

func main() {
	for i := 1; i <= 10; i++ {
		if _, err := download(i); err != nil {
			fmt.Printf("error %q\n", err)
		} else {
			fmt.Printf("%d done\n", i)
		}
	}
}
