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
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"strings"
)

func getSystemKeyPairs() map[string]string {
	systemPairs := make(map[string]string)
	for _, info := range systemIdDescList {
		if info.Id >= 0 && info.Id < 300 {
			if isInvalidConflict(info.Id) {
				continue
			}
			shortcut := getSystemKeyValue(info.Name, false)
			action := getSystemKeyValue(info.Name, true)
			systemPairs[shortcut] = action
		}
	}
	PrevSystemPairs = systemPairs
	return systemPairs
}

func getCustomKeyPairs() map[string]string {
	customPairs := make(map[string]string)
	customList := getCustomIdList()

	for _, i := range customList {
		if isInvalidConflict(i) {
			logger.Warningf("%d is invalid conflict", i)
			continue
		}

		gs := getSettingsById(i)
		if gs == nil {
			logger.Warningf("Get Settings For '%d' Failed", i)
			continue
		}
		shortcut := gs.GetString(CUSTOM_KEY_SHORTCUT)
		action := gs.GetString(CUSTOM_KEY_ACTION)
		customPairs[shortcut] = action
	}

	PrevCustomPairs = customPairs
	return customPairs
}

func grabKeyPress(wid xproto.Window, shortcut string) bool {
	if len(shortcut) < 1 {
		logger.Warning("grabKeyPress args error...")
		return false
	}

	mod, keys, err := keybind.ParseString(X, shortcut)
	if err != nil {
		logger.Warning("In GrabKey Parse shortcut failed:", err)
		return false
	}

	for _, k := range keys {
		if err := keybind.GrabChecked(X, wid, mod, k); err != nil {
			logger.Warning("GrabChecked failed:", err)
			ungrabKey(wid, shortcut)
			return false
		}
	}

	return true
}

func ungrabKey(wid xproto.Window, shortcut string) bool {
	if len(shortcut) < 1 {
		logger.Warning("Ungrab args failed...")
		return false
	}

	mod, keys, err := keybind.ParseString(X, shortcut)
	if err != nil {
		logger.Warning("In UngrabKey Parse shortcut failed:", err)
		return false
	}

	for _, k := range keys {
		keybind.Ungrab(X, wid, mod, k)
	}

	return true
}

func grabKeyPairs(pairs map[string]string, isGrab bool) {
	for k, v := range pairs {
		if len(k) < 1 {
			continue
		}

		logger.Infof("Grab key pair: %v --- %v", k, v)
		if strings.ToLower(k) == "super" {
			if isGrab {
				keyInfo, ok := newKeycodeInfo("Super_L")
				if !ok {
					continue
				}
				if grabKeyPress(X.RootWin(), "Super_L") {
					grabKeyBindsMap[keyInfo] = v
				}

				keyInfo, ok = newKeycodeInfo("Super_R")
				if !ok {
					continue
				}
				if grabKeyPress(X.RootWin(), "Super_R") {
					grabKeyBindsMap[keyInfo] = v
				}
			} else {
				keyInfo, ok := newKeycodeInfo("Super_L")
				if !ok {
					continue
				}
				if ungrabKey(X.RootWin(), "Super_L") {
					delete(grabKeyBindsMap, keyInfo)
				}

				keyInfo, ok = newKeycodeInfo("Super_R")
				if !ok {
					continue
				}
				if ungrabKey(X.RootWin(), "Super_R") {
					delete(grabKeyBindsMap, keyInfo)
				}
			}
			continue
		}

		shortcut := formatXGBShortcut(formatShortcut(k))
		keyInfo, ok := newKeycodeInfo(shortcut)
		if !ok {
			logger.Warningf("New Keycode Info Failed. Key: %s, Value: %s", k, v)
			continue
		}

		if isGrab {
			if grabKeyPress(X.RootWin(), shortcut) {
				grabKeyBindsMap[keyInfo] = v
			}
		} else {
			if ungrabKey(X.RootWin(), shortcut) {
				delete(grabKeyBindsMap, keyInfo)
			}
		}
	}
}

func grabMediaKeys() {
	keyList := mediaGSettings.ListKeys()
	for _, key := range keyList {
		value := mediaGSettings.GetString(key)
		grabKeyPress(X.RootWin(), convertKeysToMods(value))
	}
}

var keyPressStr string

func grabKeyboardAndMouse() {
	X, err := xgbutil.NewConn()
	if err != nil {
		logger.Info("Get New Connection Failed:", err)
		return
	}
	keybind.Initialize(X)
	mousebind.Initialize(X)

	err = keybind.GrabKeyboard(X, X.RootWin())
	if err != nil {
		logger.Info("Grab Keyboard Failed:", err)
		return
	}

	grabAllMouseButton(X)

	xevent.ButtonPressFun(
		func(X *xgbutil.XUtil, e xevent.ButtonPressEvent) {
			GetManager().KeyReleaseEvent("")
			ungrabAllMouseButton(X)
			keybind.UngrabKeyboard(X)
			logger.Info("Button Press Event")
			xevent.Quit(X)
		}).Connect(X, X.RootWin())

	xevent.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			value := parseKeyEnvent(X, e.State, e.Detail)
			keyPressStr = value
			GetManager().KeyPressEvent(value)
		}).Connect(X, X.RootWin())

	xevent.KeyReleaseFun(
		func(X *xgbutil.XUtil, e xevent.KeyReleaseEvent) {
			GetManager().KeyReleaseEvent(keyPressStr)
			keyPressStr = ""
			ungrabAllMouseButton(X)
			keybind.UngrabKeyboard(X)
			xevent.Quit(X)
		}).Connect(X, X.RootWin())

	xevent.Main(X)
}

func grabAllMouseButton(X *xgbutil.XUtil) {
	mousebind.Grab(X, X.RootWin(), 0, 1, false)
	mousebind.Grab(X, X.RootWin(), 0, 2, false)
	mousebind.Grab(X, X.RootWin(), 0, 3, false)
}

func ungrabAllMouseButton(X *xgbutil.XUtil) {
	mousebind.Ungrab(X, X.RootWin(), 0, 1)
	mousebind.Ungrab(X, X.RootWin(), 0, 2)
	mousebind.Ungrab(X, X.RootWin(), 0, 3)
}

func parseKeyEnvent(X *xgbutil.XUtil, state uint16, detail xproto.Keycode) string {
	modStr := keybind.ModifierString(state)
	keyStr := keybind.LookupString(X, state, detail)

	if detail == 65 {
		keyStr = "Space"
	}

	if keyStr == "L1" {
		keyStr = "F11"
	}

	if keyStr == "L2" {
		keyStr = "F12"
	}

	value := ""
	modStr = deleteSpecialMod(modStr)

	if len(modStr) > 0 {
		value = convertModsToKeys(modStr) + "-" + keyStr
	} else {
		value = keyStr
	}

	return value
}

func parseModifierKey(mod, key string) string {
	tmp := strings.ToLower(key)
	if strings.Contains(tmp, "control") {
		return "control"
	} else if strings.Contains(tmp, "alt") {
		return "alt"
	} else if strings.Contains(tmp, "shift") {
		return "shift"
	} else if strings.Contains(tmp, "super") {
		if len(mod) > 0 {
			return "super"
		}
	}

	return key
}
