package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

const rootDirectory = "docs"
const templateFile = "docs/template.html"

var markdownProcessor = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithXHTML(),
	),
)

func main() {
	files := make([]string, 0)
	if err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && isMarkdown(path) {
			files = append(files, path)
		}

		return nil

	}); err != nil {
		panic(err)
	}

	// sort the file array
	sort.Strings(files)

	websiteBuffer := new(bytes.Buffer)

	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := websiteBuffer.WriteString("<section id=\"" + getName(file) + "\">\n"); err != nil {
			log.Fatal(err)
		}

		if err := markdownProcessor.Convert(content, websiteBuffer); err != nil {
			log.Fatal(err)
		}

		if _, err := websiteBuffer.WriteString("</section>\n"); err != nil {
			log.Fatal(err)
		}
	}

	buf := strings.ReplaceAll(string(templateContent), "{{ body }}", websiteBuffer.String())

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, err := f.WriteString(buf)
	if err != nil {
		log.Fatal(err)
	}

	f.Sync()

	log.Println("Finished. Wrote " + strconv.FormatInt(int64(n), 10) + " bytes.")
}

// based on CommonMark
func isMarkdown(filename string) bool {
	extName := filepath.Ext(filename)
	return extName == ".md" || extName == ".markdown"
}

func getName(filename string) string {
	return strings.Split(strings.Split(filename, "/")[1], ".")[0]
}
