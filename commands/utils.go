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
	key   string
	value string
}

type contentObject struct {
	name       string
	attributes []attribute
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

func getHeaders(emlContent eml.Message) error {

	// show all the headers
	// v := reflect.ValueOf(emlContent.HeaderInfo)
	// typeOfS := v.Type()

	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Printf("%s :\t %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	// }
	headerInfo := emlContent.HeaderInfo
	sentTo, _ := json.Marshal(headerInfo.To)
	headerContent := contentObject{
		name: "header",
		attributes: []attribute{
			{
				key:   "date",
				value: headerInfo.Date.String(),
			},
			{
				key:   "sender",
				value: headerInfo.Sender.String(),
			},
			{
				key:   "subject",
				value: headerInfo.Subject,
			},
			{
				key:   "Sent to",
				value: string(sentTo),
			},
			{
				key:   "subject",
				value: headerInfo.Subject,
			},
			{
				key:   "subject",
				value: headerInfo.Subject,
			},
		},
	}
	_prettyTable(headerContent)
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

func _prettyTable(c contentObject) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{fmt.Sprintf("%s Field", c.name), fmt.Sprintf("%v Value", c.name)})
	for _, attribute := range c.attributes {
		t.AppendRow([]interface{}{attribute.key, attribute.value})
		t.AppendRow([]interface{}{" ", " "})
	}
	t.Render()
}
