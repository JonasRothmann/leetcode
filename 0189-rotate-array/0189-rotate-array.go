

func rotate(nums []int, k int)  {
    if len(nums) == 1 {
        return
    }

    k %= len(nums)

    start := nums[len(nums)-k:]
    end := nums[:len(nums)-k]
    result := append(start, end...)
    
    copy(nums, result)
}