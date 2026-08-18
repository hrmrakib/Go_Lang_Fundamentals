package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["PY"] = "Python"

	fmt.Println(languages)

	delete(languages, "PY")

	fmt.Println(languages)

	for k, v := range languages {
		fmt.Printf("For key %v, value is %v\n", k, v)
	}
}
