package main

import "fmt"

type Frac struct {
    n, d int
}

func gcd(a int, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}

func (f Frac) ToString() string {
    c := gcd(f.n, f.d)
    if c == 0 {
        return "undefined"
    }
    return fmt.Sprintf("%d/%d", f.n / c, f.d / c)
}

func main() {
    // 6 sides per dice
    totalPossibilities := 6 * 6 
    // cases with sum = 6 and faces are different are (1,5), (2,4), (4,2), and (5,1)
    totalCases := 4

    soln := Frac{n: totalCases, d: totalPossibilities}
    fmt.Println("Probability =", soln.ToString())
}
