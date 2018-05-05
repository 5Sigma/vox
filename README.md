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

Vox provides a `Printlnc` function for printing a line of text in a given color

```go
vox.Printlnc(vox.Red, "Hello")
```

You can also the color constants directly in any other function:

```go
fmt.Printf("%sHello%s\n", vox.Red, vox.ColorReset)
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
```

## Other helper functions

### PrintProperty(key, value)
Prints a key and value and pads them to align on the edges of the screen.

```
Key:                                   some value
```

### PrintResult(name, error)
Prints a key and a result string depending on if the error value is nil.

```
Task:                                   [OK]
Task2:                                  [FAIL]
 - Error messsage
Task3:                                  [OK]
```


## Loglevel functions

The loglevel functions print standard types of messages
(debug, error, alert, info) colored and formatted. See the
package documentation for their definitions.
