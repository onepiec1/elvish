// Code generated by "stringer -type=feature -output=feature_string.go"; DO NOT EDIT.

package lscolors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[featureInvalid-0]
	_ = x[featureOrphanedSymlink-1]
	_ = x[featureSymlink-2]
	_ = x[featureMultiHardLink-3]
	_ = x[featureNamedPipe-4]
	_ = x[featureSocket-5]
	_ = x[featureDoor-6]
	_ = x[featureBlockDevice-7]
	_ = x[featureCharDevice-8]
	_ = x[featureWorldWritableStickyDirectory-9]
	_ = x[featureWorldWritableDirectory-10]
	_ = x[featureStickyDirectory-11]
	_ = x[featureDirectory-12]
	_ = x[featureCapability-13]
	_ = x[featureSetuid-14]
	_ = x[featureSetgid-15]
	_ = x[featureExecutable-16]
	_ = x[featureRegular-17]
}

const _feature_name = "featureInvalidfeatureOrphanedSymlinkfeatureSymlinkfeatureMultiHardLinkfeatureNamedPipefeatureSocketfeatureDoorfeatureBlockDevicefeatureCharDevicefeatureWorldWritableStickyDirectoryfeatureWorldWritableDirectoryfeatureStickyDirectoryfeatureDirectoryfeatureCapabilityfeatureSetuidfeatureSetgidfeatureExecutablefeatureRegular"

var _feature_index = [...]uint16{0, 14, 36, 50, 70, 86, 99, 110, 128, 145, 180, 209, 231, 247, 264, 277, 290, 307, 321}

func (i feature) String() string {
	if i < 0 || i >= feature(len(_feature_index)-1) {
		return "feature(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _feature_name[_feature_index[i]:_feature_index[i+1]]
}
