package util

import (
	"log"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// TrimSpace 删除字符串两头空格字符
func TrimSpace(value string) string {
	return strings.Trim(value, " ")
}

// StringToDecimal Strings to Decimal (金额)
func StringToDecimal(value string) decimal.Decimal {
	val, err := decimal.NewFromString(strings.Trim(value, " "))
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// StringToFloat Strings to Float64
func StringToFloat(value string) float64 {
	val, err := strconv.ParseFloat(strings.Trim(value, " "), 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// FloatToString Float64 to Strings
func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
