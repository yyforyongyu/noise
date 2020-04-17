package noise

import (
	"testing"

	"github.com/yyforyongyu/noise/cipher"
	noiseCipher "github.com/yyforyongyu/noise/cipher"
	noiseCurve "github.com/yyforyongyu/noise/dh"
	noiseHash "github.com/yyforyongyu/noise/hash"
)

var (
	cipherA = noiseCipher.FromString("AESGCM")
	cipherB = noiseCipher.FromString("AESGCM")

	hashA = noiseHash.FromString("SHA256")
	hashB = noiseHash.FromString("SHA256")

	curveA = noiseCurve.FromString("25519")
	curveB = noiseCurve.FromString("25519")

	key = [CipherKeySize]byte{
		0xa8, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab,
		0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab,
		0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab,
		0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0x6b,
	}
	ad = []byte{
		0xa8, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab,
		0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab,
	}
	message  = []byte("Noise Protocol Framework")
	maxNonce = cipher.MaxNonce
)

func TestStates(t *testing.T) {

}
