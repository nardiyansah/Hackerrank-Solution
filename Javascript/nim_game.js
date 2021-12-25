/**
 * @param {number} n
 * @return {boolean}
 */
 var canWinNim = function (n) {
  if (n <= 3) {
    return true;
  }
  if (n % 4 === 0) {
    return false;
  } else {
    return true;
  }
};