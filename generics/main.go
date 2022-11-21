package main

import (
	"fmt"

	"strings"
)

func main() {
	// // Initialize a map for the integer values
	// ints := map[string]int64{
	// 	"first":  34,
	// 	"second": 12,
	// }

	// // Initialize a map for the float values
	// floats := map[string]float64{
	// 	"first":  35.98,
	// 	"second": 26.99,
	// }

	// fmt.Printf("Non-Generic Sums: %v and %v\n",
	// 	SumInts(ints),
	// 	SumFloats(floats))
	// str := "你是中共"

	// sA := strings.Split(str, "")

	// fmt.Println(sA)
	// fmt.Println(isValid("(("))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	// fmt.Println(searchInsert([]int{1, 3}, 2))
	// var str = "   fly me   to   the moon  "
	// for _, v := range str {
	// 	fmt.Println(string(v))
	// 	fmt.Println()
	// }
	var str = "Today is a nice day"
	// fmt.Println(InterceptString(str))
	fmt.Println(lengthOfLastWord(str))
	// fmt.Print(strings.TrimSpace("tes"))

}

func lengthOfLastWord(s string) int {
	strArray := strings.Split(s, " ")

	max := 0
	for _, v := range strArray {

		ss := strings.TrimSpace(v)
		if ss != "" {
			max = len(ss)
		}

	}

	return max

}
func InterceptString(resStr string) int {

	r := []rune(resStr)
	return len(r)
	// if len(r) >= 10 {
	// 	result = string(r[:10])
	// }
	// return result
}
func searchInsert(nums []int, target int) int {

	if nums[len(nums)-1] < target {
		return len(nums)
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	if nums[0] >= target {
		return 0
	}
	var middle int
	for i, j := 0, len(nums); i <= j; {
		if middle == (i+j)/2 {
			break
		}
		middle = (i + j) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			j = middle - 1
		} else if nums[middle] < target {

			i = middle + 1

		}

	}

	if nums[middle] < target {
		return middle + 1
	} else if nums[middle] > target {
		return middle
	} else {
		return middle - 1
	}

}
func searchInsert3(nums []int, target int) int {

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	if nums[0] > target {
		return 0
	}
	ch := make(chan int)

	go searchInsert2(nums, target, 0, len(nums)-1, ch)
	for i := range ch {
		return i
	}

	return 0

}
func searchInsert2(nums []int, target int, start int, end int, c chan int) {
	if end < start {
		return
	}
	var middle = (start + end) / 2
	if nums[middle] == target {
		c <- middle
		close(c)

	} else if nums[middle] > target {
		if end-start == 1 || end == start {
			c <- end
			close(c)

		} else {
			go searchInsert2(nums, target, start, middle, c)
		}

	} else {
		if end-start == 1 || end == start {
			c <- end
			close(c)

		} else {
			go searchInsert2(nums, target, middle, end, c)
		}

	}

}

func removeElement(nums []int, val int) int {

	var temp []int
	for i, v := range nums {
		if val == v {
			if len(nums) > i {
				temp = append(nums[:i], nums[(i+1):]...)
			}
			break
		}
	}
	if temp == nil {
		return len(nums)
	}
	return removeElement(temp, val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {

		return list1
	}

	list1Node := list1
	list2Node := list2

	var res *ListNode
	var end *ListNode

	for list1Node != nil || list2Node != nil {

		if list1Node != nil && list2Node != nil {
			if list1Node.Val > list2Node.Val {
				var temp = &ListNode{Val: list2Node.Val}
				if end == nil {
					res = temp
					end = temp
					list2Node = list2Node.Next
					continue
				}
				end.Next = temp
				end = temp
				list2Node = list2Node.Next
			} else if list1Node.Val < list2Node.Val {
				var temp = &ListNode{Val: list1Node.Val}
				if end == nil {
					res = temp
					end = temp
					list1Node = list1Node.Next
					continue
				}
				end.Next = temp
				end = temp
				list1Node = list1Node.Next
			} else {
				var temp2 = &ListNode{Val: list2Node.Val}
				if end == nil {
					res = temp2
					end = temp2
				} else {
					end.Next = temp2
					end = temp2
				}
				var temp1 = &ListNode{Val: list1Node.Val}
				end.Next = temp1
				end = temp1
				list2Node = list2Node.Next
				list1Node = list1Node.Next
			}
		} else if list1Node != nil {

			var temp = &ListNode{Val: list1Node.Val}
			if end == nil {
				res = temp
				end = temp
				list1Node = list1Node.Next
				continue
			}
			end.Next = temp
			end = temp
			list1Node = list1Node.Next

		} else {
			var temp = &ListNode{Val: list2Node.Val}
			if end == nil {
				res = temp
				end = temp
				list2Node = list2Node.Next
				continue
			}
			end.Next = temp
			end = temp
			list2Node = list2Node.Next
		}

	}

	return res
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var kv = map[byte]byte{41: 40, 93: 91, 125: 123}
	originArray := []byte(s)

	tempArray := []byte{}

	for _, v := range originArray {
		if v == 40 || v == 91 || v == 123 {
			tempArray = append(tempArray, v)
		} else {

			var last byte
			if len(tempArray) > 0 {
				last = tempArray[len(tempArray)-1]
			} else {
				last = 100
			}
			if kv[v] != last {
				return false
			}
			le := len(tempArray)
			tempArray = tempArray[:le-1]
		}
	}
	if len(tempArray) != 0 {
		return false
	}

	return true

}
func romanToInt(s string) int {
	kv := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sArray := strings.Split(s, "")
	var res = []int{}
	for i := 0; i < len(sArray); {
		first := sArray[i]
		var second = ""
		if i+1 < len(sArray) {
			second = sArray[i+1]
		}
		if first == "I" {
			if second == "V" {
				res = append(res, 4)
				i++
			} else if second == "X" {
				res = append(res, 9)
				i++
			} else {
				res = append(res, 1)

			}

		} else if first == "X" {
			if second == "L" {
				res = append(res, 40)
				i++
			} else if second == "C" {
				res = append(res, 90)
				i++
			} else {
				res = append(res, 10)
			}

		} else if first == "C" {
			if second == "D" {
				res = append(res, 400)
				i++
			} else if second == "M" {
				res = append(res, 900)
				i++
			} else {
				res = append(res, 100)
			}
		} else {
			res = append(res, kv[first])

		}
		i++
	}

	var r = 0
	for _, v := range res {
		r = r + v
	}
	return r

}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
