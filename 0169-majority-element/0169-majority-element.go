import "sort"

func majorityElement(nums []int) int {
    max := len(nums) / 2
    y := 0

    sort.Ints(nums)
    for x := range nums {
        // if these are the same num, count up, else reset
        if x != 0 && nums[x] == nums[x-1] {
            y++
            if y == max {
                return nums[x]
            }
        } else {
            y = 0
        }
    }

    return nums[0]
}