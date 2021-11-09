package runtime

import "unsafe"

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

type TheType _type

type TheEface eface

func (self TheEface) GetData() unsafe.Pointer {
	return self.data
}
func (self TheEface) GetType() TheType {
	return TheType(*self._type)
}
func (self TheEface) GetTypeName() string {
	return self._type.string()
}

type TheIface iface
