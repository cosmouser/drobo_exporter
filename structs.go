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
