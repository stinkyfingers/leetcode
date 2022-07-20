/**
 * @param {string} s
 * @return {boolean}
 */
var repeatedSubstringPattern = function(s) {
   let cursor = s.length / 2;
   while (cursor > 0) {
       if (checkRepeatedSubstrings(s.slice(0, cursor), s.slice(cursor))) {
           return true;
       }
       cursor--;
   }
   return false;
};

const checkRepeatedSubstrings = (pattern, s) => {
    if (s.length % pattern.length !== 0) {
        return false;
    }
    let i = 0;
    while (i <= s.length - pattern.length) {
        if (pattern != s.slice(i, i + pattern.length)) {
            return false;
        }
        i += pattern.length;
    }
    return true;
}

const tests = [
    ['abab', true],
    ['aba', false],
    ['abcabcabcabc', true],
    ['ababab', true],
    ['aaa', true]
]
describe('repeatedSubstring', () => {
    test.each(tests)('can %s be made from its repeated substrings', (s, expected) => {
        expect(repeatedSubstringPattern(s)).toBe(expected);
    });
});