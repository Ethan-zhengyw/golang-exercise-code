package main

import "fmt"
import "math"
import "runtime"
import "reflect"

func apply(op func(int, int) int, a, b int) int {

    p := reflect.ValueOf(op).Pointer()
    opName := runtime.FuncForPC(p).Name()

    fmt.Printf("Calling function %s with args: (%d, %d)\n", opName, a, b)

    return op(a, b)
}

func pow(a, b int) int {
    return int(math.Pow(float64(a), float64(b)))
}

func flexibleArgsSum(numbers ...int) int {
    sum := 0

    for i := range numbers {
        sum += numbers[i]
    }
    
    return sum
}

func main() {
    fmt.Println(apply(pow, 3, 4))
    
    fmt.Println(apply(
        func(a int, b int) int { 
            return int(math.Pow(
                float64(a), float64(b)))
        }, 3, 4))
    
    fmt.Println(flexibleArgsSum(1, 2, 3, 4, 5))
}
