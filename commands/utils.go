package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sg3des/eml"
)

func readEmlFile(filePath string) (eml.Message, error) {
	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return eml.Message{}, fmt.Errorf("Couldn't read the eml file %v", color.RedString(err.Error()))
	}

	email, err := eml.Parse(rawData)
	if err != nil {
		return eml.Message{}, fmt.Errorf("Couldn't parse the eml file %v", color.RedString(err.Error()))
	}
	return email, nil
}

func getHeaders(emlContent eml.Message) error {
	//fmt.Print(emlContent.Text)
	fmt.Print(emlContent.HeaderInfo.Subject)
	return nil
}

func getHash() error {
	// stuff here
	return nil
}

func getDigest() error {
	// stuff here
	return nil
}

func getContent() error {
	// stuff here
	return nil
}

func getAttachements() error {
	// stuff here
	return nil
}

func getLinks() error {
	// stuff here
	return nil
}
