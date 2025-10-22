package main

import (
	"fmt"

	phonebook "github.com/tsuruevimran17/Phonebook/phoneBook"
)

func main() {
	//user := phonebook.Contact{ID: 89}
	list := phonebook.ContactList{}

	list.Add("y","z","x","Imran@mail.com",7,true)
	//list.AddEmail("Imran@mail.com")
	list.Print(7)
	fmt.Println("add: ", list.BlockList())
	list.UpdateEmail("Tsuruev@gmail.com",7)
	list.UpdateID(7,8)
	list.Print(8)
}
