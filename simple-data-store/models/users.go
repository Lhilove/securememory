package models

// this is seperating users and admins into different structs, this is to demonstrate composition and separation of concerns, it also allows us to add more fields to the admin struct without affecting the user struct
type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
	Password  []string `json:"-"` // this will hide the password field
	Role      Role
}

var User1 = User{
	FirstName: "Jeff",
	LastName:  "Alex",
	Email:     "jeff.alex@example.com",
	Age:       21,
	Password:  []string{"password123"},
	Role:      UserRole,
}

var User2 = User{
	FirstName: "Tobi",
	LastName:  "Alabi",
	Email:     "tobi.alabi@gmail.com",
	Age:       29,
	Password:  []string{"password001"},
	Role:      UserRole,
}

var User3 = User{
	FirstName: "John",
	LastName:  "Sanjay",
	Email:     "john.sanjay@apple.com",
	Age:       18,
	Password:  []string{"password202"},
	Role:      AdminRole,
}
