package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
)

func main() {
	filePtr := flag.String("flag", "defaultValue", "Pass the file to open.")
	flag.Parse()
	FileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(&b, string(FileContents))

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*filePtr, b.Bytes(), 777)
	if err != nil {
		panic(err)
	}

}
