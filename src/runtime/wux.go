package runtime

func GetCurG() MyG {
	a := MyG{}
	a.G = getg()
	return a
}

func GetCurM() interface{} {
	return getg().m
}

func GetCurP() MyP {
	myp := MyP{}
	p := getg().m.p.ptr()
	myp.P = p
	return myp
}

func GetAllGs() MyGs {
	MyGs := MyGs{}
	for _, g := range allgs {
		a := &MyG{}
		a.G = g
		MyGs = append(MyGs, a)
	}
	return MyGs
}

func GetGCount() int {
	return len(allgs)
}

type MyG struct {
	G *g
}

type MyP struct {
	P *p
}

type MyGs []*MyG
