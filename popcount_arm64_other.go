// +build arm64,!gccgo,!appengine,!darwin

package popcount

func hasPOPCNT() bool {
	return getProcFeatures()&(0xF<<20) != 15<<20
}
