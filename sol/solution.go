package sol

type Record struct {
	sIndex, pIndex int
}

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	cache := make(map[Record]bool)
	var dfs func(sIndex, pIndex int) bool
	dfs = func(sIndex, pIndex int) bool {
		if sIndex >= sLen && pIndex >= pLen {
			return true
		}
		if pIndex >= pLen {
			return false
		}
		record := Record{sIndex: sIndex, pIndex: pIndex}
		if val, ok := cache[record]; ok {
			return val
		}
		// character match
		match := sIndex < sLen && (s[sIndex] == p[pIndex] || p[pIndex] == '.')
		result := false
		// for next character == '*'
		if pIndex+1 < pLen && p[pIndex+1] == '*' {
			result = (match && dfs(sIndex+1, pIndex)) || dfs(sIndex, pIndex+2)
			cache[record] = result
			return result
		}
		if match {
			// found next match
			result = dfs(sIndex+1, pIndex+1)
			cache[record] = result
			return result
		}
		cache[record] = false
		return false
	}
	return dfs(0, 0)
}
