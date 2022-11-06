/*
 * @lc app=leetcode.cn id=148 lang=cpp
 *
 * [148] 排序链表
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
    ListNode* sortList(ListNode* head) {
         return gbsort(head,nullptr);
    }

    ListNode* gbsort(ListNode* head,ListNode* tail){
        if(head==nullptr)return nullptr;
        if(head->next==tail){
            head->next=nullptr;
            return head;
        }
        ListNode *f=head,*s=head;
        while(f!=tail){
            f=f->next;
            if(f==tail)break;
            else f=f->next;
            s=s->next;
        }
        ListNode* x=gbsort(head,s);
        ListNode* y=gbsort(s,tail);
        ListNode* ans;
        if(y==nullptr||x->val<y->val){
            ans=x;
            x=x->next;
        }else{
            ans=y;
            y=y->next;
        }
        ListNode*temp=ans;
        while(x!=nullptr&&y!=nullptr){
            if(x->val<y->val){
                temp->next=x;
                x=x->next;
            }else{
                temp->next=y;
                y=y->next;
            }
            temp=temp->next;
        }
        while(x!=nullptr){
            temp->next=x;
            temp=temp->next;
            x=x->next;
        }
        while(y!=nullptr){
            temp->next=y;
            temp=temp->next;
            y=y->next;
        }
        temp->next=nullptr;
        return ans;

    }
};
// @lc code=end

