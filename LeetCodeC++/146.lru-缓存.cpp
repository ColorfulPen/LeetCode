/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU 缓存
 */

// @lc code=start
class LRUCache {
private:
    unordered_map<int ,list<pair<int,int>>::iterator> cache;
    list<pair<int,int>>lru;
    int capacity;
public:
    LRUCache(int _capacity):capacity(_capacity){

    }
    
    int get(int key) {
        auto it=cache.find(key);
        if(it!=cache.end()){
            lru.splice(lru.begin(),lru,it->second);
            return it->second->second;
        }
        return -1;
    }
    
    void put(int key, int value) {
        auto it=cache.find(key);
        if(it!=cache.end()){
            lru.splice(lru.begin(),lru,it->second);
            it->second->second=value;
            return ;
        }
        lru.emplace_front(key,value);
        cache[key]=lru.begin();
        if(cache.size()>capacity){
            cache.erase(lru.back().first);
            lru.pop_back();
        }
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end

