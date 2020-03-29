package greetings

var greetingstring = "Hello Its my world"

// PrintGreetings fine; I am commenting
func PrintGreetings(name string) string {
	return greetingstring + "-" + name
}
