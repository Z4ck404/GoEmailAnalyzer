package commands

import (
	"fmt"

	"github.com/fatih/color"
)

type ParseCLI struct {
	Filename string `name:"filename" short:"f" help:"The eml file name to scan" required:""`

	Headers      bool `help:"Get the headers of the email"`
	Content      bool `help:"Get the content of the email"`
	Hash         bool `help:"Get the hash of the email"`
	Links        bool `help:"Get the links in the email file"`
	Digests      bool `help:"Get the digests of the eml file"`
	Attachements bool `help:"Get the attachments from the eml file"`
}

func (cmd *ParseCLI) Run() error {
	//email,err := readEmlFile(cmd.Filename);
	_, err := readEmlFile(cmd.Filename)
	if err != nil {
		return fmt.Errorf("Could not read the eml file %v \n", color.RedString(err.Error()))
	}
	if cmd.Headers {
		err := getHeaders()
		if err != nil {
			return fmt.Errorf("Could not get the headers %v \n", color.RedString(err.Error()))
		}
	}
	if cmd.Hash {
		err := getHash()
		if err != nil {
			return fmt.Errorf("Could not get the hash %v \n", color.RedString(err.Error()))
		}
	}
	if cmd.Links {
		err := getLinks()
		if err != nil {
			return fmt.Errorf("Could not get the Links %v \n", color.RedString(err.Error()))
		}
	}
	if cmd.Digests {
		err := getDigest()
		if err != nil {
			return fmt.Errorf("Could not get the digest %v \n", color.RedString(err.Error()))
		}
	}
	if cmd.Attachements {
		err := getAttachements()
		if err != nil {
			return fmt.Errorf("Could not get the attachements %v \n", color.RedString(err.Error()))
		}
	}
	if cmd.Content {
		err := getContent()
		if err != nil {
			return fmt.Errorf("Could not get the content %v \n", color.RedString(err.Error()))
		}
	}
	return nil
}
