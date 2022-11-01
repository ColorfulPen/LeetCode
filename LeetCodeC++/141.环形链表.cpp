/*
 * @lc app=leetcode.cn id=141 lang=cpp
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *f=head,*s=head;
        while(s!=nullptr){
            f=f->next;
            s=s->next;
            if(s)s=s->next;else return false;
            if(f==s)return true;
        }
        return false;
    }
};
// @lc code=end

