package protocol

import (
	"encoding/hex"
	"testing"
)

func TestTokenToString(t *testing.T) {
	bs := []byte{52, 48, 182, 128, 220, 221, 144, 132}
	s1 := hex.EncodeToString(bs)
	s2 := string(bs)
	t.Log("s1: ", s1)
	t.Log("s2: ", s2)
}
