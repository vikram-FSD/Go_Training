package main

func find_oddOrEven_PrimeorNot(num int) (int, string, string) {
	count := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}
	if (count > 2) && (num%2 == 0) {
		return num, "Not a Prime", "Even_Number"
	}
	return num, "Prime_Number", "Odd_Number"
}
