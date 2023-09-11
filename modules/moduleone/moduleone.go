/*
Go code is grouped into packages, and packages are grouped into modules.
Module specifies dependencies needed to run the code, including
the Go version and the set of other modules it requires.
*/
package moduleone

import "fmt"

// DoSomething returns a greeting for the named person.
func DoSomething(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}