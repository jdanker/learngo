package iteration

func Repeat(a string, multiple int) string {
	var repeated string

	for i := 0; i < multiple; i++ {
		repeated += a
	}

	return repeated
}

