package com.problem.solving.linkedList.removeDuplicate;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.LinkedList;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    @DisplayName("empty list should return empty list")
    void removeDuplicate1() {
        LinkedList<Integer> list1 = new LinkedList<Integer>();
        LinkedList<Integer> list2 = new LinkedList<Integer>();
        Solution solution = new Solution(list1);
        solution.removeDuplicate();
        assertEquals(list2 , solution.getLinkedList());
    }

    @Test
    @DisplayName("list with one element")
    void removeDuplicate2() {
        LinkedList<Integer> list1 = new LinkedList<Integer>();
        list1.add(2);
        LinkedList<Integer> list2 = new LinkedList<Integer>();
        list2.add(2);
        Solution solution = new Solution(list1);
        solution.removeDuplicate();
        assertEquals(list2 , solution.getLinkedList());
    }

    @Test
    @DisplayName("case with no duplicate")
    void removeDuplicate3() {
        LinkedList<Integer> list1 = new LinkedList<Integer>();
        list1.add(1);
        list1.add(2);
        list1.add(3);
        LinkedList<Integer> list2 = new LinkedList<Integer>();
        list2.add(1);
        list2.add(2);
        list2.add(3);
        Solution solution = new Solution(list1);
        solution.removeDuplicate();
        assertEquals(list2 , solution.getLinkedList());
    }

    @Test
    @DisplayName("case with duplicate")
    void removeDuplicate4() {
        LinkedList<Integer> list1 = new LinkedList<Integer>();
        list1.add(1);
        list1.add(2);
        list1.add(1);
        list1.add(3);
        list1.add(2);
        LinkedList<Integer> list2 = new LinkedList<Integer>();
        list2.add(1);
        list2.add(2);
        list2.add(3);
        Solution solution = new Solution(list1);
        solution.removeDuplicate();
        assertEquals(list2 , solution.getLinkedList());
    }
}