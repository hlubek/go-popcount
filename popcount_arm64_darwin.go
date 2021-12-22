// +build arm64,!gccgo,!appengine,darwin

package popcount

func hasPOPCNT() bool {
	// We can assume arm64 on darwin means Apple M1 which supports the instruction
	return true
}
