package utils

// CheckErr is yes
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MustErr(i int) {
	panic(i)
}
