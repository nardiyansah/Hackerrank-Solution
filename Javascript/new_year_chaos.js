function minimumBribes(q) {
  // Write your code here
  let bribes = 0;
  let correct_order = [1, 2, 3];
  let index = 2;
  while (index < q.length) {
    correct_order.push(index + 2);
    if (correct_order[0] === q[index - 2]) {
      correct_order.splice(0, 1);
    } else if (correct_order[1] === q[index - 2]) {
      bribes += 1;
      correct_order.splice(1, 1);
    } else if (correct_order[2] === q[index - 2]) {
      bribes += 2;
      correct_order.splice(2, 1);
    } else {
      console.log('Too chaotic');
      return;
    }
    index += 1
  }
  if (correct_order[0] === q[q.length - 1]) {
    bribes += 1;
  }
  console.log(bribes);
}
