package array

/*
	Problem link: https://leetcode.com/problems/remove-element/
	Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.

	Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

	Return k after placing the final result in the first k slots of nums.

	Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

	Input: nums = [3,2,2,3], val = 3
	Output: 2, nums = [2,2,_,_]
	Explanation: Your function should return k = 2, with the first two elements of nums being 2.
	It does not matter what you leave beyond the returned k (hence they are underscores).

	Input: nums = [0,1,2,2,3,0,4,2], val = 2
	Output: 5, nums = [0,1,4,0,3,_,_,_]
	Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
	Note that the five elements can be returned in any order.
	It does not matter what you leave beyond the returned k (hence they are underscores).

*/

/*
Observations:
	- Elements have to be removed in-place.
	- Relative order of elements changed can be different from what they were before. This is explained well by second example above.

Approach:
	- we can have two pointers begin and i.
	- we update the value of array(nums) when nums[i]!= val:
		- if nums[i]!=val:
			nums[begin]=nums[i]
			begin++
	- So, we're essentially moving the array elements to the front of array if nums[i]!=val
	- return begin

Example: Consider the first example listed with problem above-
	[3,2,2,3], val = 3
	begin,i=0,0
	-> first iteration: begin=0 & i=0
		if nums[0]!=3: false
	-> second iteration begin=0 & i=1
		if nums[1]!=3: true
			nums[0(begin)]=nums[1(i)] // here we'll swap nums[0] with nums[1]
			begin++
	-> third iteration begin=1 & i=2
		if nums[2]!=3: true
			nums[1(begin)]=nums[2(i)] // here we'll swap nums[1] with nums[2]
			begin++
	-> fourth iteration begin=2 & i=3
		if nums[3]!=3: false

	at the end of fourth iteration nums now looks like this: nums[2,2,2,3]
	Since we don't care anything after k(elements left after removing `val`) elements, this solution is correct
*/

func removeElement(nums []int, val int) int {
	begin, len := 0, len(nums)
	for i := 0; i < len; i++ {
		if nums[i] != val {
			nums[begin] = nums[i]
			begin += 1
		}
	}
	return begin
}
