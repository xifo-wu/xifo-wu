/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let l = 0,
    r = height.length - 1,
    max = 0;
  while (l < r) {
    let t = 0;
    if (height[l] > height[r]) {
      t = (r - l) * height[r];
      r--;
    } else {
      t = (r - l) * height[l];
      l++;
    }

    if (t > max) {
      max = t;
    }
  }

  return max;
};

console.log(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]), 49);
