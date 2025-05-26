package structs

import (
	"fmt"
	"strconv"
)

type Account struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       uint8
}

func (a *Account) editProfile(edit map[string]string) error {

	var err uint8

	for elements := range edit {
		if edit[elements] == "" {
			err = 1
			break
		}

		switch elements {
		case "firstName":
			a.firstName = edit[elements]
		case "lastName":
			a.lastName = edit[elements]
		case "email":
			a.email = edit[elements]
		case "password":
			a.password = edit[elements]
		case "age":
			u, err := strconv.ParseUint(edit[elements], 10, 8)
			if err != nil {
				fmt.Printf("Error Conversion to uint8 %v", err)
			} else {
				a.age = uint8(u)
			}

		default:
			fmt.Print("No Such Data")
		}
	}
	if err != 1 {
		return nil
	} else {
		return fmt.Errorf("error, please try again")
	}
}

func Index() {

	acc := []Account{
		{
			firstName: "XQ",
			lastName:  "C",
			email:     "Xqc@gmail.com",
			password:  "123",
			age:       22,
		},
	}

	err := acc[0].editProfile(map[string]string{"password": "admin123", "age": "100"})

	if err == nil {
		fmt.Print(acc)
	} else {
		fmt.Print(err)
	}
}
