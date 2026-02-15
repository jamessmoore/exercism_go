package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11 
	case card == "king":
		return 10
	case card == "queen":
		return 10
	case card == "jack":
		return 10
	case card == "ten":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	myCardsSum := int(ParseCard(card1) + ParseCard(card2))

	switch {
	// If you have a pair of aces you must always split them.
	case (card1 =="ace" && card2 == "ace"):
		return "P"
	/* If you have a Blackjack (two cards that sum up to a value of 21), 
	and the dealer does not have an ace, a face card (Jack/Queen/King) 
	or a ten then you automatically win. If the dealer does have any of 
	those cards then you'll have to stand and wait for the reveal of the 
	other card.
	*/
	case (myCardsSum == 21):
		switch {
		case (ParseCard(dealerCard) < 10):
			return "W"
		default:
			return "S"
		}
	/*
	If your cards sum up to a value within the range [17, 20] you 
	should always stand.
	*/
	case (myCardsSum > 16 && myCardsSum < 21):
		return "S"
	/* 
	If your cards sum up to a value within the range [12, 16] you 
	should always stand unless the dealer has a 7 or higher, in which 
	case you should always hit.
	*/
	case (myCardsSum > 11 && myCardsSum < 17):
		switch {
		case (ParseCard(dealerCard) >= 7):
			return "H"
		default:
			return "S"
		}
	// If your cards sum up to 11 or lower you should always hit.
	case (myCardsSum <= 11):
		return "H"
	default:
		return "S"
	}
}
