/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package keybinding

import (
	"strings"
)

func convertKeyToMod(key string) string {
	if v, ok := keyToModMap[key]; ok {
		return v
	}

	return key
}

func convertModToKey(mod string) string {
	if v, ok := modToKeyMap[mod]; ok {
		return v
	}

	return mod
}

func convertKeysToMods(keys string) string {
	array := strings.Split(keys, "-")
	ret := ""
	for i, v := range array {
		if i != 0 {
			ret += "-"
		}
		tmp := convertKeyToMod(v)
		ret += tmp
	}

	return ret
}

func convertModsToKeys(mods string) string {
	array := strings.Split(mods, "-")
	ret := ""
	for i, v := range array {
		if i != 0 {
			ret += "-"
		}
		tmp := convertModToKey(v)
		ret += tmp
	}

	return ret
}

/**
 * Input: <control><alt>t
 * Output: modx-modx-t
 */
func formatXGBShortcut(shortcut string) string {
	if len(shortcut) < 1 {
		return ""
	}

	ret := formatShortcut(shortcut)
	return convertKeysToMods(ret)
}

/**
 * Input: <control><alt>t
 * Output: control-alt-t
 */
func formatShortcut(shortcut string) string {
	l := len(shortcut)
	if l < 1 {
		logger.Warning("formatShortcut args error")
		return ""
	}

	ret := ""
	flag := false
	start := 0
	end := 0

	for i, ch := range shortcut {
		if ch == '<' {
			flag = true
			start = i
			continue
		}

		if ch == '>' && flag {
			end = i
			flag = false

			for j := start + 1; j < end; j++ {
				ret += string(shortcut[j])
			}
			ret += "-"
			continue
		}

		if !flag {
			ret += string(ch)
		}
	}

	// parse 'primary' to 'control'
	array := strings.Split(ret, "-")
	ret = ""
	for i, v := range array {
		str := strings.ToLower(v)
		if str == "primary" || str == "control" {
			// multi control
			if !strings.Contains(ret, "control") {
				if i != 0 {
					ret += "-"
				}
				ret += "control"
			}
			continue
		}

		if i != 0 {
			ret += "-"
		}
		ret += v
	}

	return ret
}

/**
 * delete Num_Lock and Caps_Lock
 */
func deleteSpecialMod(modStr string) string {
	if !strings.Contains(modStr, "-") {
		if modStr == "lock" || modStr == "mod2" {
			return ""
		}

		return modStr
	}

	ret := ""
	strs := strings.Split(modStr, "-")
	l := len(strs)
	for i, s := range strs {
		if s == "lock" || s == "mod2" {
			continue
		}

		if i == l-1 {
			ret += s
			break
		}
		ret += s + "-"
	}

	if length := len(ret); length > 1 {
		if ret[length-1] == '-' {
			ret = ret[0 : length-1]
		}
	}

	return ret
}
