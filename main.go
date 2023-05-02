package main

import(
"fmt"

)
func divide(dividendo, divisor int) (int, int) {
    cociente := dividendo / divisor
    resto := dividendo % divisor
    return cociente, resto
}

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    cociente, resto := divide(17, 5)
    fmt.Println("Cociente:", cociente)
    fmt.Println("Resto:", resto)

    resultado := sum(1, 2, 3, 4, 5)
    fmt.Println("Resultado:", resultado)
}