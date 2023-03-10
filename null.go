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

// ptrVal returns the zero value of a type
// on nil and the real value otherwise.
func ptrVal[T any](p *T) (v T) {
	if p == nil {
		return
	}
	return *p
}

func ptrIsSet[T any](p *T) bool {
	return p != nil
}
