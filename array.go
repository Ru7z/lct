package lct

import (
	"math/rand"
	"time"
)

// GenRandomArray 生成测试随机列
func GenRandomArray() (array []int) {
	var (
		minLength = 4
		maxLength = 7
		randStep  = 10
		m         = make(map[int]struct{}, maxLength)
		randN     int
	)
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength-minLength) + minLength
	array = make([]int, length)
	for i := 0; i < length; i++ {
		var e bool
		for !e {
			randN = rand.Intn(randStep) + 1
			if _, ok := m[randN]; !ok {
				m[randN] = struct{}{}
				e = true
			} else {
				randN = rand.Intn(randStep) + 1
			}
		}
		array[i] = randN
	}
	return array
}

// GetSortedArray 生成有序列
func GetSortedArray() (array []int) {
	var (
		minLength = 3
		maxLength = 7
		randStep  = 3
	)
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength-minLength) + minLength
	array = make([]int, length+1)
	for i := 1; i <= length; i++ {
		randN := rand.Intn(randStep) + 1
		array[i] = randN + array[i-1]
	}
	return array[1:]
}
