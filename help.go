package cli

import "fmt"

var HelpCommand = Command{
	Name:        "help",
	Description: "Helps you with a command",
	Run:         RunHelp,
	Args: []Arg{
		{
			Name:        "command",
			Description: "Command to get help for",
			Required:    false,
			Default:     nil,
		},
	},
}

func RunHelp(args RuntimeArgs, c *CLI) error {
	if args["command"] != nil {
		command := args["command"].(string)
		for _, c := range c.Commands {
			if c.Name == command {
				fmt.Printf("Help for command %s:\n", command)
				fmt.Printf("Description: %s\n", c.Description)
				if c.Args != nil {
					fmt.Println("Arguments:")
					for _, a := range c.Args {
						fmt.Printf("	%s - %s\n", a.Name, a.Description)
					}
				}
				return nil
			}
		}
		return fmt.Errorf("Command %s not found", command)
	}
	fmt.Printf("Supported commands for %s are:\n", c.Name)
	for _, c := range c.Commands {
		fmt.Printf("%s - %s\n", c.Name, c.Description)
		if c.Args != nil {
			fmt.Println("	Arguments:")
			for _, a := range c.Args {
				fmt.Printf("	%s - %s\n", a.Name, a.Description)
			}
		}
		fmt.Println("------------------")
	}

	return nil
}
