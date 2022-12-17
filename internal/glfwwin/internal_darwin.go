package glfwwin

import (
	"github.com/ebitengine/purego/objc"
	"github.com/hajimehoshi/ebiten/v2/internal/cocoa"
	"github.com/hajimehoshi/ebiten/v2/internal/corefoundation"
)

type platformContextState struct {
}

type platformWindowState struct {
	object   objc.ID
	delegate objc.ID
	view     objc.ID
	layer    objc.ID

	maximized bool
	occluded  bool
	retina    bool

	// Cached window properties to filter out duplicate events
	width, height     int
	fbWidth, fbHeight int
	xscale, yscale    float32

	// The total sum of the distances the cursor has been warped
	// since the last cursor motion event was processed
	// This is kept to counteract Cocoa doing the same internally
	cursorWarpDeltaX, cursorWarpDeltaY float64
}

type platformMonitorState struct {
	displayID           _CGDirectDisplayID
	previousMod         _CGDisplayModeRef
	unitNumber          uint32
	screen              cocoa.NSScreen
	fallbackRefreshRate float64
}

type platformCursorState struct {
	object objc.ID
}

type platformTLSState struct {
	allocated bool
	key       pthread_key
}

type platformLibraryState struct {
	//    CGEventSourceRef    eventSource;
	delegate objc.ID
	//    GLFWbool            cursorHidden;
	//    TISInputSourceRef   inputSource;
	//    IOHIDManagerRef     hidManager;
	//    id                  unicodeData;
	helper objc.ID
	//    id                  keyUpMonitor;
	//    id                  nibObjects;
	//
	//    char                keynames[GLFW_KEY_LAST + 1][17];
	keycodes  [256]Key
	scancodes [KeyLast + 1]int
	//    short int           scancodes[GLFW_KEY_LAST + 1];
	//    char*               clipboardString;
	//    CGPoint             cascadePoint;
	//    // Where to place the cursor when re-enabled
	//    double              restoreCursorPosX, restoreCursorPosY;
	//    // The window whose disabled cursor mode is active
	disableCursorWindow *Window
	//
	//    struct {
	//        CFBundleRef     bundle;
	//        PFN_TISCopyCurrentKeyboardLayoutInputSource CopyCurrentKeyboardLayoutInputSource;
	//        PFN_TISGetInputSourceProperty GetInputSourceProperty;
	//        PFN_LMGetKbdType GetKbdType;
	//        CFStringRef     kPropertyUnicodeKeyLayoutData;
	//    } tis;
}

type platformLibraryContextState struct {
	// dlopen handle for OpenGL.framework (for glfwGetProcAddress)
	framework corefoundation.CFBundleRef
}
