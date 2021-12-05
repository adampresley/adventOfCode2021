package lifesupport

func CalculateLifeSupportRating(oxygenGeneratorRating, co2ScrubberRating int) int {
	return oxygenGeneratorRating * co2ScrubberRating
}
