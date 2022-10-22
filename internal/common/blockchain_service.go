package common

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

type BlockchainService interface {
	Hash(src []byte) ([]byte, error)
}

type BlockchainServiceImpl struct {
	rand *rand.Rand
}

func NewBlockchainServiceImpl() *BlockchainServiceImpl {
	src := rand.NewSource(time.Now().UnixMicro())
	return &BlockchainServiceImpl{
		rand: rand.New(src),
	}
}

func (svc *BlockchainServiceImpl) Hash(src []byte) ([]byte, error) {
	buf := make([]byte, len(src))
	hash := sha256.New()
	_, err := hash.Write(buf)
	if err != nil {
		return nil, fmt.Errorf("create hash: %v", err)
	}
	return hash.Sum(nil), nil
}
