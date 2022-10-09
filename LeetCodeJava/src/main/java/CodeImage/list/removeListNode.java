package CodeImage.list;

public class removeListNode {
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
        printNode(node1);
        printNode(removeNode(1, node1));
    }

    public static ListNode removeNode(int val, ListNode head) {
//        清除头节点开始相同的节点
        while (head != null && head.val == val) {
            head = head.nextNode;
        }
        if (head == null) {
            return null;
        }
//        清除头节点以外相同的节点
        ListNode pre = head;
        ListNode curr = head.nextNode;
        while (curr != null) {
            if (curr.val == val) {
                pre.nextNode = curr.nextNode;
            } else {
                pre = pre.nextNode;
            }
            curr = curr.nextNode;
        }
        return head;
    }

    public static void printNode(ListNode head) {
        while (head.nextNode != null) {
            System.out.print(head.val + " ");
            head = head.nextNode;
        }
        System.out.println(head.val);
    }
}
