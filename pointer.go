package null

import (
	"encoding/json"
	"time"
)

// BoolP acts as bool in *BoolP,
// but has convenient methods.
type BoolP bool

func NewBoolP(v bool) *BoolP  { return (*BoolP)(&v) }
func (p *BoolP) Val() bool    { return ptrVal((*bool)(p)) }
func (p *BoolP) IsSet() bool  { return ptrIsSet((*bool)(p)) }
func (p *BoolP) ToNull() Bool { return Bool{NewJ((*bool)(p))} }

// ByteP acts as byte in *ByteP,
// but has convenient methods.
type ByteP byte

func NewByteP(v byte) *ByteP  { return (*ByteP)(&v) }
func (p *ByteP) Val() byte    { return ptrVal((*byte)(p)) }
func (p *ByteP) IsSet() bool  { return ptrIsSet((*byte)(p)) }
func (p *ByteP) ToNull() Byte { return Byte{NewJ((*byte)(p))} }

// Int16P acts as int16 in *Int16P,
// but has convenient methods.
type Int16P int16

func NewInt16P(v int16) *Int16P { return (*Int16P)(&v) }
func (p *Int16P) Val() int16    { return ptrVal((*int16)(p)) }
func (p *Int16P) IsSet() bool   { return ptrIsSet((*int16)(p)) }
func (p *Int16P) ToNull() Int16 { return Int16{NewJ((*int16)(p))} }

// Int32P acts as int32 in *Int32P,
// but has convenient methods.
type Int32P int32

func NewInt32P(v int32) *Int32P { return (*Int32P)(&v) }
func (p *Int32P) Val() int32    { return ptrVal((*int32)(p)) }
func (p *Int32P) IsSet() bool   { return ptrIsSet((*int32)(p)) }
func (p *Int32P) ToNull() Int32 { return Int32{NewJ((*int32)(p))} }

// Int64P acts as int64 in *Int64P,
// but has convenient methods.
type Int64P int64

func NewInt64P(v int64) *Int64P { return (*Int64P)(&v) }
func (p *Int64P) Val() int64    { return ptrVal((*int64)(p)) }
func (p *Int64P) IsSet() bool   { return ptrIsSet((*int64)(p)) }
func (p *Int64P) ToNull() Int64 { return Int64{NewJ((*int64)(p))} }

// StrP acts as string in *StrP,
// but has convenient methods.
type StrP string

func NewStrP(v string) *StrP { return (*StrP)(&v) }
func (p *StrP) Val() string  { return ptrVal((*string)(p)) }
func (p *StrP) IsSet() bool  { return ptrIsSet((*string)(p)) }
func (p *StrP) ToNull() Str  { return Str{NewJ((*string)(p))} }

// TimeP acts as time.Time in *TimeP,
// but has convenient methods.
type TimeP time.Time

func NewTimeP(v time.Time) *TimeP { return (*TimeP)(&v) }
func (p *TimeP) Val() time.Time   { return ptrVal((*time.Time)(p)) }
func (p *TimeP) IsSet() bool      { return ptrIsSet((*time.Time)(p)) }
func (p *TimeP) ToNull() Time     { return Time{NewJ((*time.Time)(p))} }

// CustomTP acts as DBType in *CustomTP,
// but has convenient methods.
type CustomTP[
	T any,
	pT interface {
		*T
		ScannerValuer
	},
] struct {
	Internal T
}

func NewDBTypeP[
	T any,
	pT interface {
		*T
		ScannerValuer
	},
](v T) *CustomTP[T, pT] {
	return &CustomTP[T, pT]{v}
}

func (p *CustomTP[T, pT]) Val() T      { return ptrVal(p).Internal }
func (p *CustomTP[T, pT]) IsSet() bool { return ptrIsSet(p) }

func (v *CustomTP[T, pT]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Internal)
}

func (v *CustomTP[T, pT]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.Internal)
}

func (p *CustomTP[T, pT]) ToNull() (v CustomT[T, pT]) {
	if p == nil {
		return
	}
	v.Val = p.Internal
	v.IsSet = true
	return
}
