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
    // 7 balls in X, 9 in Y, and 8 in Z
    totalPossibilities := 7 * 9 * 8
    // case 1: 4 red in X, 5 red in Y, and 4 black in Z
    totalCases := 4 * 5 * 4
    // case 2: 4 red in X, 4 black in Y, and 4 red in Z
    totalCases += 4 * 4 * 4
    // case 3: 3 black in X, 5 red in Y, and 4 red in Z
    totalCases += 3 * 5 * 4

    soln := Frac{n: totalCases, d: totalPossibilities}
    fmt.Println("Probability =", soln.ToString())
}
