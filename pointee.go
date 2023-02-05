package null

import (
	"encoding/json"
	"time"
)

// BoolP is a pointee to bool.
type BoolP bool

func (p *BoolP) Val() bool { return pVal((*bool)(p)) }
func (p *BoolP) Set() bool { return pSet((*bool)(p)) }
func (p *BoolP) V() Bool   { return Bool{pV((*bool)(p))} }

// ByteP is a pointee to byte.
type ByteP byte

func (p *ByteP) Val() byte { return pVal((*byte)(p)) }
func (p *ByteP) Set() bool { return pSet((*byte)(p)) }
func (p *ByteP) V() Byte   { return Byte{pV((*byte)(p))} }

// Int16P is a pointee to int16.
type Int16P int16

func (p *Int16P) Val() int16 { return pVal((*int16)(p)) }
func (p *Int16P) Set() bool  { return pSet((*int16)(p)) }
func (p *Int16P) V() Int16   { return Int16{pV((*int16)(p))} }

// Int32P is a pointee to int32.
type Int32P int32

func (p *Int32P) Val() int32 { return pVal((*int32)(p)) }
func (p *Int32P) Set() bool  { return pSet((*int32)(p)) }
func (p *Int32P) V() Int32   { return Int32{pV((*int32)(p))} }

// Int64P is a pointee to int64.
type Int64P int64

func (p *Int64P) Val() int64 { return pVal((*int64)(p)) }
func (p *Int64P) Set() bool  { return pSet((*int64)(p)) }
func (p *Int64P) V() Int64   { return Int64{pV((*int64)(p))} }

// StrP is a pointee to string.
type StrP string

func (p *StrP) Val() string { return pVal((*string)(p)) }
func (p *StrP) Set() bool   { return pSet((*string)(p)) }
func (p *StrP) V() Str      { return Str{pV((*string)(p))} }

// TimeP is a pointee to time.Time.
type TimeP time.Time

func (p *TimeP) Val() time.Time { return pVal((*time.Time)(p)) }
func (p *TimeP) Set() bool      { return pSet((*time.Time)(p)) }
func (p *TimeP) V() Time        { return Time{pV((*time.Time)(p))} }

// DBTypeP is a pointee to DBType.
type DBTypeP[T ScannerValuer] struct {
	Internal T
}

func (p *DBTypeP[T]) Val() T    { return pVal(p).Internal }
func (p *DBTypeP[T]) Set() bool { return pSet(p) }

func (v *DBTypeP[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Internal)
}

func (v *DBTypeP[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.Internal)
}

func (p *DBTypeP[T]) V() (v DBType[T]) {
	if p == nil {
		return
	}
	v.Val = p.Internal
	v.Set = true
	return
}
