/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    let total = 0;
    const vals = {
        'I': 1,
        'V': 5,
        'X': 10, 
        'L': 50,
        'C': 100,
        'D': 500,
        'M':1000
    }
    let skip = false;

    for (let i = 0; i < s.length; i++) {
        if (skip) {
            skip = false;
            continue;
        }
        skip = false;
        let current = vals[s[i]];
        if (i < s.length - 1) {
            const next = vals[s[i+1]];
            if (next > current) {
                current = next - current;
                skip = true;
            }
        }
        total += current;
    }
    return total;
};

const tests = [
    ['III', 3],
    ['IV', 4],
    ['LVIII', 58],
    ['MCMXCIV', 1994]
];

test.each(tests)('converts roman %s to int %d', (roman, expected) => {
    expect(romanToInt(roman)).toBe(expected);
});