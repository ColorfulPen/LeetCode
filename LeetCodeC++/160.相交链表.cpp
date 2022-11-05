/*
 * @lc app=leetcode.cn id=160 lang=cpp
 *
 * [160] 相交链表
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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        ListNode*a=headA,*b=headB;
        // while(a!=nullptr&&b!=nullptr){
        //     a=a->next;
        //     b=b->next;
        // }
        // if(b!=nullptr){
        //     a=headB;
        //     while(b!=nullptr){
        //         b=b->next;
        //         a=a->next;
        //     }
        //     b=headA;
        // }else if(a!=nullptr){
        //     b=headA;
        //     while(a!=nullptr){
        //         a=a->next;
        //         b=b->next;
        //     }
        //     a=headB;
        // }else{
        //     a=headB;
        //     b=headA;
        // }
        // while(a!=b){
        //     a=a->next;
        //     b=b->next;
        // }
        while(a!=b){
            a=(a==nullptr)?headB:a->next;
            b=(b==nullptr)?headA:b->next;
        }
        return a;
    }
};
// @lc code=end

