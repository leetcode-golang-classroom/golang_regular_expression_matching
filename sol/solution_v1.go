package sol

func isMatchV1(s, p string) bool {
	sLen, pLen := len(s), len(p)
	// dp[i][j] == s[i:] match p[j:]
	dp := make([][]bool, sLen+1)
	for row := range dp {
		dp[row] = make([]bool, pLen+1)
	}
	// empty match empty case
	dp[sLen][pLen] = true
	for sStart := sLen; sStart >= 0; sStart-- {
		for pStart := pLen - 1; pStart >= 0; pStart-- {
			match := sStart < sLen && (s[sStart] == p[pStart] || p[pStart] == '.')
			if pStart+1 < pLen && p[pStart+1] == '*' {
				dp[sStart][pStart] = dp[sStart][pStart+2] || (match && dp[sStart+1][pStart])
			} else {
				dp[sStart][pStart] = match && dp[sStart+1][pStart+1]
			}
		}
	}
	return dp[0][0]
}
