package main

func main() {
	//loop

	/*
		i := 0
		for i < 3 {
			if i == 1 {
				continue
			}
			fmt.Println(i)
			break
		}
	*/
	/*
		nums := []int{1, 3, 2, 4, 0}
		for index, value := range nums {
			fmt.Println(index)
			fmt.Println(value)
		}
	*/
	/*
		nums := []int{1, 3, 2, 4, 0}
		for _, value := range nums {
			fmt.Println(value)
		}
	*/
	/*
		nums := []int{1, 3, 2, 4, 0}
		for _, value := range nums {
			if value == 3 {
				continue //or break //
			}
			fmt.Println(value)
		}
	*/

	/*
		nums := []string{"Divyanhsu", "Singh"}
		for _, value := range nums {
			// if value == "u" {
			// 	break
			// }
			fmt.Println(value)
		}
	*/

	//function

	/*
	   	w := new(int)
	   	//name := new(string)
	   	t := "hjbjk"
	   	result, x := c(*w, &t)
	   	fmt.Println(result)
	   	fmt.Println(x)
	   	//fmt.Println(*name)
	   	//fmt.Println(*w)
	   	r, _ := b(1, 2, 3, 4, 5, 6, 6)
	   	fmt.Print(r)
	   }
	   func a() (int, string) {
	   	return 122, "weqrr3"
	   }
	   func b(args ...int) (bool, int) {
	   	for _, v := range args {
	   		fmt.Println(v)
	   	}
	   	return true, 11
	   }
	   func c(w int, name *string) (i int, j string) {
	   	i = 10
	   	j = "rahul"
	   	w = 100
	   	*name = "code"
	   	log.Print("j")
	   	return
	   }
	*/

	// Pointer

	/*

		//i := 10
		//j := &i
		//*j = 100

		var i int
		i = 10
		//var j *int    //declaration..j is empyty
		//j := new(int) //declaration + assignment(j is not empty)
		j = &i
		*j = 100

		fmt.Println(i)
		fmt.Println(j)
		fmt.Println(*j)

		name := new(string)
		*name = "golang"
		fmt.Println(*name)

	*/

	//CLOUSER

	/*
	   	number := 10
	   	var number int
	   	number = 10
	   	fmt.Println(number)

	   	var getInt func(int) int
	   	getInt = func(x int) int {
	   		fmt.Println("In a 1st function")
	   		return 20 + x
	   	}
	   	g(getInt, 8)
	   	getInt = func(x int) int {
	   		fmt.Println("In a 2nd function")
	   		return 10 + x
	   	}
	   	var y func()
	   	y = func() {
	   		fmt.Print(9)
	   	}
	   	y()

	   	j := func(x int) int { //nameless function
	   		fmt.Println("In a 2st function")
	   		return 20 + x
	   	}(10)
	   	fmt.Println(j)
	   }

	   func g(getInt func(int) int, u int) {
	   	j := getInt(78)
	   	fmt.Print(j)
	   }

	*/

}
