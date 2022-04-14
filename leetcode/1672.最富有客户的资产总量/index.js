const arrSum = (arr) => {
  let sum = 0;
  arr.forEach(i => {
    sum += i
  });

  return sum
}

/**
 * @param {number[][]} accounts
 * @return {number}
 */
var maximumWealth = function (accounts) {
  let max = 0;
  accounts.forEach(account => {
    const sum = arrSum(account);
    if (max < sum) {
      max = sum
    }
  });

  return max
};

console.log(maximumWealth([[1,2,3],[3,2,1]]), 6)
console.log(maximumWealth([[1,5],[7,3],[3,5]]), 10)
console.log(maximumWealth([[2,8,7],[7,1,3],[1,9,5]]), 17)