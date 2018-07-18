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

There are a number of functions that allow printing arbitrary text with colors such as `Printc`, `Printlnc`, and `Sprintc`. These all take a color as their first argument and printable objects after them. They will also append a color reset to the end.

```go
  vox.Printlnc(vox.Red, "Hello, I am red")
```

Printing complex color coded text can be slightly more complicated. The basic printing functions will remove their color codes for output Pipelines that are considered _Plain_, such as the FilePipeline. To print complex color coded text something like this could be used:

```go
vox.Print("I am")
vox.Printc(vox.Green, "green ")
vox.Print(" and ")
vox.Printc(vox.Red, " red", vox.ResetColor)
vox..Print("\n")
```

Although a better solution maybe to use the `PrintRich` and `PrintPlain` functions which will only output to plain or non plain Pipelines.

```go
vox.PrintRich("I am ", vox.Green, "green ", vox.ResetColor, " and ", vox.Red, "red", vox.ResetColor)
vox.PrintPlain("I am green and red)
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

```go
vox.Debug("hai")
vox.Alert("oh no")
vox.Info("this thing happened")
vox.Error("whoops")
```


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

A testing pipeline is provided that directs all output into an internal string
slice. It also provides utility functions to make accessing values easier.

```go
v := vox.New()
pipline = vox.TestPipeline{}
v.SetPipelines(pipeline)
v.Print("test")
if (pipeline.Last() != "test" {
  t.Error("not test")
}
```

A Test helper function is provided to make this easier:

```go
pipeline := vox.Test()
```

## An example test

```go
func AskForFile() string {
	return vox.Prompt("Enter a file", "")
}

func TestReadConfig(t *testing.T) {
	// setup for testing
	pipeline = vox.Test()

	// verify an error output
	err := checkFile()
	if pipeline.Last() != fmt.Sprint(vox.Red, "No config file found.") {
		t.Errorf("Error message not printed: %s", pipeline.Last())
	}

	// Asks user for a file path
	SendInput("test.txt")
	name := AskForFile()

	// Builds a config file
	SetupFile(name)


	err = checkFile()
	if err != nil {
		t.Errorf("Could not load config file: %s", err.Error())
	}
	if pipeline.Last() != "Config file read" {
		t.Error("Value missmatch: %s", pipeline.Last())
	}
}
```
