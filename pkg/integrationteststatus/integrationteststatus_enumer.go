// Code generated by "enumer -type=IntegrationTestStatus -linecomment -json"; DO NOT EDIT.

package integrationteststatus

import (
	"encoding/json"
	"fmt"
)

const _IntegrationTestStatusName = "PendingInProgressDeletedEnvironmentProvisionErrorDeploymentErrorTestFailTestPassedTestInvalidBuildPLRInProgressSnapshotCreationFailedBuildPLRFailedGroupSnapshotCreationFailed"

var _IntegrationTestStatusIndex = [...]uint8{0, 7, 17, 24, 49, 64, 72, 82, 93, 111, 133, 147, 174}

func (i IntegrationTestStatus) String() string {
	i -= 1
	if i < 0 || i >= IntegrationTestStatus(len(_IntegrationTestStatusIndex)-1) {
		return fmt.Sprintf("IntegrationTestStatus(%d)", i+1)
	}
	return _IntegrationTestStatusName[_IntegrationTestStatusIndex[i]:_IntegrationTestStatusIndex[i+1]]
}

var _IntegrationTestStatusValues = []IntegrationTestStatus{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

var _IntegrationTestStatusNameToValueMap = map[string]IntegrationTestStatus{
	_IntegrationTestStatusName[0:7]:     1,
	_IntegrationTestStatusName[7:17]:    2,
	_IntegrationTestStatusName[17:24]:   3,
	_IntegrationTestStatusName[24:49]:   4,
	_IntegrationTestStatusName[49:64]:   5,
	_IntegrationTestStatusName[64:72]:   6,
	_IntegrationTestStatusName[72:82]:   7,
	_IntegrationTestStatusName[82:93]:   8,
	_IntegrationTestStatusName[93:111]:  9,
	_IntegrationTestStatusName[111:133]: 10,
	_IntegrationTestStatusName[133:147]: 11,
	_IntegrationTestStatusName[148:174]: 12,
}

// IntegrationTestStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func IntegrationTestStatusString(s string) (IntegrationTestStatus, error) {
	if val, ok := _IntegrationTestStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to IntegrationTestStatus values", s)
}

// IntegrationTestStatusValues returns all values of the enum
func IntegrationTestStatusValues() []IntegrationTestStatus {
	return _IntegrationTestStatusValues
}

// IsAIntegrationTestStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i IntegrationTestStatus) IsAIntegrationTestStatus() bool {
	for _, v := range _IntegrationTestStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for IntegrationTestStatus
func (i IntegrationTestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for IntegrationTestStatus
func (i *IntegrationTestStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("IntegrationTestStatus should be a string, got %s", data)
	}

	var err error
	*i, err = IntegrationTestStatusString(s)
	return err
}
