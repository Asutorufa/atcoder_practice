package main

import "testing"

func TestCount(t *testing.T) {
	t.Log(count([]byte("101001000001111")))
	t.Log(countN([]byte("101001000001111")))

	t.Log(count([]byte("101001001000")))
	t.Log(countN([]byte("101001001000")))

	t.Log(count([]byte("000110010101000")))
	t.Log(countN([]byte("000110010101000")))

	t.Log(count([]byte("0001101111111000")))
	t.Log(countN([]byte("0001101111111000")))
}
