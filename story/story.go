package story

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const storyExt = ".story"

type Story struct {
	Description string   `yaml:"description"`
	Given       string   `yaml:"given"`
	When        string   `yaml:"when"`
	Then        string   `yaml:"then"`
	And         []string `yaml:"and"`
}

type Stories struct {
	Scenarios   []*Story `yaml:"stories"`
	Package     string
	GeneratedAt time.Time
}

func Parse(filename string, pkg string) (*Stories, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s%s", filename, storyExt))
	if err != nil {
		return nil, err
	}

	var s Stories
	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, err
	}

	s.GeneratedAt = time.Now()
	s.Package = pkg
	return &s, nil
}

func (s *Story) TestName() string {
	return strings.Replace(strings.Title(strings.ToLower(s.Description)), " ", "", -1)
}

func (s *Story) Comments() string {
	comments := []string{
		fmt.Sprintf("// Scenario: %s", s.Description),
		fmt.Sprintf("// Given %s", s.Given),
		fmt.Sprintf("// When %s", s.When),
		fmt.Sprintf("// Then %s", s.Then),
	}

	if len(s.And) > 0 {
		for _, and := range s.And {
			comments = append(comments, fmt.Sprintf("// And %s", and))
		}
	}
	return strings.Join(comments, "\n")
}
