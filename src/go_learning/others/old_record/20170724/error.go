package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot Proceed,the divider is zero.
	dividee:%d
	divider:0
	`
	return fmt.Sprintf(strFormat, de.dividee)

}
func Divide(varDividee int, varDivider int) (result int, errorMSG string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMSG = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if result, errorMSG := Divide(100, 10); errorMSG == "" {
		fmt.Println("100/10=", result)
	}
	if _, errorMSG := Divide(100, 0); errorMSG != "" {
		fmt.Println("error message is:", errorMSG)
	}
}
