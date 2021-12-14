function superDigit(n, k) {
  // Write your code here
  if (n.length === 1) {
    return parseInt(n);
  }
  const split_digit = n.split('');
  let sum = 0;
  split_digit.forEach(element => {
    sum += parseInt(element);
  });
  return superDigit((sum * k).toString(), 1);
}
