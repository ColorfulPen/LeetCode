/*
 * @lc app=leetcode.cn id=206 lang=cpp
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    // ListNode* ret=nullptr;
    ListNode* reverseList(ListNode* head) {
        // dg(head);
        // return ret;
        if(!head)return head;
        ListNode* x=head,*y=head->next;
        x->next=nullptr;
        while(y){
            ListNode*temp=y->next;
            y->next=x;
            x=y;
            y=temp;
        }
        return x;
    }
    // ListNode* dg(ListNode* head){
    //     if(head==nullptr||head->next==nullptr){
    //         ret=head;
    //         return head;
    //     }
    //     ListNode*temp=dg(head->next);
    //     temp->next=head;
    //     head->next=nullptr;
    //     return head;
    // }
};
// @lc code=end

