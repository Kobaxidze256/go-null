/*
	Copyright (C) 2023 Giorgi Kobakhidze

	This file is part of go-null.

	go-null is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, version 3 of the License.

	go-null is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.

	You should have received a copy of the GNU Lesser General Public License along with go-null. If not, see <https://www.gnu.org/licenses/>.
*/

// Package null provides types and functions
// useful for representing absent values
// without pointers, while being compatible
// with pointers, encoding/json and database/sql.
//
// Null values of these types shall be null.
//
// Warning: deep copies are not made.
package null

import (
	"encoding/json"
)

// V can be used to create new nullable types
// that work with pointers and encoding/json.
// T must support marshaling and unmarshaling.
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

func (v *V[T]) Ptr() *T {
	if !v.Set {
		return nil
	}
	return &v.Val
}

func NewV[T any](ptr *T) (v V[T]) {
	if ptr == nil {
		return
	}
	v = V[T]{
		Val: *ptr,
		Set: true,
	}
	return
}
