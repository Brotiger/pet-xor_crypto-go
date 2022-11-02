package xor

func EncryptDecrypt(input, key string, shift int) (output string) {
	kl := len(key) - 1
	q := 0
	for i := range input {
		if q > kl {
			q = 0
		}
		output += string(input[i] ^ (key[q] + byte(shift)))
		q++
		shift++
	}

	return output
}
