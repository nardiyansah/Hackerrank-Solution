function merge(left, right) {
  let arr = [];
  while (left.length && right.length) {
    if (left[0] < right[0]) {
      arr.push(left.shift())
    } else {
      arr.push(right.shift())
    }
  }
  if (left.length > 0) {
    arr = arr.concat(left)
  }
}

function mergeSort(params) {

}