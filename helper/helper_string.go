package helper

import (
	"math/rand"
	"time"
)

func StringWithMerchantset(Merchantset string) string {

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 10)
	for i := range b {
		b[i] = Merchantset[seededRand.Intn(len(Merchantset))]
	}
	return string(b)
}

func MapData(merchant interface{}) interface{} {
	data := map[string]interface{}{
		"merchant": merchant,
	}

	return data
}
