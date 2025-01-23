package main

import "fmt"

func convertToTitle(columnNumber int) string {
	next := columnNumber / 26
	remainder := columnNumber % 26
	if remainder == 0 {
		next -= 1
	}
	if next == 0 {
		return excelConvert(remainder)
	}
	title := convertToTitle(next)
	return title + excelConvert(remainder)
}

func excelConvert(remainder int) string {
	switch remainder {
	case 0:
		return "Z"
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "E"
	case 6:
		return "F"
	case 7:
		return "G"
	case 8:
		return "H"
	case 9:
		return "I"
	case 10:
		return "J"
	case 11:
		return "K"
	case 12:
		return "L"
	case 13:
		return "M"
	case 14:
		return "N"
	case 15:
		return "O"
	case 16:
		return "P"
	case 17:
		return "Q"
	case 18:
		return "R"
	case 19:
		return "S"
	case 20:
		return "T"
	case 21:
		return "U"
	case 22:
		return "V"
	case 23:
		return "W"
	case 24:
		return "X"
	case 25:
		return "Y"

	}
	return ""
}
func main() {
	columnNumber := 1
	//columnNumber := 28
	//columnNumber := 701
	//columnNumber := 52
	fmt.Println(convertToTitle(columnNumber))
}
