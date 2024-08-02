//暴力解法
func removeElement(nums []int, val int) int {
    size:=len(nums)
for i:=0;i<len(nums);i++{
    if nums[i]==val{
       for j:=i+1;j<size;j++{//画图 
        nums[j-1]=nums[j]//防止超范围所以设置j-1 
       }
       size--
       i--//确保在删除元素后不会跳过检查下一个元素
    }
    }
    return size
}
