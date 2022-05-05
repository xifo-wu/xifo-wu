/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  let p = 0,
    q = 0,
    r = 1;

  for (let i = 0; i < n; ++i) {
    p = q;
    q = r;
    r = p + q;
  }

  return r;
};

console.log(climbStairs(1));
console.log(climbStairs(2));
console.log(climbStairs(3));
console.log(climbStairs(4));
