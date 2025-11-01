package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	// step 1: create private/public key pair (32/64 bytes)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey

	// step 2: perform sha256 hashing on the public key -> 32 bytes
	h2 := sha256.New()
	h2.Write(publicKey.X.Bytes())
	h2.Write(publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// step 3: perform ripemd160 hash on the sha256 result -> 20 bytes
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	// step 4: add version byte in front of ripemd160 hash (0x00 for main network) -> 21 bytes
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	// step 5: perform sha256 over result from step 4

	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// step 6: perform sha256 over result from step 5
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	// step 7: take first 4 bytes of result from step 6 for checksum
	checksum := digest6[:4]

	// step 8: add the checksum bytes at the end of extended ripemd160 hash from step 4
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], checksum)

	// step 9: convert the result into a base58 string
	blockchainAddress := base58.Encode(dc8)

	return &Wallet{
		privateKey:        privateKey,
		publicKey:         publicKey,
		blockchainAddress: blockchainAddress,
	}
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
