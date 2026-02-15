package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var rate float64
	switch {
	case balance < 0:
		rate = 3.213
	case balance >= 0 && balance < 1000:
		rate = 0.5
	case balance >= 1000 && balance < 5000:
		rate = 1.621
	case balance >= 5000:
		rate = 2.475
	}
	return float32(rate)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	if targetBalance <= balance {
		// The target balance has already been reached
		return years
	} else {
		for i:=0; i>=0; i++ {
			// increment the balance for each year until
			// target is reached
			years += 1
			balance = AnnualBalanceUpdate(balance)
			if balance >= targetBalance {
				break
			}
		}
	}
	return years
}
