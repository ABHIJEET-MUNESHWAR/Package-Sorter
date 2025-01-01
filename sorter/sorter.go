package sorter

import "github.com/ABHIJEET-MUNESHWAR/Package-Sorter/utils"

// Sort determines the stack for the package based on its dimensions and mass.
func Sort(width, height, length, mass int) string {
	volume := utils.CalculateVolume(width, height, length)
	isBulky := volume >= 1000000 || width >= 150 || height >= 150 || length >= 150
	isHeavy := mass >= 20
	switch {
	case isBulky && isHeavy:
		return "REJECTED"
	case isBulky || isHeavy:
		return "SPECIAL"
	default:
		return "STANDARD"
	}
}