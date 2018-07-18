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


Pipelines

Vox offers pipelines as a way of configuring one or more output streams. Four
built in pipelines are provided with the package:

- ConsolePipeline - This is the default Pipeline set for any vox instance. This
pipeline will present colored output to standard out.

- FilePipeline - This pipeline will redirect all data to a local file. This
pipeline uses plain output, without color codes.

- TestPipeline - All output will be internally stored in a string slice and
utility functions are provided to make accessing values easier. This pipeline
should be used for unit tests.

- WriterPipeline - This is a generic pipeline that allows you to specifiy any
writer that implements the io.Writer interface.


Testing

A testing pipeline is provided that directs all output into an internal string
slice. It also provides utility functions to make accessing values easier.

				v := vox.New()
				pipline = vox.TestPipeline{}
				v.SetPipelines(pipeline)
				v.Print("test")
				if (pipeline.Last() != "test" {
					t.Error("not test")
				}

A Test helper function is provided to make this easier:

				pipeline := vox.Test()

You can use the `SendInput` function.  SendInput must be called before
any prompt function, so that the data is ready in the buffer when `Prompt`
is called.

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
*/
package vox
