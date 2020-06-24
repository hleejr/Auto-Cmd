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

func createFileFromTemp(filename string) {

	inputFile := flag.String("file", "", "what file are you trying to convert?")
	inputDir := flag.String("dir", "", "In what directory are your text files?")
	flag.Parse()

	if *inputFile != "" {
		data := readFile(*inputFile)
		content := content{Text: data}
		temp := template.Must(template.New("template.tmpl").ParseFiles(filename))

		var err error
		err = temp.Execute(os.Stdout, content)

		if err != nil {
			panic(err)
		}

		inputName := strings.Split(*inputFile, ".")
		newName := inputName[0] + ".html"

		file, err := os.Create(newName)

		if err != nil {
			panic(err)
		}

		err = temp.Execute(file, content)

		if err != nil {
			panic(err)
		}
	} else if *inputDir != "" {
		files, err := ioutil.ReadDir(*inputDir)

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), ".txt") {
				data := readFile(file.Name())
				content := content{Text: data}
				temp := template.Must(template.New("template.tmpl").ParseFiles(filename))

				var err error
				err = temp.Execute(os.Stdout, content)

				if err != nil {
					panic(err)
				}

				inputName := strings.Split(file.Name(), ".")
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
		}
	} else {
		fmt.Println("You must specify a file or directory using the --file or --dir flags")
	}
}

func main() {
	createFileFromTemp("template.tmpl")
}
