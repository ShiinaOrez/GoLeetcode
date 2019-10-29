package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	leftFactor, rightFactor := 1, 1
	n, d := numerator, denominator
	if numerator < 0 {
		leftFactor = -1
		n = -1 * numerator
	}
	if denominator < 0 {
		rightFactor = -1
		d = -1 * denominator
	}
	res := ""
	if leftFactor*rightFactor < 0 {
		res += "-"
	}
	if numerator == 0 {
		return "0"
	}
	if denominator == 1 { // 整除的特殊情况
		return strconv.Itoa(numerator)
	}
	if n%d == 0 { // 整除的一般情况
		return res + strconv.Itoa(numerator/denominator)
	}

	m := n % d       // 小数部分的分子
	z := (n - m) / d // 整数部分

	twoNum, fiveNum := 0, 0
	for d%5 == 0 {
		d /= 5
		fiveNum += 1
	}
	for d%2 == 0 {
		d /= 2
		twoNum += 1
	}
	if d == 1 { // 分母仅有2和5两质因子
		rightStr := ""
		denominator = rightFactor * denominator
		for m > 0 {
			rightStr += strconv.Itoa(m * 10 / denominator)
			m = m * 10 % denominator
		}
		return res + strconv.Itoa(z) + "." + rightStr
	}

	e := 1
	for getE(10, e, d) != 1 {
		e += 1
	}

	fiveFactor, twoFactor := 0, 0
	if twoNum > fiveNum {
		fiveFactor = twoNum - fiveNum
	} else if fiveNum > twoNum {
		twoFactor = fiveNum - twoNum
	}

	str := strconv.Itoa(n * int(math.Pow(2, float64(twoFactor))) * int(math.Pow(5, float64(fiveFactor))) / d)
	right := (n * int(math.Pow(2, float64(twoFactor))) * int(math.Pow(5, float64(fiveFactor)))) % d
	rightStr := ""
	for j := 0; j < e; j++ {
		rightStr += strconv.Itoa(right * 10 / d)
		right = right * 10 % d
	}
	maxNum := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}(twoNum, fiveNum)

	if len(str) > maxNum {
		res += str[:len(str)-maxNum] + "." + str[len(str)-maxNum:]
	} else {
		res += "0."
		res += strings.Repeat("0", maxNum-len(str))
		res += str
	}
	res += "(" + rightStr + ")"
	return res
}

func getE(a, b, mod int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = ret * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return ret
}

func main() {
	fmt.Println(fractionToDecimal(-2147483648, -1))
}
