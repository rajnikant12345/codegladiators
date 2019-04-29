package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func comparenums(s1, s2 []int) []int {
	l := len(s1)
	if l > len(s2) {
		l = len(s2)
	}
	for i :=l-1 ;i>=0;i-- {
		if s1[i] > s2[i] {
			return s1
		}
		if s1[i] < s2[i] {
			return s2
		}
	}
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func printstr(i []int) {
	for j := len(i)-1;j>=0;j-- {
		fmt.Print(i[j])
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func setval(val int) int {
	if val < 0 {
		return 0
	}
	return val
}


func CreateOutputArray(arr []int, prev []int, offset int) []int {
	var out []int
	if arr[offset] > 0 {
		out = append(out,arr[offset])
	}
	idx := prev[offset]
	for {
		if idx == -1 {
			break
		}
		if arr[idx] > 0 {
			out = append(out,arr[idx])
		}
		idx = prev[idx]
	}
	return out
}


func execute(reader *bufio.Reader, n int) {

	var sum  []int
	var prev []int

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	sum = make([]int,n)
	prev = make([]int,n)

	var maxelem = -1000000

	for i := 0; i < n ; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
		if arrItem > maxelem {
			maxelem = arrItem
		}
		sum[i] = 0
		prev[i] = -1
	}
	if maxelem <= 0 {
		fmt.Println(maxelem)
		return
	}
	if n == 1 {
		fmt.Println(arr[0])
		return
	}
	if n == 2 {
		fmt.Println( max(arr[0],arr[1]) )
		return
	}
	sum[n-1] = setval(arr[n-1])
	sum[n-2] = setval(arr[n-2])
	sum[n-3] = setval(arr[n-3]) + sum[n-1]
	if sum[n-1] > 0 {
		prev[n-3] = n-1
	}

	for i := n-4;i>=0;i-- {
		idx2 := i+2
		idx3 := i+3
		sum[i] = setval(arr[i]) + max(sum[idx2],sum[idx3])
		if sum[i] != 0 {
			if max(sum[idx2],sum[idx3]) != 0 {
				if sum[idx2] > sum[idx3] {
					prev[i] = idx2
				} else if sum[idx2] < sum[idx3] {
					prev[i] = idx3
				} else {
					if arr[idx2] > arr[idx3] {
						prev[i] = idx2
					} else {
						prev[i] = idx3
					}
				}
			}
		}
	}
	var out []int
	if sum[0] > sum[1] {
		out = CreateOutputArray(arr , prev, 0)
	} else if  sum[0] < sum[1] {
		out = CreateOutputArray(arr , prev, 1)
	} else {
		out1 := CreateOutputArray(arr , prev, 0)
		out2 := CreateOutputArray(arr , prev, 1)
		out = comparenums(out1,out2)
	}
	printstr(out)
}



func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	for i:=0;i<n;i++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)
		execute(reader , n)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
