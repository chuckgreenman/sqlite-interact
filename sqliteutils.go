package sqliteutils

import "fmt"

func hello(name string) string {
	CreateDatabase("test.db")
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}
