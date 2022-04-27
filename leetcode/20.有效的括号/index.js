/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  const stack = [];
  const len = s.length;

  const m = {
    '(': ')',
    '{': '}',
    '[': ']',
  }
  for (let i = 0; i < len; ++i) {
    const c = s[i];

    if (c === '(' || c === '{' || c === '[') {
      stack.unshift(m[c]);
    }

    if (c === ')' || c === '}' || c === ']') {
      const x = stack.shift();
      if (x !== c) {
        return false
      }
    }
  }

  if (stack.length > 0) {
    return false;
  }

  return true;
};


console.log(isValid("[}[}"))
console.log(isValid("[[]]"))
console.log(isValid("][]["))
console.log(isValid("["))