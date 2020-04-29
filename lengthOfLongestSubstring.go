func lengthOfLongestSubstring1(s string) int {
    sChars := []byte(s)
    m := make([]int, 128)
    maxLen := 0
    left := 0
    right := 0
    for ;right < len(s); right ++{
        char := sChars[right]
        if  left < m[char] {
            left = m[char]
        }
        m[char] = right + 1 
        if maxLen < right - left + 1 {
            maxLen = right - left + 1
        } 
    }
    return maxLen
}
func lengthOfLongestSubstring2(s string) int {
    sChars := []byte(s)
    m := make([]int, 128)
    maxLen := 0
    left := 0
    right := 0
    for ;right < len(s); right ++{
        char := sChars[right]
        if  left < m[char] {
            left = m[char]
        }
        m[char] = right + 1 
        if maxLen < right - left + 1 {
            maxLen = right - left + 1
        } 
    }
    return maxLen
}
