// @program:     pokemon
// @author:      sunsun

package pokemon

import (
	"fmt"

	"pokemon/move"
)

//喜闻乐见的不会还手的木头人
//怎么让他动起来?
type WoodenMan struct {
	BasePokemon
}

func NewWoodenMan() Pokemon {
	//return &BasePokemon{
	//	name:  "wooden man",
	//	hp:    300,
	//	mp:    0,
	//	moves: [4]move.Move{
	//		move.NewMove("只需要1招",20,0),
	//	},
	//}
	return &WoodenMan{BasePokemon{
		name: "wooden man",
		hp:   300,
		mp:   0,
		moves: [4]move.Move{
			move.NewMove("只需要1招", 20, 0),
		},
	}}
}

// Attack WoodenMan 的攻击方式
func (wm *WoodenMan) Attack(other Pokemon) {
	move:=wm.moves[0]
	fmt.Println(wm.Name(),"对",other.Name(),"使用了", move.MoveName())
	other.ReduceHP(move.MoveHurt())
}
