package main

func main() {
	var balance int32 = 15_000_000_00
	var transfer int32 = 10_000_000_00
	total := int64(balance) + int64(transfer)
	println(total)
}