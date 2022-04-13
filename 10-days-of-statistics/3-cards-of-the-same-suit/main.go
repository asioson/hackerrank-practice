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
    // 52 cards in first pick, 51 cards in second pick 
    totalPossibilities := 52 * 51
    // In a suit, 13 cards in first pick, 12 cards in second pick (do this per suit)
    totalCases := 13 * 12 * 4

    soln := Frac{n: totalCases, d: totalPossibilities}
    fmt.Println("Probability =", soln.ToString())
}
