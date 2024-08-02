27. 移除元素
//暴力解法
func removeElement(nums []int, val int) int {
    size := len(nums)
    for i := 0; i < size; i ++ {//首先遍历数组
        if nums[i] == val {//当遇到删除元素时
            for j := i + 1; j < size; j ++ {//画图写一下就好了
                nums[j - 1] = nums[j]
            }
            i --//防止刚替换的数字没有遍历，所以i--
            size --
        }
    }
    return size
}
//快慢双指针
//设置slow fast 快慢指针，当fast碰到目标值不录入到 slow 就好了
func removeElement(nums []int, val int) int {
a:=0
for i:=0;i<len(nums);i++{
if nums[i]!=val{
nums[a]=nums[i]
a++
}
}
return a
}
//相向双指针
//避免与（&&）短路情况 设置 left right双指针向中间跑，保证右边和左边换的不算是目标值
func removeElement(nums []int, val int) int {
    right := len(nums) - 1
    left := 0
    for left <= right {
        // 确保先检查left <= right，再检查nums[left] != val，避免数组越界
        for left <= right && nums[left] != val {
            left++
        }
        // 确保先检查left <= right，再检查nums[right] == val，避免数组越界
        for left <= right && nums[right] == val {
            right--
        }
        if left < right {
            nums[left] = nums[right]
            right--
            left++
        }
    }
    return left
}
