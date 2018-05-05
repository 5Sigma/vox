/*
Vox is a Go package designed to help make terminal/console applications more attractive.

It is a collection of small helper functions that aid in printing various pieces of information to the console.

- Various predefined and common printing tasks like printing property key/value pairs, result responses, etc.
- Print JSON data with syntax highlighting
- Easily print colorized output
- Display real time progress bars for tasks
- Easy helper functions for printing various types of messages: Alerts, errors, debug messages, etc.
- Control the output and input streams to help during application testing.

Printing to the screen

There are a number of output functions to print data to the screen with out without coloring. 
Most of the output functions accept an a series of string parts that are combined together. 
Color constants can be interlaced between these parts to color the output. 

        vox.Println(vox.Red, "Some read text", vox.ResetColor)

There are also a number of "LogLevel" type functions that easily color the output.

        vox.Alert("A warning")
        vox.Errorf("An error occured: %s", err.Error())
        vox.Debug(""A debug message")

Prompting for input

There are several helper functions for gathering input from the console.

        strResponse := vox.Prompt("a message", "a default value")
        boolResponse := vox.PromptBool("a message", true)
        choices := []string{"option 1", "option 2"}
        choiceIndex := vox.PromptChoice("A message", choices, 1)

Testing
The output and input from for vox can be redirected to memory to make it easy to test 
the input and output for CLI applications. To reroute the library simply call the 
Test function.

        vox.Test()

Now data will be read/written to in memory stores instead of STDIN/STDOUT.  
You can use `GetOutput` to get the current output in // the buffer. 
To send user input for functions like `Prompt` you can use the `SendInput` function. 
NOTE: SendInput must be called before any prompt function, 
so that the data is ready in the buffer when `Prompt` is called.

Vox also provides an AssertOutput helper for tests that checks the current 
output against the passed string. It calls testing.Error if it does not match.

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
*/
package vox