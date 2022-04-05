const isPrime = (x) => {
  for (let i = 2; i * i <= x; i++) {
    if (x % i == 0) {
      return false;
    }
  }
  return true;
};

/**
 * @param {number} left
 * @param {number} right
 * @return {number}
 */
var countPrimeSetBits = function (left, right) {
  let result = 0;
  for (let i = left; i <= right; ++i) {
    const r = i.toString(2);
    let oneCount = 0
    for (let j = 0; j < r.length; ++j) {
      if (r[j] === '1') {
        oneCount++;
      }
    }

    if (oneCount > 1 && isPrime(oneCount)) {
      result++;
    }
  }

  return result;
};

console.log(countPrimeSetBits(6, 10));
console.log(countPrimeSetBits(10, 15));
