package CodeImage.HashTable;

import java.util.HashSet;
import java.util.Set;

public class HappyNumber {
    public boolean isHappy(int n){
        Set<Integer> set=new HashSet<>();
        set.add(n);
        while (true){
            if (n==1){
                return true;
            }
            n=getSum(n);
            if (set.contains(n)){
                return false;
            }else {
                set.add(n);
            }
        }
    }
    public int getSum(int n){
        int res=0;
        while (n!=0){
            int temp=n%10;
            n=n/10;
            res+=temp*temp;
        }
        return res;
    }
}
