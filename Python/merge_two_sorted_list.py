# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
        if list1 == None and list2 == None:
            return None
        elif list1 == None:
            return list2
        elif list2 == None:
            return list1
        head = ListNode()
        pointer1 = list1
        pointer2 = list2
        if pointer1.val <= pointer2.val:
            head = ListNode(val=pointer1.val)
            pointer1 = pointer1.next
        else:
            head = ListNode(pointer2.val)
            pointer2 = pointer2.next
        tail = head
        while pointer1 != None and pointer2 != None:
            if pointer1.val <= pointer2.val:
                tail.next = ListNode(val=pointer1.val)
                tail = tail.next
                pointer1 = pointer1.next
            else:
                tail.next = ListNode(pointer2.val)
                tail = tail.next
                pointer2 = pointer2.next
        while pointer1 != None:
            tail.next = ListNode(val=pointer1.val)
            tail = tail.next
            pointer1 = pointer1.next
        while pointer2 != None:
            tail.next = ListNode(pointer2.val)
            tail = tail.next
            pointer2 = pointer2.next
        return head
