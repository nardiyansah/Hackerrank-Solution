/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var rotateRight = function (head, k) {
  if (head === null || head.next === null) {
    return head;
  }
  let answerNode = head;
  let beforeLastNode = head;
  let lastNode = head.next;
  let numList = 1;
  while (head.next !== null) {
    head = head.next;
    numList += 1;
  }
  while (k % numList > 0) {
    while (lastNode.next !== null) {
      beforeLastNode = beforeLastNode.next;
      lastNode = lastNode.next;
    }
    lastNode.next = answerNode;
    beforeLastNode.next = null;
    answerNode = lastNode;
    beforeLastNode = lastNode;
    lastNode = beforeLastNode.next;
    k -= 1;
  }
  return answerNode;
};
