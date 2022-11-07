/*
 * @lc app=leetcode.cn id=208 lang=cpp
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
 struct Node{
    int end;
    Node* next[26];
    Node(){
        end=0;
        for(int i=0;i<26;i++){
            next[i]=nullptr;
        }
    }
 };


class Trie {
private:
    Node *head;
public:
    Trie() {
        head=new Node();
    }
    
    void insert(string word) {
        Node* n=head;
        for(auto x:word){
            if(n->next[x-'a']==nullptr){
                n->next[x-'a']=new Node();
            }
            n=n->next[x-'a'];
        }
        n->end++;
    }
    
    bool search(string word) {
        Node *n=head;
        for(auto x:word){
            if(n->next[x-'a']!=nullptr){
                n=n->next[x-'a'];
            }else{
                return false;
            }
        }
        if(n->end){
            return true;
        }
        return false;
    }
    
    bool startsWith(string prefix) {
        Node *n=head;
        for(auto x:prefix){
            if(n->next[x-'a']!=nullptr){
                n=n->next[x-'a'];
            }else{
                return false;
            }
        }
        return true;
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
// @lc code=end

