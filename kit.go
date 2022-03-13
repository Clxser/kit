package kit

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
)

type action interface {
	slot() int
}

func Set(slot int) set { return set{s: slot} }

type set struct {
	s int
}

func (s set) slot() int { return s.s }

func Add() add { return add{} }

type add struct{}

func (add) slot() int { return 0 }

type Items = map[action]item.Stack

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

	for act, i := range kit.Items() {
		switch a := act.(type) {
		case add:
			inv.AddItem(i)
		case set:
			inv.SetItem(a.slot(), i)
		}
	}

}
