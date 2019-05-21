package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/daiLlew/BDDTest/story"
	"github.com/pkg/errors"
)

const (
	genPath      = "/src/github.com/daiLlew/BDDTest"
	templatePath = "template/test_template.tmpl"
	goFileEnv    = "GOFILE"
	goExt        = ".go"
)

func main() {
	callerPath := path.Join(build.Default.GOPATH, genPath)
	templatePath := path.Join(callerPath, templatePath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(errors.Wrap(err, "error loading template file"))
	}

	filename := getFilename()

	storyFile, err := story.Parse(filename, os.Getenv("GOPACKAGE"))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, storyFile)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s_test.go", filename), buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func getFilename() string {
	filename := os.Getenv(goFileEnv)
	if filename == "" {
		panic(errors.New("could not find GOFILE"))
	}
	return strings.Replace(filename, goExt, "", 1)
}
