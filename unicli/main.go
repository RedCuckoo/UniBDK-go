package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/RedCuckoo/UniBDK-go/unikeys"
	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
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
	{"sign", "sign message"},
	{"help", "diplay help"},
	{"exit", "exit application"},
}

var helpSuggestions = []prompt.Suggest{
	{"help", "get help on command"},
}

func getHelp() string {
	return fmt.Sprintf("unibdk available commands:\n")
}

func passwordPrompt() ([]byte, error) {
	fmt.Printf("enter password:")
	var err error
	bytePassword := make([]byte, 0)
	exitCounter := 0
	for len(string(bytePassword)) < 8 && exitCounter < 3 {
		exitCounter++
		bytePassword, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			fmt.Printf("\ninvalid password: %s", err.Error())
			continue
		}
		if len(string(bytePassword)) < 8 {
			fmt.Printf("\ninvalid password length\ntry again:")
		}
	}

	if exitCounter == 3 {
		fmt.Println("\nthree invalid tries to set password")
		return nil, errors.New("three invalid tries to set password")
	}

	return bytePassword, nil
}

func createCommand() {
	err := os.Mkdir(STORAGE_DIR, 0777)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			fmt.Print(err)
			return
		}
	}

	var storageFile *os.File
	defer storageFile.Close()

	storageFile, err = os.Open(STORAGE_FILE)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file") {
			fmt.Print(err)
			return
		}
		// create new
		storageFile, err = os.Create(STORAGE_FILE)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("initializing new storage...")
		fmt.Print("set password (min 8 characters len):")

		bytePassword, err := passwordPrompt()
		if err != nil {
			return
		}

		fmt.Println()
		mnemonic, err := unikeys.GenerateMnemonic()
		if err != nil {
			fmt.Printf("failed to generate mnemonic: %s\n", err.Error())
			return
		}

		fmt.Printf("memorize or write down in secure place generated mnemonic:\n" +
			mnemonic)
		_, err = fmt.Scanln()
		if err != nil {
			fmt.Printf("failed to read input: %s\n", err.Error())
			return
		}

		_, err = unikeys.NewKey(mnemonic)
		if err != nil {
			fmt.Printf("failed to create key from generated mnemonic: %s\n", err.Error())
			return
		}

		encrypt, err := Encrypt(mnemonic, string(bytePassword))
		if err != nil {
			fmt.Printf("failed to encrypt wallet, needed better password: %s\n", err.Error())
			return
		}

		_, err = storageFile.Write([]byte(encrypt))
		if err != nil {
			fmt.Printf("failed to save encrypted wallet to file: %s\n", err.Error())
			return
		}

		fmt.Println("successfully created new wallet")
	} else {
		line, b, err := bufio.NewReader(storageFile).ReadLine()
		if err != nil {
			// empty file
			if strings.Contains(string(err.Error()), "EOF") {
				err := os.Remove(STORAGE_FILE)
				if err != nil {
					fmt.Printf("failed to remove empty storage file: %s\n", err.Error())
					return
				}
				createCommand()
			}
			return
		}
		if b {
			fmt.Println("removing corrupted storage")
			err := os.Remove(STORAGE_FILE)
			if err != nil {
				fmt.Printf("failed to remove empty storage file: %s\n", err.Error())
				return
			}
			createCommand()
		}

		bytePassword, err := passwordPrompt()
		if err != nil {
			return
		}

		decrypt, err := Decrypt(string(line), string(bytePassword))
		if err != nil {
			fmt.Printf("failed to decrypt wallet: %s\n", err.Error())
			return
		}

		fmt.Println(decrypt)

	}
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
	case "post", "put", "patch":
		if len(blocks) < 2 {
			fmt.Println("please set request body.")
			return
		}
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
