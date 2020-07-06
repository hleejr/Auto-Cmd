package main

// GO111MODULE=""
// GOPATH="/Users/hennybowejr/go"
// GOROOT="/usr/local/go"

// Afrikaans	af
// Irish	ga
// Albanian	sq
// Italian	it
// Arabic	ar
// Japanese	ja
// Azerbaijani	az
// Kannada	kn
// Basque	eu
// Korean	ko
// Bengali	bn
// Latin	la
// Belarusian	be
// Latvian	lv
// Bulgarian	bg
// Lithuanian	lt
// Catalan	ca
// Macedonian	mk
// Chinese Simplified	zh-CN
// Malay	ms
// Chinese Traditional	zh-TW
// Maltese	mt
// Croatian	hr
// Norwegian	no
// Czech	cs
// Persian	fa
// Danish	da
// Polish	pl
// Dutch	nl
// Portuguese	pt
// English	en
// Romanian	ro
// Esperanto	eo
// Russian	ru
// Estonian	et
// Serbian	sr
// Filipino	tl
// Slovak	sk
// Finnish	fi
// Slovenian	sl
// French	fr
// Spanish	es
// Galician	gl
// Swahili	sw
// Georgian	ka
// Swedish	sv
// German	de
// Tamil	ta
// Greek	el
// Telugu	te
// Gujarati	gu
// Thai	th
// Haitian Creole	ht
// Turkish	tr
// Hebrew	iw
// Ukrainian	uk
// Hindi	hi
// Urdu	ur
// Hungarian	hu
// Vietnamese	vi
// Icelandic	is
// Welsh	cy
// Indonesian	id
// Yiddish	yi

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bregydoc/gtranslate"
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

func translate(content string, input string, desired string) string {
	translated, err := gtranslate.TranslateWithParams(
		content,
		gtranslate.TranslationParams{
			From: input,
			To:   desired,
		},
	)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s: %s | %s: %s \n", input, content, desired, translated)
	// en: Hello World | ja: こんにちは世界
	return translated
}

func createHTML(filename string, input string, desired string) {
	data := readFile(filename)
	if input != "" {
		data = translate(data, input, desired)
	}
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
	inputLang := flag.String("lang", "", "Do you need to translate these files? What language are they in now? See language abbreviations for appropriate entries")
	desired := flag.String("trans", "", "What language do you want ot translate to? See language abbreviations for appropriate entries. EX) ja for Japanese")
	flag.Parse()

	if *inputFile != "" {

		createHTML(*inputFile, *inputLang, *desired)

	} else if *inputDir != "" {

		files, err := ioutil.ReadDir(*inputDir)

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), ".txt") {
				createHTML(file.Name(), *inputLang, *desired)
			}
		}
	} else {
		fmt.Println("You must specify a file or directory using the --file or --dir flags")
	}
}

func main() {
	createFileFromTemp("template.tmpl")
}
