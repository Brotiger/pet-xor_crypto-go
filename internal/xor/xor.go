package xor

func EncryptDecrypt(input, key string, shift int) (output string) {
	kl := len(key)

	for i := range input {
		output += string(input[i] ^ (key[i%kl] + byte(shift)))
		shift++
	}

	return output
}
