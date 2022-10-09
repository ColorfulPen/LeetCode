package CodeImage.list;

public class reverseList {
    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
//        ListNode node2=new ListNode(2);
//        ListNode node3=new ListNode(6);
//        ListNode node4=new ListNode(1);
//        ListNode node5=new ListNode(4);
//        ListNode node6=new ListNode(5);
//        ListNode node7=new ListNode(6);
//        node1.nextNode=node2;
//        node2.nextNode=node3;
//        node3.nextNode=node4;
//        node4.nextNode=node5;
//        node5.nextNode=node6;
//        node6.nextNode=node7;
//        node7.nextNode=null;
        ListUtil.printNode(reverse(new ListNode()));

    }

    public static ListNode reverse(ListNode head) {
        ListNode pre = null;
        ListNode curr = head;
        ListNode next = head.nextNode;
        while (next != null) {
            curr.nextNode = pre;
            pre = curr;
            curr = next;
            next = next.nextNode;
        }
        curr.nextNode = pre;
        return curr;
    }
}
