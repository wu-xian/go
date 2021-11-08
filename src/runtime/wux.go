package runtime

func GetCurG() MyG {
	a := MyG{}
	a.G = getg()
	return a
}

func GetCurM() interface{} {
	return getg().m
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

type MyG struct {
	G *g
}

type MyGs []*MyG
