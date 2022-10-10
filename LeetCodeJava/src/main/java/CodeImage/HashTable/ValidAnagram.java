package CodeImage.HashTable;

import java.util.Arrays;

public class ValidAnagram {
    public static void main(String[] args) {
        String str1="rat";
        String str2="car";
        isAnagram(str1,str2);
        System.out.println(isAnagram(str1,str2));
    }
    public static boolean isAnagram(String s, String t) {
        int[] letters=new int[26];
        for (int i=0;i<s.length();i++){
            letters[s.charAt(i)-'a']+=1;
        }
        for (int i=0;i<t.length();i++){
            letters[t.charAt(i)-'a']-=1;
        }
        for (int i=0;i<26;i++){
            if (letters[i]!=0){
                return false;
            }
        }
        return true;
    }
}
