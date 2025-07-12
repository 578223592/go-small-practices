package another

import "fmt"

type inner struct {
}

func NewInner() inner {
	return inner{}
}

func (i inner) Name() {
	fmt.Println("inner")
}
