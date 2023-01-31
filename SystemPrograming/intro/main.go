package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("-------header---------")
	content := bytes.NewBufferString("-------content---------")
	footer := bytes.NewBufferString("-------fotter---------")

	reader := io.MultiReader(header, content, footer)

	io.Copy(os.Stdout, reader)
	// -------header----------------content----------------fotter---------%
}
