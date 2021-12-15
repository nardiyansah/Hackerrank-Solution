/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function (nums) {
  let pointer1 = 0;
  let pointer2 = nums.length - 1;
  let squared_nums = [];
  while (pointer1 <= pointer2) {
    if (pointer1 === pointer2) {
      squared_nums.unshift(Math.pow(nums[pointer1], 2))
      pointer1 += 1
    } else if (Math.pow(nums[pointer1], 2) >= Math.pow(nums[pointer2], 2)) {
      squared_nums.unshift(Math.pow(nums[pointer1], 2))
      pointer1 += 1
    } else {
      squared_nums.unshift(Math.pow(nums[pointer2], 2))
      pointer2 -= 1
    }
  }
  return squared_nums;
};
