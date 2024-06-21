// three-legged-crow.zero
//
// @(#)color_test.go  May 29th, 2022
// Copyright(c) 2022, Leyton Goth. All rights reserved.

package color

import "testing"

func TestRandomRGB(t *testing.T) {
	t.Log(RandomRGB())
	t.Log(RandomRGB())
	t.Log(RandomRGB())
	t.Log(RandomRGB())
	t.Log(RandomRGB())
}

func TestRandomHex(t *testing.T) {
	t.Log(RandomHex())
	t.Log(RandomHex())
	t.Log(RandomHex())
	t.Log(RandomHex())
	t.Log(RandomHex())
}

func TestHashRGB(t *testing.T) {
	t.Log(HashLightRGB(""))
	t.Log(HashLightRGB(" "))
	t.Log(HashLightRGB("hello"))
	t.Log(HashLightRGB("world"))
	t.Log(HashLightRGB("Fowler–Noll–Vo (or FNV) is a non-cryptographic hash function created by Glenn Fowler, Landon Curt Noll, and Kiem-Phong Vo."))

	t.Log(HashDarkRGB(""))
	t.Log(HashDarkRGB(" "))
	t.Log(HashDarkRGB("hello"))
	t.Log(HashDarkRGB("world"))
	t.Log(HashDarkRGB("Fowler–Noll–Vo (or FNV) is a non-cryptographic hash function created by Glenn Fowler, Landon Curt Noll, and Kiem-Phong Vo."))
}
