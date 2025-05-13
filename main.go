package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCmd struct {
	name     string
	desc     string
	callBack func() error
}

var cmdMap = map[string]cliCmd{}

func cmdExit() error {
	os.Exit(0)
	return fmt.Errorf("Closing the Pokedex... Goodbye!")
}

func cmdHelp() error {
	var cmds string
	for _, cmd := range cmdMap {
		cmds += fmt.Sprintf("%s: %s\n", cmd.name, cmd.desc)
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Print(cmds)
	return nil
}

func init() {
	cmdMap["exit"] = cliCmd{
		name:     "exit",
		desc:     "Exit the Pokedex",
		callBack: cmdExit,
	}

	cmdMap["help"] = cliCmd{
		name:     "help",
		desc:     "Displays a help message",
		callBack: cmdHelp,
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}

func regCmd(cleanText []string) (cliCmd, error) {
	for _, word := range cleanText {
		cmd, exists := cmdMap[word]
		if exists {
			return cmd, nil
		}
	}
	return cliCmd{}, fmt.Errorf("Unknown command")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		cmd, err := regCmd(cleanText)
		if err == nil {
			cmd.callBack()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
