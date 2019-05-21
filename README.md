# BDDTest

## Example

`hello_world.go` contains the following generate statement:
```
//go:generate go run ../gen.go
```
In the same package we have `hello_world.story` which contains:
```
scenario: Hello world success
given: something happens
when: some condition is met
then: something happens
and: so do something else
```
Running `go generate` will generate the following test skeleton:
```go
// Code template generated by go generate: DO NOT EDIT
// Generated at 2019-05-21 12:57:28.612662 +0100 BST m=+0.001110684
package example

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Scenario: Hello world success
// Given something happens
// When some condition is met
// Then something happens
// And so do something else
func Test_HelloWorldSuccess(t *testing.T) {

	Convey("Given something happens", t, func() {

		Convey("When some condition is met", func() {

			Convey("Then something happens", func() {
			})

			Convey("And so do something else", func() {
			})
		})
	})
}
```
