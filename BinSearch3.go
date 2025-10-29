package main

import "fmt"

const NMAX int = 10
type tabInt [NMAX]int

func main(){
	var data tabInt
	var nData, x1 int
	
	fmt.Scan(&nData)
	bacaData(&data, nData)
	fmt.Scan(&x1)
	
	cetakData(data, nData)
	
	idx := binSearch(data, nData, x1)
	if idx != -1 {
		fmt.Printf("Data ditemukan pada indeks ke-%d\n", idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}	
}

func bacaData(A *tabInt, n int){
	var i int
	
	
	if n > NMAX {
		n = NMAX
	}
	
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int){
	var i int
	
	fmt.Print("Data dalam array:")
	for i = 0; i < n; i++{
		fmt.Printf(" %d", A[i])
	}
	fmt.Println()
}

func binSearch(A tabInt, n, x int) int {
	var left, right, mid int

	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if x < A[mid] {
			right = mid - 1
		} else if x > A[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}