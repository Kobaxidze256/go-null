/*
	Copyright (C) 2023 Giorgi Kobakhidze

	This file is part of go-null.

	go-null is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, version 3 of the License.

	go-null is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.

	You should have received a copy of the GNU Lesser General Public License along with go-null. If not, see <https://www.gnu.org/licenses/>.
*/

// Package null provides types and functions
// for representing absent values
// without pointers, while being compatible
// with pointers, encoding/json, database/sql
// and github.com/go-playground/validator.
//
// Null values of these types shall be null.
//
// Warning: deep copies are not made.
package null

import (
	"encoding/json"
)

func pVal[T any](p *T) (v T) {
	if p == nil {
		return
	}
	return *p
}

func pIsSet[T any](p *T) bool {
	return p != nil
}

func pV[T any](p *T) (v V[T]) {
	if p == nil {
		return
	}
	v.Val = *p
	v.Set = true
	return
}

// V can be used to create new nullable values
// that work with pointers and encoding/json.
// T must support json marshaling and unmarshaling.
type V[T any] struct {
	Val T
	Set bool
}

func (v *V[T]) MarshalJSON() ([]byte, error) {
	if !v.Set {
		return []byte("null"), nil
	}
	return json.Marshal(v.Val)
}

func (v *V[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		v.Set = false
		return nil
	}
	v.Set = true
	return json.Unmarshal(data, &v.Val)
}

func (v *V[T]) P() *T {
	if !v.Set {
		return nil
	}
	return &v.Val
}
