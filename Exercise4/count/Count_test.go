package count

import "testing"

func TestCountwords(t *testing.T) {

	str := []struct {
		input  string
		output int
	}{
		{"Hai ganesh", 2},
		{"     ", 0},
	}
	for i := 0; i < len(str); i++ {
		actualOutput := Countwords(str[i].input)
		if actualOutput != str[i].output {
			t.Errorf("Test case %d failed, expected %d but got %d", i+1, str[i].output, actualOutput)
		}
	}
}
