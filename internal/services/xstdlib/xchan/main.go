package xchan

type hchan struct {
	buf   []rune
	sendx uint64
	recvx uint64
	sendq interface{}
	recvq interface{}
}
