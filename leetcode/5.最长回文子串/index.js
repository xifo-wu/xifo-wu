// 给你一个字符串 s，找到 s 中最长的回文子串。

// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。

// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

const isHuiWenStr = (str) => {
  const len = str.length;
  for (let i = 0; i < len / 2; i++) {
    if (str[i] !== str[len - 1 - i]) {
      return false
    }
  }

  return true
}

/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
  const len = s.length;
  let res = '';
  for (let i = 0; i < len; i++) {
    for (let j = i; j < len; j++) {
      const ts = s.slice(i, j+1);
      if (isHuiWenStr(ts) && res.length <= ts.length) {
        res = ts
      }
    }
  }

  return res
};


console.log(longestPalindrome("babad"), "aba")
console.log(longestPalindrome("cbbd"), "bb")
