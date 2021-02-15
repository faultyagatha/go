package main


func arrays() {
	//initialised to 0 values
	var arr [5]int
	arr[0] = 2

}

//array literal
func arrInteration() {
	var x[5] int = [5] int{1, 2, 3, 4, 5}
	for i, v range x {
		fmt.Printf("ind %d, val %d", i, v)
	}
}


