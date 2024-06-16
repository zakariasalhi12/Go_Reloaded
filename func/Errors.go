package reroaded

import "fmt"

const (
	Warning = "\033[33;1m Warning :\033[0m "
	Ex = "\033[0;1m Example :\033[0m "
	Error   = "\033[31;1m Error :\033[0m "
)

func PrintWarning(warn string) {
	fmt.Println(Warning + warn + "  ")
}

func PrintEror(erro string) {
	fmt.Println(Error + erro + "  ")
}

func PrintEX(ex string) {
	fmt.Println(Ex + ex)
}
