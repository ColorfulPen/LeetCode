/*
 * @lc app=leetcode.cn id=234 lang=cpp
 *
 * [234] 回文链表
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
    bool isPalindrome(ListNode* head) {
        if(head==nullptr||head->next==nullptr)return true;
        ListNode* newhead=new ListNode();
        newhead->next=head;
        ListNode *f=head->next,*s=head;
        bool tag=false;
        while(f){
            f=f->next;
            if(f)f=f->next;
            else{
                tag=true;
                break;
            }
            ListNode *temp=s->next;
            s->next=temp->next;
            temp->next=newhead->next;
            newhead->next=temp;
        }
        if(tag){
            f=newhead->next;
        }else{
            f=newhead->next->next;
        }
        s=s->next;
        while(s){
            if(s->val!=f->val)return false;
            s=s->next;
            f=f->next;
        }
        return true;
    }
};
// @lc code=end

