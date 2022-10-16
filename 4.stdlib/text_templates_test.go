package main

import (
	"bytes"
	"testing"
	"text/template"
)

// начало решения

//var templateText = `{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. {{if ge .Balance 100}}Все в порядке{{else if eq .Balance 0}}Доступ заблокирован{{else}}Пора пополнить{{end}}.`
var templateText = `{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. {{if eq .Balance 0}}Доступ заблокирован{{else if lt .Balance 100}}Пора пополнить{{else}}Все в порядке{{end}}.`

// конец решения

type User struct {
	Name    string
	Balance int
}

// renderToString рендерит данные по шаблону в строку
func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

func Test(t *testing.T) {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 500}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		t.Errorf("%v: got '%v'", user, got)
	}
}

func main() {
	Test(&testing.T{})
}