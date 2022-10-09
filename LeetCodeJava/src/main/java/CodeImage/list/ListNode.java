package CodeImage.list;

public class ListNode {
    int val;
    ListNode nextNode;

    public ListNode() {

    }

    public ListNode(int val) {
        this.val = val;
    }

    public ListNode(int val, ListNode node) {
        this.val = val;
        this.nextNode = node;
    }

}
