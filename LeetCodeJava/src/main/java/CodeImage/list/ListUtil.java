package CodeImage.list;

public class ListUtil {
    public static void main(String[] args) {
        test();

    }

    public static void test() {
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
        System.out.println(get(-1, node1));
        ListNode head = addAtHead(100, node1);
        printNode(head);
    }

    public static boolean get(int index, ListNode head) {
        int count = 0;
        while (head != null) {
            if (count == index) {
                return true;
            }
            head = head.nextNode;
            count++;
        }
        return false;
    }

    public static ListNode addAtHead(int val, ListNode head) {
        return new ListNode(val, head);
    }

    public static void printNode(ListNode head) {
        while (head.nextNode != null) {
            System.out.print(head.val + " ");
            head = head.nextNode;
        }
        System.out.println(head.val);
    }
}
