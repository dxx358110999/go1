package my_util

func PageToStartEnd(page int64, size int64) (int64, int64) {
	/*
		计算分页偏移量
	*/
	start := (page - 1) * size
	end := start + size - 1 //redis是左闭右闭
	return start, end
}

func PageToStartSize(page int64, size int64) (start int64) {
	/*
		计算分页偏移量
	*/
	start = (page - 1) * size
	return start
}
