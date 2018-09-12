package ir

import (
	"math/big"

	"github.com/llir/l/ir/types"
)

// ConstInt is an integer constant.
type ConstInt struct {
	// Integer type.
	Typ *types.IntType
	// Integer constant.
	X *big.Int
}

// NewIntFromInt64 returns a new integer constant based on the given integer
// type and 64-bit interger value.
func NewIntFromInt64(typ *types.IntType, x int64) *ConstInt {
	panic("not yet implemented")
}

// NewIntFromString returns a new integer constant based on the given integer
// type and string.
//
// The integer string may be expressed in one of the following forms.
//
//    * boolean literal
//         true | false
//    * integer literal
//         [-]? [0-9]+
//    * hexadecimal integer literal
//         TODO: add support for hexadecimal integer literal notation.
func NewIntFromString(typ *types.IntType, s string) *ConstInt {
	panic("not yet implemented")
}

// Type returns the type of the constant.
func (c *ConstInt) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstInt) Ident() string {
	panic("not yet implemented")
}