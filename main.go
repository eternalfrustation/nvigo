package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"errors"
)

func main() {
	if len(os.Args) < 3 {
		log.Print("Usage: nvigo install|update <package list path>")
		log.Fatal("Not enough arguments")
	}
	confFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	shouldInstall := os.Args[1] == "install"
	listContents, err := io.ReadAll(confFile)
	listContentsStr := string(listContents)
	packages := strings.Split(listContentsStr, "\n")
	for _, module := range packages {
		if shouldInstall {
			installPackage(module)
		} else {
			updatePackage(module)
		}
	}
}

func installPackage(module string) error {
	if len(module) == 0 {
		return errors.New("Empty module name")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Print(err)
		return err
	}
	split_mod := strings.Split(module, "/")
	package_name := split_mod[len(split_mod)-1]
	cmd := exec.Command("git", "clone", "--recursive", module, fmt.Sprintf("%s/.local/share/nvim/site/pack/frustrated/start/%s", home, package_name))
	if cmd.Run() != nil {
		log.Print(err)
		return err
	}
	return nil
}

func updatePackage(module string) error {
	if len(module) == 0 {
		return errors.New("Empty module name")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Print(err)
		return err
	}
	split_mod := strings.Split(module, "\n")
	package_name := split_mod[len(split_mod)-1]
	cmd := exec.Command("git", "pull")
	cmd.Path = fmt.Sprintf("%s/.local/share/nvim/site/pack/frustated/start/%s", home, package_name)
	if cmd.Run() != nil {
		log.Print(err)
		return err
	}
	return nil
}
