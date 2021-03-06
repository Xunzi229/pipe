// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package model

// Navigation model.
type Navigation struct {
	Model

	Title      string `gorm:"size:128" json:"title"`
	URL        string `gorm:"size:255" json:"url"`
	IconURL    string `gorm:"size:255" json:"iconURL"`
	OpenMethod string `gorm:"size:32" json:"openMethod"`
	Number     int    `json:"number"` // for sorting

	BlogID uint64 `sql:"index" json:"blogID"`
}

// Navigation open methods.
const (
	NavigationOpenMethodBlank = "_blank"
	NavigationOpenMethodSelf  = "_self"
)
