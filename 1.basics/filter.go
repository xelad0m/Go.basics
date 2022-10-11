package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
    // отфильтруйте `iterable` с помощью `predicate`
    // и верните отфильтрованный срез
    var res []int
    for _, num := range iterable {
        if predicate(num) {
            res = append(res, num)
        }
    }
    return res
}

func main() {
    src := readInput()
    // отфильтруйте `src` так, чтобы остались только четные числа
    // и запишите результат в `res`
    res := filter(func (i int) bool { return i%2 == 0 }, src)
    fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
// Ctrl+D закончить ввод (!)
func readInput() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)	// Set the split function for the scanning operation.
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        nums = append(nums, num)
    }
    return nums
}