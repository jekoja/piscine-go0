package main

import (
	"os"
)

const (
	maxInt64 = 1<<63 - 1
	minInt64 = -1 << 63
)

func parseInt64(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	i := 0
	sign := int64(1)
	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		sign = -1
		i = 1
	}
	var u uint64 = 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := uint64(c - '0')
		// check overflow for unsigned accumulation:
		// for positive numbers, limit is maxInt64
		// for negative numbers, limit is uint64(maxInt64)+1
		limit := uint64(maxInt64)
		if sign < 0 {
			limit = uint64(maxInt64) + 1
		}
		if u > (limit-d)/10 {
			return 0, false
		}
		u = u*10 + d
	}
	if sign > 0 {
		if u > uint64(maxInt64) {
			return 0, false
		}
		return int64(u), true
	}
	// negative: allow upto 2^63
	if u > uint64(maxInt64)+1 {
		return 0, false
	}
	// convert unsigned absolute to int64 negative safely:
	if u == uint64(maxInt64)+1 {
		return minInt64, true
	}
	return -int64(u), true
}

func itoaInt64(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	var abs uint64
	if neg {
		// compute absolute value as uint64 safely
		u := uint64(n)
		abs = ^u + 1 // two's complement negation -> absolute value
	} else {
		abs = uint64(n)
	}
	var buf [20]byte
	pos := len(buf)
	for abs > 0 {
		pos--
		buf[pos] = byte('0' + (abs % 10))
		abs /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}

func addOK(a, b int64) (int64, bool) {
	if b > 0 && a > maxInt64-b {
		return 0, false
	}
	if b < 0 && a < minInt64-b {
		return 0, false
	}
	return a + b, true
}

func subOK(a, b int64) (int64, bool) {
	// a - b == a + (-b)
	if b == minInt64 {
		// -b would overflow, treat specially:
		// a - minInt64 overflows unless a >= 0 (result >= 2^63) — easier to compute with checks:
		// compute result using uint64 trick and verify bounds via parse limits is complex; instead:
		// use addition with checks by flipping signs where safe:
		if a >= 0 {
			// result >= 2^63 -> overflow
			return 0, false
		}
		// a < 0: a - minInt64 = a + 2^63; check if <= maxInt64
		// but since a < 0 and 2^63 + a <= maxInt64 is always true when a >= -1<<63? safe to compute via uint64:
		res := a - b
		// verify using division trick
		if a != 0 && res/b != -1 { // not a general check but b==minInt64 so skip
			// fallback: accept this as ok (rare edge)
		}
		return res, true
	}
	return addOK(a, -b)
}

func mulOK(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	// Try multiplication and verify by division (wraps on overflow, but check detects it)
	res := a * b
	if res/b != a {
		return 0, false
	}
	return res, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, okA := parseInt64(aStr)
	b, okB := parseInt64(bStr)
	if !okA || !okB {
		return
	}

	switch op {
	case "+":
		r, ok := addOK(a, b)
		if !ok {
			return
		}
		os.Stdout.Write([]byte(itoaInt64(r) + "\n"))
	case "-":
		r, ok := subOK(a, b)
		if !ok {
			return
		}
		os.Stdout.Write([]byte(itoaInt64(r) + "\n"))
	case "*":
		r, ok := mulOK(a, b)
		if !ok {
			return
		}
		os.Stdout.Write([]byte(itoaInt64(r) + "\n"))
	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		os.Stdout.Write([]byte(itoaInt64(a/b) + "\n"))
	case "%":
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		os.Stdout.Write([]byte(itoaInt64(a%b) + "\n"))
	default:
		// invalid operator: print nothing
		return
	}
}
