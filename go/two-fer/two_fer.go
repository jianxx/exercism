/*
Package twofer implements a simple function for two for one. One for you and one for me..

Given a name, return a string with the message: One for X, one for me.
Where X is the given name.

However, if the name is missing, return the string: One for you, one for me.
*/
package twofer

import "fmt"

// ShareWith implements the function: Given a name, return a string with the message.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
