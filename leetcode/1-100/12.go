package main

import "fmt"

func intToRoman(num int) string {
	res := ""
	//FIXME: range map will get random result
	//valueMap := map[int]string{
	//	1000: "M",
	//	900:  "CM",
	//	500:  "D",
	//	400:  "CD",
	//	100:  "C",
	//	90:   "XC",
	//	50:   "L",
	//	40:   "XL",
	//	10:   "X",
	//	9:    "IX",
	//	5:    "V",
	//	4:    "IV",
	//	1:    "I"}

	//for intValue, romanValue := range valueMap {
	//	for num >= intValue {
	//		res = res + romanValue
	//		num = num - intValue
	//	}
	//}
	//return res
	intValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanValues := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(intValues); i++ {
		for num >= intValues[i] {
			res = res + romanValues[i]
			num = num - intValues[i]
		}
	}
	return res

}

func main() {
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
