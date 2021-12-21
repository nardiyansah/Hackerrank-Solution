/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumDifference = function (nums) {
  let max = -1;
  for (let i = 0; i < nums.length - 1; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[j] > nums[i]) {
        let diff = nums[j] - nums[i];
        if (diff > max) {
          max = diff;
        }
      }
    }
  }
  return max;
};
