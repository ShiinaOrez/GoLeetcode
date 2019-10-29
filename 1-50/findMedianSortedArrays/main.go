package main

import "fmt"

/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length1 := len(nums1)
    length2 := len(nums2)
    re := float64(2)
    mid1 := int(len(nums1) / 2)
    mid2 := int(len(nums2) / 2)
    for mid1-mid2 <= 1 && mid1-mid2 >= -1 {
        fmt.Println(mid1, "<->", mid2)
        re *= 2
        if nums1[mid1] > nums2[mid2] {
            if float64(length1) / re < 1 {
                mid1 -= 1
            } else {
                mid1 = int(float64(mid1)-float64(length1) / re)
            }
            if mid1 < 0 {
                mid1 = 0
            }
            if float64(length2) / re < 1 {
                mid2 += 1
            } else {
                mid2 = int(float64(mid2)+float64(length2) / re)
            }
            if mid2 > length2-1 {
                mid2 = length2 -1
            }
        } else if nums1[mid1] < nums2[mid2] {
            if float64(length1) / re < 1 {
                mid1 += 1
            } else {
                mid1 = int(float64(mid1)+float64(length1) / re)
            }
            if mid1 > length1-1 {
                mid1 = length1 - 1
            }
            if float64(length2) / re < 1 {
                mid2 -= 1
            } else {
                mid2 = int(float64(mid2)-float64(length2) / re)
            }
            if mid2 < 0 {
                mid2 = 0
            }
        } else {
            break;
        }
    }
    return (float64(nums1[mid1]) + float64(nums2[mid2])) / float64(2)
}*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	index1 := 0
	index2 := 0
	length1 := len(nums1)
	length2 := len(nums2)
	if length1 == 0 {
		nums = nums2
	} else if length2 == 0 {
		nums = nums1
	} else {
		for index1 < length1 && index2 < length2 {
			if nums1[index1] <= nums2[index2] {
				nums = append(nums, nums1[index1])
				index1 += 1
			} else {
				nums = append(nums, nums2[index2])
				index2 += 1
			}
		}
		for index1 < length1 {
			nums = append(nums, nums1[index1])
			index1 += 1
		}
		for index2 < length2 {
			nums = append(nums, nums2[index2])
			index2 += 1
		}
	}
	if (length1+length2)%2 == 1 {
		return float64(nums[(length1+length2-1)/2])
	}
	fmt.Println(nums[(length1+length2)/2], nums[(length1+length2)/2-1], nums)
	return (float64(nums[(length1+length2)/2]) + float64(nums[(length1+length2)/2-1])) / 2
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3, 4}))
}
