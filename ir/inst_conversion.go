package ir

import (
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstTrunc is an LLVM IR trunc instruction.
type InstTrunc struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewTrunc returns a new trunc instruction based on the given source value and
// target type.
func NewTrunc(from value.Value, to types.Type) *InstTrunc {
	return &InstTrunc{From: from, To: to}
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstZExt is an LLVM IR zext instruction.
type InstZExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewZExt returns a new zext instruction based on the given source value and
// target type.
func NewZExt(from value.Value, to types.Type) *InstZExt {
	return &InstZExt{From: from, To: to}
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSExt is an LLVM IR sext instruction.
type InstSExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewSExt returns a new sext instruction based on the given source value and
// target type.
func NewSExt(from value.Value, to types.Type) *InstSExt {
	return &InstSExt{From: from, To: to}
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPTrunc is an LLVM IR fptrunc instruction.
type InstFPTrunc struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPTrunc returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTrunc(from value.Value, to types.Type) *InstFPTrunc {
	return &InstFPTrunc{From: from, To: to}
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPExt is an LLVM IR fpext instruction.
type InstFPExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPExt returns a new fpext instruction based on the given source value and
// target type.
func NewFPExt(from value.Value, to types.Type) *InstFPExt {
	return &InstFPExt{From: from, To: to}
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToUI is an LLVM IR fptoui instruction.
type InstFPToUI struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPToUI returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUI(from value.Value, to types.Type) *InstFPToUI {
	return &InstFPToUI{From: from, To: to}
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToSI is an LLVM IR fptosi instruction.
type InstFPToSI struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPToSI returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSI(from value.Value, to types.Type) *InstFPToSI {
	return &InstFPToSI{From: from, To: to}
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUIToFP is an LLVM IR uitofp instruction.
type InstUIToFP struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewUIToFP returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFP(from value.Value, to types.Type) *InstUIToFP {
	return &InstUIToFP{From: from, To: to}
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSIToFP is an LLVM IR sitofp instruction.
type InstSIToFP struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewSIToFP returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFP(from value.Value, to types.Type) *InstSIToFP {
	return &InstSIToFP{From: from, To: to}
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPtrToInt is an LLVM IR ptrtoint instruction.
type InstPtrToInt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewPtrToInt returns a new ptrtoint instruction based on the given source
// value and target type.
func NewPtrToInt(from value.Value, to types.Type) *InstPtrToInt {
	return &InstPtrToInt{From: from, To: to}
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstIntToPtr is an LLVM IR inttoptr instruction.
type InstIntToPtr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewIntToPtr returns a new inttoptr instruction based on the given source
// value and target type.
func NewIntToPtr(from value.Value, to types.Type) *InstIntToPtr {
	return &InstIntToPtr{From: from, To: to}
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstBitCast is an LLVM IR bitcast instruction.
type InstBitCast struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewBitCast returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCast(from value.Value, to types.Type) *InstBitCast {
	return &InstBitCast{From: from, To: to}
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAddrSpaceCast is an LLVM IR addrspacecast instruction.
type InstAddrSpaceCast struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// source value and target type.
func NewAddrSpaceCast(from value.Value, to types.Type) *InstAddrSpaceCast {
	return &InstAddrSpaceCast{From: from, To: to}
}