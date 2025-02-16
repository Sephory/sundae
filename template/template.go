package template

import (
	"regexp"
	"strings"
	"text/template"
)

const template_parser = "{{([0-9A-Za-z]+)(\\|[^}]+)?}}"

type Template struct {
	Fields []*Field
	Text   string
}

func (t *Template) GetValues() map[string]string {
	values := make(map[string]string, len(t.Fields))
	for _, f := range t.Fields {
		values[f.Name] = f.Value
	}
	return values
}

func (t *Template) Execute() (string, error) {
	var builder strings.Builder
	textTemplate, err := template.New("Sundae Template").Parse(t.Text)
	if err != nil {
		return "", err
	}
	err = textTemplate.Execute(&builder, t.GetValues())
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func New(s string) *Template {
	parser := regexp.MustCompile(template_parser)
	matches := parser.FindAllStringSubmatch(s, -1)
	fields := makeFieldsFromMatches(matches)
	text := parser.ReplaceAllString(s, "{{.$1}}")
	return &Template{
		Fields: fields,
		Text:   text,
	}
}
