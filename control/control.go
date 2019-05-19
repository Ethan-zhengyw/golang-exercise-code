package main

import "fmt"
import "io/ioutil"
import "strconv"
import "os"
import "bufio"

func readfile() {
    const filename = "tmp.txt"
    contents, err := ioutil.ReadFile(filename)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}

func readfile2() {
    const filename = "tmp.txt"
    if contents, err := ioutil.ReadFile(filename); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }

    // fmt.Printf("%s\n", contents)
}

func eval(a, b int, op string) (int, error) {
    switch op {
        case "+":
            return a + b, nil
        case "-":
            return a - b, nil
        case "*":
            return a * b, nil
        case "/":
            return a / b, nil
        default:
            return 0, fmt.Errorf("unsupported operator: %s", op)
    }
}

func countsum() int {
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    return sum
}

func convertToBin(n int) string {
    if n == 0 {
        return "0"
    }

    result := ""

    for ; n > 0; n /= 2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
    }
    return result
}

func printFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

func forever() {
    for {
        fmt.Println("abc")
    }
}

func main() {
    readfile()
    readfile2()

    fmt.Println(eval(1, 1, "+"))
    fmt.Println(eval(2, 1, "-"))
    fmt.Println(eval(2, 3, "*"))  
    fmt.Println(eval(4, 2, "/"))  
    fmt.Println(eval(4, 3, "/"))
    fmt.Println(eval(4, 3, "x"))

    fmt.Println(
        countsum(),
        convertToBin(5),  // 101
        convertToBin(13),  // 1101
        convertToBin(12312393),
        convertToBin(0),
    )
    printFile("tmp.txt")
    // forever()
}

