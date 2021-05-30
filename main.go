package main

//import "fmt"

func main()  {
run()
}


//
//var choose int
//var name string
//var phone string

//
//
//func contact()  {
//
//	names := []string{}
//	phones := []string{}
//
//	for true {
//
//		fmt.Println("Choose One:")
//		fmt.Println("1:addContact")
//		fmt.Println("2:showContact")
//		fmt.Println("3:searchContact")
//		fmt.Println("4:deleteContact")
//		fmt.Println("5:Exit")
//		fmt.Scanf("%d\n", &choose)
//		switch choose {
//		case 1:
//			fmt.Println("name: ")
//			fmt.Scanf("%s\n", &name)
//			names = append(names, name)
//			fmt.Println("phone: ")
//			fmt.Scanf("%s\n", &phone)
//			phones = append(phones, phone)
//			break
//
//		case 2:
//			for _, name := range names {
//				fmt.Println(name)
//			}
//			break
//		case 3:
//			fmt.Println("input name")
//			fmt.Scanf("%s\n" , &name)
//			for i, v := range names{
//				if name == v {
//					fmt.Println(phones[i])
//				}
//			}
//			break
//		case 4:
//			fmt.Println("which one is delete ?")
//			for _, name := range names {
//				fmt.Println(name)
//			}
//			fmt.Scanf("%s\n" , &name)
//
//			for i,v := range names {
//				if v == name{
//					names=append(names[ :i] , names[i+1:]...)
//					phones = append(phones[:i] , phones[i+1:]...)
//			}
//
//		}
//			break
//
//		case 5:
//			return
//		}
//
//	}
//}
