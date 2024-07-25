package main

import (
	"testing"
)

// Test register user command
func TestRegisterUser(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("testuser")
    if _, exists := users["testuser"]; !exists {
        t.Errorf("Expected user 'testuser' to be registered.")
    }
}

// Test registering exist username
func TestRegisterUser_Duplicate(t *testing.T) {
	users = make(map[string]*User) // Reset state for the test
	registerUser("testuser")
	registerUser("testuser") // Duplicate registration
	if _, exists := users["testuser"]; !exists {
		t.Errorf("Expected user 'testuser' to be registered.")
	}
}

// Test create folder command
func TestCreateFolder(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("testuser")
    createFolder("testuser", "testfolder", "description")
    user := users["testuser"]
    if _, exists := user.Folders["testfolder"]; !exists {
        t.Errorf("Expected folder 'testfolder' to be created.")
    }

	// test wrong user name
	createFolder("testuser1", "testfolder1", "description")
    if _, exists := user.Folders["testfolder1"]; exists {
        t.Errorf("Cannot create folder because there is no user 'testuser1'")
    }
}

// Test delete folder command
func TestDeleteFolder(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("testuser")
    createFolder("testuser", "testfolder", "description")
	user := users["testuser"]

	// test wrong user name
	deleteFolder("testuser1", "testfolder")
    if _, exists := user.Folders["testfolder"]; !exists {
        t.Errorf("cannot delete the folder because there is no user 'testuser1'")
    }

	// test wrong folder name
	deleteFolder("testuser", "testfolder1")
    if _, exists := user.Folders["testfolder"]; !exists {
        t.Errorf("cannot delete the folder because there is no folder 'testfolder1'")
    }
	
	// test delete
    deleteFolder("testuser", "testfolder")
    if _, exists := user.Folders["testfolder"]; exists {
        t.Errorf("Expected folder 'testfolder' to be deleted.")
    }
}

// Test Rename folder command
func TestRenameFolder(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("testuser")
    createFolder("testuser", "testfolder", "description")
	
    user := users["testuser"]

	// test wrong user name
	renameFolder("testuser1", "testfolder", "newtestfolder")
    if _, exists := user.Folders["testfolder"]; !exists {
        t.Errorf("cannot rename the folder because there is no user 'testuser1'")
    }

	// test wrong folder name
	renameFolder("testuser", "testfolder1", "newtestfolder")
    if _, exists := user.Folders["testfolder"]; !exists {
        t.Errorf("cannot delete the folder because there is no folder 'testfolder1'")
    }

	// rename folder 
	renameFolder("testuser", "testfolder", "newtestfolder")
    if _, exists := user.Folders["testfolder"]; exists {
        t.Errorf("Expected folder 'testfolder' to be deleted.")
    }
}

// Test create file command
func TestCreateFile(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("testuser")
    createFolder("testuser", "testfolder", "description")
    createFile("testuser", "testfolder", "testfile", "description")
    user := users["testuser"]
    folder := user.Folders["testfolder"]
    if _, exists := folder.Files["testfile"]; !exists {
        t.Errorf("Expected file 'testfile' to be created.")
    }

	// test wrong user name
	createFile("testuser1", "testfolder", "testfile", "description")
    if _, exists := folder.Files["testfile1"]; exists {
        t.Errorf("cannot create the file because there is no user 'testuser1'")
    }

	// test wrong folder name
	createFile("testuser", "testfolder1", "testfile1", "description")
    if _, exists := folder.Files["testfile1"]; exists {
        t.Errorf("cannot create the file because there is no folder 'testfolder1'")
    }
}

// Test delete file command
func TestDeleteFile(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("testuser")
    createFolder("testuser", "testfolder", "description")
    createFile("testuser", "testfolder", "testfile", "description")
	user := users["testuser"]
    folder := user.Folders["testfolder"]

	// test wrong user name
	deleteFile("testuser1", "testfolder", "testfile")
    if _, exists := folder.Files["testfile"]; !exists {
        t.Errorf("cannot delete the file because there is no user 'testuser1'")
    }

	// test wrong folder name
	deleteFile("testuser", "testfolder1", "testfile")
    if _, exists := folder.Files["testfile"]; !exists {
        t.Errorf("cannot delete the file because there is no folder 'testfolder1'")
    }

	// test wrong file name
	deleteFile("testuser", "testfolder", "testfile1")
    if _, exists := folder.Files["testfile"]; !exists {
        t.Errorf("cannot delete the file because there is no file 'testfile1'")
    }

	// delete the file
    deleteFile("testuser", "testfolder", "testfile")
    if _, exists := folder.Files["testfile"]; exists {
        t.Errorf("Expected file 'testfile' to be deleted.")
    }
}

// Test case insensitive
func TestCaseInsensitive(t *testing.T) {
    users = make(map[string]*User) // Reset state for the test
    registerUser("TESTUSER")
	if _, exists := users["testuser"]; !exists {
        t.Errorf("User 'TESTUSER' and 'testuser' should be the same user.")
    }
    createFolder("testuser", "TESTFOLDER", "description")
	user := users["testuser"]
	if _, exists := user.Folders["testfolder"]; !exists {
        t.Errorf("Folder 'TESTFOLDER' and 'testfolder' should be the same folder.")
    }
    createFile("testuser", "testfolder", "TESTFILE", "description")
    folder := user.Folders["testfolder"]
    if _, exists := folder.Files["testfile"]; !exists {
        t.Errorf("File 'TESTFILE' and 'testfile' should be the same file.")
    }
}

// Test the user command(register, length and invalid char)
func TestUserCommand(t *testing.T) {
	users = make(map[string]*User) // Reset state for the test

	// test unsupport command
	handleCommand("create user1")
	if _, exists := users["user1"]; exists {
        t.Errorf("Unrecognized command")
    }

	// test register
	handleCommand("register user1")
	if _, exists := users["user1"]; !exists {
        t.Errorf("Expected user 'user1' to be registered.")
    }

	// test length of register
	handleCommand("register user2 user2")
	if _, exists := users["user2"]; exists {
        t.Errorf("Usage: register [username]")
    }

	// test whether user contain invalid char 
	handleCommand("register user2@ ")
	if _, exists := users["user2@"]; exists {
        t.Errorf("user@ should not contain any invalid chars.")
    }
}

// Test the folder command(create, length and invalid char)
func TestFolderCommand(t *testing.T) {
	users = make(map[string]*User) // Reset state for the test
	handleCommand("register user1")

	// test unsupport command
	handleCommand("create user1 folder1")
	if _, exists := users["user1"].Folders["folder1"]; exists {
        t.Errorf("Unrecognized command")
    }

	// test create new folder 
	handleCommand("create-folder user1 folder1")
    if _, exists := users["user1"].Folders["folder1"]; !exists {
        t.Errorf("Expected folder 'folder1' to be created.")
    }

	// test length of create
	handleCommand("create-folder user1 folder1 folder2 description")
	if _, exists := users["user1"].Folders["folder2"]; exists {
        t.Errorf("Usage: create-folder [username] [foldername] [description]?")
    }

	// test the invalid folder name
	handleCommand("create-folder user1 folder.1")
	if _, exists := users["user1"].Folders["folder.1"]; exists {
        t.Errorf("folder.1 should not contain any invalid chars.")
    }

	// test length of rename 
	handleCommand("rename-folder user1 folder1 folder2 folder1")
	if _, exists := users["user1"].Folders["folder1"]; !exists {
        t.Errorf("Usage: rename-folder [username] [foldername] [new-folder-name]")
    }

	// test rename 
	handleCommand("rename-folder user1 folder1 folder2")
	if _, exists := users["user1"].Folders["folder1"]; exists {
        t.Errorf("Expected folder 'folder1' to be renamed.")
    }
	if _, exists := users["user1"].Folders["folder2"]; !exists {
        t.Errorf("Expected the renamed folder 'folder2' to exist.")
    }

	// test length of delete
	handleCommand("delete-folder user1 folder1 folder2")
	if _, exists := users["user1"].Folders["folder2"]; !exists {
        t.Errorf("Usage: delete-folder [username] [foldername]")
    }

	// test delete
	handleCommand("delete-folder user1 folder2")
	if _, exists := users["user1"].Folders["folder2"]; exists {
        t.Errorf("Expected folder 'folder2' to be deleted.")
    }

	handleCommand("create-folder user1 folder1")
	handleCommand("create-folder user1 folder2")
	// test length of list 
	handleCommand("list-folders user1 --sort-name asc desc")
	if _, exists := users["user1"].Folders["folder1"]; !exists {
        t.Errorf("list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
    }

	// test list
	handleCommand("list-folders user1 --sort-name asc")
	if _, exists := users["user1"].Folders["folder1"]; !exists {
        t.Errorf("List folders should not change anything")
    }
}

// Test the folder command(create, length and invalid char)
func TestFileCommand(t *testing.T) {
	users = make(map[string]*User) // Reset state for the test
	handleCommand("register user1")
	handleCommand("create-folder user1 folder1")

	// test unsupport command
	handleCommand("create user1 folder1 file1")
	if _, exists := users["user1"].Folders["folder1"].Files["file1"]; exists {
        t.Errorf("Unrecognized command")
    }

	// test create new file 
	handleCommand("create-file user1 folder1 file1 this-is-file1")
    if _, exists := users["user1"].Folders["folder1"].Files["file1"]; !exists {
        t.Errorf("Expected folder 'folder1' to be created.")
    }

	// test length of create
	handleCommand("create-file user1 folder1 file1 file2 this-is-file1")
	if _, exists := users["user1"].Folders["folder1"].Files["file2"]; exists {
        t.Errorf("Usage: delete-folder [username] [foldername]")
    }

	// test the invalid file name
	handleCommand("create-file user1 folder1 file??")
	if _, exists := users["user1"].Folders["folder1"].Files["file??"]; exists {
        t.Errorf("file?? should not contain any invalid chars.")
    }
      
	// test length of delete command
	handleCommand("delete-folder user1 folder1 file1 file2")
	if _, exists := users["user1"].Folders["folder1"].Files["file1"]; !exists {
        t.Errorf("Usage: delete-file [username] [foldername] [filename]")
    }

	// test delete
	handleCommand("delete-file user1 folder1 file1")
	if _, exists := users["user1"].Folders["folder1"].Files["file1"]; exists {
        t.Errorf("Expected file 'file1' to be deleted.")
    }

	handleCommand("create-file user1 folder1 file1")
	handleCommand("create-file user1 folder1 file2")

	// test length of list command
	handleCommand("list-files user1 folder1 --sort-name asc desc")
	if _, exists := users["user1"].Folders["folder1"].Files["file1"]; !exists {
        t.Errorf("Usage: list files [username] [foldername] [--sort-name|--sort-created][asc|desc]")
    }

	// test list
	handleCommand("list-files user1 folder1 --sort-name asc")
	if _, exists := users["user1"].Folders["folder1"].Files["file1"]; !exists {
        t.Errorf("List files should not change anything")
    }
}
