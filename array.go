package main

import("fmt")

func main() {
	var weekWorkDays [5]string = [5]string{"mo", "di", "mi", "do", "fr"}
	weekendDays := [2]string{"sa", "so"}

	for _, day := range weekWorkDays {
		fmt.Println(day)
	}

	fmt.Printf("Number of weekend days: %d\n", len(weekendDays))
}
