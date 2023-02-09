/*
	Copyright (C) 2023 Giorgi Kobakhidze

	This file is part of go-null.

	go-null is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, version 3 of the License.

	go-null is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.

	You should have received a copy of the GNU Lesser General Public License along with go-null. If not, see <https://www.gnu.org/licenses/>.
*/

package null

import (
	"encoding/json"
)

// J can be used to create new nullable values
// that work with encoding/json and pointers.
// *T must support json marshaling and unmarshaling.
type J[T any] struct {
	Val   T
	IsSet bool
}

func (v *J[T]) MarshalJSON() ([]byte, error) {
	if !v.IsSet {
		return []byte("null"), nil
	}
	return json.Marshal(&v.Val)
}

func (v *J[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		v.IsSet = false
		return nil
	}
	v.IsSet = true
	return json.Unmarshal(data, &v.Val)
}

func (v *J[T]) Ptr() *T {
	if !v.IsSet {
		return nil
	}
	return &v.Val
}

func NewJ[T any](p *T) (v J[T]) {
	if p == nil {
		return
	}
	v.Val = *p
	v.IsSet = true
	return
}
