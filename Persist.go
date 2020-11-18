package main

import "fmt"

func main(){
    var num int
	
	fmt.Printf("Write a number: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil{
	fmt.Printf("Write a valid number")
	}else{
	pers := num
	count := 0
	
	    for pers >= 10{
	      accum := 0
		  for pers != 0{
		    accum = accum + pers %10
			pers = pers / 10
		  }
		  count = count + 1
		  pers = accum
	    }
	    fmt.Printf("The persistance of %d is: %d\n", num, count)
		fmt.Printf("The final sum is: %d", pers)
	}

}