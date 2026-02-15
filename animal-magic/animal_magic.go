package chance
import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
// d := RollADie() // d will be assigned a random int, 1 <= d <= 20
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
// f := GenerateWandEnergy()  // f will be assigned a random float64, 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return rand.Float64() * (12 - 0)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	critters := []string{
	"ant",
	"beaver",
	"cat",
	"dog",
	"elephant",
	"fox",
	"giraffe",
	"hedgehog",
	}

	rand.Shuffle(len(critters), func(i, j int) {
		critters[i], critters[j] = critters[j], critters[i]
	})
	
	return critters
}
