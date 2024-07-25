# Virtual File System in GoLang

This project implements a virtual file system using GoLang with REPL capabilities. The system allows users to register, create folders, delete folders, list folders, rename folders, create files, delete files, and list files.

## Setup

1. Install GoLang 1.20+.
2. Clone this repository.
3. Navigate to the project directory.
4. Run `go mod tidy` to install dependencies.

## Usage

Run the following command to start the REPL:

```sh
go run main.go user.go folder.go file.go
```

## Command
You can use the following commands:

- register [username]
- create-folder [username] [foldername] [description]?
- delete-folder [username] [foldername]
- list-folders [username] [--sort-name|--sort-created] [asc|desc]
- rename-folder [username] [foldername] [new-folder-name]
- create-file [username] [foldername] [filename] [description]?
- delete-file [username] [foldername] [filename]
- list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]

## Input Validation and Restrictions
for Usernames, Folder Names, and File Names
- Maximum Length: 30 characters
- Valid Characters: Alphanumeric characters (a-z, A-Z, 0-9), hyphens (-), and underscores (_)
- No Whitespace Allowed: Usernames, folder names, and file names must not contain whitespace characters
- Case-Insensitive: All inputs are converted to lowercase for uniformity


## Testing Strategy

### Unit Testing

This project includes a comprehensive set of unit tests to ensure code quality and correctness. I aimed for high code coverage and included tests for both positive and negative scenarios.

### Input Validation

I have implemented input validation to ensure that usernames, folder names, and file names adhere to specified rules. These rules are enforced and tested to prevent invalid inputs from causing issues.

### Advanced Testing Techniques

In addition to basic unit testing, I employed advanced testing techniques such as:
- **Mocking**: To isolate units of code and test them independently.
- **Error Handling Tests**: To ensure the program recovers gracefully from unexpected inputs and states.
- **Table-Driven Tests**: To efficiently cover multiple test scenarios in a compact form.
- **Edge Case Testing**: To ensures the program can handle unusual but possible inputs.

### Code Coverage

I aimed for 75% code coverage to ensure that most of our code is tested. I used the `go test -cover` tool to measure and report code coverage.

### Running Tests

To run the tests, use the following command:
```sh
go test ./...
```
Or to run specific test, use the following command:
`go test -v -run [testfunction]`  
Example: 
```sh
go test -v -run TestRegisterUser
```
