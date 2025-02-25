package aether

import (
	"github.com/genshinsim/gcsim/internal/characters/traveler/common/electro"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/keys"
)

func init() {
	core.RegisterCharFunc(keys.AetherElectro, electro.NewChar(0))
}
