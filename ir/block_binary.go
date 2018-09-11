package ir

import (
	"github.com/llir/l/ir/instruction"
	"github.com/llir/l/ir/value"
)

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAdd returns a new add instruction based on the given operands.
func (block *BasicBlock) NewAdd(x, y value.Value) *instruction.Add {
	inst := instruction.NewAdd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFAdd returns a new fadd instruction based on the given operands.
func (block *BasicBlock) NewFAdd(x, y value.Value) *instruction.FAdd {
	inst := instruction.NewFAdd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSub returns a new sub instruction based on the given operands.
func (block *BasicBlock) NewSub(x, y value.Value) *instruction.Sub {
	inst := instruction.NewSub(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFSub returns a new fsub instruction based on the given operands.
func (block *BasicBlock) NewFSub(x, y value.Value) *instruction.FSub {
	inst := instruction.NewFSub(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewMul returns a new mul instruction based on the given operands.
func (block *BasicBlock) NewMul(x, y value.Value) *instruction.Mul {
	inst := instruction.NewMul(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFMul returns a new fmul instruction based on the given operands.
func (block *BasicBlock) NewFMul(x, y value.Value) *instruction.FMul {
	inst := instruction.NewFMul(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewUDiv returns a new udiv instruction based on the given operands.
func (block *BasicBlock) NewUDiv(x, y value.Value) *instruction.UDiv {
	inst := instruction.NewUDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSDiv returns a new sdiv instruction based on the given operands.
func (block *BasicBlock) NewSDiv(x, y value.Value) *instruction.SDiv {
	inst := instruction.NewSDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFDiv returns a new fdiv instruction based on the given operands.
func (block *BasicBlock) NewFDiv(x, y value.Value) *instruction.FDiv {
	inst := instruction.NewFDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewURem returns a new urem instruction based on the given operands.
func (block *BasicBlock) NewURem(x, y value.Value) *instruction.URem {
	inst := instruction.NewURem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSRem returns a new srem instruction based on the given operands.
func (block *BasicBlock) NewSRem(x, y value.Value) *instruction.SRem {
	inst := instruction.NewSRem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFRem returns a new frem instruction based on the given operands.
func (block *BasicBlock) NewFRem(x, y value.Value) *instruction.FRem {
	inst := instruction.NewFRem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}
