package main

import (
	"fmt"
	"sort"
	"strings"
)

var choose int
var name string
var phone string

var (
	names  = []string{}
	phones = []string{}
)

//Run Program
func run() {
	for true {
		fmt.Println("Choose One:")
		fmt.Println("---------")
		fmt.Println("1:addContact")
		fmt.Println("2:showContact")
		fmt.Println("3:searchContact")
		fmt.Println("4:deleteContact")
		fmt.Println("5:Exit")
		fmt.Scanf("%d\n", &choose)
		switch choose {
		case 1:
			addC()
			break
		case 2:
			showC()
			break
		case 3:
			searchC()
		case 4:
			deleteC()
		case 5:
			return
		case 6:
			searchWord()

		}
	}
}

//Add Contact
func addC() {
	fmt.Println("name: ")
	fmt.Scanf("%s\n", &name)
	for _,v := range names{
		if v == name {
			fmt.Println("invalid")
			return
		}

	}
	names = append(names, name)
	fmt.Println("phone: ")
	fmt.Scanf("%s\n", &phone)
	phones = append(phones, phone)

}

//Show Contact
func showC() {
	sort.Strings(names)

	for i, v := range names {
		fmt.Println(v, phones[i])
	}

}

//Search Contact
func searchC() {
	fmt.Println("input name")
	fmt.Scanf("%s\n", &name)
	for i, v := range names {
		if name == v {
			fmt.Println(phones[i])
		}
	}
}
func searchWord()  {
	fmt.Println("input word")
	fmt.Scanf("%s\n", &name)
	for _,v :=range names {
		compare := strings.Contains(v, name)
			if compare ==true {
				fmt.Println(v)

			}
		}
}

//Delete Contact
func deleteC() {
	fmt.Println("which one is delete ?")
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Scanf("%s\n", &name)
	for i, v := range names {
		if v == name {
			names = append(names[:i], names[i+1:]...)
			phones = append(phones[:i], phones[i+1:]...)
		}

	}
}
