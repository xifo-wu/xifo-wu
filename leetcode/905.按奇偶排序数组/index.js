/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArrayByParity = function (nums) {
  const len = nums.length;
  const j = [];
  const o = [];

  for (let i = 0; i < len; i++) {
    if (nums[i] % 2 === 0) {
      o.push(nums[i])
    } else {
      j.push(nums[i])
    }
  }

  return o.concat(j)
};


console.log(sortArrayByParity([3,1,2,4]))