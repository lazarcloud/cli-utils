package cli

type Command struct {
	Name        string
	Description string
	Run         func(args RuntimeArgs, c *CLI) error
	CleanUp     func(args RuntimeArgs, c *CLI) error
	Args        Args
}

type Commands []Command
