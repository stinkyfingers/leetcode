/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
    for (let i = 0; i < haystack.length; i++) {
        if (findStr(haystack.slice(i), needle)) {
            return i;
        }
    }
    return -1;
};

const findStr = (haystack, needle) => {
    if (needle.length === 0) return true;
    if (haystack.length === 0) return false;
    if (haystack[0] !== needle[0]) return false;
    return findStr(haystack.slice(1), needle.slice(1));
};

const tests = [
    ['hello', 'll', 2],
    ['aaaa', 'bba', -1],
    ['test','', 0],
]
describe('strStr', () => {
    test.each(tests)('finds in %s needle %s at %d', (haystack, needle, expected) => {
        expect(strStr(haystack, needle)).toBe(expected);
    });
});