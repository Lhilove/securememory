package main

import (
	"fmt"
	"simple-data-store/models"
	"simple-data-store/store"
	"strings"
)

func main() {

	fmt.Printf("Welcome to %s!\n", store.AppName)
	// this is to create users and print them out
	users := store.CreateUsers()
	store.PrintUsers(users)

	// filter by domain
	filtered := store.Filter(users, func(u models.User) bool {
		return strings.Contains(u.Email, "@apple.com")

	})
	fmt.Println("\nApple employees:")
	store.PrintUsers(filtered)

	// filter by age
	filtered = store.Filter(users, func(u models.User) bool {
		return u.Age < 20
	})
	fmt.Println("\nUsers younger than 20:")
	store.PrintUsers(filtered)

	// filter by name (findByName already does this ones job, this is just to demostrate firstclass function)
	// filtered = models.Filter(users, func(u models.User) bool {
	// 	return u.FirstName == "Jeff"
	// })
	// fmt.Println("\nUsers with the first name Jeff:")
	// models.Print	Users(filtered)

	// this is to delete user from the store
	fmt.Println("Enter email to delete: ")
	fmt.Scan(&store.Email)
	record, err := store.DeleteUser(users, models.User3, store.Email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("User with email %s deleted: %s %s, Email: %s, Age: %d\n", store.Email, record.FirstName, record.LastName, record.Email, record.Age)
	}

	// this is to update user email

	readded := store.UpdateUser(users, "jeff.alex@example.com", "jeff.okodua@example.com")
	fmt.Println("\nthe new user has been added:")
	store.PrintUsers(readded)

	// this is to demonstrate user input for name search
	fmt.Println("Enter name to search: ")
	fmt.Scan(&store.Name)
	result, err := store.FindByName(users, store.Name)
	if err != nil {
		fmt.Printf("User with name %s not found\n", store.Name)
	} else {
		fmt.Printf("User with name %s found: %s %s, Email: %s, Age: %d\n", store.Name, result.FirstName, result.LastName, result.Email, result.Age)
	}

}
