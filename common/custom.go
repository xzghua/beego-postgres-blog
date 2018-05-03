package common


func Rem(divisor int) (bool) {
	if (divisor+1) % 4 == 0 {
		return true
	} else {
		return false
	}
}