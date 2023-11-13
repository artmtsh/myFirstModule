package main

import contain "github.com/artmtsh/myFirstModule/pkg"

func main() {
	filePath := "text.txt"
	searchString := "в файле. В качестве"
	contains, err := contain.Contains(filePath, searchString)
	if err != nil {
		panic(err)
	}
	if contains {
		println("Подстрока найдена в файле.")
	} else {
		println("Подстрока НЕ найдена в файле.")
	}
}
