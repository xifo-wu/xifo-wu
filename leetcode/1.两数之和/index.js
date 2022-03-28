/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  const obj = {};
  for (let i = 0; i < nums.length; i++) {
    const c = target - nums[i];
    if (typeof obj[c] !== 'undefined') {
      return [obj[c], i];
    }

    obj[nums[i]] = i;
  }
};

console.log(twoSum([2,7,11,15] , 9))
console.log(twoSum([3,2,4] , 6))
console.log(twoSum([3,3] , 6))
