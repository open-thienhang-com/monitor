package main

import "fmt"

// Base type with a method
type Base struct {
}

func (b *Base) SomeMethod() {
	fmt.Println("Base method")
}

// Derived type embedding the Base type
type Derived struct {
	Base
}

// Override the SomeMethod in the Derived type
func (d *Derived) SomeMethod() {
	fmt.Println("Derived method")
}

func main() {
	// Create an instance of the Derived type
	d := &Derived{}

	// Call the overridden method
	d.SomeMethod() // Output: Derived method
}
