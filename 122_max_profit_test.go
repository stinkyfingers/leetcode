package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.



Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.

Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.

Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.



Constraints:

    1 <= prices.length <= 3 * 104
    0 <= prices[i] <= 104

*/

/*
For each day, if there is a price increase over yesterday, add it to the
running total. Return total.
*/
func maxProfit(prices []int) int {
	var total int
	for i := 1; i < len(prices); i++ {
		today := prices[i] - prices[i-1]
		if today < 1 {
			continue
		}
		total += today
	}
	return total
}

/* testing */

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{1, 2, 4},
			expected: 3,
		},
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			prices:   []int{7, 1, 5, 3, 8, 7, 6, 10},
			expected: 13,
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
		res := maxProfit(test.prices)
		fmt.Println(res, "...", test.expected)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("got %v, expected %v", res, test.expected)
		}
	}
}
