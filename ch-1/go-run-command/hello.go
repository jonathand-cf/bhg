package main

import "fmt"

// structs:
type Person struct {
	name string
	age  int
}

func (p *Person) SayHello() {
	fmt.Println("hello", p.name)
}

func main() {
	// slices and maps:
	m := map[string]string{
		"name": "Alice",
		"city": "Oslo",
	}
	fmt.Println("Map:", m)

	s := []string{"apple", "banana", "cherry"}
	fmt.Println("Slice:", s)

	// other way:
	var d = make([]string, 0)
	var c = make(map[string]string)
	d = append(d, "apple")
	c["name"] = "Apple"
	fmt.Println(d)
	fmt.Println(c["name"])

	// pointers:
	var count = int(42)
	ptr := &count    //the pointer
	fmt.Println(ptr) // returns address of the count variable
	*ptr = 100       // deference the address
	fmt.Println(count)

	//Structs
	var guy = new(Person)
	guy.name = "Bob"
	guy.SayHello()
}
