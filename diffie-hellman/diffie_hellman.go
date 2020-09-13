// Package diffiehellman provides functions for Diffie-Hellman key exchange.
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey generates a key value greater than 2 and less than p.
func PrivateKey(p *big.Int) *big.Int {
	limit := new(big.Int).Sub(p, big.NewInt(2))
	seed := rand.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	return new(big.Int).Rand(seed, limit)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0)
}

// NewPair creates a key pair using prime numbers p and g.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0)
}
