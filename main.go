package main

func main() {
	var balance uint32 = 15_000_000_00
	var transfer uint32 = 10_000_000_00
	total := uint64(balance) + uint64(transfer)
	println(total)
}
