package CodeImage.list;

public class pairwiseExchange {
    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(6);
        ListNode node4 = new ListNode(1);
        ListNode node5 = new ListNode(4);
        ListNode node6 = new ListNode(5);
        ListNode node7 = new ListNode(6);
        node1.nextNode = node2;
        node2.nextNode = node3;
        node3.nextNode = node4;
        node4.nextNode = node5;
        node5.nextNode = node6;
        node6.nextNode = node7;
        node7.nextNode = null;
        ListUtil.printNode(exchange(node1));

    }

    public static ListNode exchange(ListNode head) {
        if (head == null || head.nextNode == null) {
            return head;
        }

        ListNode curr = head;
        ListNode dummyNode = new ListNode(-1, curr);
        ListNode preNode = dummyNode;
        ListNode next = head.nextNode;
        ListNode nextCurr = head.nextNode.nextNode;
        while (curr != null && next != null) {
            preNode.nextNode = next;
            next.nextNode = curr;
            curr.nextNode = nextCurr;
            preNode = curr;
            curr = preNode.nextNode;
            if (curr != null)
                next = curr.nextNode;
            if (next != null)
                nextCurr = next.nextNode;
        }
        return dummyNode.nextNode;
    }
}
