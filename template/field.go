package template

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const field_arg_pattern = "[0-9A-Za-z]=.*"

type FormType string

const (
	Input  = "input"
	Text   = "text"
	Select = "select"
)

type Field struct {
	Name     string
	FormType FormType
	Choices  []string
	Value    string
	Prompt   string
	Order    int
	IsSet    bool
}

func makeFieldsFromMatches(matches [][]string) []*Field {
	fields := make(map[string]*Field)
	for _, m := range matches {
		name := m[1]
		if _, ok := fields[name]; ok {
			continue
		}
		f := Field{
			Name:  name,
			IsSet: false,
		}
		options := strings.Split(m[2], "|")
		for i, o := range options[1:] {
			switch i {
			case 0:
				if f.FormType == "" {
					switch strings.ToLower(o) {
					case Input:
						f.FormType = Input
					case Text:
						f.FormType = Text
					case Select:
						f.FormType = Select
					default:
						f.FormType = Input
					}
				}
			case 1:
				if f.FormType == Select {
					f.Choices = strings.Split(o, ";")
				}
			case 2:
				if f.Value == "" {
					f.Value = o
				}
			case 3:
				if f.Prompt == "" {
					f.Prompt = o
				}
			case 4:
				if f.Order == 0 {
					order, err := strconv.Atoi(o)
					if err != nil {
						panic(err)
					}
					f.Order = order
				}
			}
		}
		if f.Prompt == "" {
			f.Prompt = f.Name + ":"
		}
		fields[name] = &f

	}
	var fieldSlice []*Field
	for _, f := range fields {
		fieldSlice = append(fieldSlice, f)
	}
	sort.Slice(fieldSlice, func(i, j int) bool {
		return fieldSlice[i].Order < fieldSlice[j].Order
	})
	return fieldSlice
}

func IsFieldArg(arg string) bool {
	matched, err := regexp.MatchString(field_arg_pattern, arg)
	if err != nil {
		return false
	}
	return matched
}
