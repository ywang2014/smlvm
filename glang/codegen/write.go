package codegen

import (
	"e8vm.io/e8vm/link"
)

func writeBlock(f *link.Func, b *Block) {
	for _, inst := range b.insts {
		f.AddInst(inst.inst)
		if inst.sym != nil {
			s := inst.sym
			f.AddLink(s.fill, &link.PkgSym{s.pkg, s.sym})
		}
	}
}

func writeFunc(p *Pkg, f *Func) {
	lfunc := link.NewFunc()

	for b := f.prologue; b != nil; b = b.next {
		writeBlock(lfunc, b)
	}

	p.lib.DefineFunc(f.name, lfunc)
}
