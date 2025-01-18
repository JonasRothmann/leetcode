func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int, len(nums))

    for i, num := range nums {
        if i2, ok := hashmap[target-num]; ok {
            return []int{i, i2}
        }
        hashmap[num] = i
    }

    panic("unexpected")
}