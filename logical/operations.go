package logical

// bool 异或操作
func XOR(src bool, target bool) bool {
	return (src || target) && !(src && target)
}
