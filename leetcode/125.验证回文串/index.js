/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function (s) {
  if (s === " ") {
    return true;
  }

  const r = s.toLowerCase();

  const n = r.length;
  const a = [];
  for (let i = 0; i < n; ++i) {
    const c = r[i].charCodeAt();
    if ((c >= 48 && c <= 57) || (c >= 97 && c <= 122)) {
      a.push(r[i]);
    }
  }

  const al = a.length;
  for (let i = 0; i < al / 2; ++i) {
    if (a[i] !== a[al - i - 1]) {
      return false;
    }
  }

  return true;
};

console.log(isPalindrome("A man, a plan, a canal: Panama"));
