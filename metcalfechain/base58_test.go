package metcalfechain

import (
	"encoding/hex"
	"fmt"
	"github.com/blocktree/go-owcrypt"
	"testing"
)

func TestDecode(t *testing.T) {
	secret := "MsHWvJfDZ3RBqKUWmiuPA1RRfZfMEFRY29"
	pri, _ := Decode(secret, BitcoinAlphabet)

	fmt.Println(hex.EncodeToString(pri))

	chk := owcrypt.Hash(pri[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)

	fmt.Println(hex.EncodeToString(chk))


	pub,_ := hex.DecodeString("02f8a9bc9453df1c9e40f3337e8fe05f20d11da1537f19a9f75171f37a183a58fb")

	phash := owcrypt.Hash(pub, 20, owcrypt.HASH_ALG_HASH160)

	fmt.Println(hex.EncodeToString(phash))
}
