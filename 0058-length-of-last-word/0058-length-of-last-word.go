func lengthOfLastWord(s string) int {
    var count int
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == ' ' && count == 0 {
			continue
		} else {
            if s[i] == ' ' {
                break
            } 
            count++
        }
	}

    return count
}
