package helpers

import "strconv"

func ConvertStringToInt32(str string) int32 {

	strInt, _ := strconv.ParseInt(str, 10, 32)

	return (int32(strInt))
}
