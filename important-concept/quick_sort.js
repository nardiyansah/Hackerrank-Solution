function swap(arr, leftIndex, rightIndex) {
  let temp = arr[leftIndex];
  arr[leftIndex] = arr[rightIndex];
  arr[rightIndex] = temp;
}

function partition(arr, left, right) {
  let pivot = arr[Math.floor((left + right) / 2)];
  let i = left;
  let j = right;
  while (i <= j) {
    while (arr[i] < pivot) {
      i++;
    }
    while (arr[j] > pivot) {
      j--;
    }
    if (i <= j) {
      swap(arr, i, j);
      i++;
      j--;
    }
  }
  return i;
}

function quicksort(arr, left, right) {
  let index;
  if (arr.length > 1) {
    index = partition(arr, left, right);

    if (left < index - 1) {
      quicksort(arr, left, index - 1);
    }
    if (right > index) {
      quicksort(arr, index, right);
    }
  }
  return arr;
}

let arr = [5, 3, 7, 6, 2, 9];
console.log(arr);
quicksort(arr, 0, arr.length - 1);
console.log(arr);
