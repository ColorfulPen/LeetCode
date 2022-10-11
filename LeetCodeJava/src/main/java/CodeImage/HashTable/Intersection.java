package CodeImage.HashTable;

import java.util.HashSet;
import java.util.Set;

public class Intersection {
    public int[] intersection(int[] nums1, int[] nums2) {
        if (nums1.length == 0 || nums2.length == 0){
            return new int[0];
        }

        Set<Integer> set1=new HashSet<>();
        Set<Integer> set2=new HashSet<>();
        for (int j : nums1) {
            set1.add(j);
        }
        for (int i:nums2){
            if (set1.contains(i)){
                set2.add(i);
            }
        }

        return set2.stream().mapToInt(x->x).toArray();

    }
}
