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
    // cases with sum > 9 are (4,6), (5,5), (5,6), (6,4), (6,5), and (6,6)
    // hence cases with sum <= 9 is
    totalCases := totalPossibilities - 6

    soln := Frac{n: totalCases, d: totalPossibilities}
    fmt.Println("Probability =", soln.ToString())
}
