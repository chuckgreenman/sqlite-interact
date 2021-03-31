package sqliteutils

import "testing"

func TestHello(t *testing.T) {
	input := "Chuck"
	expectedOutput := "Hi, Chuck. Welcome!"

	output := hello(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}