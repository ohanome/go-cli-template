package command

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type DefaultOptions struct {
	// Place your default options here.
}

type Description struct {
	CommandName      string
	ShortDescription string
	LongDescription  string
	Command          any
}

var defaultOptions DefaultOptions
var parser *flags.Parser

func init() {
	parser = flags.NewParser(&defaultOptions, flags.Default)

	AddCommandFromDescriptions([]Description{
		// Place your command descriptions here.
		// E.g. you have a command called "my-command" that does something in ./my-command.go
		// there you declared an object of type Description 'myCommandDescription' and set its fields.
		// Here you will just add 'myCommandDescription' to the list of descriptions.
		sampleCommandDescription,
	})

	_, err := parser.Parse()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}

func AddCommandFromDescription(description Description) {
	_, err := parser.AddCommand(description.CommandName, description.ShortDescription, description.LongDescription, description.Command)
	if err != nil {
		panic(err)
	}
}

func AddCommandFromDescriptions(descriptions []Description) {
	for _, description := range descriptions {
		AddCommandFromDescription(description)
	}
}
