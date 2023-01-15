package count

import (
	"reflect"
)

// Countwords returns number of words in the string
func Countwords(sentence string) int {
	count := 0
	flag := 0
	spaces := 0
	for i := 0; i < len(sentence); i++ {
		if reflect.DeepEqual(string(sentence[i]), " ") {
			spaces++
			if flag == 1 {
				count++
				flag = 0
			}
		} else {
			flag = 1
		}
	}
	if spaces == len(sentence) {
		return 0
	}
	return count + 1
}
