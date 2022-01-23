package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	// читаем имя из паспорта
	scanner.Scan()
	passportName := scanner.Text()

	// читаем имя из базы данных
	scanner.Scan()
	databaseName := scanner.Text()

	// сравниваем строки
	// допустима ситуация, когда имя человека в базе отличается от имени в паспорте
	// на одну замену, одно удаление или одну вставку символа
	passportSymbols := []rune(passportName)
	databaseSymbols := []rune(databaseName)

	for i := 0; i < len(passportSymbols); i++ {
		passportSymbol := passportSymbols[i]
		databaseSymbol := databaseSymbols[i]

		if passportSymbol != databaseSymbol {
		}
	}

	fmt.Print("OK")
	fmt.Print("FAIL")
}
