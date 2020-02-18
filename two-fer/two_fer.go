/*
Package twofer implements Two Fer function
*/
package twofer

import "fmt"

// ShareWith returns Two Fer string as per given name
func ShareWith(name string) string {
	var who string
	if len(name) > 0 {
		who = name
	} else {
		who = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", who)
}
