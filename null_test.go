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
	"github.com/go-playground/validator/v10"
)

func Example() {
	var request struct {
		PhoneNumber  *StrP  `json:"phone_number" validate:"omitempty,e164"`
		EmailAddress *StrP  `json:"email_address" validate:"required_without=PhoneNumber,excluded_with=PhoneNumber,omitempty,email"`
		Password     string `json:"password"`
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
	if err := validator.New().Struct(request); err != nil {
		return
	}

	type UserInfo struct {
		FirstName Str
		LastName  Str
	}
	user1 := UserInfo{
		NewStr("Giorgi"),
		NewStr("Kobakhidze"),
	}

	userByPhoneNumber := map[string]UserInfo{"+12065550100": user1}
	userByEmail := map[string]UserInfo{"postmaster@example.com": user1}

	var (
		userInfo UserInfo
		ok       bool
	)
	if request.PhoneNumber.Set() == request.EmailAddress.Set() {
		return
	} else if request.PhoneNumber.Set() {
		userInfo, ok = userByPhoneNumber[request.PhoneNumber.Val()]
	} else {
		userInfo, ok = userByEmail[request.EmailAddress.Val()]
	}
	if !ok {
		return
	}

	responseBytes, err := json.Marshal(
		struct {
			FirstName *StrP `json:"first_name,omitempty"`
			LastName  *StrP `json:"last_name,omitempty"`
		}{
			userInfo.FirstName.P(),
			userInfo.LastName.P(),
		},
	)
	if err != nil {
		return
	}

	fmt.Println(string(responseBytes))
	// Output:
	// {"first_name":"Giorgi","last_name":"Kobakhidze"}
}
