class MinHeap {
  constructor() {
    // initialing the array heap and adding a dummy element at index 0
    this.heap = [null]
  }

  getMin() {
    // accessing the min element at index 1 in the heap array
    return this.heap[1]
  }

  insert(node) {
    // inserting the new node at the end of the heap array
    this.heap.push(node)

    // finding the correct position for the new node
    if (this.heap.length > 1) {
      let current = this.heap.length - 1

      // traversing up the parent node until the current node (current) is greater than the parent (current/2)
      // swapping the two nodes by using the ES6 destructuring syntax
      if (current > 1 && this.heap[Math.floor(current / 2)] > this.heap[current]) {
        let temp = this.heap[Math.floor(current / 2)]
        this.heap[Math.floor(current / 2)] = this.heap[current]
        this.heap[current] = temp
        current = Math.floor(current / 2)
      }
    }
  }

  remove() {
    // smallest element is at the index 1 in the heap array
    let smallest = this.heap[1]

    if (this.heap.length > 2) {
      this.heap[1] = this.heap[this.heap.length - 1]
      this.heap.splice(this.heap.length - 1)

      if (this.heap.length === 3) {
        if (this.heap[1] > this.heap[2]) {
          let temp = this.heap[1]
          this.heap[1] = this.heap[2]
          this.heap[2] = temp
        }
        return smallest
      }

      let current = 1
      let leftChildIndex = current * 2
      let rightChildIndex = current * 2 + 1

      while (this.heap[leftChildIndex] && this.heap[rightChildIndex] && (this.heap[current] > this.heap[leftChildIndex] || this.heap[current] > this.heap[rightChildIndex])) {
        if (this.heap[leftChildIndex] < this.heap[rightChildIndex]) {
          let temp = this.heap[current]
          this.heap[current] = this.heap[leftChildIndex]
          this.heap[leftChildIndex] = temp
          current = leftChildIndex
        } else {
          let temp = this.heap[current]
          this.heap[current] = this.heap[rightChildIndex]
          this.heap[rightChildIndex] = temp
          current = rightChildIndex
        }
        leftChildIndex = current * 2
        rightChildIndex = current * 2 + 1
      }
    }
    // if there are only two element in the array, we directly splice out the first element
    else if (this.heap.length === 2) {
      this.heap.splice(1, 1)
    } else {
      return null
    }

    return smallest
  }
}

let minHeap = new MinHeap();

minHeap.insert(10);
minHeap.insert(23);
minHeap.insert(36);
minHeap.insert(18);
minHeap.insert(45);
minHeap.insert(57);
minHeap.insert(32);

console.log(minHeap.heap)

minHeap.remove();

console.log(minHeap.heap)