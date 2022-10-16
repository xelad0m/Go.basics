package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	{
		fmt.Println("---")

		const txt = "Алиса: - {{.}}\n"

		tpl := template.New("value")
		tpl = template.Must(tpl.Parse(txt))

		tpl.Execute(os.Stdout, "Привет!")
		tpl.Execute(os.Stdout, "Как дела?")
		tpl.Execute(os.Stdout, "Пока!")
	}

	{
		fmt.Println("---")

		const txt = `Сейчас {{.Time}}, {{.Day}}.
{{if .Sunny -}} Солнечно! {{- else -}} Пасмурно :-/ {{- end}}
`

		tpl := template.New("greeting")
		tpl = template.Must(tpl.Parse(txt))

		type State struct {
			Time  string
			Day   string
			Sunny bool
		}

		state := State{"9:00", "четверг", true}
		tpl.Execute(os.Stdout, state)

		fmt.Println()

		state = State{"21:00", "пятница", false}
		tpl.Execute(os.Stdout, state)
	}

	{
		fmt.Println("---")

		const txt = "{{range .}}- {{ . }}\n{{end}}"

		tpl := template.New("list")
		tpl = template.Must(tpl.Parse(txt))

		list := []string{"Купить молоко", "Погладить кота", "Вынести мусор"}
		tpl.Execute(os.Stdout, list)
	}

	r := 'A'
	r += 32
	fmt.Println(string(r))
}
