package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ011756_+"

// Tools is the type need to instantiate this module
type Tools struct {
}

//RandomString returns a string of random characters of length n using randomStringSource
//as the source for the string
func (t *Tools) RandomString(n int) string {
	s := make([]rune, n)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
