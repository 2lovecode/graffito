package modg

type HashMap struct {
	pool map[PkgID]*Pkg
}

func NewHashMap() *HashMap {
	o := new(HashMap)
	o.pool = make(map[PkgID]*Pkg)
	return o
}

func (hm *HashMap) Add(p *Pkg) {
	hm.pool[p.ID()] = p
}

func (hm *HashMap) Find(id PkgID) *Pkg {
	if _, ok := hm.pool[id]; ok {
		return hm.pool[id]
	}
	return nil
}
