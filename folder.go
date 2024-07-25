// create, delete and list the folders

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Folder struct {
	Name        string
	Description string
	Files       map[string]*File
	CreatedAt   time.Time
}

// create a new folder with [username]
func createFolder(username, foldername, description string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)]
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}

	// check whether the folder exist
	foldername = strings.ToLower(foldername) // case insensitive
	if _, exists := user.Folders[foldername]; exists {
		fmt.Fprintln(os.Stderr, "Error: The", foldername, "has already existed.")
		return
	}

	// creater a new folder 
	user.Folders[foldername] = &Folder{
		Name:        foldername,
		Description: description,
		Files:       make(map[string]*File),
		CreatedAt:   time.Now(),
	}
	fmt.Printf("Create %s successfully.\n", foldername)
}

// delete the folder with [username]
func deleteFolder(username, foldername string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)] // case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}

	// check whether the folder exist
	foldername = strings.ToLower(foldername) // case insensitive
	if _, exists := user.Folders[foldername]; !exists {
		fmt.Fprintln(os.Stderr, "Error: The", foldername, "doesn't exist.")
		return
	}

	// delete the folder 
	delete(user.Folders, foldername)
	fmt.Printf("Delete %s successfully.\n", foldername)
}

// list the folders with [username]
func listFolders(username, sortBy, order string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)]  // case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}

	// check whether the user have any folders
	if len(user.Folders) == 0 {
		fmt.Println("Warning: The", user, "doesn't have any folders.")
		return
	}

	folders := make([]*Folder, 0, len(user.Folders))
	for _, folder := range user.Folders {
		folders = append(folders, folder)
	}
	// sort the folders by name/createdtime in asc/desc
	switch sortBy {
	// sort with name
	case "name":
		sort.SliceStable(folders, func(i, j int) bool {
			if order == "asc" {    // sort in asc
				return folders[i].Name < folders[j].Name
			} else {               // sort in desc
				return folders[i].Name > folders[j].Name
			}
		})
	// sort with created time
	case "created":
		sort.SliceStable(folders, func(i, j int) bool {
			if order == "asc" {    // sort in asc
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			} else {                // sort in desc
				return folders[i].CreatedAt.After(folders[j].CreatedAt)
			}
		})
	default:
		sort.SliceStable(folders, func(i, j int) bool {
			return folders[i].Name < folders[j].Name
		})
	}
	for _, folder := range folders {
		fmt.Printf("%s %s %s %s\n", folder.Name, folder.Description, folder.CreatedAt.Format(time.DateTime), username)
	}
}

// rename the folder name from [foldername] to [newfoldername]
func renameFolder(username, foldername, newfoldername string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)]
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}
	// case insensitive
	foldername = strings.ToLower(foldername)
	newfoldername = strings.ToLower(newfoldername)
	// check whether the folder exist
	folder, exists := user.Folders[foldername]
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", foldername, "doesn't exist.")
		return
	}
	// check whether the new folder exist
	if _, exists := user.Folders[newfoldername]; exists {
		fmt.Fprintln(os.Stderr, "Error: The", newfoldername, "already exists.")
		return
	}
	// delete the old foldername and rename the folder with new name
	delete(user.Folders, foldername)
	folder.Name = newfoldername
	user.Folders[newfoldername] = folder
	fmt.Printf("Rename %s to %s successfully.\n", foldername, newfoldername)
}
