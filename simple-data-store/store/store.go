package store

import (
	"fmt"
	"simple-data-store/models"
	"strings"
)

// this variable is for findByName, deleteUser
const MinEmailLength = 6 // minimum length for a valid email (e.g., a@b.c)
const AppName = "SecureStore"

func CreateUsers() []models.User {
	users := []models.User{models.User1, models.User2, models.User3}
	return users
}

func PrintUsers(users []models.User) {
	for _, value := range users {
		fmt.Printf("Name: %s %s, Email: %s, Age: %d\n", value.FirstName, value.LastName, value.Email, value.Age)
	}
}

// this is a firstclass function that allows you filter by name, domain or age instead of writing every filter function
func Filter(users []models.User, fn func(models.User) bool) []models.User {
	var filtered []models.User
	// using range instead of index loop since we do not need the postion of each slice
	for _, value := range users {
		if fn(value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}

// func printAppleUsers(users []user) {
// 	for _, value := range users {
// 		fmt.Printf("user with email %s is an apple employee\n", value.email)
// 	}
// }

// only admins can delete users, this is to demonstrate role based access control, it also prevents auth bypass via email manipulation since only admins can delete users and the email input is validated
func DeleteUser(users []models.User, User models.User, Email string) (models.User, error) {

	//user input validation to prevent auth bypass via email manipulation.
	isAdmin := User.Role == models.AdminRole

	if !isAdmin {
		return models.User{}, fmt.Errorf("only admins can delete users")
	}
	isValidEmail := strings.Contains(Email, "@") && len(Email) >= MinEmailLength
	if !isValidEmail {
		return models.User{}, fmt.Errorf("invalid email format")
	}
	for _, u := range users {
		if u.Email == Email {
			return u, nil // return the deleted user
		}

	}
	return models.User{}, fmt.Errorf("user not found")
}

// updates the email of an existing user identified by their current email
func UpdateUser(users []models.User, Email string, NewEmail string) []models.User {
	var readded []models.User
	for _, x := range users {
		if x.Email == Email {
			x.Email = NewEmail
		}
		readded = append(readded, x)
	}
	return readded
}

// added pointer receiver to updateUser function to modify the original user in the slice instead of creating a new slice with updated email
func UpdateUserPointer(users []models.User, Email string, NewEmail string) {
	for i := range users {
		if users[i].Email == Email {
			users[i].Email = NewEmail
		}
	}
}

// returns the user if found, returns an error if not found
// caller must check err before using the returned user
func FindByName(users []models.User, Name string) (models.User, error) {
	var names []models.User
	for _, value := range users {
		// fix case insensitive name search, prevents auth bypass via case manipulation
		if strings.EqualFold(value.FirstName, Name) || strings.EqualFold(value.LastName, Name) {
			names = append(names, value)
		}
	}

	//user input validation to verify the inputed name is not empty
	if len(names) == 0 {
		return models.User{}, fmt.Errorf("user not found")
	}
	return names[0], nil // return the first user found, you can modify this to return all users found if needed
}
