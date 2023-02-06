// SPDX-License-Identifier: Zlib
// SPDX-FileCopyrightText: 2002-2006 Marcus Geelnard
// SPDX-FileCopyrightText: 2006-2019 Camilla Löwy
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

package goglfw

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/hajimehoshi/ebiten/v2/internal/microsoftgdk"
)

func createKeyTables() {
	for i := range _glfw.windowState.keycodes {
		_glfw.windowState.keycodes[i] = -1
	}
	for i := range _glfw.windowState.scancodes {
		_glfw.windowState.keycodes[i] = -1
	}

	_glfw.windowState.keycodes[0x00B] = Key0
	_glfw.windowState.keycodes[0x002] = Key1
	_glfw.windowState.keycodes[0x003] = Key2
	_glfw.windowState.keycodes[0x004] = Key3
	_glfw.windowState.keycodes[0x005] = Key4
	_glfw.windowState.keycodes[0x006] = Key5
	_glfw.windowState.keycodes[0x007] = Key6
	_glfw.windowState.keycodes[0x008] = Key7
	_glfw.windowState.keycodes[0x009] = Key8
	_glfw.windowState.keycodes[0x00A] = Key9
	_glfw.windowState.keycodes[0x01E] = KeyA
	_glfw.windowState.keycodes[0x030] = KeyB
	_glfw.windowState.keycodes[0x02E] = KeyC
	_glfw.windowState.keycodes[0x020] = KeyD
	_glfw.windowState.keycodes[0x012] = KeyE
	_glfw.windowState.keycodes[0x021] = KeyF
	_glfw.windowState.keycodes[0x022] = KeyG
	_glfw.windowState.keycodes[0x023] = KeyH
	_glfw.windowState.keycodes[0x017] = KeyI
	_glfw.windowState.keycodes[0x024] = KeyJ
	_glfw.windowState.keycodes[0x025] = KeyK
	_glfw.windowState.keycodes[0x026] = KeyL
	_glfw.windowState.keycodes[0x032] = KeyM
	_glfw.windowState.keycodes[0x031] = KeyN
	_glfw.windowState.keycodes[0x018] = KeyO
	_glfw.windowState.keycodes[0x019] = KeyP
	_glfw.windowState.keycodes[0x010] = KeyQ
	_glfw.windowState.keycodes[0x013] = KeyR
	_glfw.windowState.keycodes[0x01F] = KeyS
	_glfw.windowState.keycodes[0x014] = KeyT
	_glfw.windowState.keycodes[0x016] = KeyU
	_glfw.windowState.keycodes[0x02F] = KeyV
	_glfw.windowState.keycodes[0x011] = KeyW
	_glfw.windowState.keycodes[0x02D] = KeyX
	_glfw.windowState.keycodes[0x015] = KeyY
	_glfw.windowState.keycodes[0x02C] = KeyZ

	_glfw.windowState.keycodes[0x028] = KeyApostrophe
	_glfw.windowState.keycodes[0x02B] = KeyBackslash
	_glfw.windowState.keycodes[0x033] = KeyComma
	_glfw.windowState.keycodes[0x00D] = KeyEqual
	_glfw.windowState.keycodes[0x029] = KeyGraveAccent
	_glfw.windowState.keycodes[0x01A] = KeyLeftBracket
	_glfw.windowState.keycodes[0x00C] = KeyMinus
	_glfw.windowState.keycodes[0x034] = KeyPeriod
	_glfw.windowState.keycodes[0x01B] = KeyRightBracket
	_glfw.windowState.keycodes[0x027] = KeySemicolon
	_glfw.windowState.keycodes[0x035] = KeySlash
	_glfw.windowState.keycodes[0x056] = KeyWorld2

	_glfw.windowState.keycodes[0x00E] = KeyBackspace
	_glfw.windowState.keycodes[0x153] = KeyDelete
	_glfw.windowState.keycodes[0x14F] = KeyEnd
	_glfw.windowState.keycodes[0x01C] = KeyEnter
	_glfw.windowState.keycodes[0x001] = KeyEscape
	_glfw.windowState.keycodes[0x147] = KeyHome
	_glfw.windowState.keycodes[0x152] = KeyInsert
	_glfw.windowState.keycodes[0x15D] = KeyMenu
	_glfw.windowState.keycodes[0x151] = KeyPageDown
	_glfw.windowState.keycodes[0x149] = KeyPageUp
	_glfw.windowState.keycodes[0x045] = KeyPause
	_glfw.windowState.keycodes[0x039] = KeySpace
	_glfw.windowState.keycodes[0x00F] = KeyTab
	_glfw.windowState.keycodes[0x03A] = KeyCapsLock
	_glfw.windowState.keycodes[0x145] = KeyNumLock
	_glfw.windowState.keycodes[0x046] = KeyScrollLock
	_glfw.windowState.keycodes[0x03B] = KeyF1
	_glfw.windowState.keycodes[0x03C] = KeyF2
	_glfw.windowState.keycodes[0x03D] = KeyF3
	_glfw.windowState.keycodes[0x03E] = KeyF4
	_glfw.windowState.keycodes[0x03F] = KeyF5
	_glfw.windowState.keycodes[0x040] = KeyF6
	_glfw.windowState.keycodes[0x041] = KeyF7
	_glfw.windowState.keycodes[0x042] = KeyF8
	_glfw.windowState.keycodes[0x043] = KeyF9
	_glfw.windowState.keycodes[0x044] = KeyF10
	_glfw.windowState.keycodes[0x057] = KeyF11
	_glfw.windowState.keycodes[0x058] = KeyF12
	_glfw.windowState.keycodes[0x064] = KeyF13
	_glfw.windowState.keycodes[0x065] = KeyF14
	_glfw.windowState.keycodes[0x066] = KeyF15
	_glfw.windowState.keycodes[0x067] = KeyF16
	_glfw.windowState.keycodes[0x068] = KeyF17
	_glfw.windowState.keycodes[0x069] = KeyF18
	_glfw.windowState.keycodes[0x06A] = KeyF19
	_glfw.windowState.keycodes[0x06B] = KeyF20
	_glfw.windowState.keycodes[0x06C] = KeyF21
	_glfw.windowState.keycodes[0x06D] = KeyF22
	_glfw.windowState.keycodes[0x06E] = KeyF23
	_glfw.windowState.keycodes[0x076] = KeyF24
	_glfw.windowState.keycodes[0x038] = KeyLeftAlt
	_glfw.windowState.keycodes[0x01D] = KeyLeftControl
	_glfw.windowState.keycodes[0x02A] = KeyLeftShift
	_glfw.windowState.keycodes[0x15B] = KeyLeftSuper
	_glfw.windowState.keycodes[0x137] = KeyPrintScreen
	_glfw.windowState.keycodes[0x138] = KeyRightAlt
	_glfw.windowState.keycodes[0x11D] = KeyRightControl
	_glfw.windowState.keycodes[0x036] = KeyRightShift
	_glfw.windowState.keycodes[0x15C] = KeyRightSuper
	_glfw.windowState.keycodes[0x150] = KeyDown
	_glfw.windowState.keycodes[0x14B] = KeyLeft
	_glfw.windowState.keycodes[0x14D] = KeyRight
	_glfw.windowState.keycodes[0x148] = KeyUp

	_glfw.windowState.keycodes[0x052] = KeyKP0
	_glfw.windowState.keycodes[0x04F] = KeyKP1
	_glfw.windowState.keycodes[0x050] = KeyKP2
	_glfw.windowState.keycodes[0x051] = KeyKP3
	_glfw.windowState.keycodes[0x04B] = KeyKP4
	_glfw.windowState.keycodes[0x04C] = KeyKP5
	_glfw.windowState.keycodes[0x04D] = KeyKP6
	_glfw.windowState.keycodes[0x047] = KeyKP7
	_glfw.windowState.keycodes[0x048] = KeyKP8
	_glfw.windowState.keycodes[0x049] = KeyKP9
	_glfw.windowState.keycodes[0x04E] = KeyKPAdd
	_glfw.windowState.keycodes[0x053] = KeyKPDecimal
	_glfw.windowState.keycodes[0x135] = KeyKPDivide
	_glfw.windowState.keycodes[0x11C] = KeyKPEnter
	_glfw.windowState.keycodes[0x059] = KeyKPEqual
	_glfw.windowState.keycodes[0x037] = KeyKPMultiply
	_glfw.windowState.keycodes[0x04A] = KeyKPSubtract

	for scancode := 0; scancode < 512; scancode++ {
		if _glfw.windowState.keycodes[scancode] > 0 {
			_glfw.windowState.scancodes[_glfw.windowState.keycodes[scancode]] = scancode
		}
	}
}

func updateKeyNamesWin32() {
	// MapVirtualKeyW is not implemented in Xbox.
	if microsoftgdk.IsXbox() {
		return
	}

	for i := range _glfw.windowState.keynames {
		_glfw.windowState.keynames[i] = ""
	}

	var state [256]byte

	for key := KeySpace; key <= KeyLast; key++ {
		scancode := _glfw.windowState.scancodes[key]
		if scancode == -1 {
			continue
		}

		var vk uint32
		if key >= KeyKP0 && key <= KeyKPAdd {
			vks := []uint32{
				_VK_NUMPAD0, _VK_NUMPAD1, _VK_NUMPAD2, _VK_NUMPAD3,
				_VK_NUMPAD4, _VK_NUMPAD5, _VK_NUMPAD6, _VK_NUMPAD7,
				_VK_NUMPAD8, _VK_NUMPAD9, _VK_DECIMAL, _VK_DIVIDE,
				_VK_MULTIPLY, _VK_SUBTRACT, _VK_ADD,
			}
			vk = vks[key-KeyKP0]
		} else {
			vk = _MapVirtualKeyW(uint32(scancode), _MAPVK_VSC_TO_VK)
		}

		var chars [16]uint16
		length := _ToUnicode(vk, uint32(scancode), state[:], chars[:], int32(len(chars)), 0)
		if length == -1 {
			// This is a dead key, so we need a second simulated key press
			// to make it output its own character (usually a diacritic)
			length = _ToUnicode(vk, uint32(scancode), state[:], chars[:], int32(len(chars)), 0)
		}
		if length < 1 {
			continue
		}

		_glfw.windowState.keynames[key] = windows.UTF16ToString(chars[:length])
	}
}

func createHelperWindow() error {
	h, err := _CreateWindowExW(_WS_EX_OVERLAPPEDWINDOW, _GLFW_WNDCLASSNAME, "GLFW message window", _WS_CLIPSIBLINGS|_WS_CLIPCHILDREN, 0, 0, 1, 1, 0, 0, _glfw.windowState.instance, nil)
	if err != nil {
		return err
	}

	_glfw.windowState.helperWindowHandle = h

	// HACK: The command to the first ShowWindow call is ignored if the parent
	//       process passed along a STARTUPINFO, so clear that with a no-op call
	_ShowWindow(_glfw.windowState.helperWindowHandle, _SW_HIDE)

	// Register for HID device notifications
	if !microsoftgdk.IsXbox() {
		_GUID_DEVINTERFACE_HID := windows.GUID{
			Data1: 0x4d1e55b2,
			Data2: 0xf16f,
			Data3: 0x11cf,
			Data4: [...]byte{0x88, 0xcb, 0x00, 0x11, 0x11, 0x00, 0x00, 0x30},
		}

		var dbi _DEV_BROADCAST_DEVICEINTERFACE_W
		dbi.dbcc_size = uint32(unsafe.Sizeof(dbi))
		dbi.dbcc_devicetype = _DBT_DEVTYP_DEVICEINTERFACE
		dbi.dbcc_classguid = _GUID_DEVINTERFACE_HID
		notify, err := _RegisterDeviceNotificationW(windows.Handle(_glfw.windowState.helperWindowHandle), unsafe.Pointer(&dbi), _DEVICE_NOTIFY_WINDOW_HANDLE)
		if err != nil {
			return err
		}
		_glfw.windowState.deviceNotificationHandle = notify
	}

	var msg _MSG
	for _PeekMessageW(&msg, _glfw.windowState.helperWindowHandle, 0, 0, _PM_REMOVE) {
		_TranslateMessage(&msg)
		_DispatchMessageW(&msg)
	}

	return nil
}

func isWindowsVersionOrGreaterWin32(major, minor, sp uint16) bool {
	osvi := _OSVERSIONINFOEXW{
		dwMajorVersion:    uint32(major),
		dwMinorVersion:    uint32(minor),
		wServicePackMajor: sp,
	}
	osvi.dwOSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	var mask uint32 = _VER_MAJORVERSION | _VER_MINORVERSION | _VER_SERVICEPACKMAJOR
	cond := _VerSetConditionMask(0, _VER_MAJORVERSION, _VER_GREATER_EQUAL)
	cond = _VerSetConditionMask(cond, _VER_MINORVERSION, _VER_GREATER_EQUAL)
	cond = _VerSetConditionMask(cond, _VER_SERVICEPACKMAJOR, _VER_GREATER_EQUAL)
	// HACK: Use RtlVerifyVersionInfo instead of VerifyVersionInfoW as the
	//       latter lies unless the user knew to embed a non-default manifest
	//       announcing support for Windows 10 via supportedOS GUID
	return _RtlVerifyVersionInfo(&osvi, mask, cond) == 0
}

func isWindows10BuildOrGreaterWin32(build uint16) bool {
	osvi := _OSVERSIONINFOEXW{
		dwMajorVersion: 10,
		dwMinorVersion: 0,
		dwBuildNumber:  uint32(build),
	}
	osvi.dwOSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	var mask uint32 = _VER_MAJORVERSION | _VER_MINORVERSION | _VER_BUILDNUMBER
	cond := _VerSetConditionMask(0, _VER_MAJORVERSION, _VER_GREATER_EQUAL)
	cond = _VerSetConditionMask(cond, _VER_MINORVERSION, _VER_GREATER_EQUAL)
	cond = _VerSetConditionMask(cond, _VER_BUILDNUMBER, _VER_GREATER_EQUAL)

	// HACK: Use RtlVerifyVersionInfo instead of VerifyVersionInfoW as the
	//       latter lies unless the user knew to embed a non-default manifest
	//       announcing support for Windows 10 via supportedOS GUID
	return _RtlVerifyVersionInfo(&osvi, mask, cond) == 0
}

func platformInit() error {
	// Changing the foreground lock timeout was removed from the original code.
	// See https://github.com/glfw/glfw/commit/58b48a3a00d9c2a5ca10cc23069a71d8773cc7a4

	m, err := _GetModuleHandleExW(_GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS|_GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT, unsafe.Pointer(&_glfw))
	if err != nil {
		return err
	}
	_glfw.windowState.instance = _HINSTANCE(m)

	createKeyTables()
	updateKeyNamesWin32()

	if isWindows10CreatorsUpdateOrGreaterWin32() {
		if !microsoftgdk.IsXbox() {
			// Ignore the error as SetProcessDpiAwarenessContext returns an error on Steam Deck (#2113).
			// This seems an issue in Wine and/or Proton, but there is nothing we can do.
			_ = _SetProcessDpiAwarenessContext(_DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2)
		}
	} else if _IsWindows8Point1OrGreater() {
		if err := _SetProcessDpiAwareness(_PROCESS_PER_MONITOR_DPI_AWARE); err != nil && !errors.Is(err, handleError(windows.E_ACCESSDENIED)) {
			return err
		}
	} else if _IsWindowsVistaOrGreater() {
		_SetProcessDPIAware()
	}

	if err := registerWindowClassWin32(); err != nil {
		return err
	}

	if err := createHelperWindow(); err != nil {
		return err
	}
	if microsoftgdk.IsXbox() {
		// On Xbox, APIs to get monitors are not available.
		// Create a pseudo monitor instance instead.
		w, h := microsoftgdk.MonitorResolution()
		mode := &VidMode{
			Width:       w,
			Height:      h,
			RedBits:     8,
			GreenBits:   8,
			BlueBits:    8,
			RefreshRate: 0, // TODO: Is it possible to get an appropriate refresh rate?
		}
		m := &Monitor{
			name:  "Xbox Monitor",
			modes: []*VidMode{mode},
		}
		if err := inputMonitor(m, Connected, _GLFW_INSERT_LAST); err != nil {
			return err
		}
	} else {
		if err := pollMonitorsWin32(); err != nil {
			return err
		}
	}
	return nil
}

func platformTerminate() error {
	if _glfw.windowState.deviceNotificationHandle != 0 {
		if err := _UnregisterDeviceNotification(_glfw.windowState.deviceNotificationHandle); err != nil {
			return err
		}
	}

	if _glfw.windowState.helperWindowHandle != 0 {
		if !microsoftgdk.IsXbox() {
			if err := _DestroyWindow(_glfw.windowState.helperWindowHandle); err != nil {
				return err
			}
		}
	}

	if err := unregisterWindowClassWin32(); err != nil {
		return err
	}

	terminateWGL()

	return nil
}
