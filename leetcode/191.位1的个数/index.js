/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function (n) {
  const s = n.toString(2);

  const l = s.length;
  let sum = 0;
  for (let i = 0; i < l; ++i) {
    if (s[i] === "1") {
      sum += 1;
    }
  }

  return sum;
};

console.log(hammingWeight(11))