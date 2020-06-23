package main

import (
	"html/template"
	"io/ioutil"
	"os"
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

func createFileFromTemp(filename string, data string) {
	content := content{Text: data}
	temp := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = temp.Execute(os.Stdout, content)

	if err != nil {
		panic(err)
	}

	file, err := os.Create("first-post.html")

	if err != nil {
		panic(err)
	}

	err = temp.Execute(file, content)

	if err != nil {
		panic(err)
	}
}

func main() {
	createFileFromTemp("template.tmpl", readFile("first-post.txt"))
}
