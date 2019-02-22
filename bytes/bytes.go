package bytes

import "bytes"

func Combine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
