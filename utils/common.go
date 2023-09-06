package utils

import "strconv"

func GetFileName(ID, offset int) string {
	fileName := strconv.Itoa(ID) + "_" + strconv.Itoa(offset) + ".png"
	return fileName
}
