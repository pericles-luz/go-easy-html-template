package easy_html_template

import (
	"bytes"
	"os"
	"runtime/debug"
	"strings"
	"text/template"

	"github.com/pericles-luz/go-converter/pkg/converter"
)

type EasyHTMLTemplate struct {
	ID     uint16
	Name   string
	Text   string
	Type   uint8
	Order  uint8
	data   map[string]string
	assets map[string]string
}

func NewTemplate() (*EasyHTMLTemplate, error) {
	script := &EasyHTMLTemplate{}
	return script, nil
}

func (c *EasyHTMLTemplate) GetId() uint16 {
	return c.ID
}

func (c *EasyHTMLTemplate) GetName() string {
	return c.Name
}

func (c *EasyHTMLTemplate) GetText() string {
	return c.Text
}

func (c *EasyHTMLTemplate) GetType() uint8 {
	return c.Type
}

func (c *EasyHTMLTemplate) GetOrder() uint8 {
	return c.Order
}

func (c *EasyHTMLTemplate) GetData() map[string]string {
	return c.data
}

func (c *EasyHTMLTemplate) GetAssets() map[string]string {
	return c.assets
}

func (c *EasyHTMLTemplate) SetId(id uint16) error {
	c.ID = id
	return nil
}

func (c *EasyHTMLTemplate) SetName(name string) error {
	c.Name = name
	return nil
}

func (c *EasyHTMLTemplate) SetText(text string) error {
	c.Text = text
	return nil
}

func (c *EasyHTMLTemplate) SetType(t uint8) error {
	c.Type = t
	return nil
}

func (c *EasyHTMLTemplate) SetOrder(order uint8) error {
	c.Order = order
	return nil
}

func (c *EasyHTMLTemplate) SetData(data map[string]string) error {
	c.data = data
	return nil
}

func (c *EasyHTMLTemplate) SetAssets(assets map[string]string) {
	c.assets = assets
}

func (c *EasyHTMLTemplate) GetTranslated() (string, error) {
	var variables string
	onlyMap := make(map[string]string)
	for key, value := range c.GetData() {
		if !strings.Contains(c.Text, "{{$"+key+"}}") {
			onlyMap[key] = value
			continue
		}
		variables += "{{$" + key + ":=" + `"` + value + `"` + "}}"
	}
	for key, value := range c.GetAssets() {
		data := converter.AssetToBase64(value)
		if data == "" {
			continue
		}
		if !strings.Contains(c.Text, "{{$"+key+"}}") {
			onlyMap[key] = value
			continue
		}
		variables += "{{$" + key + ":=" + `"` + value + `"` + "}}"
	}

	text, err := template.New("script").Parse(variables + c.Text)
	if err != nil {
		return "", err
	}
	result := &bytes.Buffer{}
	err = text.Execute(result, onlyMap)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func LoadDynamicTemplate(templateName string, data map[string]string) (string, error) {
	text, err := LoadTemplate(templateName)
	if err != nil {
		return "", err
	}
	template, err := NewTemplate()
	if err != nil {
		debug.PrintStack()
		return "", err
	}
	template.SetText(string(text))
	template.SetData(data)
	return template.GetTranslated()
}

func LoadDynamicTemplateWithAssets(templateName string, assets, data map[string]string) (string, error) {
	text, err := LoadTemplate(templateName)
	if err != nil {
		return "", err
	}
	template, err := NewTemplate()
	if err != nil {
		debug.PrintStack()
		return "", err
	}
	template.SetText(string(text))
	template.SetData(data)
	template.SetAssets(assets)
	return template.GetTranslated()
}

func GetTemplate(templateName string) (string, error) {
	result, err := LoadTemplate(templateName)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func LoadTemplate(filePath string) ([]byte, error) {
	if _, err := os.Stat(filePath); nil != err {
		return nil, err
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
