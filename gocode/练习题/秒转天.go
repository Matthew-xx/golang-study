package main

import "fmt"

func main() {
	const givenSecond int = 107653
	const daySecond int = 24 * 60 * 60
	const hourSecond int = 60 * 60
	const minuteSecond int = 60
	day := givenSecond / daySecond
	hour := (givenSecond - day*daySecond) / hourSecond
	minute := (givenSecond - day*daySecond - hour*hourSecond) / minuteSecond
	second := (givenSecond - day*daySecond - hour*hourSecond) % minuteSecond
	fmt.Printf("一共是%d天%d时%d分%d秒\n", day, hour, minute, second)
}
