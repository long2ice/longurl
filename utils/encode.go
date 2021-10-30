package utils

var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(num uint64) string {
	var bytes []byte
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
