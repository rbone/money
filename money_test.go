package money

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	testUSDMoneyParser = moneyParser{decimalIndicator: '.', groupingIndicator: ',', minorUnits: 2}
	testEURMoneyParser = moneyParser{decimalIndicator: ',', groupingIndicator: '.', minorUnits: 2}
	testJPYMoneyParser = moneyParser{decimalIndicator: '.', groupingIndicator: ',', minorUnits: 0}
	testBHDMoneyParser = moneyParser{decimalIndicator: '.', groupingIndicator: ',', minorUnits: 3}
)

func TestManualParsingWorks(t *testing.T) {
	// USD Parsing
	testSuccess(t, testUSDMoneyParser, "-123", -12300)
	testSuccess(t, testUSDMoneyParser, "-123.00", -12300)
	testSuccess(t, testUSDMoneyParser, "-123.12", -12312)
	testSuccess(t, testUSDMoneyParser, "-123.02", -12302)
	testSuccess(t, testUSDMoneyParser, "+321", 32100)
	testSuccess(t, testUSDMoneyParser, "+321.00", 32100)
	testSuccess(t, testUSDMoneyParser, "+321.12", 32112)
	testSuccess(t, testUSDMoneyParser, "+321.02", 32102)
	testSuccess(t, testUSDMoneyParser, "456", 45600)
	testSuccess(t, testUSDMoneyParser, "456.00", 45600)
	testSuccess(t, testUSDMoneyParser, "456.12", 45612)
	testSuccess(t, testUSDMoneyParser, "456.02", 45602)
	testSuccess(t, testUSDMoneyParser, "0.02", 2)
	testSuccess(t, testUSDMoneyParser, "+0.02", 2)
	testSuccess(t, testUSDMoneyParser, "-0.02", -2)
	testSuccess(t, testUSDMoneyParser, "-1,123", -112300)
	testSuccess(t, testUSDMoneyParser, "-1,123.00", -112300)
	testSuccess(t, testUSDMoneyParser, "-1,123.12", -112312)
	testSuccess(t, testUSDMoneyParser, "-1,123.02", -112302)
	testSuccess(t, testUSDMoneyParser, "+1,321", 132100)
	testSuccess(t, testUSDMoneyParser, "+1,321.00", 132100)
	testSuccess(t, testUSDMoneyParser, "+1,321.12", 132112)
	testSuccess(t, testUSDMoneyParser, "+1,321.02", 132102)
	testSuccess(t, testUSDMoneyParser, "1,456", 145600)
	testSuccess(t, testUSDMoneyParser, "1,456.00", 145600)
	testSuccess(t, testUSDMoneyParser, "1,456.12", 145612)
	testSuccess(t, testUSDMoneyParser, "1,456.02", 145602)
	testSuccess(t, testUSDMoneyParser, "0.02", 2)
	testSuccess(t, testUSDMoneyParser, "+0.02", 2)
	testSuccess(t, testUSDMoneyParser, "-0.02", -2)

	testError(t, testUSDMoneyParser, "00.2")
	testError(t, testUSDMoneyParser, "0.0000")
	testError(t, testUSDMoneyParser, "0.000")
	testError(t, testUSDMoneyParser, "0.000")
	testError(t, testUSDMoneyParser, "01.00")
	testError(t, testUSDMoneyParser, "123.5")
	testError(t, testUSDMoneyParser, "123,42")
	testError(t, testUSDMoneyParser, "99.99.99")
	testError(t, testUSDMoneyParser, ".121.1")
	testError(t, testUSDMoneyParser, ",121,1")
	testError(t, testUSDMoneyParser, "12 110")
	testError(t, testUSDMoneyParser, ".1211")
	testError(t, testUSDMoneyParser, "100.1211")
	testError(t, testUSDMoneyParser, "some.value")
	testError(t, testUSDMoneyParser, "1,,456")
	testError(t, testUSDMoneyParser, "1,4,56")
	testError(t, testUSDMoneyParser, ",1456")
	testError(t, testUSDMoneyParser, "14,56")
	testError(t, testUSDMoneyParser, "145,6")

	// EUR parsing
	testSuccess(t, testEURMoneyParser, "-123", -12300)
	testSuccess(t, testEURMoneyParser, "-123,00", -12300)
	testSuccess(t, testEURMoneyParser, "-123,12", -12312)
	testSuccess(t, testEURMoneyParser, "-123,02", -12302)
	testSuccess(t, testEURMoneyParser, "+321", 32100)
	testSuccess(t, testEURMoneyParser, "+321,00", 32100)
	testSuccess(t, testEURMoneyParser, "+321,12", 32112)
	testSuccess(t, testEURMoneyParser, "+321,02", 32102)
	testSuccess(t, testEURMoneyParser, "456", 45600)
	testSuccess(t, testEURMoneyParser, "456,00", 45600)
	testSuccess(t, testEURMoneyParser, "456,12", 45612)
	testSuccess(t, testEURMoneyParser, "456,02", 45602)
	testSuccess(t, testEURMoneyParser, "0,02", 2)
	testSuccess(t, testEURMoneyParser, "+0,02", 2)
	testSuccess(t, testEURMoneyParser, "-0,02", -2)
	testSuccess(t, testEURMoneyParser, "-1.123", -112300)
	testSuccess(t, testEURMoneyParser, "-1.123,00", -112300)
	testSuccess(t, testEURMoneyParser, "-1.123,12", -112312)
	testSuccess(t, testEURMoneyParser, "-1.123,02", -112302)
	testSuccess(t, testEURMoneyParser, "+1.321", 132100)
	testSuccess(t, testEURMoneyParser, "+1.321,00", 132100)
	testSuccess(t, testEURMoneyParser, "+1.321,12", 132112)
	testSuccess(t, testEURMoneyParser, "+1.321,02", 132102)
	testSuccess(t, testEURMoneyParser, "1.456", 145600)
	testSuccess(t, testEURMoneyParser, "1.456,00", 145600)
	testSuccess(t, testEURMoneyParser, "1.456,12", 145612)
	testSuccess(t, testEURMoneyParser, "1.456,02", 145602)
	testSuccess(t, testEURMoneyParser, "0,02", 2)
	testSuccess(t, testEURMoneyParser, "+0,02", 2)
	testSuccess(t, testEURMoneyParser, "-0,02", -2)

	testError(t, testEURMoneyParser, "00,2")
	testError(t, testEURMoneyParser, "0,0000")
	testError(t, testEURMoneyParser, "0,000")
	testError(t, testEURMoneyParser, "0,000")
	testError(t, testEURMoneyParser, "01,00")
	testError(t, testEURMoneyParser, "123,5")
	testError(t, testEURMoneyParser, "123.42")
	testError(t, testEURMoneyParser, "99,99,99")
	testError(t, testEURMoneyParser, ",121,1")
	testError(t, testEURMoneyParser, ".121.1")
	testError(t, testEURMoneyParser, "12 110")
	testError(t, testEURMoneyParser, ",1211")
	testError(t, testEURMoneyParser, "100,1211")
	testError(t, testEURMoneyParser, "some,value")
	testError(t, testEURMoneyParser, "1..456")
	testError(t, testEURMoneyParser, "1.4.56")
	testError(t, testEURMoneyParser, ".1456")
	testError(t, testEURMoneyParser, "14.56")
	testError(t, testEURMoneyParser, "145.6")

	// JPY money parsing
	testSuccess(t, testJPYMoneyParser, "-123", -123)
	testSuccess(t, testJPYMoneyParser, "+321", 321)
	testSuccess(t, testJPYMoneyParser, "456", 456)
	testSuccess(t, testJPYMoneyParser, "-1,123", -1123)
	testSuccess(t, testJPYMoneyParser, "+1,321", 1321)
	testSuccess(t, testJPYMoneyParser, "1,456", 1456)

	testError(t, testJPYMoneyParser, "1,456.00")
	testError(t, testJPYMoneyParser, "1,456.12")
	testError(t, testJPYMoneyParser, "1,456.02")
	testError(t, testJPYMoneyParser, "0.02")
	testError(t, testJPYMoneyParser, "+0.02")
	testError(t, testJPYMoneyParser, "-0.02")
	testError(t, testJPYMoneyParser, "+1,321.00")
	testError(t, testJPYMoneyParser, "+1,321.12")
	testError(t, testJPYMoneyParser, "+1,321.02")
	testError(t, testJPYMoneyParser, "-1,123.00")
	testError(t, testJPYMoneyParser, "-1,123.12")
	testError(t, testJPYMoneyParser, "-1,123.02")
	testError(t, testJPYMoneyParser, "456.00")
	testError(t, testJPYMoneyParser, "456.12")
	testError(t, testJPYMoneyParser, "456.02")
	testError(t, testJPYMoneyParser, "0.02")
	testError(t, testJPYMoneyParser, "+0.02")
	testError(t, testJPYMoneyParser, "-0.02")
	testError(t, testJPYMoneyParser, "+321.00")
	testError(t, testJPYMoneyParser, "+321.12")
	testError(t, testJPYMoneyParser, "+321.02")
	testError(t, testJPYMoneyParser, "-123.00")
	testError(t, testJPYMoneyParser, "-123.12")
	testError(t, testJPYMoneyParser, "-123.02")
	testError(t, testJPYMoneyParser, "00.2")
	testError(t, testJPYMoneyParser, "0.0000")
	testError(t, testJPYMoneyParser, "0.000")
	testError(t, testJPYMoneyParser, "0.000")
	testError(t, testJPYMoneyParser, "01.00")
	testError(t, testJPYMoneyParser, "123.5")
	testError(t, testJPYMoneyParser, "123,42")
	testError(t, testJPYMoneyParser, "99.99.99")
	testError(t, testJPYMoneyParser, ".121.1")
	testError(t, testJPYMoneyParser, ",121,1")
	testError(t, testJPYMoneyParser, "12 110")
	testError(t, testJPYMoneyParser, ".1211")
	testError(t, testJPYMoneyParser, "100.1211")
	testError(t, testJPYMoneyParser, "some.value")
	testError(t, testJPYMoneyParser, "1,,456")
	testError(t, testJPYMoneyParser, "1,4,56")
	testError(t, testJPYMoneyParser, ",1456")
	testError(t, testJPYMoneyParser, "14,56")
	testError(t, testJPYMoneyParser, "145,6")

	// BHD Parsing
	testSuccess(t, testBHDMoneyParser, "-123", -123000)
	testSuccess(t, testBHDMoneyParser, "-123.000", -123000)
	testSuccess(t, testBHDMoneyParser, "-123.312", -123312)
	testSuccess(t, testBHDMoneyParser, "-123.002", -123002)
	testSuccess(t, testBHDMoneyParser, "+321", 321000)
	testSuccess(t, testBHDMoneyParser, "+321.000", 321000)
	testSuccess(t, testBHDMoneyParser, "+321.312", 321312)
	testSuccess(t, testBHDMoneyParser, "+321.002", 321002)
	testSuccess(t, testBHDMoneyParser, "456", 456000)
	testSuccess(t, testBHDMoneyParser, "456.000", 456000)
	testSuccess(t, testBHDMoneyParser, "456.312", 456312)
	testSuccess(t, testBHDMoneyParser, "456.002", 456002)
	testSuccess(t, testBHDMoneyParser, "0.002", 2)
	testSuccess(t, testBHDMoneyParser, "+0.002", 2)
	testSuccess(t, testBHDMoneyParser, "-0.002", -2)
	testSuccess(t, testBHDMoneyParser, "-1,123", -1123000)
	testSuccess(t, testBHDMoneyParser, "-1,123.000", -1123000)
	testSuccess(t, testBHDMoneyParser, "-1,123.312", -1123312)
	testSuccess(t, testBHDMoneyParser, "-1,123.002", -1123002)
	testSuccess(t, testBHDMoneyParser, "+1,321", 1321000)
	testSuccess(t, testBHDMoneyParser, "+1,321.000", 1321000)
	testSuccess(t, testBHDMoneyParser, "+1,321.312", 1321312)
	testSuccess(t, testBHDMoneyParser, "+1,321.002", 1321002)
	testSuccess(t, testBHDMoneyParser, "1,456", 1456000)
	testSuccess(t, testBHDMoneyParser, "1,456.000", 1456000)
	testSuccess(t, testBHDMoneyParser, "1,456.312", 1456312)
	testSuccess(t, testBHDMoneyParser, "1,456.002", 1456002)
	testSuccess(t, testBHDMoneyParser, "0.002", 2)
	testSuccess(t, testBHDMoneyParser, "+0.002", 2)
	testSuccess(t, testBHDMoneyParser, "-0.002", -2)

	testError(t, testBHDMoneyParser, "00.2")
	testError(t, testBHDMoneyParser, "0.0000")
	testError(t, testBHDMoneyParser, "0.00")
	testError(t, testBHDMoneyParser, "01.000")
	testError(t, testBHDMoneyParser, "123.5")
	testError(t, testBHDMoneyParser, "123,42")
	testError(t, testBHDMoneyParser, "99.99.999")
	testError(t, testBHDMoneyParser, ".121.1")
	testError(t, testBHDMoneyParser, ",121,1")
	testError(t, testBHDMoneyParser, "12 110")
	testError(t, testBHDMoneyParser, ".1211")
	testError(t, testBHDMoneyParser, "100.1211")
	testError(t, testBHDMoneyParser, "some.value")
	testError(t, testBHDMoneyParser, "1,,456")
	testError(t, testBHDMoneyParser, "1,4,56")
	testError(t, testBHDMoneyParser, ",1456")
	testError(t, testBHDMoneyParser, "14,56")
	testError(t, testBHDMoneyParser, "145,6")

}

func testError(t *testing.T, parser Parser, input string) {
	t.Helper()
	actualValue, err := parser.Parse(input)
	if err == nil {
		t.Error(fmt.Errorf("expected parsing %s to fail but it didn't, received value %d", input, actualValue))
	}
}

func testSuccess(t *testing.T, parser Parser, input string, expectedValue int64) {
	t.Helper()
	actualValue, err := parser.Parse(input)
	if err != nil {
		t.Error(fmt.Errorf("failed converting %q received error : %s", input, err.Error()))
	} else if expectedValue != actualValue {
		t.Error(fmt.Errorf("failed converting %q result was %d, expected %d", input, actualValue, expectedValue))
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
