package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func get_bmr(height, weight, age float64) float64 {
	weight_float := 13.397
	height_float := 4.799
	age_float := 5.677
	f := 88.362
	bmr := (weight_float * weight) + (height_float * height) - (age_float * age) + f

	fmt.Printf("【あなたの1日基礎代謝量】\n %s\n", strconv.Itoa(int(bmr)))
	return bmr
}

func get_one_day_cal(bmr, active_level float64) float64 {
	one_day_cal := bmr * active_level

	fmt.Printf("【減量時の1日最適摂取カロリー】\n %s\n", strconv.Itoa(int(one_day_cal)))
	return one_day_cal
}

func get_pfc_balance(one_day_cal float64) {
	protein := strconv.Itoa(int(0.25 * one_day_cal))
	fat := strconv.Itoa(int(0.15 * one_day_cal))
	carbs := strconv.Itoa(int(0.6 * one_day_cal))

	fmt.Println("【1日のPFCバランス】")
	fmt.Println("タンパク質:", protein)
	fmt.Println("脂質:", fat)
	fmt.Println("炭水化物:", carbs)
}

var height, weight, age, active_level string

var cmd = &cobra.Command{
	Use: "app",
	Run: func(c *cobra.Command, args []string) {
		const bitSize = 64
		height, _ := strconv.ParseFloat(height, bitSize)
		weight, _ := strconv.ParseFloat(weight, bitSize)
		age, _ := strconv.ParseFloat(age, bitSize)
		active_level, _ := strconv.ParseFloat(active_level, bitSize)

		// 1日基礎代謝量
		bmr := get_bmr(height, weight, age)
		// 減量時の1日最適摂取カロリー
		one_day_cal := get_one_day_cal(bmr, active_level)
		// 1日のPFCバランス
		get_pfc_balance(one_day_cal)
	},
}

// ファイル読み込みのタイミングでフラグを定義する
func init() {
	// フラグの値を変数にセットする場合
	// 第1引数: 変数のポインタ
	// 第2引数: フラグ名
	// 第3引数: デフォルト値
	// 第4引数: 説明
	cmd.PersistentFlags().StringVar(&height, "height", "0", "現在の身長")
	cmd.PersistentFlags().StringVar(&weight, "weight", "0", "現在の体重")
	cmd.PersistentFlags().StringVar(&age, "age", "0", "現在の年齢")
	cmd.PersistentFlags().StringVar(&active_level, "active_level", "0", "日常の活動量")
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
