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

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	count, total := sum(1, 2, 3, 4, 5)

	println(count, total)

	asum :=
		func(n ...int) int {
			s := 0
			for _, i := range n {
				s += i
			}
			return s
		}
	result := asum(1, 2, 3, 4, 5, 6, 7)
	println("asum result =", result)
	next := nextValue()
	println(next())
	println(next())
	println(next())

}
