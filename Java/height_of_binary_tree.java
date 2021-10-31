public static int height(Node root) {
    // Write your code here.
    if (root == null) return 0;
    if (root.left == null && root.right == null) return 0;
    if (height(root.left) >= height(root.right)) {
        return 1 + height(root.left);
    } else {
        return 1 + height(root.right);
    }
}
