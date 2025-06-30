package h264

type IBox interface {
	parse()
	write()
}

type Box struct {
	size int
	name string
}

type Fytp struct {
	Box
	majorBrand      uint
	minorVersion    uint
	compatibleBrand []uint
}
