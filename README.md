[![GoDoc](https://godoc.org/github.com/5Sigma/vox?status.svg)](https://godoc.org/github.com/5Sigma/vox) 
[![Go Report Card](https://goreportcard.com/badge/github.com/5sigma/vox)](https://goreportcard.com/report/github.com/5sigma/vox) 
[![Build Status](https://travis-ci.org/5Sigma/vox.svg?branch=master)](https://travis-ci.org/5Sigma/vox)

Vox is a Go package designed to help make terminal/console applications more
attractive.

It is a collection of small helper functions that aid in printing various
pieces of information to the console.

- Various predefined and common printing tasks like printing property key/value
    pairs, result responses, etc.
- Print JSON data with syntax highlighting
- Easily print colorized output
- Display real time progress bars for tasks
- Easy helper functions for printing various types of messages: Alerts, errors,
    debug messages, etc.
 - Control the output and input streams to help during application testing.


# Usage

```go
import "github.com/5sigma/vox"
```

## Printing with color

Vox provides a numbr of functions and patterns for colored output.

```go
vox.Printlnc(vox.Red, "Hello, I am red")
vox.PrintLn("I am ", vox.Green, "green ", vox.ResetColor, "and ", 
  vox.Red, " red", vox.ResetColor)
```

### Colors

The following color constants exist, along with a *ResetColor* constant:

- Black
- Red
- Green
- Yellow
- Blue
- Magenta
- Cyan
- White

## Loglevel functions

The loglevel functions print standard types of messages
(debug, error, alert, info) colored and formatted. See the
package documentation for their definitions.


## Printing results
Prints a key and a result string depending on if the error value is nil.

```go
err := writeFile()
vox.PrintResult("Writing file", err)
```
Task:                                   [OK]
Task2:                                  [FAIL]
 - Error messsage
Task3:                                  [OK]


## Printing property lists
Prints a key and value and pads them to align on the edges of the screen.

```go
vox.PrintResult("Name", user.Name)
vox.PrintResult("Email", user.Email)
```


```
Key:                                   some value
Key:                                   some value
Key:                                   some value
```


## Prompting for input

Prompting for a basic string response from the user:

```go
result := vox.Prompt("Enter a response", "none")
vox.Println("You entered ", result)
```

Prompting for a boolean response:

```go
result := vox.PromptBool("Are you sure", false)
if !result {
  return
}
```

Prompting for a choice of options:

```go
choices := []string{"Option 1", "Option 2", "Option 3"}
resultIndex := vox.PromptChoice("Choose an option", choices, 0)
```

## Displaying progress

The progress bar is controlled using `StartProgress`, `IncProgress`, and
`SetProgress`.

```go
vox.StartProgress(0, len(myTasks))
for _, t := range myTasks {
  err := doTask(myTask)
  if err != nil {
    vox.Error(err.Error())
    break
  }
  vox.IncProgress()
}
vox.StopProgress()
```


# Testing

The output and input from for vox can be redirected to memory to make it easy to test the input and output for CLI applications. To reroute the library simply call the `Test` function.

```go
vox.Test()
```

Now data will be read/written to in memory stores instead of STDIN/STDOUT.  You can use `GetOutput` to get the current output in the buffer. To send user input for functions like `Prompt` you can use the `SendInput` function. NOTE: SendInput must be called before any prompt function, so that the data is ready in the buffer when `Prompt` is called.

Vox also provides an `AssertOutput` helper for tests that checks the current output against the passed string. It calls testing.Error if it does not match.


## An example test

```go
func AskForFile() string {
  return vox.Prompt("Enter a file", "") 
}


func TestReadConfig(t *testing.T) {
  vox.Test()
  
  err := checkFile()
  vox.AssertOutput(t, vox.Red, "No config file found.")
  SetupFile("test.txt") // Builds a config file
  SendInput("test.txt")
  AskForFile() // Asks user for a file path  
  err = checkFile()
  if err != nil {
    t.Errorf("Could not load config file: %s", err.Error())
  }
  AssertOutput(t, "Config file read")
}
```
