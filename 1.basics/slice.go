package mainasjdals

import (
    "fmt"
)

func main() {
    var text string
    var width int
    fmt.Scanf("%s %d", &text, &width)

    // Возьмите первые `width` байт строки `text`,
    // допишите в конце `...` и сохраните результат
    // в переменную `res`
    tmp := make([]byte, len([]byte(text)))
    copy(tmp, []byte(text))
    
    var res string
    if width < len(text) {  
        res = string(tmp[:width]) + "..."
    } else {
        res = string(tmp)
    }

    fmt.Println(res)
}