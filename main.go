package main

import (
	"io"
	"os"

	"github.com/diop/indigo/primitive"
)

func main() {
	inFile, err := os.Open("images/in/afro-futurism.png")
	if err != nil {
		panic(err)
	}
	out, err := primitive.Transform(inFile, 50)
	if err != nil {
		panic(err)
	}
	os.Remove("out.png")
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	io.Copy(outFile, out)
}
