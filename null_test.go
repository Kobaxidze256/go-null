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
	"fmt"
)

func Example() {
	var x struct {
		X *StrP `json:"x,omitempty"`
		Y *StrP `json:"y,omitempty"`
	}
	_ = json.Unmarshal([]byte(`{"x": "x"}`), &x)
	fmt.Println("x.X.IsSet():", x.X.IsSet())
	fmt.Printf("x.X.Val(): %#v\n", x.X.Val())
	fmt.Println("x.X.V().Set:", x.X.V().Set)
	fmt.Printf("x.X.V().Val: %#v\n", x.X.V().Val)
	fmt.Println("x.Y.IsSet():", x.Y.IsSet())
	fmt.Println("x.Y.V().Set:", x.Y.V().Set)

	xXV := x.X.V()
	xYV := x.Y.V()
	yBytes, _ := json.Marshal(
		struct {
			X *StrP `json:"x,omitempty"`
			Y *StrP `json:"y,omitempty"`
			Z *StrP `json:"z,omitempty"`
		}{
			xXV.P(),
			xYV.P(),
			nil,
		},
	)

	fmt.Printf("string(yBytes): %#q\n", string(yBytes))
	// Output:
	// x.X.IsSet(): true
	// x.X.Val(): "x"
	// x.X.V().Set: true
	// x.X.V().Val: "x"
	// x.Y.IsSet(): false
	// x.Y.V().Set: false
	// string(yBytes): `{"x":"x"}`
}
