package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
)

func reverse(s string) string { // O(n)
	rs := []rune(s)

	for l, r := 0, len(rs)-1; l < r; l, r = l+1, r-1 {
		rs[l], rs[r] = rs[r], rs[l]
	}

	return string(rs)
}

type BigInt struct {
	repr string
	sign bool
}

var Zero = BigInt{"0", true}

func (bi BigInt) Equals(other BigInt) bool {
	return bi.sign == other.sign && bi.repr == other.repr
}

func (bi *BigInt) FromString(s string) {
	if s[0] == '-' {
		s = s[1:]
		bi.sign = false
	} else {
		bi.sign = true
	}

	bi.repr += reverse(s)
}

func (bi *BigInt) FromInt(n int) {
	bi.FromString(fmt.Sprint(n))
}

func (bi BigInt) Neg() BigInt {
	if bi.repr == "0" {
		return BigInt{bi.repr, true}
	}

	return BigInt{bi.repr, !bi.sign}
}

func (bi BigInt) LessThen(other BigInt) bool {
	if !bi.sign && !other.sign {
		return other.Neg().LessThen(bi.Neg())
	}

	if bi.sign && !other.sign {
		return false
	}

	if !bi.sign && other.sign {
		return true
	}

	if len(bi.repr) > len(other.repr) {
		return false
	}

	if len(bi.repr) < len(other.repr) {
		return true
	}

	return reverse(bi.repr) < reverse(other.repr)
}

func (bi BigInt) LessThanOrEqual(other BigInt) bool {
	return bi.Equals(other) || bi.LessThen(other)
}

func (bi BigInt) Sub(other BigInt) BigInt {
	// x and y are both positive

	if !bi.sign && !other.sign { // -x - -y = y-x
		return bi.Add(other.Neg())
	}

	if bi.sign && !other.sign { // x - -y = x+y
		return bi.Add(other.Neg())
	}

	if !bi.sign && other.sign { // -x - y = -(x+y)
		return bi.Neg().Add(other).Neg()
	}

	if bi.LessThen(other) { // small - big = -(big-small)
		return other.Sub(bi).Neg()
	}

	length := max(len(bi.repr), len(other.repr))

	ret := strings.Builder{}

	carry := 0

	last_nonzero_index := 0

	for i := 0; i < length; i++ {
		digit := 0

		if i < len(bi.repr) {
			digit += int(bi.repr[i]) - '0'
		}

		if i < len(other.repr) {
			digit -= int(other.repr[i]) - '0'
		}

		digit += carry
		if digit < 0 {
			carry = -1
			digit += 10
		} else {
			carry = 0
		}

		if digit != 0 {
			last_nonzero_index = ret.Len()
		}

		fmt.Fprint(&ret, digit)
	}

	return BigInt{ret.String()[:last_nonzero_index+1], true}
}

func (bi BigInt) Add(other BigInt) BigInt {
	// x and y are both positive

	if !bi.sign && !other.sign { // -x + -y = -(x+y)
		return bi.Neg().Add(other.Neg()).Neg()
	}

	if bi.sign && !other.sign { // x + -y = x-y
		return bi.Sub(other.Neg())
	}

	if !bi.sign && other.sign { // -x + y = y-x
		return other.Sub(bi.Neg())
	}

	length := max(len(bi.repr), len(other.repr))

	ret := strings.Builder{}

	carry := 0

	for i := 0; i < length; i++ {
		digit := 0

		if i < len(bi.repr) {
			digit += int(bi.repr[i]) - '0'
		}

		if i < len(other.repr) {
			digit += int(other.repr[i]) - '0'
		}

		digit += carry
		carry = digit / 10
		digit = digit % 10

		fmt.Fprint(&ret, digit)
	}

	if carry != 0 {
		fmt.Fprint(&ret, carry)
	}

	return BigInt{ret.String(), true}
}

func (bi BigInt) times_10_to_n(n int) BigInt {
	if bi.repr == "0" {
		return Zero
	}

	return BigInt{strings.Repeat("0", n) + bi.repr, bi.sign}
}

func (bi BigInt) mult_by_digit(n int) BigInt {
	ret := strings.Builder{}

	carry := 0

	for i := 0; i < len(bi.repr); i++ {
		digit := (int(bi.repr[i]) - '0') * n

		digit += carry
		carry = digit / 10
		digit = digit % 10

		fmt.Fprint(&ret, digit)
	}

	if carry != 0 {
		fmt.Fprint(&ret, carry)
	}

	return BigInt{ret.String(), bi.sign}
}

func (bi BigInt) Mult(other BigInt) BigInt {
	if !bi.sign {
		return bi.Neg().Mult(other).Neg()
	}

	if !other.sign {
		return bi.Mult(other.Neg()).Neg()
	}

	if other.Equals(Zero) || bi.Equals(Zero) {
		return Zero
	}

	ret := Zero

	for i := 0; i < len(other.repr); i++ {
		digit := int(other.repr[i]) - '0'

		product := bi.mult_by_digit(digit)
		ret = ret.Add(product.times_10_to_n(i))
	}

	return ret
}

func (bi BigInt) Div(other BigInt) (BigInt, BigInt) {
	if other.Equals(Zero) {
		panic("Division by zero")
	}

	if !bi.sign {
		q, r := bi.Neg().Div(other)
		return q.Neg(), r.Neg()
	}

	if !other.sign {
		q, r := bi.Div(other.Neg())
		return q.Neg(), r
	}

	quotient := Zero
	remainder := Zero

	this := reverse(bi.repr)

	for _, digit := range this {
		remainder = remainder.times_10_to_n(1)
		remainder = remainder.Add(BigInt{string(digit), true})

		quotient = quotient.times_10_to_n(1)

		if other.LessThanOrEqual(remainder) {
			for i := 1; i <= 10; i++ {
				multiple := other.mult_by_digit(i)

				if remainder.LessThen(multiple) {
					quotient = quotient.Add(BigInt{strconv.Itoa(i - 1), true})

					remainder = remainder.Sub(multiple.Sub(other))
					break
				}

				if remainder.Equals(multiple) {
					quotient = quotient.Add(BigInt{strconv.Itoa(i), true})

					remainder = remainder.Sub(multiple)
					break
				}
			}
		}
	}

	return quotient, remainder
}

func (bi BigInt) String() string {
	ret := strings.Builder{}

	if !bi.sign {
		fmt.Fprint(&ret, "-")
	}

	fmt.Fprint(&ret, reverse(bi.repr))

	return ret.String()
}

func test_add() {
	for test := 0; test < 10000; test++ {
		a := BigInt{}
		b := BigInt{}

		aint := rand.Intn(2000) - 1000
		bint := rand.Intn(2000) - 1000

		a.FromInt(aint)
		b.FromInt(bint)

		expected := fmt.Sprint(aint + bint)

		if expected != a.Add(b).String() {
			fmt.Println(test, a, b, a.Add(b), aint+bint)
		}
	}
}

func test_sub() {
	for test := 0; test < 10000; test++ {
		a := BigInt{}
		b := BigInt{}

		aint := rand.Intn(2000) - 1000
		bint := rand.Intn(2000) - 1000

		a.FromInt(aint)
		b.FromInt(bint)

		expected := fmt.Sprint(bint - aint)

		if expected != b.Sub(a).String() {
			fmt.Println(test, b, a, b.Sub(a), bint-aint)
		}
	}
}

func test_mult() {
	for test := 0; test < 10000; test++ {
		a := BigInt{}
		b := BigInt{}

		aint := rand.Intn(2000) - 1000
		bint := rand.Intn(2000) - 1000

		a.FromInt(aint)
		b.FromInt(bint)

		expected := fmt.Sprint(bint * aint)
		result := b.Mult(a)

		if expected != result.String() {
			fmt.Println(test, b, a, result, bint*aint)
		}
	}
}

func test_div() {
	for test := 0; test < 10000; test++ {
		a := BigInt{}
		b := BigInt{}

		aint := rand.Intn(2000) - 1000
		bint := rand.Intn(2000) - 1000

		if aint == 0 { // chosen by a fair 200-sided dice
			aint = 134
		}

		a.FromInt(aint)
		b.FromInt(bint)

		expected := fmt.Sprint(bint/aint, bint%aint)

		q, r := b.Div(a)

		result := fmt.Sprint(q, r)

		if expected != result {
			fmt.Println(test, b, a, result, expected)
		}
	}
}

func test_corner_cases() {
	one := BigInt{}
	one.FromInt(1)

	for test := 0; test < 10000; test++ {
		a := BigInt{}

		aint := rand.Intn(2000) - 1000

		a.FromInt(aint)

		if !a.Mult(Zero).Equals(Zero) {
			panic(fmt.Errorf("%d * 0 is %s", aint, a.Mult(Zero)))
		}

		if !a.Add(Zero).Equals(a) {
			panic(fmt.Errorf("%d + 0 is %s", aint, a.Add(Zero)))
		}

		if !a.Mult(one).Equals(a) {
			panic(fmt.Errorf("%d * 1 is %s", aint, a.Mult(one)))
		}

		if a.Equals(Zero) {
			continue
		}

		q, r := a.Div(a)

		if !(q == one && r == Zero) {
			panic(fmt.Errorf("%s / %s is %s", a, a, a.Mult(a)))
		}
	}
}

func test_big() {
	for test := 0; test < 10000; test++ {
		aint := rand.Intn(100000) + 1<<31
		bint := rand.Intn(100000) + 1<<32

		bia := big.Int{}
		bib := big.Int{}

		bia.SetInt64(int64(aint))
		bib.SetInt64(int64(bint))

		mya := BigInt{}
		myb := BigInt{}

		mya.FromInt(aint)
		myb.FromInt(bint)

		res := big.Int{}

		res.Mul(&bia, &bib)

		bmult := fmt.Sprint(res.String())
		mmult := fmt.Sprint(mya.Mult(myb))

		if bmult != mmult {
			panic(fmt.Errorf("%d * %d should be %s, got %s", aint, bint, bmult, mmult))
		}

		res.Add(&bia, &bib)

		badd := fmt.Sprint(res.String())
		madd := fmt.Sprint(mya.Add(myb))

		if badd != madd {
			panic(fmt.Errorf("%d + %d should be %s, got %s", aint, bint, badd, madd))
		}

		res.Sub(&bia, &bib)

		bsub := fmt.Sprint(res.String())
		msub := fmt.Sprint(mya.Sub(myb))

		if bsub != msub {
			panic(fmt.Errorf("%d - %d should be %s, got %s", aint, bint, bsub, msub))
		}
	}
}

func main() {
	test_add()
	test_sub()
	test_mult()
	test_div()
	test_corner_cases()
	test_big()

	aint := rand.Intn(100000) + 1<<31
	bint := rand.Intn(100000) + 1<<32

	a := BigInt{}
	b := BigInt{}

	a.FromInt(aint)
	b.FromInt(bint)

	a = a.Mult(a)
	b = b.Mult(b)

	fmt.Printf("%s * %s = %s\n", a, b, a.Mult(b))
	fmt.Printf("%s + %s = %s\n", a, b, a.Add(b))
	fmt.Printf("%s - %s = %s\n", a, b, a.Sub(b))

	q, r := a.Div(b)
	fmt.Printf("%s / %s = %s (rem %s)\n", a, b, q, r)
}
