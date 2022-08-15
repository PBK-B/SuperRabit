package utils

type ArrayUtil struct{}

func (arr ArrayUtil) IndexOf(slice []interface{}, el interface{}) int {
	for k, v := range slice {
		if v == el {
			return k
		}
	}
	return -1
}
