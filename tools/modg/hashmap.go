package modg

type hashmap struct {
	pool map[pkgID]*pkg
}

func newHashmap() *hashmap {
	o := new(hashmap)
	o.pool = make(map[pkgID]*pkg)
	return o
}
