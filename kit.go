package kit

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

type action interface {
	slot() int
}

type Set struct {
	int
}

func (s Set) slot() int { return s.int }

type Add struct{}

func (Add) slot() int { return 0 }

type Kit interface {
	Name() string
	Items() map[action]item.Stack
	Armour() [4]item.Stack
}

var armrFuncs = [4]func(armr *inventory.Armour, i item.Stack){
	func(armr *inventory.Armour, i item.Stack) { armr.SetHelmet(i) },
	func(armr *inventory.Armour, i item.Stack) { armr.SetChestplate(i) },
	func(armr *inventory.Armour, i item.Stack) { armr.SetLeggings(i) },
	func(armr *inventory.Armour, i item.Stack) { armr.SetBoots(i) },
}

func GiveKit(p *player.Player, kit Kit) {
	inv := p.Inventory()
	armr := p.Armour()

	for n, i := range kit.Armour() {
		if !i.Empty() {
			armrFuncs[n](armr, i)
		}
	}

	for act, i := range kit.Items() {
		switch a := act.(type) {
		case Add:
			inv.AddItem(i)
		case Set:
			inv.SetItem(a.slot(), i)
		}
	}

}
