/**
 * @param {string} s
 * @return {number}
 */
var myAtoi = function (s) {
  const len = s.length;

  // 空格 charCode 32
  // + charCode 43
  // - charCode 45
  // 0 charCode 48
  // ..
  // 9 charCode 57

  let r = '';
  let z = 1;
  let f = 0;
  for (let i = 0; i < len; ++i) {
    const c = s[i].charCodeAt();
    if (c === 32 && !r) {
      if (f) break;
      continue;
    }

    if (c === 32 && r) {
      break;
    }

    if (c === 43 && r) {
      break;
    }

    if (c === 43 && !r && !f) {
      f = 1;
      continue;
    }

    if (c === 45 && !r && !f) {
      f = 1;
      z = -1;
      continue;
    }

    if (c === 45 && r) {
      break;
    }

    if (c < 48 || c > 57) {
      break;
    }

    r += s[i]
  }

  r *= z
  if (!r) {
    return 0
  }
  if (r < -2147483648) {
    return -2147483648;
  }
  if (r > 2147483647) {
    return 2147483647
  }

  return r;
};


console.log(myAtoi("   -42"), -42)
console.log(myAtoi("4193 with words"), 4193)
console.log(myAtoi("42"), 42)
console.log(myAtoi("+42"), 42)
console.log(myAtoi("-42"), -42)
console.log(myAtoi("-+42"), 0)
console.log(myAtoi("-+-42"), 0)
console.log(myAtoi("  +  413"), 0)
