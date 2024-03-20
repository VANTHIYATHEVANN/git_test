package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencyIsNegative(t *testing.T) {
	_, err := NewDollar(-3)

	if err == nil {
		t.Error("Invalid Dollar denomimation")
	}
}

func TestOneDollarEquals82Rupees(t *testing.T) {
	rupee, _ := NewRupee(82.47)
	dollar, _ := NewDollar(1)

	isEqual := compareRupeeAndDollar(rupee, dollar)

	assert.Equal(t, true, isEqual)
}

func TestCurrencyConversionFromRupeeToDollar(t *testing.T) {
	rupee, _ := NewRupee(82.47)

	want, _ := NewDollar(1)
	got := rupee.ToDollar()

	assert.Equal(t, want, got)
}

func TestCurrencyConversionFromDollarToRupee(t *testing.T) {
	dollar, _ := NewDollar(1)
	want, _ := NewRupee(82.47)
	got := dollar.ToRupee()
	assert.Equal(t, want, got)
}

func Test_WalletShouldAtleastHave10RsWhen10RsIsAdded(t *testing.T) {
	wallet, _ := NewWallet(5, "INR")

	err := wallet.AddAmount(5, "INR", 1)
	isSame := wallet.Compare(10, "INR", 1)

	assert.Equal(t, true, isSame)
	assert.Equal(t, nil, err)
}

func Test_WalletShouldNotBeAddedWithNegativeDenominations(t *testing.T) {
	wallet, _ := NewWallet(5, "INR")

	err := wallet.AddAmount(-5, "INR", 1)

	assert.NotEqual(t, nil, err)

}

func Test_WalletShouldNotBeAddedWithZeroDenominations(t *testing.T) {
	wallet, _ := NewWallet(5, "INR")

	err := wallet.AddAmount(0, "INR", 1)

	assert.NotEqual(t, nil, err)

}

func TestAmountShouldBeDebited(t *testing.T) {
	wallet, _ := NewWallet(1000, "INR")

	isPossible := wallet.ToDebit("USD", 8, 82.47)

	assert.Equal(t, true, isPossible)
}

func TestShouldNotBeDebitedAsRequestedAmountIsGreaterThanBalance(t *testing.T) {
	wallet, _ := NewWallet(1000, "INR")

	isPossible := wallet.ToDebit("USD", 15, 82.47)

	assert.Equal(t, false, isPossible)
}

func TestShouldNotBeDebitedAsRequestedAmountIsNegative(t *testing.T) {
	wallet, _ := NewWallet(1000, "INR")

	isPossible := wallet.ToDebit("USD", -15, 82.47)

	assert.Equal(t, false, isPossible)
}

func TestShouldNotBeDebitedAsRequestedAmountIsZero(t *testing.T) {
	wallet, _ := NewWallet(1000, "INR")

	isPossible := wallet.ToDebit("USD", 0, 82.47)

	assert.Equal(t, false, isPossible)
}

func TestAmountShouldBeCredited(t *testing.T) {
	wallet, _ := NewWallet(1000, "INR")

	isPossible := wallet.ToCredit("USD", 8, 82.47)

	assert.Equal(t, true, isPossible)
}

func TestDebitTransactionMaintainsCorrectBalance(t *testing.T) {
	wallet, _ := NewWallet(1000, "INR")

	isPossible := wallet.ToCredit("USD", 8, 82.47)
	got := wallet.GetCurrentBalance()
	want := float64(1000) - (82.47 * 8)

	assert.Equal(t, want, got)
	assert.Equal(t, true, isPossible)
}
