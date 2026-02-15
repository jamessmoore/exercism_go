package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64){
	var noodles int
	var sauce float64

	for layer := 0; layer < len(layers); layer++ {
		switch {
		case layers[layer] == "noodles":
			noodles = noodles + 50
		case layers[layer] == "sauce":
			sauce = sauce + 0.20
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[(len(myList) - 1)] = friendsList[(len(friendsList) - 1)]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))
	for amount := 0; amount < len(amounts); amount++ {
		scaledAmounts[amount] = (amounts[amount] / float64(2)) * float64(portions)
	}
	return scaledAmounts
}

