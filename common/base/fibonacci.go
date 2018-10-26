package base

import (
	"math/big"
	"fmt"
)

func init()  {
	var r *big.Int
	f := fibonacci()
	for i := 0; i < 10000; i++ {
		r = f()
	}
	fmt.Println(r)
}

func fibonacci() func() *big.Int {
	v, s := big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		var temp big.Int
		temp.Set(s)
		s.Add(s, v)
		v = &temp
		return s
	}
}
