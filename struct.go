package main

import "fmt"

// Define a struct
type Assign1 struct {
	Name  string
	Value int
}

// Method to modify the struct
func (ms *Assign1) ModifyStruct(newName string, newValue int) {
	ms.Name = "Chibuoke"
	ms.Value = 50
}

// Method to access the struct without modifying it
func (ms Assign1) AccessStruct() {
	fmt.Printf("Name: %s, Value: %d\n", ms.Name, ms.Value)
}

func main() {
	// Create an instance of the struct
	Database := Assign1{Name: "Uchechukwu", Value: 42}

	// Print the initial values
	fmt.Println("Initial values:")
	Database.AccessStruct()
	// Modify the struct using the method
	Database.ModifyStruct("Updated", 99)

	// Print the modified values
	fmt.Println("Modified values:")
	Database.AccessStruct()
}
