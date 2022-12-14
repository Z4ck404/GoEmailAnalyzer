package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/sg3des/eml"
)

type attribute struct {
	Key   string
	Value string
}

type contentObject struct {
	Name       string
	Attributes []eml.Header
}

func readEmlFile(filePath string) (eml.Message, error) {
	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return eml.Message{}, fmt.Errorf("%v", color.RedString(err.Error()))
	}

	email, err := eml.Parse(rawData)
	if err != nil {
		return eml.Message{}, fmt.Errorf("%v", color.RedString(err.Error()))
	}
	return email, nil
}

func getHeaders(emlContent eml.Message, getMoreDetails bool) error {

	//show all the headers
	headerInfo := emlContent.HeaderInfo
	fullHeaderAttributes := contentObject{
		Name:       "header",
		Attributes: headerInfo.FullHeaders,
	}
	sentTo, _ := json.Marshal(headerInfo.To)
	shortHeaderAttributes := contentObject{
		Name: "header",
		Attributes: []eml.Header{
			{
				Key:   "date",
				Value: headerInfo.Date.String(),
			},
			{
				Key:   "sender",
				Value: headerInfo.Sender.String(),
			},
			{
				Key:   "subject",
				Value: headerInfo.Subject,
			},
			{
				Key:   "Sent to",
				Value: string(sentTo),
			},
			{
				Key:   "subject",
				Value: headerInfo.Subject,
			},
			{
				Key:   "subject",
				Value: headerInfo.Subject,
			},
		},
	}

	if getMoreDetails {
		_prettyTable(fullHeaderAttributes)
	} else {
		_prettyTable(shortHeaderAttributes)
	}

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

func getContent(emlContent eml.Message) error {
	// contentAttributes := contentObject{
	// 	Name: "content",
	// 	Attributes: []eml.Header{
	// 		{
	// 			Key:   "Message Body",
	// 			Value: emlContent.Text,
	// 		},
	// 	},
	// }
	//_prettyTable(contentAttributes)
	fmt.Print(emlContent.Text)
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

func _prettyTable(c contentObject) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{fmt.Sprintf("%s Field", c.Name), fmt.Sprintf("%v Value", c.Name)})
	for _, attribute := range c.Attributes {
		t.AppendRow([]interface{}{attribute.Key, attribute.Value})
		t.AppendRow([]interface{}{" ", " "})
	}
	t.Render()
}
