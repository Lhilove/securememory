# Simple Data Store - Go Learning Project

This is my first golang project, the goal is to put everything I have learnt into a working software. 

## What it does
- Stores user records using structs.
- Groups users into a slice.
- Loops through and prints each user.
- Filters users by email domain and age using string matching.
- Every logical block became its own funtiion
- Creates, updates and delete users
- Find user by name (using fmt.Scan(&name))
- Input validation
- Renamed to SecureMemory
- Added minEmailLength in deleteUser
- Project Structure:
    securememory/
    ├── main.go
    ├── models/
    │   └── user.go
    └── store/
        └── store.go

- Implemented Role Based Access Control; only admin can delete users

## Concepts demonstrated
- Structs and slices
- For Loops and range 
- Conditionals
- String packages
- Error handling
- First class functions, anonymous functions
- Pointers
- User Input
- Logical Operators
- Package level variables
- Go packages

## Part of
SecureMemory is the first building block of SecureFlow
an end-to-end secure fintech microservices platform on AWS 
with real-world SSDLC.