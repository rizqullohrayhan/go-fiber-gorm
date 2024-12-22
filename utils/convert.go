package utils

func ConvertToUintArray(input []int) []uint {
	output := make([]uint, len(input))
	for i, v := range input {
		output[i] = uint(v)
	}
	return output
}