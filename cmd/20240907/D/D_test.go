package main

import "testing"

func TestXxx(t *testing.T) {
	var s set

	s.add(1)
	s.add(2)
	s.add(3)
	s.add(5)

	t.Log(s, cap(s.slice))

	t.Log(s.LowerBound(2))
	t.Log(s.ReverseLowerBound(2))
	t.Log(s.ReverseLowerBound(4))
	t.Log(s.ReverseLowerBound(0))

	s.removeByIndex(1)

	t.Log(s)

	t.Log(s.LowerBound(6))
	t.Log(s.ReverseLowerBound(6))

	t.Log(s)
	s.removeByValue(3)
	s.removeByValue(5)
	s.removeByValue(2)
	t.Log(s)
}

func Test2(t *testing.T) {
	var s set

	s.add(0)
	s.add(3)
	s.add(4)
	s.add(7)
	s.add(9)
	s.add(22)

	t.Log(s.findByValue(9))
}

func Test3(t *testing.T) {
	w := newWalls(4, 3)

	w.Print()
}
