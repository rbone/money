package money

import "fmt"

const (
	stateFirstChar        int8 = 1
	statePreDecimal       int8 = 2
	statePreDecimalGroup3 int8 = 3
	statePreDecimalGroup2 int8 = 4
	statePreDecimalGroup1 int8 = 5
	statePostDecimal      int8 = 6
	stateExpectingDecimal int8 = 7
)

type Parser interface {
	Parse(amount string) (int64, error)
}

type moneyParser struct {
	decimalIndicator  rune
	groupingIndicator rune
	minorUnits        int64
}

var usdMoneyParser = moneyParser{decimalIndicator: '.', groupingIndicator: ',', minorUnits: 2}

func (p moneyParser) Parse(amount string) (int64, error) {
	if len(amount) == 0 {
		return 0, fmt.Errorf("expected int string, got empty string")
	}
	var value int64 = 0

	state := stateFirstChar

	for pos, char := range amount {
		switch state {
		case stateFirstChar:
			switch char {
			case '+':
			case '-':
			case '0':
				value = 0
				state = stateExpectingDecimal
			case '1':
				value = 1
				state = statePreDecimal
			case '2':
				value = 2
				state = statePreDecimal
			case '3':
				value = 3
				state = statePreDecimal
			case '4':
				value = 4
				state = statePreDecimal
			case '5':
				value = 5
				state = statePreDecimal
			case '6':
				value = 6
				state = statePreDecimal
			case '7':
				value = 7
				state = statePreDecimal
			case '8':
				value = 8
				state = statePreDecimal
			case '9':
				value = 9
				state = statePreDecimal
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d", char, pos)
			}
		case stateExpectingDecimal:
			switch char {
			case p.decimalIndicator:
				if int64(len(amount)-pos) != p.minorUnits+1 {
					return 0, fmt.Errorf("insufficient characters after %q", p.decimalIndicator)
				} else {
					state = statePostDecimal
				}
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d, expecting %q", char, pos, p.decimalIndicator)
			}
		case statePreDecimal:
			switch char {
			case p.decimalIndicator:
				if int64(len(amount)-pos) != p.minorUnits+1 {
					return 0, fmt.Errorf("insufficient characters after %q", p.decimalIndicator)
				} else {
					state = statePostDecimal
				}
			case p.groupingIndicator:
				if len(amount)-pos < 4 {
					return 0, fmt.Errorf("insufficient characters after %q", p.groupingIndicator)
				}
				state = statePreDecimalGroup3
			case '0':
				value = value * 10
			case '1':
				value = value*10 + 1
			case '2':
				value = value*10 + 2
			case '3':
				value = value*10 + 3
			case '4':
				value = value*10 + 4
			case '5':
				value = value*10 + 5
			case '6':
				value = value*10 + 6
			case '7':
				value = value*10 + 7
			case '8':
				value = value*10 + 8
			case '9':
				value = value*10 + 9
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d", char, pos)
			}
		case statePreDecimalGroup3:
			switch char {
			case '0':
				value = value * 10
			case '1':
				value = value*10 + 1
			case '2':
				value = value*10 + 2
			case '3':
				value = value*10 + 3
			case '4':
				value = value*10 + 4
			case '5':
				value = value*10 + 5
			case '6':
				value = value*10 + 6
			case '7':
				value = value*10 + 7
			case '8':
				value = value*10 + 8
			case '9':
				value = value*10 + 9
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d", char, pos)
			}
			state = statePreDecimalGroup2
		case statePreDecimalGroup2:
			switch char {
			case '0':
				value = value * 10
			case '1':
				value = value*10 + 1
			case '2':
				value = value*10 + 2
			case '3':
				value = value*10 + 3
			case '4':
				value = value*10 + 4
			case '5':
				value = value*10 + 5
			case '6':
				value = value*10 + 6
			case '7':
				value = value*10 + 7
			case '8':
				value = value*10 + 8
			case '9':
				value = value*10 + 9
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d", char, pos)
			}
			state = statePreDecimalGroup1
		case statePreDecimalGroup1:
			switch char {
			case '0':
				value = value * 10
			case '1':
				value = value*10 + 1
			case '2':
				value = value*10 + 2
			case '3':
				value = value*10 + 3
			case '4':
				value = value*10 + 4
			case '5':
				value = value*10 + 5
			case '6':
				value = value*10 + 6
			case '7':
				value = value*10 + 7
			case '8':
				value = value*10 + 8
			case '9':
				value = value*10 + 9
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d", char, pos)
			}
			state = statePreDecimal
		case statePostDecimal:
			switch char {
			case '0':
				value = value * 10
			case '1':
				value = value*10 + 1
			case '2':
				value = value*10 + 2
			case '3':
				value = value*10 + 3
			case '4':
				value = value*10 + 4
			case '5':
				value = value*10 + 5
			case '6':
				value = value*10 + 6
			case '7':
				value = value*10 + 7
			case '8':
				value = value*10 + 8
			case '9':
				value = value*10 + 9
			default:
				return 0, fmt.Errorf("unexpected character %q at position %d", char, pos)
			}
		default:
			panic(fmt.Errorf("we should never get here!"))
		}
	}

	switch state {
	case statePreDecimalGroup3, statePreDecimalGroup2, statePreDecimalGroup1:
		return 0, fmt.Errorf("insufficient digits after %q", p.groupingIndicator)
	case statePreDecimal:
		for i := int64(0); i < p.minorUnits; i++ {
			value = value * 10
		}
	}

	if amount[0] == '-' {
		value = value * -1
	}
	return value, nil
}

func ParseAmountForCurrency(amount string, currency *Currency) (int64, error) {
	if currency != USD {
		return 0, fmt.Errorf("only USD is supported at the moment")
	}

	return usdMoneyParser.Parse(amount)
}
