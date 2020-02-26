package main

import "fmt"

func main() {
	typeCountItems := map[string][]string{}
	typeCountItems["uagent"] = append(typeCountItems["uagent"], "server-1 : uagent-1\n")
	// typeCountItems["uagent"] = append(typeCountItems["uagent"], "server-2 : uagent-2\n")
	// typeCountItems["ustats"] = append(typeCountItems["uagent"], "server-1 : ustats-1\n")
	// typeCountItems["ustats"] = append(typeCountItems["uagent"], "server-2 : uagent-2\n")

	fmt.Println(len(typeCountItems["uagent"]))
	fmt.Printf("Uagent Server is %v\n", typeCountItems["uagent"])

	typeCountItems["uagent"] = typeCountItems["uagent"][1:]
	fmt.Println(len(typeCountItems["uagent"]))
	fmt.Printf("Uagent Server is %v\n", typeCountItems["uagent"])

	// 	typeCountItems["uagent"] = append(typeCountItems["uagent"], typeCountItems["uagent"][0])
	// 	fmt.Println(len(typeCountItems["uagent"]))
	// 	fmt.Printf("Uagent Server is %v\n", typeCountItems["uagent"])
}
