# Money

A parser for converting string monetary values into integer values.

For example

```golang
amount, err := money.ParseAmountForCurrency("432.12", money.USD)
if err != nil {
  panic(err)
}
fmt.Printf("Value as an integer is: %d", amount)
// Prints: "Value as an integer is: 43212"
```

Currently the only supported currency is USD.

## Possible Future Improvements

  - Supporting currencies other than USD.
  - Supporting non-english locales.
  - Supporting formatting monetary values as well.

## Performance

For now parsing is fairly quick, and beats naive alternatives like `strconv.ParseFloat`,
however that may change as more currencies & locales are supported and complexity
increases.

```
[~/go/src/github.com/rbone/money]: go test -bench=. --benchmem
goos: darwin
goarch: amd64
pkg: github.com/rbone/money
BenchmarkParsing-4                       	50000000	        33.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkParseAmountForCurrencyFloat-4   	30000000	        39.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rbone/money	3.936s
```
