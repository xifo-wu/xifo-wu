function isMatch(left, right) {
  if (left === "[" && right === "]") return true;
  if (left === "(" && right === ")") return true;
  if (left === "{" && right === "}") return true;
  return false;
}

/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  const stack = [];
  let sLen = s.length;

  const left = "[{(";
  const right = "]})";

  for (let i = 0; i < sLen; ++i) {
    const currentS = s[i];
    // 如果左括号就入栈
    if (left.includes(currentS)) {
      stack.push(currentS);
    }

    // 如果是右括号就出栈
    if (right.includes(currentS)) {
      const r = stack[stack.length - 1];
      if (isMatch(r, currentS)) {
        stack.pop();
      } else {
        return false;
      }
    }
  }

  return stack.length === 0;
};

console.log(isValid("[}[}"));
console.log(isValid("[[]]"));
console.log(isValid("][]["));
console.log(isValid("["));
