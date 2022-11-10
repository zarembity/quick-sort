/*
	При решении задачи использовал быструю сортировку, но в части partition применил логику перемещения по указателям из описания к задаче:
	Логика: Находим опорный элемент, слева от него перемещаем данные, которые меньше, в правой части, которые больше,
	возвращаем индекс опорного элемента(на момент преобразовавния массива его индекс может измениться),
	далее делим(перемещаем укзатели) массив на две части: start..pivot_idx, pivot_idx..end,
	далее рекурсиввно перестраиваем кажду часть пока кол-во элементов не станет < 2

	Пример преобразований:
	4,8,9,10,1,5,3,20 pivot_index=7 pivot_value=20
	4,8,9,3,1,5,10 pivot_index=6 pivot_value=10
	4,8,5,3,1,9 pivot_index=5 pivot_value=9
	4,1,3,5,8 pivot_index=3 pivot_value=5
	1,4,3 pivot_index=1 pivot_value=1
	3,4 pivot_index=2 pivot_value=4
	5,8 pivot_index=4 pivot_value=5
 */
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	less  = "<"
	more  = ">"
	equal = "="
)

type (
	studentInfo struct {
		name   string
		wBalls int
		mBalls int
	}
)

var (
	scanner *bufio.Scanner
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024*5)
	scanner.Buffer(buf, 1024*1024*5)
}

func main() {
	var (
		n    = scanInt()
		data []string
	)

	for i := 0; i < n; i++ {
		data = append(data, scanStr())
	}

	result := sortStudents(data)

	writer := bufio.NewWriter(os.Stdout)
	for i := len(result) - 1; i >= 0; i-- {
		writer.WriteString(result[i].name + "\n")
	}

	writer.Flush()
}

func sortStudents(in []string) []studentInfo {
	data := prepareData(in)
	return quickSort(prepareData(in), 0, len(data)-1)
}

func quickSort(arr []studentInfo, left, right int) []studentInfo {
	if (right - left) < 2 {
		return arr
	}

	pivot := arr[calcMid(left, right)]
	pivotIdx := partition(arr, left, right, pivot)

	quickSort(arr, left, pivotIdx)
	quickSort(arr, pivotIdx, right)

	return arr
}

func partition(arr []studentInfo, left, right int, pivotStudent studentInfo) int {
	var (
		leftBroken, rightBroken bool
		pivotIndex              int
	)

	for {
		if leftBroken && rightBroken {
			arr = swap(arr, left, right)
			leftBroken, rightBroken = false, false
		}

		if compareTwoStudents(arr[left], pivotStudent) == "=" {
			pivotIndex = left
		} else if compareTwoStudents(arr[right], pivotStudent) == "=" {
			pivotIndex = right
		}

		if left >= right {
			break
		}

		if compareTwoStudents(arr[left], pivotStudent) == "<" {
			left++
		} else {
			leftBroken = true
		}

		if compareTwoStudents(arr[right], pivotStudent) == ">" {
			right--
		} else {
			rightBroken = true
		}
	}

	return pivotIndex
}

func compareTwoStudents(left, right studentInfo) string {
	if left.wBalls == right.wBalls && left.mBalls == right.mBalls && left.name == right.name {
		return equal
	}

	if left.wBalls > right.wBalls {
		return more
	}

	if left.wBalls < right.wBalls {
		return less
	}

	if left.mBalls < right.mBalls {
		return more
	}

	if left.mBalls > right.mBalls {
		return less
	}

	cr := strings.Compare(left.name, right.name)
	if cr == -1 {
		return more
	}

	if cr == 1 {
		return less
	}

	return equal
}

func swap(arr []studentInfo, left, right int) []studentInfo {
	leftVal := arr[left]
	arr[left] = arr[right]
	arr[right] = leftVal

	return arr
}

func prepareData(in []string) []studentInfo {
	var result = make([]studentInfo, 0, len(in))
	for _, item := range in {
		info := strings.Split(item, " ")
		result = append(result, studentInfo{
			name:   info[0],
			wBalls: strToInt(info[1]),
			mBalls: strToInt(info[2]),
		})
	}

	return result
}

func strToInt(val string) int {
	intVar, _ := strconv.Atoi(val)
	return intVar
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	return strToInt(scanner.Text())
}

func calcMid(left, right int) int {
	return left + (right-left)/2
}
