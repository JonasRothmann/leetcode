func removeDuplicates(nums []int) int {
    y := 0
    for x := 0; len(nums) > x; x++ {
        if x == 0 || nums[x] != nums[x-1] {
            nums[y] = nums[x]
            y++
        }
    }

    return y
}