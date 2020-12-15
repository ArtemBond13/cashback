// cashback.go выводит бонус в размере 1% от стоимости покупки
package main

func main() {
	purchase := 3333_33
	totalPercent := 100
	percent := 1
	limit := 100
	penToRub := 1_00

	bonus := (purchase / totalPercent * percent) / penToRub
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
