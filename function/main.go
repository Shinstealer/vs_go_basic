package main

func sum(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func main() {

	count, total := sum(1, 2, 3, 4, 5)

	println(count, total)

	println("hello")
}
