/*
 * @lc app=leetcode.cn id=817 lang=cpp
 *
 * [817] 链表组件
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
    int numComponents(ListNode* head, vector<int>& nums) {
        unordered_set<int>s(nums.begin(),nums.end());
        ListNode *x=head;
        int sum=0;
        while(x){
            if(s.count(x->val)){
                sum++;
                x=x->next;
                while(x&&s.count(x->val))
                    x=x->next;
            }else{
                x=x->next;
            }
        }
        return sum;
    }
};
// @lc code=end

