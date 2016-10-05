package bip44_test

import (
	"testing"

	"github.com/FactomProject/go-bip32"
	"github.com/FactomProject/go-bip39"
	. "github.com/FactomProject/go-bip44"
)

func TestTest(t *testing.T) {
	//var factoidHex uint32 = 0x80000083
	//var ecHex uint32 = 0x80000084

	mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		panic(err)
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	child, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		panic(err)
	}
	t.Logf("%v", child.String())

	child, err = child.NewChildKey(bip32.FirstHardenedChild + 131)
	if err != nil {
		panic(err)
	}
	t.Logf("%v", child.String())

	child, err = child.NewChildKey(bip32.FirstHardenedChild)
	if err != nil {
		panic(err)
	}
	t.Logf("%v", child.String())

	child, err = child.NewChildKey(0)
	if err != nil {
		panic(err)
	}
	t.Logf("%v", child.String())

	t.Logf("%x", TypeFactomFactoids)

	t.FailNow()
}
