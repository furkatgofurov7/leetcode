// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
    tmp := make(map[int]int)
    
    for i, num := range nums {            
        if index, ok := tmp[target - num]; ok {
            return []int{index, i}
        } else {
            tmp[num] = i
        }
    }

    return nil
}