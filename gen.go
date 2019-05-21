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
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	genPath      = "/src/github.com/daiLlew/BDDTest"
	templatePath = "template/test_template.tmpl"
	goFileEnv    = "GOFILE"
	storyExt     = ".story"
	goExt        = ".go"
)

type StoryFile struct {
	Scenario    string `yaml:"scenario"`
	Given       string `yaml:"given"`
	When        string `yaml:"when"`
	Then        string `yaml:"then"`
	And         string `yaml:"and"`
	Package     string
	GeneratedAt time.Time
	TestName    string
}

func main() {
	callerPath := path.Join(build.Default.GOPATH, genPath)
	templatePath := path.Join(callerPath, templatePath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(errors.Wrap(err, "error loading template file"))
	}

	filename := getFilename()

	story, err := ParseStory(filename)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, story)
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
		panic(errors.New("could not file GOFILE"))
	}
	return strings.Replace(filename, goExt, "", 1)
}

func ParseStory(filename string) (*StoryFile, error) {
	b, err := ioutil.ReadFile(filename + storyExt)
	if err != nil {
		return nil, err
	}

	var s StoryFile
	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	s.Package = os.Getenv("GOPACKAGE")
	s.GeneratedAt = time.Now()
	s.TestName = strings.Replace(strings.Title(strings.ToLower(s.Scenario)), " ", "", -1)

	return &s, nil
}
