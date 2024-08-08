package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	// "virtualFileSystem/fileops"
)

var validNamePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,30}$`)

// FileOperations defines methods for file operations.
type FileOperations interface {
	createFile(username, foldername, filename, description string) error
	deleteFile(username, foldername, filename string) error
}
var fileOps FileOperations

func main() {
	fileOps = &RealFileOperations{}  
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		scanner.Scan()
		command := scanner.Text()
		handleCommand(command)
	}
}

func handleCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	// check the command
	switch parts[0] {
	case "register":
		if len(parts) != 2 {
			fmt.Fprintln(os.Stderr, "Usage: register [username]")
			return
		}
		// chech whether username contain invalid chars
		if !isValidName(parts[1]) {
			fmt.Fprintln(os.Stderr, "Error: The", parts[1], "contain invalid chars.")
			return
		}
		registerUser(parts[1])

	case "create-folder":
		if len(parts) != 3 && len(parts) != 4 {
			fmt.Fprintln(os.Stderr, "Usage: create-folder [username] [foldername] [description]?")
			return
		}
		// chech whether folder contain invalid chars
		if !isValidName(parts[2]) {
			fmt.Fprintln(os.Stderr, "Error: The", parts[2], "contain invalid chars.")
			return
		}
		// description is optional
		description := ""
		if len(parts) > 3 {
			description = strings.Join(parts[3:], " ")
		}
		createFolder(parts[1], parts[2], description)

	case "delete-folder":
		if len(parts) != 3 {
			fmt.Fprintln(os.Stderr, "Usage: delete-folder [username] [foldername]")
			return
		}
		deleteFolder(parts[1], parts[2])

	case "list-folders":
		if len(parts) < 2 || len(parts) > 4 {
			fmt.Fprintln(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
			return
		}
		// Default: sort the list by [foldername] in ascending order.
		sortBy := "name"
		order := "asc"
		if len(parts) > 2 {
			// if there is an invalid flag
			if parts[2] != "--sort-name" && parts[2] != "--sort-created" {
				fmt.Fprintln(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
				return
			}
			sortBy = strings.TrimPrefix(parts[2], "--sort-")
		}
		if len(parts) > 3 {
			// if there is an invalid flag
			if parts[3] != "asc" && parts[3] != "desc" {
				fmt.Fprintln(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
				return
			}
			order = parts[3]
		}
		listFolders(parts[1], sortBy, order)

	case "rename-folder":
		if len(parts) != 4 {
			fmt.Fprintln(os.Stderr, "Usage: rename-folder [username] [foldername] [new-folder-name]")
			return
		}
		// chech whether folder contain invalid chars
		if !isValidName(parts[3]) {
			fmt.Fprintln(os.Stderr, "Error: The", parts[3], "contain invalid chars.")
			return
		}
		renameFolder(parts[1], parts[2], parts[3])

	case "create-file":
		if len(parts) != 4 && len(parts) != 5 {
			fmt.Fprintln(os.Stderr, "Usage: create-file [username] [foldername] [filename] [description]?")
			return
		}
		// chech whether file contain invalid chars
		if !isValidName(parts[3]) {
			fmt.Fprintln(os.Stderr, "Error: The", parts[3], "contain invalid chars.")
			return
		}
		// description is optional
		description := ""
		if len(parts) > 4 {
			description = strings.Join(parts[4:], " ")
		}

		// createFile(parts[1], parts[2], parts[3], description)
		// fileOps.On("createFile", parts[1], parts[2], parts[3], description).Return(nil)
		fileOps.createFile(parts[1], parts[2], parts[3], description)
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, "Error:", err)
		// }

	case "delete-file":
		if len(parts) != 4 {
			fmt.Fprintln(os.Stderr, "Usage: delete-file [username] [foldername] [filename]")
			return
		}

		// deleteFile(parts[1], parts[2], parts[3])
		err := fileOps.deleteFile(parts[1], parts[2], parts[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}

	case "list-files":
		if len(parts) < 3 || len(parts) > 5 {
			fmt.Fprintln(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
			return
		}
		sortBy := "name"
		order := "asc"
		if len(parts) > 3 {
			// if there is an invalid flag
			if parts[3] != "--sort-name" && parts[3] != "--sort-created" {
				fmt.Fprintln(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
				return
			}
			sortBy = strings.TrimPrefix(parts[3], "--sort-")
		}
		if len(parts) > 4 {
			// if there is an invalid flag
			if parts[4] != "asc" && parts[4] != "desc" {
				fmt.Fprintln(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
				return
			}
			order = parts[4]
		}
		listFiles(parts[1], parts[2], sortBy, order)

	default:
		fmt.Fprintln(os.Stderr, "Error: Unrecognized command")
	}
}

// check whether username/foldername/filename contains invalid chars
func isValidName(name string) bool {
	return validNamePattern.MatchString(name)
}
