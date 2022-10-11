package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

func main() {
    phrase := readString()

    // 1. Разбейте фразу на слова, используя `strings.Fields()`
    // 2. Возьмите первую букву каждого слова и приведите
    //    ее к верхнему регистру через `unicode.ToUpper()`
    // 3. Если слово начинается не с буквы, игнорируйте его
    //    проверяйте через `unicode.IsLetter()`
    // 4. Составьте слово из получившихся букв и запишите его
    //    в переменную `abbr`
 var abbr []rune
 
 for _, token := range strings.Fields(phrase) {
  for j, ch := range token {
   if j == 0 && unicode.IsLetter(ch) {
    abbr = append(abbr, unicode.ToUpper(ch))
   }
  }
 }
 
    fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
    rdr := bufio.NewReader(os.Stdin)
    str, _ := rdr.ReadString('\n')
    return str
}