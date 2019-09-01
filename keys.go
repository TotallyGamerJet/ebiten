// Copyright 2013 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

package ebiten

import (
	"strings"

	"github.com/hajimehoshi/ebiten/internal/driver"
)

// A Key represents a keyboard key.
// These keys represent pysical keys of US keyboard.
// For example, KeyQ represents Q key on US keyboards and ' (quote) key on Dvorak keyboards.
type Key int

// Keys.
const (
	Key0            Key = Key(driver.Key0)
	Key1            Key = Key(driver.Key1)
	Key2            Key = Key(driver.Key2)
	Key3            Key = Key(driver.Key3)
	Key4            Key = Key(driver.Key4)
	Key5            Key = Key(driver.Key5)
	Key6            Key = Key(driver.Key6)
	Key7            Key = Key(driver.Key7)
	Key8            Key = Key(driver.Key8)
	Key9            Key = Key(driver.Key9)
	KeyA            Key = Key(driver.KeyA)
	KeyB            Key = Key(driver.KeyB)
	KeyC            Key = Key(driver.KeyC)
	KeyD            Key = Key(driver.KeyD)
	KeyE            Key = Key(driver.KeyE)
	KeyF            Key = Key(driver.KeyF)
	KeyG            Key = Key(driver.KeyG)
	KeyH            Key = Key(driver.KeyH)
	KeyI            Key = Key(driver.KeyI)
	KeyJ            Key = Key(driver.KeyJ)
	KeyK            Key = Key(driver.KeyK)
	KeyL            Key = Key(driver.KeyL)
	KeyM            Key = Key(driver.KeyM)
	KeyN            Key = Key(driver.KeyN)
	KeyO            Key = Key(driver.KeyO)
	KeyP            Key = Key(driver.KeyP)
	KeyQ            Key = Key(driver.KeyQ)
	KeyR            Key = Key(driver.KeyR)
	KeyS            Key = Key(driver.KeyS)
	KeyT            Key = Key(driver.KeyT)
	KeyU            Key = Key(driver.KeyU)
	KeyV            Key = Key(driver.KeyV)
	KeyW            Key = Key(driver.KeyW)
	KeyX            Key = Key(driver.KeyX)
	KeyY            Key = Key(driver.KeyY)
	KeyZ            Key = Key(driver.KeyZ)
	KeyApostrophe   Key = Key(driver.KeyApostrophe)
	KeyBackslash    Key = Key(driver.KeyBackslash)
	KeyBackspace    Key = Key(driver.KeyBackspace)
	KeyCapsLock     Key = Key(driver.KeyCapsLock)
	KeyComma        Key = Key(driver.KeyComma)
	KeyDelete       Key = Key(driver.KeyDelete)
	KeyDown         Key = Key(driver.KeyDown)
	KeyEnd          Key = Key(driver.KeyEnd)
	KeyEnter        Key = Key(driver.KeyEnter)
	KeyEqual        Key = Key(driver.KeyEqual)
	KeyEscape       Key = Key(driver.KeyEscape)
	KeyF1           Key = Key(driver.KeyF1)
	KeyF2           Key = Key(driver.KeyF2)
	KeyF3           Key = Key(driver.KeyF3)
	KeyF4           Key = Key(driver.KeyF4)
	KeyF5           Key = Key(driver.KeyF5)
	KeyF6           Key = Key(driver.KeyF6)
	KeyF7           Key = Key(driver.KeyF7)
	KeyF8           Key = Key(driver.KeyF8)
	KeyF9           Key = Key(driver.KeyF9)
	KeyF10          Key = Key(driver.KeyF10)
	KeyF11          Key = Key(driver.KeyF11)
	KeyF12          Key = Key(driver.KeyF12)
	KeyGraveAccent  Key = Key(driver.KeyGraveAccent)
	KeyHome         Key = Key(driver.KeyHome)
	KeyInsert       Key = Key(driver.KeyInsert)
	KeyKP0          Key = Key(driver.KeyKP0)
	KeyKP1          Key = Key(driver.KeyKP1)
	KeyKP2          Key = Key(driver.KeyKP2)
	KeyKP3          Key = Key(driver.KeyKP3)
	KeyKP4          Key = Key(driver.KeyKP4)
	KeyKP5          Key = Key(driver.KeyKP5)
	KeyKP6          Key = Key(driver.KeyKP6)
	KeyKP7          Key = Key(driver.KeyKP7)
	KeyKP8          Key = Key(driver.KeyKP8)
	KeyKP9          Key = Key(driver.KeyKP9)
	KeyKPAdd        Key = Key(driver.KeyKPAdd)
	KeyKPDecimal    Key = Key(driver.KeyKPDecimal)
	KeyKPDivide     Key = Key(driver.KeyKPDivide)
	KeyKPEnter      Key = Key(driver.KeyKPEnter)
	KeyKPEqual      Key = Key(driver.KeyKPEqual)
	KeyKPMultiply   Key = Key(driver.KeyKPMultiply)
	KeyKPSubtract   Key = Key(driver.KeyKPSubtract)
	KeyLeft         Key = Key(driver.KeyLeft)
	KeyLeftBracket  Key = Key(driver.KeyLeftBracket)
	KeyMenu         Key = Key(driver.KeyMenu)
	KeyMinus        Key = Key(driver.KeyMinus)
	KeyNumLock      Key = Key(driver.KeyNumLock)
	KeyPageDown     Key = Key(driver.KeyPageDown)
	KeyPageUp       Key = Key(driver.KeyPageUp)
	KeyPause        Key = Key(driver.KeyPause)
	KeyPeriod       Key = Key(driver.KeyPeriod)
	KeyPrintScreen  Key = Key(driver.KeyPrintScreen)
	KeyRight        Key = Key(driver.KeyRight)
	KeyRightBracket Key = Key(driver.KeyRightBracket)
	KeyScrollLock   Key = Key(driver.KeyScrollLock)
	KeySemicolon    Key = Key(driver.KeySemicolon)
	KeySlash        Key = Key(driver.KeySlash)
	KeySpace        Key = Key(driver.KeySpace)
	KeyTab          Key = Key(driver.KeyTab)
	KeyUp           Key = Key(driver.KeyUp)
	KeyAlt          Key = Key(driver.KeyReserved0)
	KeyControl      Key = Key(driver.KeyReserved1)
	KeyShift        Key = Key(driver.KeyReserved2)
	KeyMax          Key = KeyShift
)

func (k Key) isValid() bool {
	switch k {
	case Key0:
		return true
	case Key1:
		return true
	case Key2:
		return true
	case Key3:
		return true
	case Key4:
		return true
	case Key5:
		return true
	case Key6:
		return true
	case Key7:
		return true
	case Key8:
		return true
	case Key9:
		return true
	case KeyA:
		return true
	case KeyB:
		return true
	case KeyC:
		return true
	case KeyD:
		return true
	case KeyE:
		return true
	case KeyF:
		return true
	case KeyG:
		return true
	case KeyH:
		return true
	case KeyI:
		return true
	case KeyJ:
		return true
	case KeyK:
		return true
	case KeyL:
		return true
	case KeyM:
		return true
	case KeyN:
		return true
	case KeyO:
		return true
	case KeyP:
		return true
	case KeyQ:
		return true
	case KeyR:
		return true
	case KeyS:
		return true
	case KeyT:
		return true
	case KeyU:
		return true
	case KeyV:
		return true
	case KeyW:
		return true
	case KeyX:
		return true
	case KeyY:
		return true
	case KeyZ:
		return true
	case KeyAlt:
		return true
	case KeyApostrophe:
		return true
	case KeyBackslash:
		return true
	case KeyBackspace:
		return true
	case KeyCapsLock:
		return true
	case KeyComma:
		return true
	case KeyControl:
		return true
	case KeyDelete:
		return true
	case KeyDown:
		return true
	case KeyEnd:
		return true
	case KeyEnter:
		return true
	case KeyEqual:
		return true
	case KeyEscape:
		return true
	case KeyF1:
		return true
	case KeyF2:
		return true
	case KeyF3:
		return true
	case KeyF4:
		return true
	case KeyF5:
		return true
	case KeyF6:
		return true
	case KeyF7:
		return true
	case KeyF8:
		return true
	case KeyF9:
		return true
	case KeyF10:
		return true
	case KeyF11:
		return true
	case KeyF12:
		return true
	case KeyGraveAccent:
		return true
	case KeyHome:
		return true
	case KeyInsert:
		return true
	case KeyKP0:
		return true
	case KeyKP1:
		return true
	case KeyKP2:
		return true
	case KeyKP3:
		return true
	case KeyKP4:
		return true
	case KeyKP5:
		return true
	case KeyKP6:
		return true
	case KeyKP7:
		return true
	case KeyKP8:
		return true
	case KeyKP9:
		return true
	case KeyKPAdd:
		return true
	case KeyKPDecimal:
		return true
	case KeyKPDivide:
		return true
	case KeyKPEnter:
		return true
	case KeyKPEqual:
		return true
	case KeyKPMultiply:
		return true
	case KeyKPSubtract:
		return true
	case KeyLeft:
		return true
	case KeyLeftBracket:
		return true
	case KeyMenu:
		return true
	case KeyMinus:
		return true
	case KeyNumLock:
		return true
	case KeyPageDown:
		return true
	case KeyPageUp:
		return true
	case KeyPause:
		return true
	case KeyPeriod:
		return true
	case KeyPrintScreen:
		return true
	case KeyRight:
		return true
	case KeyRightBracket:
		return true
	case KeyScrollLock:
		return true
	case KeySemicolon:
		return true
	case KeyShift:
		return true
	case KeySlash:
		return true
	case KeySpace:
		return true
	case KeyTab:
		return true
	case KeyUp:
		return true

	default:
		return false
	}
}

// String returns a string representing the key.
//
// If k is an undefined key, String returns an empty string.
func (k Key) String() string {
	switch k {
	case Key0:
		return "0"
	case Key1:
		return "1"
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key5:
		return "5"
	case Key6:
		return "6"
	case Key7:
		return "7"
	case Key8:
		return "8"
	case Key9:
		return "9"
	case KeyA:
		return "A"
	case KeyB:
		return "B"
	case KeyC:
		return "C"
	case KeyD:
		return "D"
	case KeyE:
		return "E"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyI:
		return "I"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeyM:
		return "M"
	case KeyN:
		return "N"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyQ:
		return "Q"
	case KeyR:
		return "R"
	case KeyS:
		return "S"
	case KeyT:
		return "T"
	case KeyU:
		return "U"
	case KeyV:
		return "V"
	case KeyW:
		return "W"
	case KeyX:
		return "X"
	case KeyY:
		return "Y"
	case KeyZ:
		return "Z"
	case KeyAlt:
		return "Alt"
	case KeyApostrophe:
		return "Apostrophe"
	case KeyBackslash:
		return "Backslash"
	case KeyBackspace:
		return "Backspace"
	case KeyCapsLock:
		return "CapsLock"
	case KeyComma:
		return "Comma"
	case KeyControl:
		return "Control"
	case KeyDelete:
		return "Delete"
	case KeyDown:
		return "Down"
	case KeyEnd:
		return "End"
	case KeyEnter:
		return "Enter"
	case KeyEqual:
		return "Equal"
	case KeyEscape:
		return "Escape"
	case KeyF1:
		return "F1"
	case KeyF2:
		return "F2"
	case KeyF3:
		return "F3"
	case KeyF4:
		return "F4"
	case KeyF5:
		return "F5"
	case KeyF6:
		return "F6"
	case KeyF7:
		return "F7"
	case KeyF8:
		return "F8"
	case KeyF9:
		return "F9"
	case KeyF10:
		return "F10"
	case KeyF11:
		return "F11"
	case KeyF12:
		return "F12"
	case KeyGraveAccent:
		return "GraveAccent"
	case KeyHome:
		return "Home"
	case KeyInsert:
		return "Insert"
	case KeyKP0:
		return "KP0"
	case KeyKP1:
		return "KP1"
	case KeyKP2:
		return "KP2"
	case KeyKP3:
		return "KP3"
	case KeyKP4:
		return "KP4"
	case KeyKP5:
		return "KP5"
	case KeyKP6:
		return "KP6"
	case KeyKP7:
		return "KP7"
	case KeyKP8:
		return "KP8"
	case KeyKP9:
		return "KP9"
	case KeyKPAdd:
		return "KPAdd"
	case KeyKPDecimal:
		return "KPDecimal"
	case KeyKPDivide:
		return "KPDivide"
	case KeyKPEnter:
		return "KPEnter"
	case KeyKPEqual:
		return "KPEqual"
	case KeyKPMultiply:
		return "KPMultiply"
	case KeyKPSubtract:
		return "KPSubtract"
	case KeyLeft:
		return "Left"
	case KeyLeftBracket:
		return "LeftBracket"
	case KeyMenu:
		return "Menu"
	case KeyMinus:
		return "Minus"
	case KeyNumLock:
		return "NumLock"
	case KeyPageDown:
		return "PageDown"
	case KeyPageUp:
		return "PageUp"
	case KeyPause:
		return "Pause"
	case KeyPeriod:
		return "Period"
	case KeyPrintScreen:
		return "PrintScreen"
	case KeyRight:
		return "Right"
	case KeyRightBracket:
		return "RightBracket"
	case KeyScrollLock:
		return "ScrollLock"
	case KeySemicolon:
		return "Semicolon"
	case KeyShift:
		return "Shift"
	case KeySlash:
		return "Slash"
	case KeySpace:
		return "Space"
	case KeyTab:
		return "Tab"
	case KeyUp:
		return "Up"
	}
	return ""
}

func keyNameToKey(name string) (Key, bool) {
	switch strings.ToLower(name) {
	case "0":
		return Key0, true
	case "1":
		return Key1, true
	case "2":
		return Key2, true
	case "3":
		return Key3, true
	case "4":
		return Key4, true
	case "5":
		return Key5, true
	case "6":
		return Key6, true
	case "7":
		return Key7, true
	case "8":
		return Key8, true
	case "9":
		return Key9, true
	case "a":
		return KeyA, true
	case "b":
		return KeyB, true
	case "c":
		return KeyC, true
	case "d":
		return KeyD, true
	case "e":
		return KeyE, true
	case "f":
		return KeyF, true
	case "g":
		return KeyG, true
	case "h":
		return KeyH, true
	case "i":
		return KeyI, true
	case "j":
		return KeyJ, true
	case "k":
		return KeyK, true
	case "l":
		return KeyL, true
	case "m":
		return KeyM, true
	case "n":
		return KeyN, true
	case "o":
		return KeyO, true
	case "p":
		return KeyP, true
	case "q":
		return KeyQ, true
	case "r":
		return KeyR, true
	case "s":
		return KeyS, true
	case "t":
		return KeyT, true
	case "u":
		return KeyU, true
	case "v":
		return KeyV, true
	case "w":
		return KeyW, true
	case "x":
		return KeyX, true
	case "y":
		return KeyY, true
	case "z":
		return KeyZ, true
	case "alt":
		return KeyAlt, true
	case "apostrophe":
		return KeyApostrophe, true
	case "backslash":
		return KeyBackslash, true
	case "backspace":
		return KeyBackspace, true
	case "capslock":
		return KeyCapsLock, true
	case "comma":
		return KeyComma, true
	case "control":
		return KeyControl, true
	case "delete":
		return KeyDelete, true
	case "down":
		return KeyDown, true
	case "end":
		return KeyEnd, true
	case "enter":
		return KeyEnter, true
	case "equal":
		return KeyEqual, true
	case "escape":
		return KeyEscape, true
	case "f1":
		return KeyF1, true
	case "f2":
		return KeyF2, true
	case "f3":
		return KeyF3, true
	case "f4":
		return KeyF4, true
	case "f5":
		return KeyF5, true
	case "f6":
		return KeyF6, true
	case "f7":
		return KeyF7, true
	case "f8":
		return KeyF8, true
	case "f9":
		return KeyF9, true
	case "f10":
		return KeyF10, true
	case "f11":
		return KeyF11, true
	case "f12":
		return KeyF12, true
	case "graveaccent":
		return KeyGraveAccent, true
	case "home":
		return KeyHome, true
	case "insert":
		return KeyInsert, true
	case "kp0":
		return KeyKP0, true
	case "kp1":
		return KeyKP1, true
	case "kp2":
		return KeyKP2, true
	case "kp3":
		return KeyKP3, true
	case "kp4":
		return KeyKP4, true
	case "kp5":
		return KeyKP5, true
	case "kp6":
		return KeyKP6, true
	case "kp7":
		return KeyKP7, true
	case "kp8":
		return KeyKP8, true
	case "kp9":
		return KeyKP9, true
	case "kpadd":
		return KeyKPAdd, true
	case "kpdecimal":
		return KeyKPDecimal, true
	case "kpdivide":
		return KeyKPDivide, true
	case "kpenter":
		return KeyKPEnter, true
	case "kpequal":
		return KeyKPEqual, true
	case "kpmultiply":
		return KeyKPMultiply, true
	case "kpsubtract":
		return KeyKPSubtract, true
	case "left":
		return KeyLeft, true
	case "leftbracket":
		return KeyLeftBracket, true
	case "menu":
		return KeyMenu, true
	case "minus":
		return KeyMinus, true
	case "numlock":
		return KeyNumLock, true
	case "pagedown":
		return KeyPageDown, true
	case "pageup":
		return KeyPageUp, true
	case "pause":
		return KeyPause, true
	case "period":
		return KeyPeriod, true
	case "printscreen":
		return KeyPrintScreen, true
	case "right":
		return KeyRight, true
	case "rightbracket":
		return KeyRightBracket, true
	case "scrolllock":
		return KeyScrollLock, true
	case "semicolon":
		return KeySemicolon, true
	case "shift":
		return KeyShift, true
	case "slash":
		return KeySlash, true
	case "space":
		return KeySpace, true
	case "tab":
		return KeyTab, true
	case "up":
		return KeyUp, true
	}
	return 0, false
}
