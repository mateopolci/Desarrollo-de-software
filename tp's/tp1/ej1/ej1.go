package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var option int
	for {
		fmt.Println("----------------------------")
		fmt.Println("1: Calculator")
		fmt.Println("2: Caesar Cipher")
		fmt.Println("----------------------------")
		fmt.Print("Select an option: ")
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			math()
		case 2:
			cipher()
		}
	}
}

func math() {
	var (
		option int
		num1, num2, res float64
		input string
	)
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("----------------------------")
	fmt.Println("1: Adding")
	fmt.Println("2: Subtracting")
	fmt.Println("3: Multiplying")
	fmt.Println("4: Dividing")
	fmt.Println("----------------------------")

	//Get calculator option
	fmt.Print("Select an option: ")
	scanner.Scan()
	input = scanner.Text()
	option, _ = strconv.Atoi(input)
	
	//Get first number
	fmt.Print("Enter the first number: ")
	scanner.Scan()
	input = scanner.Text()
	num1, _ = strconv.ParseFloat(input, 64)
	
	//Get second number
	fmt.Print("Enter the second number: ")
	scanner.Scan()
	input = scanner.Text()
	num2, _ = strconv.ParseFloat(input, 64)
	
	switch option {
	case 1:
		fmt.Print("The sum is: ")
		res=num1+num2
	case 2:
		fmt.Print("The difference is: ")
		res=num1-num2
	case 3:
		fmt.Print("The product is: ")
		res=num1*num2
	case 4:
		fmt.Print("The quotient is: ")
		res=num1/num2
	}
	fmt.Println(res)
}

func cipher(){
	var(
		option, key int
		input, phrase string
		alphabet[]rune
		shiftedAlphabet[]rune
	)
	scanner := bufio.NewScanner(os.Stdin)
	
	//Creates the original alphabet
	for i := 'a'; i <= 'z' ; i++ {
		alphabet = append(alphabet, i)
	}

	fmt.Println("Welcome to the Caesar Cipher module:")
	fmt.Println("----------------------------")
	fmt.Println("1: Cipher")
	fmt.Println("2: Decipher")
	fmt.Println("----------------------------")
	
	//Get cipher option
	fmt.Print("Select an option: ")
	scanner.Scan()
	input = scanner.Text()
	option, _ = strconv.Atoi(input)
	
	switch option {
	case 1:
		//Get key
		fmt.Print("Enter a cipher key [0-26]: ")
		scanner.Scan()
		input = scanner.Text()
		key, _ = strconv.Atoi(input)
		
		//Creates shifted alphabet based on key
		for i := 'a'; i <= 'z' ; i++ {
			shiftedLetter := alphabet[(int(i)-'a'+key)%26]
			shiftedAlphabet = append(shiftedAlphabet, shiftedLetter)
		}
		
		//Get phrase
		fmt.Printf("Enter a phrase to cipher with %d\n", key)
		scanner.Scan()
		phrase = scanner.Text()
		
		//Return ciphered phrase
		//
		var result strings.Builder
		for _, char := range phrase {
			if char >= 'a' && char <= 'z' {
				result.WriteRune(shiftedAlphabet[char-'a'])
				} else if char >= 'A' && char <= 'Z' {
					lowerChar := char + ('a' - 'A')
					result.WriteRune(shiftedAlphabet[lowerChar-'a'] - ('a' - 'A'))
					} else {
						result.WriteRune(char)
					}
				}
				fmt.Printf("Ciphered phrase: %s\n", result.String())
		//
	case 2:
		//Get key
		fmt.Print("Enter a decipher key [0-26]: ")
		scanner.Scan()
		input = scanner.Text()
		key, _ = strconv.Atoi(input)
		
		//Creates shifted alphabet based on key
		for i := 'a'; i <= 'z' ; i++ {
			shiftedLetter := alphabet[(int(i)-'a'+key)%26]
			shiftedAlphabet = append(shiftedAlphabet, shiftedLetter)
		}
		
		//Get phrase
		fmt.Printf("Enter a phrase to cipher with %d\n", key)
		scanner.Scan()
		phrase = scanner.Text()
		
		//Return deciphered phrase
		//
		var result strings.Builder
        reverseAlphabetMap := make(map[rune]rune)
        for i, char := range shiftedAlphabet {
			reverseAlphabetMap[char] = alphabet[i]
        }
        for _, char := range phrase {
			if char >= 'a' && char <= 'z' {
				result.WriteRune(reverseAlphabetMap[char])
				} else if char >= 'A' && char <= 'Z' {
					lowerChar := char + ('a' - 'A')
					result.WriteRune(reverseAlphabetMap[lowerChar] - ('a' - 'A'))
					} else {
						result.WriteRune(char)
					}
				}
				fmt.Printf("Deciphered phrase: %s\n", result.String())
		//
	}
}