package convert

func ByteToInt(in []byte) (out []int) {
	for i := 0; i < len(in); i++ {
		out = append(out, int(uint8(in[i])))
	}
	return out
}
