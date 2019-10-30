package main

import (
	"fmt"
	"math"
)

const (
	lcg_seed = 0

	max_testnum          = math.MaxInt32
	max_lag_for_analysis = 10
	max_j_for_analysis = 1
	startblag            = 2
	startslag            = startblag - 1
)

var (
        mod                  = 6 // 1001 // 999983
	initialValues []int
	curValues     []int
	occurs        []int

	lcg_x = 0

	// recommended next values: (97, 33) for AMAZING quality or (17, 5) for less memory usage.
	fib_bigger_lag  = 17
	fib_smaller_lag = 5 // should be > 0
)

func LCG() int {
	return 1
	lcg_x = (lcg_x*2416 + 374441) % 1771875
	return lcg_x
}

func initfib() {
	curValues = make([]int, 0)
	initialValues = make([]int, 0)
	occurs = make([]int, mod)
	for i := 0; i < fib_bigger_lag; i++ {
		newval := LCG() % mod
		initialValues = append(initialValues, newval)
		curValues = append(curValues, newval)
	}
}

func fibrand() int {
	b := curValues[len(curValues)-fib_smaller_lag]
	a := curValues[0]
	//fmt.Println(curValues)
	//fmt.Printf("a=%d,b=%d \n\n", a, b)
	new := a + b
	if new >= mod {
		new -= mod
	}
	for j := 0; j < len(curValues)-1; j++ {
		curValues[j] = curValues[j+1]
	}
	curValues[len(curValues)-1] = new
	return new
}

// TESTS BELOW

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func checkIfRepeat() bool {
	for i := 0; i < len(initialValues); i++ {
		if curValues[i] != initialValues[i] {
			return false
		}
	}
	return true
}

func fibtest() int {
	for i := 0; i < max_testnum; i++ {
		rnd := fibrand()
		// fmt.Printf("%d,", rnd)
		occurs[rnd]++
		if checkIfRepeat() {
			return i + 1
		}
	}
	return max_testnum
}

func analyzeOccurs(print bool) int {
	period := fibtest()

	mean := 0.0
	for i := 0; i < len(occurs); i++ {
		// fmt.Printf("%d: %d occurs\n", i, occurs[i])
		mean += float64(i * occurs[i])
	}

	res_string := fmt.Sprintf("(%d, %d): ", fib_bigger_lag, fib_smaller_lag)
	if print {
		res_string += fmt.Sprintf("mean = %.5f  of %.2f ", mean/float64(period), float64(mod-1)/2.0)
	}

	if period == max_testnum {
		res_string += ("repeat did not occur. \n")
	} else {
		res_string += fmt.Sprintf("%d numbers generated until repeat\n", period)
	}

	if print {
		fmt.Print(res_string)
	}
	return period
}

func analyzeLags(print bool) string {
	min := -1
	max := -1
	var bestblag, bestslag, worstblag, worstslag int
perebor:
	for i := startblag; i <= max_lag_for_analysis; i++ {
		for j := startslag; j < 2; j++ {
			fib_bigger_lag = i
			fib_smaller_lag = j

			if gcd(fib_bigger_lag, fib_smaller_lag) != 1 {
				// fmt.Printf("%d and %d aren't coprime, skipping...\n", fib_bigger_lag, fib_smaller_lag)
				continue
			}

			if print {
				fmt.Printf("\n\nBIG=%d SMALL=%d\n", fib_bigger_lag, fib_smaller_lag)
			}

			lcg_x = lcg_seed
			initfib()
			length := analyzeOccurs(print)
			if length < min || min == -1 {
				min = length
				worstblag = i
				worstslag = j
			}
			if length > max {
				max = length
				bestblag = i
				bestslag = j
			}
			if length == max_testnum {
				if print {
					fmt.Print("MAX_TESTNUM EXCEEDED, continuation is meaningless. Stopping now...")
				}
				break perebor
			}
			if print {
				fmt.Printf("Best yet %d (lags %d, %d) worst yet %d (lags %d, %d)\n", max, bestblag, bestslag, min, worstblag, worstslag)
			}
		}
	}
	report := fmt.Sprintf("FINAL RESULT FOR MOD %d: Best %d (lags %d, %d) worst %d (lags %d, %d)", mod, max, bestblag, bestslag, min, worstblag, worstslag)
	if print {
		fmt.Printf(report)
	}
	return report
}

func main() {
	//initfib()
	//fibtest()
	//analyzeOccurs()
        res := make([]string, 0)
        for _, mod = range []int{2, 6, 10, 12} {
            fmt.Printf("\n\n============\n=== 1D%d ===\n", mod)
            rep := analyzeLags(true)
			res = append(res, rep)
        }
        fmt.Printf("\n\n==================================================\n\n")
        for _, rep := range res {
        	fmt.Println(rep)
		}
	fmt.Printf("Finished.\n")
}
