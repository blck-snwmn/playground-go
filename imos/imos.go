package main

import (
	"fmt"
	"sort"
	"time"
)

type Period struct {
	start time.Time
	end   time.Time
}

func main() {
	periods := []Period{
		{start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), end: time.Date(2023, 3, 31, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC), end: time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), end: time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2023, 7, 10, 0, 0, 0, 0, time.UTC), end: time.Date(2023, 9, 20, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC), end: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC), end: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC)},
		{start: time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC)},
	}

	// 座標圧縮
	dates := make([]time.Time, 0)
	for _, period := range periods {
		dates = append(dates, period.start, period.end.AddDate(0, 0, 1))
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	dateToIndex := make(map[time.Time]int)
	for i, date := range dates {
		dateToIndex[date] = i
	}

	// imos法
	imos := make([]int, len(dates))
	for _, period := range periods {
		start := dateToIndex[period.start]
		end := dateToIndex[period.end.AddDate(0, 0, 1)]
		imos[start]++
		imos[end]--
	}

	// 累積和の計算
	for i := 1; i < len(imos); i++ {
		imos[i] += imos[i-1]
	}

	for i := 0; i < len(imos); i++ {
		fmt.Printf("%s: %d\n", dates[i].Format("2006-01-02"), imos[i])
	}
	return

	// 重複判定
	maxOverlap := 0
	for _, count := range imos {
		if count > maxOverlap {
			maxOverlap = count
		}
	}

	if maxOverlap > 1 {
		fmt.Println("期間が重複しています")
	} else {
		fmt.Println("期間は重複していません")
	}

	// 期間に含まれない日の判定
	for i, count := range imos {
		if count == 0 {
			fmt.Printf("%s は期間に含まれていません\n", dates[i].Format("2006-01-02"))
		} else {
			fmt.Printf("%s は期間に含まれています\n", dates[i].Format("2006-01-02"))
		}
	}
}
