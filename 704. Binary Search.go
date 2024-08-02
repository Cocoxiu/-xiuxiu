704. Binary Search
//就和数学一样，如果右就开不包含最右元素，小于。
func search(nums []int, target int) int {
right:=len(nums)//左闭右开
left:=0
for left<right{
mid:=(left+right)/2
if nums[mid]==target{
return mid
}
if target<nums[mid]{
    right=mid
}
if target>nums[mid]{
    left=mid+1
}
}
return-1
}
