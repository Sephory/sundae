package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sephory/sundae/form"
	"github.com/sephory/sundae/template"
)

func main() {
	var input string
	args := os.Args[1:]
	if len(args) > 0 && !template.IsFieldArg(args[0]) {
		input = args[0]
		args = args[1:]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input += scanner.Text() + "\n"
		}
	}
	t := template.New(input)
	for _, a := range args {
		if template.IsFieldArg(a) {
			fieldValue := strings.SplitN(a, "=", 2)
			for _, f := range t.Fields {
				if f.Name == fieldValue[0] {
					f.Value = fieldValue[1]
					f.IsSet = true
					break
				}
			}
		}
	}
	form.RunForm(t)
	final, err := t.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Print(final)
}
