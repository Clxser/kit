package kit

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

type Add = []item.Stack
type Slots = map[int]item.Stack

type Items struct {
	Slots Slots
	Add Add
}

type Armour struct {
	Helmet, Chestplate, Leggings, Boots item.Stack
}

type EffectKit interface {
	Effects() []effect.Effect
}

type Kit interface {
	Name() string
	Items() Items
	Armour() Armour
}

func GiveKit(p *player.Player, kit Kit) {
	inv := p.Inventory()
	armr := p.Armour()

	if e, ok := kit.(EffectKit); ok {
		for _, ef := range e.Effects() {
			p.AddEffect(ef)
		}
	}
	
	a := kit.Armour()
	armr.Set(a.Helmet, a.Chestplate, a.Leggings, a.Boots)

	for slot, i := range kit.Items().Slots{
		inv.SetItem(slot, i)
	}

	for _, i := range kit.Items().Add{
		inv.AddItem(i)
	}

}
