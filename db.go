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

// Bool wraps V[bool] and sql.NullBool.
type Bool struct {
	V[bool]
}

func NewBool(v bool) (nullV Bool) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Bool) P() *BoolP { return (*BoolP)(v.V.P()) }

func (v *Bool) Scan(value any) (err error) {
	var sqlV sql.NullBool
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Bool
	v.Set = sqlV.Valid
	return
}

func (v *Bool) Value() (driver.Value, error) {
	return sql.NullBool{
		Bool:  v.Val,
		Valid: v.Set,
	}.Value()
}

// Byte wraps V[byte] and sql.NullByte.
type Byte struct {
	V[byte]
}

func NewByte(v byte) (nullV Byte) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Byte) P() *ByteP { return (*ByteP)(v.V.P()) }

func (v *Byte) Scan(value any) (err error) {
	var sqlV sql.NullByte
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Byte
	v.Set = sqlV.Valid
	return
}

func (v *Byte) Value() (driver.Value, error) {
	return sql.NullByte{
		Byte:  v.Val,
		Valid: v.Set,
	}.Value()
}

// Int16 wraps V[int16] and sql.NullInt16.
type Int16 struct {
	V[int16]
}

func NewInt16(v int16) (nullV Int16) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Int16) P() *Int16P { return (*Int16P)(v.V.P()) }

func (v *Int16) Scan(value any) (err error) {
	var sqlV sql.NullInt16
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Int16
	v.Set = sqlV.Valid
	return
}

func (v *Int16) Value() (driver.Value, error) {
	return sql.NullInt16{
		Int16: v.Val,
		Valid: v.Set,
	}.Value()
}

// Int32 wraps V[int32] and sql.NullInt32.
type Int32 struct {
	V[int32]
}

func NewInt32(v int32) (nullV Int32) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Int32) P() *Int32P { return (*Int32P)(v.V.P()) }

func (v *Int32) Scan(value any) (err error) {
	var sqlV sql.NullInt32
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Int32
	v.Set = sqlV.Valid
	return
}

func (v *Int32) Value() (driver.Value, error) {
	return sql.NullInt32{
		Int32: v.Val,
		Valid: v.Set,
	}.Value()
}

// Int64 wraps V[int64] and sql.NullInt64.
type Int64 struct {
	V[int64]
}

func NewInt64(v int64) (nullV Int64) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Int64) P() *Int64P { return (*Int64P)(v.V.P()) }

func (v *Int64) Scan(value any) (err error) {
	var sqlV sql.NullInt64
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Int64
	v.Set = sqlV.Valid
	return
}

func (v *Int64) Value() (driver.Value, error) {
	return sql.NullInt64{
		Int64: v.Val,
		Valid: v.Set,
	}.Value()
}

// Str wraps V[string] and sql.NullString.
type Str struct {
	V[string]
}

func NewStr(v string) (nullV Str) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Str) P() *StrP { return (*StrP)(v.V.P()) }

func (v *Str) Scan(value any) (err error) {
	var sqlV sql.NullString
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.String
	v.Set = sqlV.Valid
	return
}

func (v *Str) Value() (driver.Value, error) {
	return sql.NullString{
		String: v.Val,
		Valid:  v.Set,
	}.Value()
}

// Time wraps V[time.Time] and sql.NullTime.
type Time struct {
	V[time.Time]
}

func NewTime(v time.Time) (nullV Time) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *Time) P() *TimeP { return (*TimeP)(v.V.P()) }

func (v *Time) Scan(value any) (err error) {
	var sqlV sql.NullTime
	if err = sqlV.Scan(value); err != nil {
		return
	}
	v.Val = sqlV.Time
	v.Set = sqlV.Valid
	return
}

func (v *Time) Value() (driver.Value, error) {
	return sql.NullTime{
		Time:  v.Val,
		Valid: v.Set,
	}.Value()
}

type ScannerValuer interface {
	sql.Scanner
	driver.Valuer
}

// DBType wraps V[ScannerValuer], sql.Scanner and driver.Valuer.
type DBType[
	T any,
	pT interface{
		*T
		ScannerValuer
	},
] struct {
	V[T]
}

func NewDBType[
	T any,
	pT interface{
		*T
		ScannerValuer
	},
](v T) (nullV DBType[T, pT]) {
	nullV.Val = v
	nullV.Set = true
	return
}

func (v *DBType[T, pT]) Scan(value any) (err error) {
	if err = (pT)(&v.Val).Scan(value); err != nil {
		return
	}
	v.Set = true
	return
}

func (v *DBType[T, pT]) Value() (driver.Value, error) {
	if !v.Set {
		return nil, nil
	}
	return (pT)(&v.Val).Value()
}
