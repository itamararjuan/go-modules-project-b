package b

import "fmt"

// SayHelloInEnglish - Hello saying function
// Use this method to return a text saying Hello in English
func SayHelloInEnglish(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// SayHelloInSpanish - Says hello in Spanish
func SayHelloInSpanish() string {
	return "Hola"
}
