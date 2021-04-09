package main

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"math"
)

var marks = []string{"btcusdt", "adausdt", "ethusdt"}

func GetPrice(client *client.MarketClient, s string) {
	resp, err := client.GetLatestTrade(s)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, trade := range resp.Data {
			applogger.Info("name=%v, Price=%v", s, trade.Price)
		}
	}
}

func Calculate(dataSource []float64) {
	var sum float64 = 0
	for _, v := range dataSource {
		sum += v
	}
	μ := float64(sum) / float64(len(dataSource))

	//标准差
	var variance float64
	for _, v := range dataSource {
		variance += math.Pow((v - μ), 2)
	}
	σ := math.Sqrt(variance / float64(len(dataSource)))
	fmt.Println("σ:", σ)
	fmt.Println("μ:", μ)

	//正态分布公式
	a := 1 / (math.Sqrt(2*math.Pi) * σ) * math.Pow(math.E, (-math.Pow((μ-μ), 2)/(2*math.Pow(σ, 2))))
	fmt.Println(a)
}

func main() {
	cli := new(client.MarketClient).Init(config.Host)
	for _, mark := range marks {
		GetPrice(cli, mark)
	}
	dataSource := []float64{1, 2, 3}

	//计算正态分布
	Calculate(dataSource)

}
