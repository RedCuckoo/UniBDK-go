package main

import (
	"fmt"
	"github.com/josephspurrier/goconsole"
	"os"
	"strings"
)

const (
	STORAGE_DIR  = "./storage"
	STORAGE_FILE = "./storage/wallet"
)

func CreateCommand(com string) {
	err := os.Mkdir(STORAGE_DIR, 0777)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			fmt.Print(err)
			return
		}
	}

	var storageFile *os.File

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

	} else {
		// already exists
	}

	defer storageFile.Close()

}

func main() {
	con := goconsole.New()
	//con.Title = "UniBDK console\n"
	con.Prompt = "unicli> "
	con.NotFound = "Command not found: "
	//con.NewLine = ""
	//con.NewLine
	con.Add("create", "generate new wallet", CreateCommand)
	con.Start()
}
