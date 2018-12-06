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

func MapDataProduct(merchant interface{}) interface{} {
	data := map[string]interface{}{
		"Product": merchant,
	}

	return data
}

func MapDataReport(report interface{}) interface{} {
	data := map[string]interface{}{
		"report": report,
	}

	return data
}
