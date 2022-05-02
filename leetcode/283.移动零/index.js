/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function (nums) {
  const n = nums.length;
  let l = 0;
  let r = 0;
  while (r < n) {
    if (nums[r] !== 0) {
      let t = nums[l];
      nums[l] = nums[r];
      nums[r] = t;
      l++;
    }

    r++;
  }

  return nums;
};

console.log(moveZeroes([0, 1, 0, 3, 12]));
