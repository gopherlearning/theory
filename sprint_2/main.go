package main

import "fmt"

type Exp interface {
	// fmt.Stringer /* only required for the workaround, see below.*/
	Eval() Ty
}

type Ty interface{}
type Lit Ty

func (lit Lit[Ty]) Eval() Ty       { return Ty(lit) }
func (lit Lit[Ty]) String() string { return fmt.Sprintf("(lit %v)", Ty(lit)) }

type Eq struct {
	a Exp[Ty]
	b Exp[Ty]
}

func (e Eq[Ty]) String() string {
	// doesn't panic if String is called explicitely
	// (needs fmt.Stringer in Exp)
	return fmt.Sprintf("(eq %v %v)", e.a /*.String()*/, e.b /*.String()*/)
}

var (
	e0 = Eq[int]{Lit[int](128), Lit[int](64)}
	e1 = Eq[bool]{Lit[bool](true), Lit[bool](true)}
)

func main() {
	// this call prints:
	// %!v(PANIC=String method: runtime error: invalid memory address or nil pointer dereference)
	fmt.Printf("%v\n", e0)

	// this crashes the program with:
	// unexpected fault address 0x101010109
	// fatal error: fault
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x101010109 pc=0x106ed0d]
	fmt.Printf("%v\n", e1)
}
