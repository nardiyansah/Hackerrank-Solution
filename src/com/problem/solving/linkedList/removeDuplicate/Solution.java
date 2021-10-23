package com.problem.solving.linkedList.removeDuplicate;

import java.util.LinkedHashSet;
import java.util.LinkedList;
import java.util.Set;

public class Solution {

    private LinkedList<Integer> linkedList = new LinkedList<Integer>();

    public Solution(LinkedList list) {
        this.linkedList = list;
    }

    public void removeDuplicate() {
        Set<Integer> set = new LinkedHashSet<Integer>();
        for (int i = 0; i<linkedList.size(); i++) {
            if (set.contains(linkedList.get(i))) {
                linkedList.remove(i);
            } else {
                set.add(linkedList.get(i));
            }
        }
    }

    public LinkedList getLinkedList() {
        return linkedList;
    }
}
