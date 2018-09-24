package tpltags

// Iterator return slice with values in range 0..cnt-1
func Iterator(cnt uint) []uint {
	var iterations []uint
	var i uint
	for i = 0; i < cnt; i++ {
		iterations = append(iterations, i)
	}
	return iterations
}
