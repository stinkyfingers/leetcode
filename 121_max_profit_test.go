package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.



Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.



Constraints:

    1 <= prices.length <= 105
    0 <= prices[i] <= 104
*/

func maxProfit1(prices []int) int {
	var total, current int
	for i := 1; i < len(prices); i++ {
		current = current + prices[i] - prices[i-1]
		if current < 0 {
			current = 0
		}
		if current > total {
			total = current
		}
	}
	return total
}

/* testing */

func TestMaxProfit1(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 2, 4, 1, 6},
			expected: 5,
		},
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, test := range tests {
		res := maxProfit1(test.prices)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
