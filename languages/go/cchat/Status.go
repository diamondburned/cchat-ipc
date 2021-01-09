// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import "strconv"

type Status byte

const (
	StatusUnknown   Status = 0
	StatusOnline    Status = 1
	StatusIdle      Status = 2
	StatusBusy      Status = 3
	StatusAway      Status = 4
	StatusOffline   Status = 5
	StatusInvisible Status = 6
)

var EnumNamesStatus = map[Status]string{
	StatusUnknown:   "Unknown",
	StatusOnline:    "Online",
	StatusIdle:      "Idle",
	StatusBusy:      "Busy",
	StatusAway:      "Away",
	StatusOffline:   "Offline",
	StatusInvisible: "Invisible",
}

var EnumValuesStatus = map[string]Status{
	"Unknown":   StatusUnknown,
	"Online":    StatusOnline,
	"Idle":      StatusIdle,
	"Busy":      StatusBusy,
	"Away":      StatusAway,
	"Offline":   StatusOffline,
	"Invisible": StatusInvisible,
}

func (v Status) String() string {
	if s, ok := EnumNamesStatus[v]; ok {
		return s
	}
	return "Status(" + strconv.FormatInt(int64(v), 10) + ")"
}
