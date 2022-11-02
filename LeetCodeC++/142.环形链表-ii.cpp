/*
 * @lc app=leetcode.cn id=142 lang=cpp
 *
 * [142] 环形链表 II
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
    ListNode *detectCycle(ListNode *head) {
        ListNode *f=head,*s=head;
        while(s!=nullptr){
            f=f->next;
            s=s->next;
            if(s!=nullptr)s=s->next;else return nullptr;
            if(s==f)break;
        }
        if(s==nullptr)return nullptr;
        f=head;
        while(s!=f){
            s=s->next;
            f=f->next;
        }
        return s;
    }
};
// @lc code=end

