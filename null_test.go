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
	var request struct {
		PhoneNumber Str `json:"phone_number"`
		EmailAddress *string `json:"email_address"`
		Password string `json:"password"`
	}
	if err := json.Unmarshal(
		[]byte(`{
			"email_address": "postmaster@example.com",
			"password": "process candidates rankings farming ministries"
		}`),
		&request,
	); err != nil {
		return
	}

	type UserInfo struct {
		FirstName Str
		LastName Str
	}
	exampleFirstName := "Giorgi"
	exampleLastName := "Kobakhidze"
	user1 := UserInfo{
		NewStr(&exampleFirstName),
		NewStr(&exampleLastName),
	}

	userByPhoneNumber := map[string]UserInfo{"+12065550100": user1}
	userByEmail := map[string]UserInfo{"postmaster@example.com": user1}

	var (
		userInfo UserInfo
		ok bool
	)
	if request.PhoneNumber.Set == (request.EmailAddress != nil) {
		return
	} else if request.PhoneNumber.Set {
		userInfo, ok = userByPhoneNumber[request.PhoneNumber.Val]
	} else {
		userInfo, ok = userByEmail[NewStr(request.EmailAddress).Val]
	}
	if !ok {
		return
	}

	responseBytes, err := json.Marshal(
		struct {
			FirstName *string `json:"first_name,omitempty"`
			LastName *string `json:"last_name,omitempty"`
		} {
			userInfo.FirstName.Ptr(),
			userInfo.LastName.Ptr(),
		},
	)
	if err != nil {
		return
	}

	fmt.Println(string(responseBytes))
	// Output:
	// {"first_name":"Giorgi","last_name":"Kobakhidze"}
}
