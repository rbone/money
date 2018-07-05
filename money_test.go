package money

import (
	"fmt"
	"strconv"
	"testing"
)

type moneyParseAmountForCurrencyr func(string, Currency) (int64, error)

func TestManualParsingWorks(t *testing.T) {
	testSuccess(t, "-123", -12300, ParseAmountForCurrency)
	testSuccess(t, "-123.00", -12300, ParseAmountForCurrency)
	testSuccess(t, "-123.12", -12312, ParseAmountForCurrency)
	testSuccess(t, "-123.02", -12302, ParseAmountForCurrency)
	testSuccess(t, "+321", 32100, ParseAmountForCurrency)
	testSuccess(t, "+321.00", 32100, ParseAmountForCurrency)
	testSuccess(t, "+321.12", 32112, ParseAmountForCurrency)
	testSuccess(t, "+321.02", 32102, ParseAmountForCurrency)
	testSuccess(t, "456", 45600, ParseAmountForCurrency)
	testSuccess(t, "456.00", 45600, ParseAmountForCurrency)
	testSuccess(t, "456.12", 45612, ParseAmountForCurrency)
	testSuccess(t, "456.02", 45602, ParseAmountForCurrency)
	testSuccess(t, "0.02", 2, ParseAmountForCurrency)
	testSuccess(t, "+0.02", 2, ParseAmountForCurrency)
	testSuccess(t, "-0.02", -2, ParseAmountForCurrency)

	testError(t, "0.0000", ParseAmountForCurrency)
	testError(t, "0.000", ParseAmountForCurrency)
	testError(t, "0.000", ParseAmountForCurrency)
	testError(t, "01.00", ParseAmountForCurrency)
	testError(t, "123.5", ParseAmountForCurrency)
	testError(t, "123,42", ParseAmountForCurrency)
	testError(t, "99.99.99", ParseAmountForCurrency)
	testError(t, ".121.1", ParseAmountForCurrency)
	testError(t, ",121,1", ParseAmountForCurrency)
	testError(t, "12 110", ParseAmountForCurrency)
	testError(t, "12,110", ParseAmountForCurrency)
	testError(t, "12,110.12", ParseAmountForCurrency)
	testError(t, ".1211", ParseAmountForCurrency)
	testError(t, "100.1211", ParseAmountForCurrency)
	testError(t, "some.value", ParseAmountForCurrency)
}

func testError(t *testing.T, input string, parseFunc moneyParseAmountForCurrencyr) {
	actualValue, err := parseFunc(input, USD)
	if err == nil {
		t.Error(fmt.Errorf("expected parsing %s to fail but it didn't, received value %d", input, actualValue))
	}
}

func testSuccess(t *testing.T, input string, expectedValue int64, parseFunc moneyParseAmountForCurrencyr) {
	actualValue, err := parseFunc(input, USD)
	if err != nil {
		t.Error(err)
	}
	if expectedValue != actualValue {
		t.Error(fmt.Errorf("failed converting %s, result was %d, expected %d", input, actualValue, expectedValue))
	}
}

func parseFloatForBenchComparison(amount string) (int64, error) {
	payoutAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}
	return int64(payoutAmount * 100), nil
}

func BenchmarkParsing(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseAmountForCurrency("512312.22", USD)
	}
}

func BenchmarkParseAmountForCurrencyFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseFloatForBenchComparison("512312.22")
	}
}
