package main

import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var data tabInt

	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)

	cetakData(data)
}

func bacaData(A *tabInt) {
	if A.n >= NMAX {
		return
	}

	fmt.Scan(&A.info[A.n])
	A.n++
}

func cetakData(A tabInt) {

	for i := 0; i < A.n; i++ {
		fmt.Printf("<%d> ", A.info[i])
	}
}
