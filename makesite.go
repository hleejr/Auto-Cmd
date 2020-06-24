package main

import (
	"flag"
	"fmt"
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

func createHTML(filename string) {
	data := readFile(filename)
	content := content{Text: data}
	temp := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	var err error
	err = temp.Execute(os.Stdout, content)

	if err != nil {
		panic(err)
	}

	inputName := strings.Split(filename, ".")
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

func createFileFromTemp(filename string) {

	inputFile := flag.String("file", "", "what file are you trying to convert?")
	inputDir := flag.String("dir", "", "In what directory are your text files?")
	flag.Parse()

	if *inputFile != "" {

		createHTML(*inputFile)

	} else if *inputDir != "" {

		files, err := ioutil.ReadDir(*inputDir)

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), ".txt") {
				createHTML(file.Name())
			}
		}
	} else {
		fmt.Println("You must specify a file or directory using the --file or --dir flags")
	}
}

func main() {
	createFileFromTemp("template.tmpl")
}
