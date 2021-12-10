import java.util.HashMap;

class Solution {
    public int totalFruit(int[] fruits) {
        int max = 0;
        int fruit_in_baskets = 0;
        boolean isThe3rdTypeTree = false;
        HashMap<Integer, Integer> basket = new HashMap<Integer, Integer>();
        int start = 0;
        int end = 0;
        while (end <= fruits.length - 1) {
            if (isThe3rdTypeTree) {
                basket.put(fruits[start], basket.get(fruits[start]) - 1);
                if (basket.get(fruits[start]) == 0) {
                    basket.remove(fruits[start]);
                }
                fruit_in_baskets -= 1;
                start += 1;
            } else {
                if (basket.containsKey(fruits[end])) {
                    basket.put(fruits[end], basket.get(fruits[end]) + 1);
                } else {
                    basket.put(fruits[end], 1);
                }
                fruit_in_baskets += 1;
                end += 1;
            }
            if (basket.values().size() == 3) {
                isThe3rdTypeTree = true;
            } else {
                isThe3rdTypeTree = false;
            }
            if (fruit_in_baskets > max && !isThe3rdTypeTree) {
                max = fruit_in_baskets;
            }
        }
        return max;
    }
}
