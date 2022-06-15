package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

const (
	STORAGE_DIR  = "./storage"
	STORAGE_FILE = "./storage/wallet"
)

var suggestions = []prompt.Suggest{
	{"create", "create new wallet"},
	{"import", "import keys from mnemonic"},
	{"balance", "get balances by blockchain"},
	{"send", "send tokens"},
	{"help", "diplay help"},
	{"exit", "exit application"},
}

var helpSuggestions = []prompt.Suggest{
	{"help", "get help on command"},
}

func executor(in string) {
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		os.Exit(0)
		return
	case "help":
		fmt.Print(getHelp())
		return
	case "create":
		createCommand()
		return
	case "balance":
		if len(blocks) > 1 {
			switch strings.ToLower(blocks[1]) {
			case "btc", "bitcoin", "0":
				balance(0)
			case "eth", "ethereum", "1":
				balance(1)
			}
		} else {
			fmt.Println("not complete command, call `balance help`")
			return
		}
	case "import":
	default:
		fmt.Println("unknown command")
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.CurrentLineBeforeCursor()
	if w == "" {
		return suggestions
	} else if w[len(w)-1] == ' ' {
		blocks := strings.Split(w, " ")
		if len(blocks) == 2 {
			return prompt.FilterHasPrefix(helpSuggestions, "", true)
		} else {
			return prompt.FilterHasPrefix([]prompt.Suggest{}, "", true)
		}
	} else {
		blocks := strings.Split(w, " ")
		if len(blocks) == 2 {
			return prompt.FilterHasPrefix(helpSuggestions, blocks[1], true)
		}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

func main() {
	var baseURL = "unibdk"

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(baseURL+"> "),
		prompt.OptionTitle("unibdk wallet cli"),
	)
	p.Run()
}
