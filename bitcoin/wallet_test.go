package bitcoin

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, errorInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("got an error but did't want one")
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but did't got one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
