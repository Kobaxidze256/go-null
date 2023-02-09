/*
	Copyright (C) 2023 Giorgi Kobakhidze

	This file is part of go-null.

	go-null is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, version 3 of the License.

	go-null is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.

	You should have received a copy of the GNU Lesser General Public License along with go-null. If not, see <https://www.gnu.org/licenses/>.
*/

package null

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// Bool wraps J[bool] and sql.NullBool.
type Bool struct {
	J[bool]
}

func NewBool(v bool) (nullV Bool) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Bool) Ptr() *BoolP { return (*BoolP)(v.J.Ptr()) }

func (v *Bool) Scan(value any) (err error) {
	var sqlV sql.NullBool
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Bool
	v.IsSet = sqlV.Valid
	return
}

func (v Bool) Value() (driver.Value, error) {
	return sql.NullBool{
		Bool:  v.Val,
		Valid: v.IsSet,
	}.Value()
}

// Byte wraps J[byte] and sql.NullByte.
type Byte struct {
	J[byte]
}

func NewByte(v byte) (nullV Byte) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Byte) Ptr() *ByteP { return (*ByteP)(v.J.Ptr()) }

func (v *Byte) Scan(value any) (err error) {
	var sqlV sql.NullByte
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Byte
	v.IsSet = sqlV.Valid
	return
}

func (v Byte) Value() (driver.Value, error) {
	return sql.NullByte{
		Byte:  v.Val,
		Valid: v.IsSet,
	}.Value()
}

// Int16 wraps J[int16] and sql.NullInt16.
type Int16 struct {
	J[int16]
}

func NewInt16(v int16) (nullV Int16) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Int16) Ptr() *Int16P { return (*Int16P)(v.J.Ptr()) }

func (v *Int16) Scan(value any) (err error) {
	var sqlV sql.NullInt16
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Int16
	v.IsSet = sqlV.Valid
	return
}

func (v Int16) Value() (driver.Value, error) {
	return sql.NullInt16{
		Int16: v.Val,
		Valid: v.IsSet,
	}.Value()
}

// Int32 wraps J[int32] and sql.NullInt32.
type Int32 struct {
	J[int32]
}

func NewInt32(v int32) (nullV Int32) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Int32) Ptr() *Int32P { return (*Int32P)(v.J.Ptr()) }

func (v *Int32) Scan(value any) (err error) {
	var sqlV sql.NullInt32
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Int32
	v.IsSet = sqlV.Valid
	return
}

func (v Int32) Value() (driver.Value, error) {
	return sql.NullInt32{
		Int32: v.Val,
		Valid: v.IsSet,
	}.Value()
}

// Int64 wraps J[int64] and sql.NullInt64.
type Int64 struct {
	J[int64]
}

func NewInt64(v int64) (nullV Int64) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Int64) Ptr() *Int64P { return (*Int64P)(v.J.Ptr()) }

func (v *Int64) Scan(value any) (err error) {
	var sqlV sql.NullInt64
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Int64
	v.IsSet = sqlV.Valid
	return
}

func (v Int64) Value() (driver.Value, error) {
	return sql.NullInt64{
		Int64: v.Val,
		Valid: v.IsSet,
	}.Value()
}

// Str wraps J[string] and sql.NullString.
type Str struct {
	J[string]
}

func NewStr(v string) (nullV Str) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Str) Ptr() *StrP { return (*StrP)(v.J.Ptr()) }

func (v *Str) Scan(value any) (err error) {
	var sqlV sql.NullString
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.String
	v.IsSet = sqlV.Valid
	return
}

func (v Str) Value() (driver.Value, error) {
	return sql.NullString{
		String: v.Val,
		Valid:  v.IsSet,
	}.Value()
}

// Time wraps J[time.Time] and sql.NullTime.
type Time struct {
	J[time.Time]
}

func NewTime(v time.Time) (nullV Time) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *Time) Ptr() *TimeP { return (*TimeP)(v.J.Ptr()) }

func (v *Time) Scan(value any) (err error) {
	var sqlV sql.NullTime
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Time
	v.IsSet = sqlV.Valid
	return
}

func (v Time) Value() (driver.Value, error) {
	return sql.NullTime{
		Time:  v.Val,
		Valid: v.IsSet,
	}.Value()
}

type ScannerValuer interface {
	sql.Scanner
	driver.Valuer
}

// CustomT wraps J[ScannerValuer], sql.Scanner and driver.Valuer.
// Internal type T has to be given as T and as *T.
// *T must support json marshaling and unmarshaling.
type CustomT[
	T any,
	pT interface {
		*T
		ScannerValuer
	},
] struct {
	J[T]
}

func NewCustomT[
	T any,
	pT interface {
		*T
		ScannerValuer
	},
](v T) (nullV CustomT[T, pT]) {
	nullV.Val = v
	nullV.IsSet = true
	return
}

func (v *CustomT[T, pT]) Scan(value any) (err error) {
	if err = (pT)(&v.Val).Scan(value); err != nil {
		return
	}
	v.IsSet = true
	return
}

func (v CustomT[T, pT]) Value() (driver.Value, error) {
	if !v.IsSet {
		return nil, nil
	}
	return (pT)(&v.Val).Value()
}
