func isPalindrome(s string) bool {
    var (
        x int
        y = len(s)-1
    )

    str := []rune(s)

    for x < y {
        xChar := str[x]
        yChar := str[y]

        if !isAlphaNumeric(xChar) {
            x++
        } else if !isAlphaNumeric(yChar) {
            y--
        } else if toLowerCase(xChar) != toLowerCase(yChar) {
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