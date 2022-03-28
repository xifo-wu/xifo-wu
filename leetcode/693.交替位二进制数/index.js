// 给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。

// 示例 1：
// 输入：n = 5
// 输出：true
// 解释：5 的二进制表示是：101

// 示例 2：
// 输入：n = 7
// 输出：false
// 解释：7 的二进制表示是：111.

// 示例 3：
// 输入：n = 11
// 输出：false
// 解释：11 的二进制表示是：1011.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-number-with-alternating-bits
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * @param {number} n
 * @return {boolean}
 */
var hasAlternatingBits = function (n) {
  const s = n.toString(2);
  let flag = null;
  for (let i = 0; i < s.length; i++) {
    if (flag === s[i]) {
      return false;
    }

    flag = s[i];
  }

  return true;
};

console.log(hasAlternatingBits(5), true);
console.log(hasAlternatingBits(7), false);
console.log(hasAlternatingBits(11), false);
console.log(hasAlternatingBits(4), false);
console.log(hasAlternatingBits(3), false);
