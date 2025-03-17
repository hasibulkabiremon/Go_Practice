package main

func fizzbuzz() {

	// ?
	
	for x :=1 ; x < 101 ; x ++{
		var linestr string
		if x % 3 == 0{
			linestr = "fizz"
		}

		if x % 5 == 0 {
			linestr += "bazz"
		}

		if x % 3 == 0 || x % 5 == 0 {
			println(linestr)
		}else{
			println(x)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
