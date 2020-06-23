// https://golang.org/doc/effective_go.html#commentary
package twofer

// If provided with a name, returns "One for <name>, one for me. If <name> does not exist, returns "One for you, one for me."
func ShareWith(name string) string {
	var response string
	switch name {
		case "": response = "One for you, one for me."
		default: response = "One for " + name + ", one for me."
	}
	return response
}
