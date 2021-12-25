var assert = require('assert')

function gameWithCells(n, m) {
  if (n % 2 === 0 && m % 2 === 0) {
    return (n * m) / 4;
  } else {
    return (Math.floor(n / 2) + (n % 2)) * (Math.floor(m / 2) + (m % 2));
  }
}

assert.equal(gameWithCells(1, 1), 1);
assert.equal(gameWithCells(2, 4), 2);
assert.equal(gameWithCells(1, 4), 2);
assert.equal(gameWithCells(2, 5), 3)
