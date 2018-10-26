package typedefs

import "fmt"

func Run()  {
	var paoA pao = new(badMan)
	var paoB pao = new(sportMan)

	paoA.pao()

	paoB.pao()
}

type badMan struct {
}

func (badMan *badMan) pao(){
	fmt.Println("BadMan.pao")
}

type sportMan struct {
}
func (sportMan *sportMan) pao(){
	fmt.Println("SportMan.pao")
}

type pao interface {
	pao()
}

