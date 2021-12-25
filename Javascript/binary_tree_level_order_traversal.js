/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrder = function (root) {
  if (!root) {
    return [];
  }
  let queue = [root];
  let answer = [];
  let values = [];
  let temp = [];
  while (queue.length > 0) {
    let node = queue.shift();
    values.push(node.val);
    if (node.left) {
      temp.push(node.left);
    }
    if (node.right) {
      temp.push(node.right)
    }
    if (queue.length === 0) {
      answer.push(values);
      values = [];
      queue = temp;
      temp = [];
    }
  }
  return answer;
};
