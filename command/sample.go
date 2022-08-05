package command

type SampleCommand struct {
	SomeOption string `long:"some-option" short:"s" alias:"so" description:"lorem ipsum" value-name:"SOME" required:"false"`
}

var sampleCommand SampleCommand

var sampleCommandDescription = Description{
	CommandName:      "sample",
	ShortDescription: "Lorem ipsum",
	LongDescription:  "Lorem ipsum",
	Command:          &sampleCommand,
}

func (c *SampleCommand) Execute(args []string) error {
	// Command logic
	return nil
}
