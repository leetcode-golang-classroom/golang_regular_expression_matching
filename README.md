# golang_regular_expression_matching

## 

Given an input string `s` and a pattern `p`, implement regular expression matching with support for `'.'` and `'*'` where:

- `'.'` Matches any single character.
- `'*'` Matches zero or more of the preceding element.

The matching should cover the **entire** input string (not partial).

## Examples

**Example 1:**

```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

```

**Example 2:**

```
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

```

**Example 3:**

```
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

```

**Constraints:**

- `1 <= s.length <= 20`
- `1 <= p.length <= 30`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'.'`, and `'*'`.
- It is guaranteed for each appearance of the character `'*'`, there will be a previous valid character to match.

## 解析

給定兩個字串 s, p , s 代表要比較字串, p 代表要比較的Pattern字串

 Pattern p的字元 p[i] , 除了一般的字串外會出現以下兩種特殊字元

‘.’：代表可以 match 任何字元

‘*’: 代表 p[i-1] 可以出現 0 次或是多次

要求寫一個演算法來判斷 s 是否能夠 match Pattern 字串 p

可以發現 如果 Pattern p 中是一般字元或是 ‘.’

都跟其他比較方式一樣

比較麻煩的是當出現了 ‘*’ 這種字元的狀況

以下舉例來思考如何比對具有 ‘*’

 

![](https://i.imgur.com/cqNZDW4.png)

值得思考的是，何時代表比較結束

 1.   i ≥ len(s) && j ≥ len(p): 代表比較到最後一步, 有 match

1.  j ≥ len(p): 代表 pattern 用光,但 s 沒有被 match 完

而對於有 * 後墜的 pattern 的每個選擇

在 s[i] == p[j]  的情況下 可以繼續選擇用 p[j] 匹配

否則就使用 p[j+1] 做下一個比對的字元

假設定義 match(i, j) 為 s[i:] 是否 match p[j:]

則會有以下關係式

match(i,j) = (s[i] == p[j] && match(i+1,j))) || match(i,j+2)  if p[j+1] == ‘*’

match(i,j) = (s[i] == p[j] && match(i+1,j+1))) if p[j+1] ≠ ‘*’

match(i,j) = false otherwise

所求是 match(0,0)

畫出決策樹如下

![](https://i.imgur.com/IdymmUo.png)

因為一定要走完(i,j) 而 對每個決策樹假設是 * 的狀況都有 2 種選擇

所以最差的狀況有  O($2^{len(s)*len(p)}$)

以 Bruce Force 的方式 時間複雜度會是 O($2^{len(s)*len(p)}$)

因為 (i,j) 起始位置可能性有 len(s) * len(p)

假設透過 memorization 方式優化

時間複雜度是 O(len(s) * len(p))

而需要暫存所有比對過的位置結果 所以空間複雜度也是 O(len(s) * len(p))

比 Bruce Force 的方式 O($2^{len(s)*len(p)}$)還要好

## 程式碼
```go
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

```
## 困難點

1. 要理解如何去 match * pattern 要思考對應關係，還有要考慮到何時比對條件結束

## Solve Point

- [x]  建立一個 cache 用來暫存比對過的位置
- [x]  從 i = 0, j =0 開始逐步依照關係式比對 s 與 p