package ir

import (
	"github.com/llir/l/ir/instruction"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewTrunc returns a new trunc instruction based on the given source value and
// target type.
func (block *BasicBlock) NewTrunc(from value.Value, to types.Type) *instruction.Trunc {
	inst := instruction.NewTrunc(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewZExt returns a new zext instruction based on the given source value and
// target type.
func (block *BasicBlock) NewZExt(from value.Value, to types.Type) *instruction.ZExt {
	inst := instruction.NewZExt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSExt returns a new sext instruction based on the given source value and
// target type.
func (block *BasicBlock) NewSExt(from value.Value, to types.Type) *instruction.SExt {
	inst := instruction.NewSExt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPTrunc returns a new fptrunc instruction based on the given source value
// and target type.
func (block *BasicBlock) NewFPTrunc(from value.Value, to types.Type) *instruction.FPTrunc {
	inst := instruction.NewFPTrunc(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPExt returns a new fpext instruction based on the given source value and
// target type.
func (block *BasicBlock) NewFPExt(from value.Value, to types.Type) *instruction.FPExt {
	inst := instruction.NewFPExt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPToUI returns a new fptoui instruction based on the given source value
// and target type.
func (block *BasicBlock) NewFPToUI(from value.Value, to types.Type) *instruction.FPToUI {
	inst := instruction.NewFPToUI(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFPToSI returns a new fptosi instruction based on the given source value
// and target type.
func (block *BasicBlock) NewFPToSI(from value.Value, to types.Type) *instruction.FPToSI {
	inst := instruction.NewFPToSI(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewUIToFP returns a new uitofp instruction based on the given source value
// and target type.
func (block *BasicBlock) NewUIToFP(from value.Value, to types.Type) *instruction.UIToFP {
	inst := instruction.NewUIToFP(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSIToFP returns a new sitofp instruction based on the given source value
// and target type.
func (block *BasicBlock) NewSIToFP(from value.Value, to types.Type) *instruction.SIToFP {
	inst := instruction.NewSIToFP(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewPtrToInt returns a new ptrtoint instruction based on the given source
// value and target type.
func (block *BasicBlock) NewPtrToInt(from value.Value, to types.Type) *instruction.PtrToInt {
	inst := instruction.NewPtrToInt(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewIntToPtr returns a new inttoptr instruction based on the given source
// value and target type.
func (block *BasicBlock) NewIntToPtr(from value.Value, to types.Type) *instruction.IntToPtr {
	inst := instruction.NewIntToPtr(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewBitCast returns a new bitcast instruction based on the given source value
// and target type.
func (block *BasicBlock) NewBitCast(from value.Value, to types.Type) *instruction.BitCast {
	inst := instruction.NewBitCast(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// source value and target type.
func (block *BasicBlock) NewAddrSpaceCast(from value.Value, to types.Type) *instruction.AddrSpaceCast {
	inst := instruction.NewAddrSpaceCast(from, to)
	block.Insts = append(block.Insts, inst)
	return inst
}