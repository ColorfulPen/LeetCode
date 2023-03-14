package dynamic
// dynamic programing
// dp[i][j] represent s[:i+1] is match p[:j+1]
// case1: s[i]==p[j] and is character
// f[i][j]=f[i-1][j-1] else false

// case2: p[j]=='*'
// match 0 - inf times,therefor:
// 0: f[i][j]=f[i][j-2]
// 1: f[i][j]=f[i-1][j-2]  if s[i]==p[j-1]
// 2: f[i][j]=f[i-2][j-2]  if s[i]==s[i-1]==p[j-1]
// etc...

// so we consider: (1) does not match characters, discard the combination
// (2) matches a character at the end of s, discards the character, and the combination continues to match;

// f[i][j]= (1) s[i] not equal p[j-1]: f[i][j-2]
//  (2) s[i] == p[j-1]:  f[i-1][j] or f[i][j-2]

func isMatch(s string,p string)bool{
	m,n:=len(s),len(p)
	matches:=func (i,j int) bool {
		if i==0{
			return false
		}
		if p[j-1]=='.'{
			return true
		}
		return s[i-1]==p[j-1]
	}
	f:=make([][]bool,m+1)
	for i:=0;i<=m;i++{
		f[i]=make([]bool, n+1)
	}
	f[0][0]=true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1]=='*'{
				if matches(i,j-1){
					f[i][j]=f[i-1][j] || f[i][j-2]
				}else{
					f[i][j]=f[i][j-2]
				}
			}else if matches(i,j){
				f[i][j]==f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
