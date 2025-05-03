/*
Macrame is a program that enables the user to create keyboard macros and button panels. 
The macros are saved as simple JSON files and can be linked to the button panels. The panels can 
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation, either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package structs

type Allowed struct {
	Local  []string
	Remote []string
	Auth   []string
}

var Endpoints = Allowed{
	Local: []string{
		"/macro/check",
		"/macro/record",
		"/macro/list",
		"/macro/open",
		"/macro/delete",
		"/macro/play",
		"/device/server/ip",
		"/device/list",
		"/device/access/check",
		"/device/access/request",
		"/device/link/ping",
		"/device/link/start",
		"/device/link/poll",
		"/device/link/remove",
		"/device/handshake",
		"/panel/get",
		"/panel/list",
		"/panel/save/json",
	},
	Remote: []string{
		"/device/access/check",
		"/device/access/request",
		"/device/server/ip",
		"/device/link/ping",
		"/device/link/end",
		"/device/handshake",
		"/device/auth",
	},
	Auth: []string{
		"/macro/play",
		"/device/link/remove",
		"/panel/get",
		"/panel/list",
	},
}
