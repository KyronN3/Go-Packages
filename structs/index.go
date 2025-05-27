package structs

import (
	"fmt"
	"strconv"
	"strings"
)

type Account struct {
	id        uint64
	firstName string
	lastName  string
	email     string
	password  string
	age       uint8
}

func (a *Account) editProfile(edit map[string]string, er error) error {

	var err uint8

	if er != nil {
		return fmt.Errorf("argument should not be empty or length(0)")
	}

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
		return fmt.Errorf("error, Please Try Again")
	}
}

func ConcatMulti(text []byte) string {
	var concat strings.Builder

	for _, elements := range text {
		concat.WriteString(strconv.FormatUint(uint64(elements), 10))
	}
	return concat.String()
}

func CreateAccount(fname string, lname string, email string, pass string, age uint8) (Account, error) {

	if len(fname) == 0 || len(lname) == 0 || len(email) == 0 || len(pass) == 0 {
		return Account{}, fmt.Errorf("error, Please Try Again")
	} else {
		u, _ := strconv.ParseUint(ConcatMulti([]byte{fname[0], lname[0], email[0], pass[0], age}), 10, 64)
		return Account{
			id:        u,
			firstName: fname,
			lastName:  lname,
			email:     email,
			password:  pass,
			age:       age,
		}, nil
	}
}

func Index() {

	acc, er := CreateAccount("Ass", "m", "ass@gmail.com", "123", 23)

	err := acc.editProfile(map[string]string{"password": "admin123", "age": "100"}, er)

	if err == nil {
		fmt.Print(acc)
	} else {
		fmt.Print(er)
	}
}
