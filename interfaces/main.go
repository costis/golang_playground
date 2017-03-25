package main

import "fmt"

type Speaker interface {
	Speak(s string) string
}

type Person struct {
	Name     string
	Lastname string
}

func (u Person) Speak(s string) string {
	return s
}

// A func with a T value receiver can receive both type values and type pointers.
func main() {
	uPointer := &Person{"Cos", "Pan"}
	uValue := Person{"Nick", "And"}

	fmt.Println(uPointer.Speak("I am a a pointer to a type value "))
	fmt.Println(uValue.Speak("I am a type value"))
}
