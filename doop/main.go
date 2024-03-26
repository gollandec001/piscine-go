package main

import (
	"os"
)

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '-' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func intToString(num int) string {
	if num == 0 {
		return "0"
	}

	var isNegative bool

	if num < 0 {
		isNegative = true
		num = -num
	}

	var result string

	for num > 0 {
		digit := num % 10
		character := '0' + rune(digit)
		result = string(character) + result
		num /= 10
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func Atoi(s string) int {
	ans := 0
	sign := 1
	for i := len(s) - 1; i >= 0; i-- {
		a := len(s) - 1 - i
		if i == 0 && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				sign = -1
			}
			continue
		}
		if rune(s[i]) >= 48 && rune(s[i]) <= 58 {
			cml := int(s[i] - '0')
			for a > 0 {
				cml *= 10
				a--
			}
			ans += cml
		} else {
			return 0
		}
	}
	return ans * sign
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 3 || len(args) == 2 {
		return
	}
	if IsNumeric(args[0]) && IsNumeric(args[2]) && (args[0] < "9223372036854775807") && (args[2] < "9223372036854775807") && (args[0][1:] < "9223372036854775808") && (args[2][1:] < "9223372036854775808") {
		if args[1] == "+" {
			os.Stdout.Write([]byte(intToString(Atoi(args[0])+Atoi(args[2])) + "\n"))
		}
		if args[1] == "-" {
			os.Stdout.Write([]byte(intToString(Atoi(args[0])-Atoi(args[2])) + "\n"))
		}
		if args[1] == "/" {
			if Atoi(args[2]) == 0 {
				os.Stdout.Write([]byte("No division by 0\n"))
				return
			}
			os.Stdout.Write([]byte(intToString(Atoi(args[0])/Atoi(args[2])) + "\n"))
		}
		if args[1] == "*" {
			os.Stdout.Write([]byte(intToString(Atoi(args[0])*Atoi(args[2])) + "\n"))
		}
		if args[1] == "%" {
			if Atoi(args[2]) == 0 {
				os.Stdout.Write([]byte("No modulo by 0\n"))
				return
			}
			os.Stdout.Write([]byte(intToString(Atoi(args[0])%Atoi(args[2])) + "\n"))
		}
	} else {
		return
	}
}
