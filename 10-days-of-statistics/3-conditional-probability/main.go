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
    // possibilities are: (G, B), (B, G), and (B, B) 
    totalPossibilities := 3
    // case with both children are boys: (B, B)
    totalCases := 1 

    soln := Frac{n: totalCases, d: totalPossibilities}
    fmt.Println("Probability =", soln.ToString())
}
