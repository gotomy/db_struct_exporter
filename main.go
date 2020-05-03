package main

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func main() {
	body, err := ioutil.ReadFile("test.md")
	if err != nil {
		panic(err)
	}
	output := blackfriday.Run(body, blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(
		blackfriday.HTMLRendererParameters{
			Flags: blackfriday.TOC,
		},
		)))

	err = ioutil.WriteFile("test.html", output, 0777)
	if err != nil {
		panic(err)
	}
}
