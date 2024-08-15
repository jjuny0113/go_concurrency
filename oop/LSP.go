package oop

type T interface {
	Something()
}

type S struct {
}

func (s *S) Something() {}

type U struct {
}

func (u *U) Something() {}

func q(t T) {

}

var y = &S{}

var u = &U{}

func test() {
	q(y)
	q(u)
}
