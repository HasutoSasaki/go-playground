package vars

import "fmt"

func Pointer() {
	fmt.Println("Pointer:")
	var p *int
	fmt.Printf(("p=%v nil=%v\n"), p, p == nil)

	v := 10
	p = &v
	fmt.Printf("p=%v *p=%d v=%d\n", p, *p, v)

	*p = 20
	fmt.Printf("After *p=20 p=%v *p=%d v=%d\n", p, *p, v)

	q := new(int)
	fmt.Printf("q=%v *q=%d\n", q, *q)
	*q = 30
	fmt.Printf("After *q=30 q=%v *q=%d\n", q, *q)

	increment := func(n *int) {
		*n += 5
	}
	increment(p)
	fmt.Printf("After increment p=%v *p=%d v=%d\n", p, *p, v)

	var r *int
	if r == nil {
		r = new(int)
	}
	fmt.Printf("r=%v *r=%d\n", r, *r)
}
