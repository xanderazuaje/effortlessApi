package main

import (
	"github.com/xanderazuaje/effortlessApi/lexer"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("syntax")
	if err != nil {
		log.Fatal(err)
	}
	tokens := lexer.Tokenize(string(b))
	for _, t := range tokens {
		t.Debug()
	}
}
