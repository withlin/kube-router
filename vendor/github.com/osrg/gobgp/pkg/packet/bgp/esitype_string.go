// Code generated by "stringer -type=ESIType"; DO NOT EDIT.

package bgp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ESI_ARBITRARY-0]
	_ = x[ESI_LACP-1]
	_ = x[ESI_MSTP-2]
	_ = x[ESI_MAC-3]
	_ = x[ESI_ROUTERID-4]
	_ = x[ESI_AS-5]
}

const _ESIType_name = "ESI_ARBITRARYESI_LACPESI_MSTPESI_MACESI_ROUTERIDESI_AS"

var _ESIType_index = [...]uint8{0, 13, 21, 29, 36, 48, 54}

func (i ESIType) String() string {
	if i >= ESIType(len(_ESIType_index)-1) {
		return "ESIType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ESIType_name[_ESIType_index[i]:_ESIType_index[i+1]]
}
