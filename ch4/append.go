package main

func main() {

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		// slice 仍有增长空间，扩展slice内容
		z = x[:zlen]
	} else {
		// slice已无空间，为他分配一个新的底层数组
		// 为了达到分摊纯属复杂性
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
