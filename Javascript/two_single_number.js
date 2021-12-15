var singleNumber = function (nums) {
  let first = 0;
  let second = 0;
  nums.forEach(num => {
    first = (first ^ num) & ~second;
    second = (second ^ num) & ~first;
  });
  return first;
};
