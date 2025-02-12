package cli_utils

import (
	"fmt"
	"os"
	"strings"
)

func (c *CLI) getRuntimeArgs(argsWithoutProg []string) (RuntimeArgs, error) {
	args := RuntimeArgs{}
	for _, item := range argsWithoutProg[1:] {
		if item[0] != '-' || item[1] != '-' || !strings.Contains(item, "=") || item[3] == '=' || len(item) < 5 {
			return nil, fmt.Errorf("invalid argument %s", item)
		}
		key := item[2:strings.Index(item, "=")]
		value := item[strings.Index(item, "=")+1:]
		if value == "" {
			return nil, fmt.Errorf("invalid value for argument %s", key)
		}
		args[key] = value
	}
	return args, nil
}

func (c *CLI) checkRuntimeArgs(commandName string, args *RuntimeArgs) error {
	command := c.GetCommand(commandName)

	for k, _ := range *args {
		isMissing := command.Args.isMissing(k)
		if isMissing {
			return fmt.Errorf("argument %s is not supported for command %s", k, command.Name)
		}
	}

	commandArgs := command.Args
	for _, a := range commandArgs {
		_, ok := (*args)[a.Name]
		if !ok && a.Required {
			return fmt.Errorf("argument %s is required", a.Name)
		}
		if !ok {
			(*args)[a.Name] = a.Default
		}
	}

	for _, a := range commandArgs {
		if a.Check != nil {
			err := a.Check(&a, (*args)[a.Name].(string))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *CLI) Run() error {
	runtimeArgs := RuntimeArgs{}
	c.Commands = append(c.Commands, *c.DefaultCommand)
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		return c.DefaultCommand.Run(runtimeArgs, c)
	}

	command := argsWithoutProg[0]

	if !c.HasCommand(command) {
		fmt.Printf("Command %s is not supported\n", command)
		return c.DefaultCommand.Run(runtimeArgs, c)
	}

	runtimeArgs, err := c.getRuntimeArgs(argsWithoutProg)
	if err != nil {
		return err
	}

	err = c.checkRuntimeArgs(command, &runtimeArgs)
	if err != nil {
		return err
	}

	fmt.Println("Running command", command)

	for _, com := range c.Commands {
		if com.Name == command {
			err := com.Run(runtimeArgs, c)
			if com.CleanUp != nil {
				fmt.Printf("Running cleanup for command %s", command)
				cleanUpErr := com.CleanUp(runtimeArgs, c)
				if err != nil && cleanUpErr != nil {
					return fmt.Errorf("run error: %v, cleanup error: %v", err, cleanUpErr)
				}
				if cleanUpErr != nil {
					return cleanUpErr
				}
			}

			if err != nil {
				return err
			}

		}
	}

	return nil
}
