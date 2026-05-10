package main

func maxMessages(thresh int) int {
	for i := 0; ; i++ {
		thresh -= 100 + i
		if thresh < 0 {
			return i
		}
	}
}
