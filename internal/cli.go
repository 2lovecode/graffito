package internal

import (
	"math/rand"

	"github.com/speps/go-hashids/v2"
)

func DefaultHash() *hashids.HashID {
	hash := hashids.NewData()
	hash.Salt = "APpAUptXsssdfffww"
	hash.MinLength = 6
	hash.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890+-*=@#$%"
	h, _ := hashids.NewWithData(hash)
	return h
}

func EncodeHash(num int64) string {
	h := DefaultHash()

	e, _ := h.EncodeInt64([]int64{num})
	return e
}

func DecodeHash(str string) int64 {
	h := DefaultHash()
	d, _ := h.DecodeInt64WithError(str)
	return d[0]
}

func consistentShuffleInPlace(alphabet []rune, salt []rune) {
	if len(salt) == 0 {
		return
	}

	for i, v, p := len(alphabet)-1, 0, 0; i > 0; i-- {
		p += int(salt[v])
		j := (int(salt[v]) + v + p) % i
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
		v = (v + 1) % len(salt)
	}
	return
}

func fisherYatesShuffleWithConsistency(arr []int, seed int64) {
	rng := rand.New(rand.NewSource(seed))

	for i := len(arr) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
