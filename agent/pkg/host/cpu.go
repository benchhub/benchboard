package host

type CPUStat struct {
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	Irq       uint64
	SoftIrq   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

func (s *CPUStat) Path() string {
	return "stat"
}
