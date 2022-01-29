package main

import (
	"fmt"
	"testing"
)

type Exam struct {
	X int
}

func getExam() Exam {
	return Exam{
		X: 1,
	}
}

func TestPointFromFunc(t *testing.T) {
	// 不能对函数返回值取指针，可能是因为不可取地址
	// fmt.Println(&getExam())
	fmt.Println(getExam())
}
