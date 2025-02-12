package cli

type CLI struct {
	Name           string
	DefaultCommand *Command
	Commands       []Command
}

func NewCLI(name string, commands Commands, defaultCommand *Command) *CLI {
	if defaultCommand == nil {
		defaultCommand = &HelpCommand
	}
	return &CLI{
		Name:           name,
		DefaultCommand: defaultCommand,
		Commands:       commands,
	}
}

func (c *CLI) HasCommand(name string) bool {
	command := c.GetCommand(name)
	return command != nil
}

func (c *CLI) GetCommand(name string) *Command {
	for _, com := range c.Commands {
		if com.Name == name {
			return &com
		}
	}
	return nil
}
