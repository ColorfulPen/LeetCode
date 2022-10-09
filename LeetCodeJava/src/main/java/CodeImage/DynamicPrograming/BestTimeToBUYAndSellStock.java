package CodeImage.DynamicPrograming;

public class BestTimeToBUYAndSellStock {
//    这个方法不是dp算法
    public int maxProfit(int[] prices) {
        int minPrice=prices[0];
        int maxProfit=0;
        for(int i=1;i<prices.length;i++){
            if(prices[i]<minPrice){
                minPrice=prices[i];
            }else{
                maxProfit=Math.max(maxProfit,prices[i]-minPrice);
            }
        }
        return maxProfit;
    }
}
