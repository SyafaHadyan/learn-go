package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Money) {
		t.Helper()
		got := wallet.GetBalance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

    assertError := func(t testing.TB, got error, want string) {
        t.Helper()

        if got == nil {
            t.Fatal("didn't get an error but wanted one")
        }

        if got.Error() != want {
            t.Errorf("got %q, want %q", got, want)
        }
    }

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Money(10))
		assertBalance(t, wallet, Money(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Money(20)}
		wallet.Withdraw(Money(10))
		assertBalance(t, wallet, Money(10))
	})

    t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Money(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Money(100))

        assertBalance(t, wallet, startingBalance)

        if err == nil {
            t.Error("wanted an error but didn't get one")
        }
    })

    t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Money(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Money(100))

        assertError(t, err, "Withdraw amount exceeds balance")
        assertBalance(t, wallet, startingBalance)
    })

}
