package form

import (
	"github.com/charmbracelet/huh"
	"github.com/sephory/sundae/template"
)

func RunForm(t *template.Template) {
	var fields []huh.Field
	for _, f := range t.Fields {
		if f.IsSet {
			continue
		}
		switch f.FormType {
		case template.Input:
			fields = append(fields, huh.NewInput().Title(f.Prompt).Value(&f.Value))
		case template.Text:
			fields = append(fields, huh.NewText().Title(f.Prompt).Value(&f.Value))
		case template.Select:
			fields = append(fields, huh.NewSelect[string]().Title(f.Prompt).Options(huh.NewOptions(f.Choices...)...).Value(&f.Value))
		}
	}
	if len(fields) == 0 {
		return
	}
	form := huh.NewForm(
		huh.NewGroup(
			fields...,
		),
	)
	form.Run()
}
