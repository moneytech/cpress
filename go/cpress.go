/*
cpress - go bindings

-------------------------------------------------------------------------------

Installation:

    go get github.com/solusipse/cpress/go

Usage:

Import package with:

    import "github.com/solusipse/cpress/go"

See self-explanatory file `examples/example.go` for more informations.

-------------------------------------------------------------------------------

License: MIT (http://www.opensource.org/licenses/mit-license.php)
Repository: https://github.com/solusipse/cpress

-------------------------------------------------------------------------------
*/

package cpress

// #include "../src/cpress.c"
import "C"

func Initialize() {
	/*
	This method is used to initialize cpress. Call it before
	doing anything.
	*/
    C.initialize()
}

func Press_key(key C.int) {
	/*
	Simple key press.
	*/
    C.press_key(key)
}

func Hold_key(key C.int) {
	/*
	This method changes key state to pressed until Release_key
	method is called for same key.
	*/
    C.hold_key(key)
}

func Release_key(key C.int) {
	/*
	This method releases previously activated key
	*/
    C.release_key(key)
}

func Press_combination(keys ... _Ctype_int) {
	/*
	This method had to be rewritten because of incompatibility
	of C variadic function and Go.
	It allows to do key combos, for example: CTRL+C, CTRL+ALT+F1.
	*/
    for i := range keys {
        C.hold_key(keys[i])
    }
    for i := range keys {
        C.release_key(keys[i])
    }
    
}

func Close() {
	/*
	Use this for clean up.
	*/
	C.finish()
}


/* Keylist */
var KEY_RESERVED C.int = 0
var KEY_ESC C.int = 1
var KEY_1 C.int = 2
var KEY_2 C.int = 3
var KEY_3 C.int = 4
var KEY_4 C.int = 5
var KEY_5 C.int = 6
var KEY_6 C.int = 7
var KEY_7 C.int = 8
var KEY_8 C.int = 9
var KEY_9 C.int = 10
var KEY_0 C.int = 11
var KEY_MINUS C.int = 12
var KEY_EQUAL C.int = 13
var KEY_BACKSPACE C.int = 14
var KEY_TAB C.int = 15
var KEY_Q C.int = 16
var KEY_W C.int = 17
var KEY_E C.int = 18
var KEY_R C.int = 19
var KEY_T C.int = 20
var KEY_Y C.int = 21
var KEY_U C.int = 22
var KEY_I C.int = 23
var KEY_O C.int = 24
var KEY_P C.int = 25
var KEY_LEFTBRACE C.int = 26
var KEY_RIGHTBRACE C.int = 27
var KEY_ENTER C.int = 28
var KEY_LEFTCTRL C.int = 29
var KEY_A C.int = 30
var KEY_S C.int = 31
var KEY_D C.int = 32
var KEY_F C.int = 33
var KEY_G C.int = 34
var KEY_H C.int = 35
var KEY_J C.int = 36
var KEY_K C.int = 37
var KEY_L C.int = 38
var KEY_SEMICOLON C.int = 39
var KEY_APOSTROPHE C.int = 40
var KEY_GRAVE C.int = 41
var KEY_LEFTSHIFT C.int = 42
var KEY_BACKSLASH C.int = 43
var KEY_Z C.int = 44
var KEY_X C.int = 45
var KEY_C C.int = 46
var KEY_V C.int = 47
var KEY_B C.int = 48
var KEY_N C.int = 49
var KEY_M C.int = 50
var KEY_COMMA C.int = 51
var KEY_DOT C.int = 52
var KEY_SLASH C.int = 53
var KEY_RIGHTSHIFT C.int = 54
var KEY_KPASTERISK C.int = 55
var KEY_LEFTALT C.int = 56
var KEY_SPACE C.int = 57
var KEY_CAPSLOCK C.int = 58
var KEY_F1 C.int = 59
var KEY_F2 C.int = 60
var KEY_F3 C.int = 61
var KEY_F4 C.int = 62
var KEY_F5 C.int = 63
var KEY_F6 C.int = 64
var KEY_F7 C.int = 65
var KEY_F8 C.int = 66
var KEY_F9 C.int = 67
var KEY_F10 C.int = 68
var KEY_NUMLOCK C.int = 69
var KEY_SCROLLLOCK C.int = 70
var KEY_KP7 C.int = 71
var KEY_KP8 C.int = 72
var KEY_KP9 C.int = 73
var KEY_KPMINUS C.int = 74
var KEY_KP4 C.int = 75
var KEY_KP5 C.int = 76
var KEY_KP6 C.int = 77
var KEY_KPPLUS C.int = 78
var KEY_KP1 C.int = 79
var KEY_KP2 C.int = 80
var KEY_KP3 C.int = 81
var KEY_KP0 C.int = 82
var KEY_KPDOT C.int = 83
var KEY_ZENKAKUHANKAKU C.int = 85
var KEY_102ND C.int = 86
var KEY_F11 C.int = 87
var KEY_F12 C.int = 88
var KEY_RO C.int = 89
var KEY_KATAKANA C.int = 90
var KEY_HIRAGANA C.int = 91
var KEY_HENKAN C.int = 92
var KEY_KATAKANAHIRAGANA C.int = 93
var KEY_MUHENKAN C.int = 94
var KEY_KPJPCOMMA C.int = 95
var KEY_KPENTER C.int = 96
var KEY_RIGHTCTRL C.int = 97
var KEY_KPSLASH C.int = 98
var KEY_SYSRQ C.int = 99
var KEY_RIGHTALT C.int = 100
var KEY_LINEFEED C.int = 101
var KEY_HOME C.int = 102
var KEY_UP C.int = 103
var KEY_PAGEUP C.int = 104
var KEY_LEFT C.int = 105
var KEY_RIGHT C.int = 106
var KEY_END C.int = 107
var KEY_DOWN C.int = 108
var KEY_PAGEDOWN C.int = 109
var KEY_INSERT C.int = 110
var KEY_DELETE C.int = 111
var KEY_MACRO C.int = 112
var KEY_MUTE C.int = 113
var KEY_VOLUMEDOWN C.int = 114
var KEY_VOLUMEUP C.int = 115
var KEY_POWER C.int = 116
var KEY_KPEQUAL C.int = 117
var KEY_KPPLUSMINUS C.int = 118
var KEY_PAUSE C.int = 119
var KEY_SCALE C.int = 120 
var KEY_KPCOMMA C.int = 121
var KEY_HANGEUL C.int = 122
var KEY_HANGUEL C.int = KEY_HANGEUL
var KEY_HANJA C.int = 123
var KEY_YEN C.int = 124
var KEY_LEFTMETA C.int = 125
var KEY_RIGHTMETA C.int = 126
var KEY_COMPOSE C.int = 127
var KEY_STOP C.int = 128
var KEY_AGAIN C.int = 129
var KEY_PROPS C.int = 130
var KEY_UNDO C.int = 131
var KEY_FRONT C.int = 132
var KEY_COPY C.int = 133
var KEY_OPEN C.int = 134
var KEY_PASTE C.int = 135
var KEY_FIND C.int = 136
var KEY_CUT C.int = 137
var KEY_HELP C.int = 138
var KEY_MENU C.int = 139
var KEY_CALC C.int = 140
var KEY_SETUP C.int = 141
var KEY_SLEEP C.int = 142
var KEY_WAKEUP C.int = 143
var KEY_FILE C.int = 144
var KEY_SENDFILE C.int = 145
var KEY_DELETEFILE C.int = 146
var KEY_XFER C.int = 147
var KEY_PROG1 C.int = 148
var KEY_PROG2 C.int = 149
var KEY_WWW C.int = 150
var KEY_MSDOS C.int = 151
var KEY_COFFEE C.int = 152
var KEY_SCREENLOCK C.int = KEY_COFFEE
var KEY_DIRECTION C.int = 153
var KEY_CYCLEWINDOWS C.int = 154
var KEY_MAIL C.int = 155
var KEY_BOOKMARKS C.int = 156
var KEY_COMPUTER C.int = 157
var KEY_BACK C.int = 158
var KEY_FORWARD C.int = 159
var KEY_CLOSECD C.int = 160
var KEY_EJECTCD C.int = 161
var KEY_EJECTCLOSECD C.int = 162
var KEY_NEXTSONG C.int = 163
var KEY_PLAYPAUSE C.int = 164
var KEY_PREVIOUSSONG C.int = 165
var KEY_STOPCD C.int = 166
var KEY_RECORD C.int = 167
var KEY_REWIND C.int = 168
var KEY_PHONE C.int = 169
var KEY_ISO C.int = 170
var KEY_CONFIG C.int = 171
var KEY_HOMEPAGE C.int = 172
var KEY_REFRESH C.int = 173
var KEY_EXIT C.int = 174
var KEY_MOVE C.int = 175
var KEY_EDIT C.int = 176
var KEY_SCROLLUP C.int = 177
var KEY_SCROLLDOWN C.int = 178
var KEY_KPLEFTPAREN C.int = 179
var KEY_KPRIGHTPAREN C.int = 180
var KEY_NEW C.int = 181
var KEY_REDO C.int = 182
var KEY_F13 C.int = 183
var KEY_F14 C.int = 184
var KEY_F15 C.int = 185
var KEY_F16 C.int = 186
var KEY_F17 C.int = 187
var KEY_F18 C.int = 188
var KEY_F19 C.int = 189
var KEY_F20 C.int = 190
var KEY_F21 C.int = 191
var KEY_F22 C.int = 192
var KEY_F23 C.int = 193
var KEY_F24 C.int = 194
var KEY_PLAYCD C.int = 200
var KEY_PAUSECD C.int = 201
var KEY_PROG3 C.int = 202
var KEY_PROG4 C.int = 203
var KEY_DASHBOARD C.int = 204
var KEY_SUSPEND C.int = 205
var KEY_CLOSE C.int = 206
var KEY_PLAY C.int = 207
var KEY_FASTFORWARD C.int = 208
var KEY_BASSBOOST C.int = 209
var KEY_PRINT C.int = 210
var KEY_HP C.int = 211
var KEY_CAMERA C.int = 212
var KEY_SOUND C.int = 213
var KEY_QUESTION C.int = 214
var KEY_EMAIL C.int = 215
var KEY_CHAT C.int = 216
var KEY_SEARCH C.int = 217
var KEY_CONNECT C.int = 218
var KEY_FINANCE C.int = 219
var KEY_SPORT C.int = 220
var KEY_SHOP C.int = 221
var KEY_ALTERASE C.int = 222
var KEY_CANCEL C.int = 223
var KEY_BRIGHTNESSDOWN C.int = 224
var KEY_BRIGHTNESSUP C.int = 225
var KEY_MEDIA C.int = 226
var KEY_SWITCHVIDEOMODE C.int = 227
var KEY_KBDILLUMTOGGLE C.int = 228
var KEY_KBDILLUMDOWN C.int = 229
var KEY_KBDILLUMUP C.int = 230
var KEY_SEND C.int = 231
var KEY_REPLY C.int = 232
var KEY_FORWARDMAIL C.int = 233
var KEY_SAVE C.int = 234
var KEY_DOCUMENTS C.int = 235
var KEY_BATTERY C.int = 236
var KEY_BLUETOOTH C.int = 237
var KEY_WLAN C.int = 238
var KEY_UWB C.int = 239
var KEY_UNKNOWN C.int = 240
var KEY_VIDEO_NEXT C.int = 241
var KEY_VIDEO_PREV C.int = 242
var KEY_BRIGHTNESS_CYCLE C.int = 243
var KEY_BRIGHTNESS_ZERO C.int = 244
var KEY_DISPLAY_OFF C.int = 245
var KEY_WIMAX C.int = 246
var KEY_RFKILL C.int = 247
var KEY_MICMUTE C.int = 248