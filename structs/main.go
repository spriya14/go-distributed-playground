package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // This is struct embedding
}

func main() {
	//alex := person{"Alex", "Cooper"} // we are completely relying on the order struct is defined.
	// person struct now has one more field for contacts which we need to defined, the above method wont automatically assign a zero variable
	alex := person{firstName: "Alex", lastName: "Cooper", contact: contactInfo{email: "hola123@gmail.com", zip: 12333}}
	// Other way
	alan := person{firstName: "Alan", lastName: "Henderson"}
	fmt.Println(alex)
	fmt.Println(alan)

	var cooper person
	fmt.Println(cooper)         // this will be empty as we didn't assign any values.
	fmt.Printf("%+v\n", cooper) // this will tell us all the fields present in the struct

	// How do we update the properties of the fields
	//cooper.firstName = "Cooper"
	//cooper.lastName = "Liz"
	cooper = person{
		firstName: "Cooper",
		lastName:  "Brown",
		contact: contactInfo{
			email: "Hola123@gmail.com",
			zip:   9000,
		},
	}

	// go is pass by value language. when we do the below go creates a copy of the struct and then updates that copy. Picture attached for reference in pocRef-for-concepts folder.
	//cooper.updateName("Scarlett")
	// The way to get it to work how we expect and not update the copy is to do below using POINTER.
	// &variable -> gives memory address of the value this variable is pointing at. (refer to the image)
	cooperPointer := &cooper // & is nothing but the memory address.
	// VIP: Turn Address to value  -> *address
	// VIP: Turn value to address -> &value
	cooperPointer.updateFirstName("Scarlett")

	// VIP : Tip to Remember : If a method needs to modify data, use a pointer receiver. Go will automatically handle value vs pointer when calling it.
	// so we don't have to pass memory address like above but instead we can use the shortcut like below: Go automatically will insert the & and pass the memory address.
	cooper.updateFirstName("Millie Bobby")

	cooper.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	//IMP : p.firstName = newFirstName // we just updated the name and then printed it but noticed that the update did not happen -> This is where POINTER comes into picture
	pointerToPerson.firstName = newFirstName
}

// Function that has receivers, the receiver will be of type person and we will print out its values
// so we have a function that has a receiver -> (p person), (p person) means you can call this function on any value of type person,
// and then inside the body of the function you can reference that person as the variable p
func (p person) print() {
	//fmt.Printf("%+v", p)
	fmt.Println(p)
}

// Slice data structure works very different from Struct - refer to the diagram. the ptr to head refers to the same address where the data is stored.
// Reference Type -> reference to another data structure in memory.
// Several other types of ref type - > slices, maps, channels, pointers, functions --> Don't worry about pointers with these.
// Value Type : structs, int, float, bool, string --> Always use pointers to modify values within these in a function.
