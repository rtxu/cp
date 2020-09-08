func isPossiblePalindrome(s string, leftEnd, rightBegin int) bool {
    n := len(s)
    for i := 0; rightBegin + i < n && leftEnd - i >= 0; i++ {
        if s[rightBegin + i] == s[leftEnd - i] {
            continue
        } else {
            return false
        }
    }
    return true
}

func shortestPalindrome(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    
    foundOrigin := false
    originPos := (n-1)/2
    takeOriginAsMirrorPoint := n % 2 == 1
    if takeOriginAsMirrorPoint {
        // 占用 originPos 位置的字符作为中轴点
        if isPossiblePalindrome(s, originPos-1, originPos+1) {
            foundOrigin = true
        } else {
            originPos--
            takeOriginAsMirrorPoint = !takeOriginAsMirrorPoint
        }
    }
    
    for !foundOrigin {
        // 不占用 originPos
        if isPossiblePalindrome(s, originPos, originPos+1) {
            foundOrigin = true
            break
        }
        takeOriginAsMirrorPoint = !takeOriginAsMirrorPoint
        
        // 占用 originPos
        if isPossiblePalindrome(s, originPos-1, originPos+1) {
            foundOrigin = true
            break
        }
        takeOriginAsMirrorPoint = !takeOriginAsMirrorPoint
        
        originPos--
    }
    fmt.Println(originPos, "takeOriginAsMirrorPoint", takeOriginAsMirrorPoint)
    
    rightPart := s[originPos+1:]
    leftPart := make([]byte, len(rightPart))
    for i := 0; i < len(rightPart); i++ {
        leftPart[i] = rightPart[len(rightPart) - 1 -i]
    }
    fmt.Println(rightPart)
    
    if takeOriginAsMirrorPoint {
        return string(leftPart) + string(s[originPos]) + rightPart
    } else {
        return string(leftPart) + rightPart
    }
}
