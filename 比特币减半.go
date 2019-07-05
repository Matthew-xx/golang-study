// 已知比特币大约10分钟1个区块，每隔21W个区块挖矿奖励减半
// 问：比特币发行总量为多少？什么时间挖完？一直到挖完挖了几轮？
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

func main01() {
	var rewardCount float64
	rewardCount = 50
	var total float64
	count := 0
	blockInterval := 210000
	for rewardCount > 0 {
		currentCircleCount := float64(blockInterval) * float64(rewardCount)
		total += currentCircleCount
		count += 1
		log.Println("当前轮数：", count, "本轮奖励：", currentCircleCount)
		rewardCount *= 0.5
	}
	fmt.Printf("比特币发行总量:%f\n挖完时间：%d\n挖了%d轮", total, count*210000*10/60/24/365+2018, count)
}

// float32
// 比特币发行总量:21000000.000000
// 挖完时间：2641
// 挖了156轮

// float64
// 比特币发行总量:21000000.000000
// 挖完时间：6337
// 挖了1081轮

func main() {
	// 设置除不尽的情况，保留的小数位数，如果能除尽，则有几位保留几位
	decimal.DivisionPrecision = 8
	// 当轮的单块奖励
	rewardCount := decimal.NewFromFloat(50)
	// 总奖励
	total := decimal.NewFromFloat(0)
	// 轮数
	count := 0
	// 区块间隔
	blockInterval := decimal.NewFromFloat(210000)
	// 比特币诞生时间
	bornTime, _ := time.Parse("20060504", "20090103")
	// 由于是精确计算，这里将精度保持在小数点后8位
	for rewardCount.GreaterThanOrEqual(decimal.NewFromFloat(0.00000001)) {
		// 本轮开始
		// 本轮总奖励=区块间隔数*单区块奖励
		currentCircleCount := blockInterval.Mul(rewardCount)
		// 累加总奖励
		total = total.Add(currentCircleCount)
		// 轮数加1
		count += 1
		log.Println("轮数：", count, "单块奖励：", rewardCount, "本轮奖励：", currentCircleCount, "累计奖励：", total, "本轮理论结束时间：", bornTime.AddDate(0, 0, count*210000*10/60/24).Format("2006年05月04日"))
		// 本轮结束，下一轮奖励减半
		// 比特币使用了精度后的小数位舍去的方式，这里利用向下取整，实现同样效果。先扩大到精度倍数，在向下取整，再缩回精度。
		rewardCount = rewardCount.Mul(decimal.NewFromFloat(100000000)).Div(decimal.NewFromFloat(2)).Floor().Div(decimal.NewFromFloat(100000000))
		// 循环结束，判断是否开始下一轮
	}
	fmt.Print("比特币发行总量:", total, "\n挖完时间：", count*210000*10/60/24/365+2008, "\n挖了", count, "轮")
}
