package main

import (
	"encoding/xml"
)

// ESATMUpdate is the status report that the drobo periodically sends
type ESATMUpdate struct {
	XMLName                   xml.Name `xml:"ESATMUpdate"`
	ESAUpdateSignature        string   `xml:"mESAUpdateSignature"`
	ESAUpdateVersion          int      `xml:"mESAUpdateVersion"`
	ESAUpdateSize             int      `xml:"mESAUpdateSize"`
	ESAID                     string   `xml:"mESAID"`
	Serial                    string   `xml:"mSerial"`
	Name                      string   `xml:"mName"`
	Version                   string   `xml:"mVersion"`
	ReleaseDate               string   `xml:"mReleaseDate"`
	Arch                      string   `xml:"mArch"`
	FirmwareFeatures          int      `xml:"mFirmwareFeatures"`
	xtFtr                     int      `xml:"extFtr"`
	FirmwareTestFeatures      int      `xml:"mFirmwareTestFeatures"`
	FirmwareTestState         int      `xml:"mFirmwareTestState"`
	FirmwareTestValue         int      `xml:"mFirmwareTestValue"`
	Status                    int      `xml:"mStatus"`
	RelayoutCount             int      `xml:"mRelayoutCount"`
	DoubleDegradedCnt         int      `xml:"mDoubleDegradedCnt"`
	LatestUELGenNumber        int      `xml:"mLatestUELGenNumber"`
	TotalCapacityProtected    int      `xml:"mTotalCapacityProtected"`
	UsedCapacityProtected     int      `xml:"mUsedCapacityProtected"`
	FreeCapacityProtected     int      `xml:"mFreeCapacityProtected"`
	TotalCapacityUnprotected  int      `xml:"mTotalCapacityUnprotected"`
	UsedCapacityOS            int      `xml:"mUsedCapacityOS"`
	TotalCapacityPT           int      `xml:"mTotalCapacityPT"`
	UsedCapacityPT            int      `xml:"mUsedCapacityPT"`
	YellowThreshold           int      `xml:"mYellowThreshold"`
	RedThreshold              int      `xml:"mRedThreshold"`
	UseUnprotectedCapacity    int      `xml:"mUseUnprotectedCapacity"`
	RealTimeIntegrityChecking int      `xml:"mRealTimeIntegrityChecking"`
	StoredFirmwareTestState   int      `xml:"mStoredFirmwareTestState"`
	StoredFirmwareTestValue   int      `xml:"mStoredFirmwareTestValue"`
	DiskPackID                int      `xml:"mDiskPackID"`
	DroboName                 string   `xml:"mDroboName"`
	ConnectionType            int      `xml:"mConnectionType"`
	SlotCountExp              int      `xml:"mSlotCountExp"`
	FirmwareFeatureStates     int      `xml:"mFirmwareFeatureStates"`
	LUNCount                  int      `xml:"mLUNCount"`
	MaxLUNs                   int      `xml:"mMaxLUNs"`
	SledName                  string   `xml:"mSledName"`
	SledStatus                int      `xml:"mSledStatus"`
	DiskPackStatus            int      `xml:"mDiskPackStatus"`
	StatusEx                  int      `xml:"mStatusEx"`
	DeviceType                int      `xml:"mDeviceType"`
	Model                     string   `xml:"mModel"`
	DNASStatus                int      `xml:"DNASStatus"`
	DNASConfigVersion         int      `xml:"DNASConfigVersion"`
	DNASDroboAppsShared       int      `xml:"DNASDroboAppsShared"`
	DNASDiskPackId            string   `xml:"DNASDiskPackId"`
	DNASFeatureTable          int      `xml:"DNASFeatureTable"`
	DNASEmailConfigEnabled    int      `xml:"DNASEmailConfigEnabled"`
	// SledVersion/
	// SledSerial/
	// LoggedinUsername/
	// DroboApps
}

var helpInfo = map[string]string{
	"ESAID":                  "The field contains the device's serial number, as indicated by the file /sys/bus/dri_dnas_fake_bus/drivers/dri_dnas_scsi/serial. In the case of the FS, the first character is replaced with a 0 (zero).",
	"Name":                   "The user-friendly name of the device, as defined in the Dashboard.",
	"Version":                "The version of the firmware.",
	"ReleaseDate":            "The release date of the firmware.",
	"Arch":                   "A string representing the CPU type (Arm) and board manufacturer (Marvell). Observed value: ArmMarvell for both DroboFS and Drobo5N.",
	"FirmwareFeatures":       "Unknown. It is 34602495 for the DroboFS, and 2456813055 for the Drobo5N.",
	"extFtr":                 "Unknown. Only available on Drobo5N. Value observed: 303.",
	"Status":                 "The current status of the Drobo.",
	"RelayoutCount":          "Number of blocks (?) that still need to be processed after a disk pack change (insertion, replacement or removal of a drive). Outside of a data relayout (a.k.a. 'data protection'), this field is always zero. Once a data relayout starts, it counts down to zero. By monitoring the rate at which this number decreases it is possible to estimate the remaining time.",
	"DoubleDegradedCnt":      "Unknown. Only available on Drobo5N. Value observed: 0.",
	"LatestUELGenNumber":     "Unknown. Only available on Drobo5N (with firmware 3.2.0+ ?). Value observed: 956301312.",
	"TotalCapacityProtected": "The total usable space in bytes. This is the number that gets reported in the Dashboard as the 'Total' capacity.",
	"UsedCapacityProtected":  "The total used space in bytes. This is the number that gets reported in the Dashboard as 'Used Space.'",
	"FreeCapacityProtected":  "This is the free space in bytes. This is the number that gets reported in the Dashboard as 'Free Space.'",
	"YellowThreshold":        "The percentage of used space to reach the yellow threshold. Format is XXYY which translates to XX.YY%. Observed value: 8500.",
	"RedThreshold":           "The percentage of used space to reach the red threshold. Format is XXYY which translates to XX.YY%. Observed value: 9500.",
	"DroboName":              "Seems to be the same value as mName.",
	"SlotCountExp":           "The number of disk slots on the device. The value is 8 for the DroboFS, and 6 for the Drobo5N.",
	"FirmwareFeatureStates":  "It seems to indicate whether or not dual-redundancy is enabled.",
	"SledName":               "Seems to be the same as mName and mDroboName.",
	"DNASDroboAppsEnabled":   "Whether DroboApps are enabled or not.",
}
