package arch

import (
	"io"
	"time"

	"shanhu.io/smlvm/arch/devs"
)

// Config contains config for constructing a machine
type Config struct {
	MemSize uint32
	Ncore   int

	Output   io.Writer
	Screen   devs.ScreenRender
	Table    devs.TableRender
	RandSeed int64

	InitPC       uint32
	InitSP       uint32
	StackPerCore uint32

	BootArg uint32

	ROM string

	PerfNow func() time.Duration
}
