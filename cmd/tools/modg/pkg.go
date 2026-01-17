package modg

import (
	"crypto/md5"
	"io"
)

type PkgID string

type Pkg struct {
	id             PkgID
	name           string
	version        string
	depends        []PkgID
	reverseDepends []PkgID
	IsVisit        bool
}

func NewPkg(name string, version string) *Pkg {
	o := new(Pkg)
	o.name = name
	o.version = version
	o.id = o.uniqID()
	return o
}

func (o *Pkg) uniqID() PkgID {
	h := md5.New()
	_, _ = io.WriteString(h, o.name)
	_, _ = io.WriteString(h, o.version)
	b := h.Sum(nil)
	return PkgID(b)
}

func (o *Pkg) ID() PkgID {
	return o.id
}

func (o *Pkg) AddDep(id PkgID) {
	hasExist := false
	for _, each := range o.depends {
		if each == id {
			hasExist = true
			break
		}
	}
	if !hasExist {
		o.depends = append(o.depends, id)
	}

}

func (o *Pkg) AddRDep(id PkgID) {
	hasExist := false
	for _, each := range o.reverseDepends {
		if each == id {
			hasExist = true
			break
		}
	}
	if !hasExist {
		o.reverseDepends = append(o.reverseDepends, id)
	}
}
