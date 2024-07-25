package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

// create a new file in [username]/[foldername]
func createFile(username, foldername, filename, description string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)]	// case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}
	// check whether the folder exist
	folder, exists := user.Folders[strings.ToLower(foldername)]	// case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", foldername, "doesn't exist.")
		return
	}
	// check whether the file exist
	filename = strings.ToLower(filename) // case insensitive
	if _, exists := folder.Files[filename]; exists {
		fmt.Fprintln(os.Stderr, "Error: The", filename, "has already existed.")
		return
	}
	folder.Files[filename] = &File{
		Name:        filename,
		Description: description,
		CreatedAt:   time.Now(),
	}
	fmt.Printf("Create %s in %s/%s successfully.\n", filename, username, foldername)
}

func deleteFile(username, foldername, filename string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)] // case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}
	// check whether the folder exist
	folder, exists := user.Folders[strings.ToLower(foldername)] // case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", foldername, "doesn't exist.")
		return
	}
	// check whether the file exist
	filename = strings.ToLower(filename)
	if _, exists := folder.Files[filename]; !exists {
		fmt.Fprintln(os.Stderr, "Error: The", filename, "doesn't exist.")
		return
	}

	// delete the file
	delete(folder.Files, filename)
	fmt.Printf("Delete %s successfully.\n", filename)
}

func listFiles(username, foldername, sortBy, order string) {
	// check whether the user exist
	user, exists := users[strings.ToLower(username)] // case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", username, "doesn't exist.")
		return
	}

	// check whether the folder exist
	folder, exists := user.Folders[strings.ToLower(foldername)] // case insensitive
	if !exists {
		fmt.Fprintln(os.Stderr, "Error: The", foldername, "doesn't exist.")
		return
	}

	// check whether The folder have any files.
	if len(folder.Files) == 0 {
		fmt.Println("Warning: The folder doesn't have any files.")
		return
	}

	files := make([]*File, 0, len(folder.Files))
	for _, file := range folder.Files {
		files = append(files, file)
	}
	//sort the files by name/createdtime in asc/desc
	switch sortBy {
	// sort with name
	case "name":
		sort.SliceStable(files, func(i, j int) bool {
			if order == "asc" {
				return files[i].Name < files[j].Name
			}
			return files[i].Name > files[j].Name
		})
	// sort with created time
	case "created":
		sort.SliceStable(files, func(i, j int) bool {
			if order == "asc" {
				return files[i].CreatedAt.Before(files[j].CreatedAt)
			}
			return files[i].CreatedAt.After(files[j].CreatedAt)
		})
	default:
		sort.SliceStable(files, func(i, j int) bool {
			return files[i].Name < files[j].Name
		})
	}
	for _, file := range files {
		fmt.Printf("%s %s %s\n", file.Name, file.Description, file.CreatedAt.Format(time.DateTime))
	}
}
