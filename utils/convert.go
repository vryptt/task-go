package utils

func BytesToMB(bytes uint64) uint64 {
	return bytes / 1024 / 1024
}

func BytesToGB(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024 / 1024
}
