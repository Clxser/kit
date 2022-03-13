package kit

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

type Items struct {
	Slots map[int]item.Stack
	Add []item.Stack
}

type EffectKit interface {
	Effects() []effect.Effect
}

type Kit interface {
	Name() string
	Items() Items
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

	if e, ok := kit.(EffectKit); ok {
		for _, ef := range e.Effects() {
			p.AddEffect(ef)
		}
	}

	for n, i := range kit.Armour() {
		if !i.Empty() {
			armrFuncs[n](armr, i)
		}
	}

	for _, i := range kit.Items().Add{
		inv.AddItem(i)
	}

	for slot, i := range kit.Items().Slots{
		inv.SetItem(slot, i)
	}

}
