package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type content struct {
	Text string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func createFileFromTemp(filename string) {

	input := flag.String("file", "first-post.txt", "what file are you trying to convert?")
	flag.Parse()

	data := readFile(*input)
	content := content{Text: data}
	temp := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = temp.Execute(os.Stdout, content)

	if err != nil {
		panic(err)
	}

	inputName := strings.Split(*input, ".")
	newName := inputName[0] + ".html"

	file, err := os.Create(newName)

	if err != nil {
		panic(err)
	}

	err = temp.Execute(file, content)

	if err != nil {
		panic(err)
	}
}

func main() {
	createFileFromTemp("template.tmpl")
}
