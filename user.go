// register new user
package main

import (
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username  string
	Folders   map[string]*Folder
}

var users = make(map[string]*User)

func registerUser(username string) {
	// check whether the user exist
	username = strings.ToLower(username) // case insensitive
	if _, exists := users[username]; exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "has already existed.")
		return
	}

	// create a new user
	users[username] = &User{
		Username:  username,
		Folders:   make(map[string]*Folder),
	}
	fmt.Printf("Add %s successfully.\n", username)
}
