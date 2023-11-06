package main

import contain "GO_PROJECT3_2/pkg"

func main() {
	filePath := "/Users/artmtsh/Desktop/GOlang/Lets-Go-Programming/Go_project_3_2/text.txt"
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
