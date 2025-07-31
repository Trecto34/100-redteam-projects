package main

import ("fmt")


func main(){
	alpha := "TEST"
	for _, letter := range alpha{
		Shift(30, letter)
	}
	fmt.Println()
}

func Shift(pace int, lett rune){	
	letter := byte(lett) + byte(pace)
	for letter > 90 {
		letter = letter - 26
	}
	fmt.Print(string(letter))
	return 
}
