package modg

type pkgID string

type pkg struct {
	id             pkgID
	name           string
	version        string
	depends        []*pkg
	reverseDepends []*pkg
}

func newPkg(name string, version string) *pkg {
	o := new(pkg)
	o.id = ""
	return o
}

func (o *pkg) addDep(p *pkg) {
	o.depends = append(o.depends, p)
}

func (o *pkg) addRDep(p *pkg) {
	o.reverseDepends = append(o.reverseDepends, p)
}
