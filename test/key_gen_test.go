package generator

import (
	"fmt"
	"testing"
)

func TestKeyGen(t *testing.T) {
	generator, err := NewKeyGenerator("INVALIDKEY")
	if err == nil {
		t.Error("should return error when key length is invalid")
	}
	testid := int64(1)
	generator, err = NewKeyGenerator("123456789066666688888888")
	if err != nil {
		t.Fatal(err)
	}
	key, err := generator.GenRandomKey(testid)
	if err != nil {
		t.Error(err)
	}
	t.Log(key)
	id, err := generator.DecodeIDFromRandomKey(key)
	if err != nil {
		t.Error(err)
	}
	if id != testid {
		t.Errorf("wrong id %d, want %d", id, testid)
	}

	id, err = generator.DecodeIDFromRandomKey("")
	if err == nil {
		t.Error("decode id from random key should return error for empty key.")
	}

	id, err = generator.DecodeIDFromRandomKey("1111111111111111111111111111111111111111")
	if err == nil {
		t.Errorf("decode id from random key should return error for bad key : %s", "1111111111111111111111111111111111111111")
	}

}

func TestGenRandomKey(t *testing.T) {
	generator, err := NewKeyGenerator("ABCDEFGHIJKLMNOPABCDEFGHIJKLMNOP")
	if err != nil {
		fmt.Println("err1:", err)
	}
	key, err := generator.GenRandomKey(1)
	if err != nil {
		fmt.Println("err2:", err)
	}
	fmt.Println("key:", key)
}

func TestLog(t *testing.T) {
	generator, err := NewKeyGenerator("ABCDEFGHIJKLMNOPABCDEFGHIJKLMNOP")
	if err != nil {
		fmt.Println("err1:", err)
	}
	key, err := generator.DecodeIDFromRandomKey("a862360cf3c8919b3425961b88c86835d6d4db1fbef5f0317c8726fe8e6b62c9")
	if err != nil {
		fmt.Println("err2:", err)
	}
	fmt.Println("key:", key)
}
