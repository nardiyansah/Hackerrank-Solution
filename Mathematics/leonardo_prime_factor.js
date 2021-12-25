var assert = require('assert')

function primeCount(n) {
  // Write your code here
  if (n === 1) {
    return 0;
  }
  if (n === 2) {
    return 1;
  }
  let prime_numbers = [2];
  let max_n_prime = 1;
  let multiply_prime = 2;
  let new_prime = 3;
  let isPrime = true;
  while (multiply_prime < n) {
    for (let i = 0; i < prime_numbers.length; i++) {
      const element = prime_numbers[i];
      if (new_prime % element === 0) {
        isPrime = false;
        break;
      }
    }
    if (isPrime) {
      multiply_prime *= new_prime;
      if (multiply_prime <= n) {
        max_n_prime += 1;
        prime_numbers.push(new_prime);
      }
    } else {
      isPrime = true;
    }
    new_prime += 1;
  }
  return max_n_prime;
}

assert.equal(primeCount(614889782588491409), 14)