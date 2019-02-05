package authash

import (
	"github.com/ShyftNetwork/go-empyrean/consensus/ethash"
	"github.com/ShyftNetwork/go-empyrean/log"
)

func New(config ethash.Config, notify []string, noverify bool) *ethash.Ethash {
	log.Info("Using authash consensus engine")
	return ethash.New(config, notify, noverify)
}
