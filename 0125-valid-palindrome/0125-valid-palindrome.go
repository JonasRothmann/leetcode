func isPalindrome(s string) bool {
    var (
        x int
        y = len(s)-1
    )

    str := []rune(s)

    for x < y {
        if !isAlphaNumeric(str[x]) {
            x++
        } else if !isAlphaNumeric(str[y]) {
            y--
        } else if toLowerCase(str[x]) != toLowerCase(str[y]) {
            return false
        } else {
            x++
            y--
        }
    }

    return true
}

func isAlphaNumeric(c rune) bool {
    return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}

func toLowerCase(c rune) rune {
    if c <= 90 {
        return c + 32
    } else {
        return c
    }
}