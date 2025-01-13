package MvCodeReaderSDK

// Device Access Mode
const (
	MvAccessExclusive                  = 1 // MvAccessExclusive: english Exclusive authority, other APP is only allowed to read the CCP register
	MvAccessExclusiveWithSwitch        = 2 // MvAccessExclusiveWithSwitch: english You can seize the authority from the 5 mode, and then open with exclusive authority
	MvAccessControl                    = 3 // MvAccessControl: english Control authority, allows other APP reading all registers
	MvAccessControlWithSwitch          = 4 // MvAccessControlWithSwitch: english You can seize the authority from the 5 mode, and then open with control authority
	MvAccessControlSwitchEnable        = 5 // MvAccessControlSwitchEnable: english Open with seized control authority
	MvAccessControlSwitchEnableWithKey = 6 // MvAccessControlSwitchEnableWithKey: english You can seize the authority from the 5 mode, and then open with seized control authority
	MvAccessMonitor                    = 7 // MvAccessMonitor: english Open with read mode and is available under control authority
)
