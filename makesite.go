package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"text/template"
)

func makeStaticPage(dir string, filename string) {
	if dir != "." {
		filename = dir + "/" + filename
	}
	fmt.Println(filename)
	FileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	sz := len(filename)
	templateFile := filename[:sz-3] + "html"

	var b bytes.Buffer
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(&b, string(FileContents))

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(templateFile, b.Bytes(), 777)
	if err != nil {
		panic(err)
	}

}

func main() {
	filePtr := flag.String("flag", "defaultValue", "Pass the file to open.")
	dirPtr := flag.String("dir", "", "Pass the directory to look in")

	flag.Parse() // parse the flags

	if *filePtr != "defaultValue" {
		makeStaticPage(".", *filePtr)
	}

	if *dirPtr != "" {
		files, _ := ioutil.ReadDir(*dirPtr)
		for _, file := range files {
			sz := len(file.Name())
			if sz > 4 {
				if file.Name()[sz-4:sz] == ".txt" {
					makeStaticPage(*dirPtr, file.Name())
				}
			}
		}
	}
}
