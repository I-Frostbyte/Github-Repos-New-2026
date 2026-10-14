# 🚀 Two Sum

-----------------------------------------------------------------------------------------------------------------------------
## The Problem:

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Example 1:

```go

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

```
 
## Constraints:

* 2 <= nums.length <= 104
* -109 <= nums[i] <= 109
* -109 <= target <= 109
* Only one valid answer exists.
 
## Follow-up
Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
-----------------------------------------------------------------------------------------------------------------------------

## Folder Structure

```
1-Two-Sum/
├── main.go        # TwoSum stub + empty main()
├── main_test.go   # Empty TestTwoSum placeholder
└── README.md
```

## Implementation Status

| Item | State |
| --- | --- |
| `main.go` | `main()` is empty; `TwoSum(nums []int, target int) []int` currently `panic("unimplemented")`. |
| `main_test.go` | `TestTwoSum` is declared but contains no assertions yet. |

## API

```go
// TwoSum returns the indices of the two numbers that add up to target.
func TwoSum(nums []int, target int) []int
```
