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
    // since one marble was drawn, only 6 marbles are left (2 red and 4 blue) 
    totalPossibilities := 6
    // 4 blue marbles remain
    totalCases := 4

    soln := Frac{n: totalCases, d: totalPossibilities}
    fmt.Println("Probability =", soln.ToString())
}
