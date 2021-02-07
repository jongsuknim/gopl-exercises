package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github/jongsuknim/gopl-exercises/ch07/eval"
)

func cal(w http.ResponseWriter, req *http.Request) {
	exprStr := req.FormValue("expr")
	if exprStr == "" {
		http.Error(w, "no expr", http.StatusBadRequest)
	}

	var env = eval.Env{}
	for k, v := range req.Form {
		if k == "expr" {
			continue
		}
		value, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		env[eval.Var(k)] = value
	}

	expr, err := eval.Parse(exprStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%.6g\n", expr.Eval(env))

	// fmt.Println(exprStr, env)
}

func main() {
	http.HandleFunc("/cal", cal)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
