package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Tuple struct {
	array  []int
	target int
}

type BinSearch func([]int, int) int

func GetFunctionName(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer())
	fnName := fn.Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fnName, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

// 在一个有序数组arr中, 寻找大于等于target的元素的第一个索引，如果存在, 则返回相应的索引index，否则, 返回arr的元素个数 n。
func main() {
	var testcase []Tuple
	testcase = append(testcase, Tuple{[]int{1, 3, 6, 13, 14, 14, 14, 14, 56, 86, 99, 111}, 14})
	testcase = append(testcase, Tuple{[]int{111}, 14})
	testcase = append(testcase, Tuple{[]int{}, 14})

	// testBinSearch(binSearchExactlyHalfClose, testcase)
	// testBinSearch(binSearchExactlyFullClose, testcase)
	testBinSearch(binSearchLowBound, testcase)

}

func testBinSearch(f BinSearch, testcase []Tuple) {

	fName := GetFunctionName(f, '.')
	fmt.Println("Function ====> ", fName)
	for i, t := range testcase {
		fmt.Println("testcaes ==> ", i)
		fmt.Print("array ==> ", t.array, "\t", "target ==> ", t.target)
		fmt.Println()
		ret := f(t.array, t.target)
		if ret == -1 {
			fmt.Println("Not found")
		} else {
			fmt.Println("Found index: ", ret)
		}
	}
	fmt.Println()
}

// wrong during byteDance interview
func binSearch(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		if a[0] == target {
			return a[0]
		} else {
			return -1
		}
	}
	lo, hi, mi := 0, len(a)-1, len(a)/2
	for ; lo < hi && a[mi] != target; mi = lo + (hi-lo)/2 {
		if a[mi] < target {
			lo = mi
		} else if a[mi] >= target {
			hi = mi + 1
		}
	}
	if a[mi] >= target {
		return a[mi]
	} else {
		return -1
	}
}

// [l, r)
// 半开半闭
// 极端情况的临界，即 l+1==r, 进入循环
// 进入循环后，mid的值是(2l+1)/2 => l
// 如果任然没有找到target, 即满足条件 target == arr[mid], 此时所在的区间是[l, l+1), 且mid==l， 数组中最后一个待对比元素arr[mid]也与target不等。
// 则无非两种情况，最终不管是 l=mid+1 或者 r=mid, 都会得到l==r 而跳出循环
func binSearchExactlyHalfClose(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if target == arr[mid] {
			return mid
		}
		// 此时[l, mid]这个区间的元素全部小于target，target只可能存在于区间[mid + 1, r)
		if target > arr[mid] {
			l = mid + 1
			// target < arr[mid]
			// 此时[mid, r)这个区间的元素全部大于target, target只可能存在于区间[l, mid-1],
			// 将全闭合区间[l, mid-1]调整为半开半闭区间[l, mid)与循环条件语义的上下文匹配
		} else {
			r = mid
		}
	}
	return -1
}

// [l, r]
// 全闭
// 极端情况的临界， 即 l==r, 进入循环
// 进入循环后， mid的值是l，同时也是r
// 如果任然没有找到target, 即满足条件 target == arr[mid], 此时所在的区间是[l, l+1), 且mid==l， 数组中最后一个待对比元素arr[mid]也与target不等。
// 则无非两种情况，最终不管是 l=mid+1 或者 r=mid-1, 都会得到l>r 而跳出循环
func binSearchExactlyFullClose(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if target == arr[mid] {
			return mid
		}
		// 此时[l, mid]这个区间的元素全部小于target，target只可能存在于区间[mid + 1, r)
		if target > arr[mid] {
			l = mid + 1
			// target < arr[mid]
			// 此时[mid, r)这个区间的元素全部大于target, target只可能存在于全闭合区间[l, mid-1],
		} else {
			r = mid - 1
		}
	}
	return -1
}

// [l, r]
// 全闭
// 在一个有序的数组arr中，寻找大于等于target的元素的第一个索引，如果存在返回相应索引，如果不存在返回-1
// 极端情况的临界， 即 l==r, 进入循环
// 此时 mid == l == r
// 如果arr[mid] <= target，那么这时可以肯定 mid == l == r 左侧的所有元素全部小于target，arr[mid]此时就是第一个大于等于target的元素
// 如果arr[mid] > target, 那么这时可以肯定 mid == l == r 左侧的所有元素全部小于target，右侧的所有元素全部大于target，而arr[mid]此时就是第一个大于target的元素

func binSearchLowBound(arr []int, target int) int {
	if len(arr) == 1 {
		if target <= arr[0] {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, len(arr)-1

	var mid = -1
	for l <= r {
		mid = (l + r) / 2
		// 当arr[mid] == target的时候，此时[mid, r]中的元素一定大于等于target，但是arr[mid]可能不是第一个等于target的元素,还需要继续在[l, mid-1]中寻找
		// 当arr[mid] > target的时候，此时[mid, r] 中的元素一定大于target，但是arr[mid]可能不是第一个大于等于target的元素,还需要继续在[l, mid-1]中寻找
		if target <= arr[mid] {
			r = mid - 1
			// 当arr[mid] < target的时候, [l, mid]中的元素全部小于target，不满足条件，
			// 此时大于等于target的元素只存在与[mid+1, r]区间中，
		} else {
			l = mid + 1
		}
	}
	return mid
}
