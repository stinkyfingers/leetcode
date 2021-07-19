package leetcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func multiply(num1 string, num2 string) string {
	arr := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		n1 := digit(num1[i])
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := digit(num2[j])

			index := i + j + 1
			value := n1 * n2

			for k := index; k >= 0; k-- {
				arr[k] += value

				if arr[k] < 9 || value == 0 {
					break
				}
				value = arr[k] / 10
				arr[k] = arr[k] % 10
			}

		}
	}

	var builder strings.Builder
	setZero := true
	for i, a := range arr {
		if setZero && a == 0 && i < len(arr)-1 {
			continue
		}
		if a > 0 {
			setZero = false
		}
		builder.WriteString(fmt.Sprintf("%d", a))
	}
	return builder.String()
}

func digit(b byte) int {
	switch b {
	case 48:
		return 0
	case 49:
		return 1
	case 50:
		return 2
	case 51:
		return 3
	case 52:
		return 4
	case 53:
		return 5
	case 54:
		return 6
	case 55:
		return 7
	case 56:
		return 8
	case 57:
		return 9
	}
	return 0
}

/* testing */

func TestMultiply(t *testing.T) {
	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			num1:     "4",
			num2:     "3",
			expected: "12",
		},
		{
			num1:     "123",
			num2:     "456",
			expected: "56088",
		},
		{
			num1:     "0",
			num2:     "0",
			expected: "0",
		},
		{
			num1:     "1",
			num2:     "1",
			expected: "1",
		},
		{
			num1:     "999",
			num2:     "999",
			expected: "998001",
		},
	}
	for _, test := range tests {
		res := multiply(test.num1, test.num2)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
