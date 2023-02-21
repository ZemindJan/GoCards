package main

type ITest interface {
	test()
}

type Test struct {
	val int
}

func (test Test) test() {

}

func main() {
	slc := make([]ITest, 0)

	test1 := Test{3}
	test2 := Test{5}

	slc = append(slc, test1, test2)

	println(slc)
}
