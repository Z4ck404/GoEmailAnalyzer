package commands

import (
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

func getHeaders(emlContent eml.Message) error {

	//show all the headers

	//ignore_fields := []string{"FullHeaders", "OptHeaders"}
	headerContent := contentObject{
		Name:       "header",
		Attributes: emlContent.HeaderInfo.FullHeaders,
	}

	// v := reflect.ValueOf(emlContent.HeaderInfo)
	// typeOfS := v.Type()

	// for i := 0; i < v.NumField(); i++ {
	// 	headerFiledName := typeOfS.Field(i).Name
	// 	headerFiledValue := v.Field(i).Interface()

	// 	for _, element := range ignore_fields {
	// 		if headerFiledName != element {
	// 			fmt.Print(headerFiledName)
	// 			headerField := attribute{
	// 				key:   headerFiledName,
	// 				value: fmt.Sprintf("%v", headerFiledValue),
	// 			}
	// 			headerContent.attributes = append(headerContent.attributes, headerField)
	// 		}
	// 	}

	// }

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
	t.AppendHeader(table.Row{fmt.Sprintf("%s Field", c.Name), fmt.Sprintf("%v Value", c.Name)})
	for _, attribute := range c.Attributes {
		t.AppendRow([]interface{}{attribute.Key, attribute.Value})
		t.AppendRow([]interface{}{" ", " "})
	}
	t.Render()
}
