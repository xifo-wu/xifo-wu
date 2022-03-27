/**
 * @param {number[]} rolls
 * @param {number} mean
 * @param {number} n
 * @return {number[]}
 */
var missingRolls = function (rolls, mean, n) {
  const result = new Array(n);
  const sum = rolls.reduce((p, c) => p + c, 0);
  const x = mean * (rolls.length + n) - sum;
  const avg = parseInt(x / n, 10);
  let yu = x % n;
  if (avg < 1 || x <= 0 || avg > 6 || (avg >= 6 && yu !== 0)) {
    return [];
  }

  console.log(
    `sum = ${sum}, x=${x}, n=${n}, mean=${mean}, avg=${avg}, yu=${yu}`
  );
  for (let i = 0; i < n; i++) {
    const zd = 6 - avg;
    if (yu === 0) {
      result[i] = avg;
      continue;
    }

    if (yu <= zd) {
      result[i] = avg + yu;
      yu = 0;
    } else {
      result[i] = 6;
      yu -= zd;
    }
  }

  return result;
};

console.log(missingRolls([3, 2, 4, 3], 4, 2), [6, 6]);
console.log(missingRolls([1,5,6], 3, 4), [2,3,2,2]);
console.log(missingRolls([1,2,3,4], 6, 4), []);
console.log(missingRolls([1], 3, 1), [5]);
console.log(missingRolls([3,5,3], 5, 3), []);
console.log(missingRolls([6,3,4,3,5,3], 1, 6), []);
