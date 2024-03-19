package wallet

import "errors"

var (
	ErrDenominationCantBeNegative = errors.New("Negative Denomination.")
	ErrDenominationCantBeZero     = errors.New("Zero Denomination.")
	ErrGreaterThanCurrentBalance  = errors.New("Greater than current balance.")
)

func NewRupee(value float64) (*Rupee, error) {
	if value > 0 {
		return &Rupee{value}, nil
	}
	return nil, ErrDenominationCantBeNegative
}

func NewDollar(value int) (*Dollar, error) {
	if value > 0 {
		return &Dollar{value}, nil
	}
	return nil, ErrDenominationCantBeNegative
}

func (rupee *Rupee) ToDollar() *Dollar {
	value := int(rupee.value / 82.47)
	return &Dollar{value}
}

func (dollar *Dollar) ToRupee() *Rupee {
	value := float64(dollar.value) * 82.47
	return &Rupee{value}
}

type Wallet struct {
	currentBalance      float64
	commonCurrencyType  string
	currentCurrencyType string
}

func compareRupeeAndDollar(rupee *Rupee, dollar *Dollar) bool {
	conversionRate := 82.47
	if rupee.value == float64(dollar.value)*conversionRate {
		return true
	}
	return false
}

func NewWallet(startingBalance int, commonCurrencyType string) (*Wallet, error) {
	if startingBalance > 0 {
		return &Wallet{float64(startingBalance), commonCurrencyType, commonCurrencyType}, nil
	}
	return nil, ErrDenominationCantBeNegative
}

func (w *Wallet) moneyConversion(toType string, amountForTransaction, conversionRate float64, transactionType string) bool {
	if transactionType == "Debit" {
		convertedValue := conversionRate * amountForTransaction
		if w.currentBalance-convertedValue >= 0 && amountForTransaction > 0 {
			w.currentCurrencyType = toType
			w.currentBalance = w.currentBalance - convertedValue
			return true
		}
		return false
	} else if transactionType == "Credit" {
		if amountForTransaction > 0 {
			w.currentCurrencyType = toType
			convertedValue := conversionRate * amountForTransaction
			w.currentBalance = w.currentBalance + convertedValue
			return true
		}
		return false
	}
	return false
}

func (w *Wallet) ToDebit(toType string, amountToBeDebited float64, conversionRate float64) bool {
	return w.moneyConversion(toType, amountToBeDebited, conversionRate, "Debit")
}

func (w *Wallet) ToCredit(toType string, amountToBeCredited float64, conversionRate float64) bool {
	return w.moneyConversion(toType, amountToBeCredited, conversionRate, "Debit")
}
func (w *Wallet) GetCurrentBalance() float64 {
	return w.currentBalance
}

func (w *Wallet) AddAmount(transactionAmount float64, typeOfCurrency string, conversionRate float64) error {
	if transactionAmount > 0 {
		w.currentBalance = w.currentBalance + (transactionAmount * conversionRate)
		return nil
	} else if transactionAmount == 0 {
		return ErrDenominationCantBeZero
	}
	return ErrDenominationCantBeNegative
}

func (w *Wallet) Compare(amount float64, typeOfCurrency string, conversionRate float64) bool {
	if amount > 0 {
		if w.currentBalance == amount {
			return true
		}
	}
	return false
}
