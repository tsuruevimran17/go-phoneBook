package phonebook

import "fmt"

type Contact struct {
	FirstName string
	LastName  string
	SurName   string
	ID        int
	Email     string
	Block     bool
}

type ContactList struct {
	contacts []Contact
}

func (r *ContactList) Add(firstName, lastName, surName, email string, id int, block bool) []Contact {
	r.contacts = append(r.contacts, Contact{FirstName: firstName, LastName: lastName, ID: id, SurName: surName, Email: email, Block: block})
	return r.contacts
}
//func (r *ContactList) AddEmail(email string) []Contact {
	//r.contacts = append(r.contacts, Contact{Email: email})
	//return r.contacts
//}

func (r *ContactList) Remove(id int) string{
	for i, v := range r.contacts {
		if v.ID == id {
			r.contacts = append(r.contacts[:i], r.contacts[i+1:]...)
			return "пользователь удален"
		}
	}
	return "пользователь заблокирован"
}

func (r ContactList) Print(id int) {
	for i, v := range r.contacts {
		if v.ID == id{
			fmt.Println(r.contacts[i])
		}
	}
}

func (r *ContactList) Block(id int) string {
	for i, v := range r.contacts {
		if v.ID == id{
			r.contacts[i].Block = true
			return "пользователь заблокирован"
		}
	}
	return "пользователь не найден"
}

func (r ContactList) Search(email string)  {
	for i, v := range r.contacts {
		if v.Email == email{
			fmt.Println(r.contacts[i])
		}
	}
}

func (r ContactList) GetFullName(firstName,lastName,surName string) bool{
	fullName := firstName + lastName + surName
	for i, v := range r.contacts {
		if v.FirstName + v.LastName +v.SurName == fullName{
			fmt.Println(r.contacts[i])
			return true
		}
	}
	return false

}

func (r *ContactList) UpdateID(id, newID int)string{
	for i , v := range r.contacts{
		if v.ID == id{
			r.contacts[i].ID = newID
			return "ID изменен на новый"
		}
	}
	return "пользователь не найден"
	
}

func (r *ContactList) UpdateEmail(newEmail string, id int) string{
	for  i, v := range r.contacts{
		if v.ID == id{
			r.contacts[i].Email = newEmail
			return "Email изменен"
		}
	}
	return "пользозователь не найден"
}

func (r *ContactList) BlockList() []string{
	result := []string{}
	for i, v := range r.contacts{
		if v.Block == true{
			result = append(result, r.contacts[i].FirstName)
		}
	} 
	return result
}




