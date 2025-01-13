package MvCodeReaderSDK

type MvCodeReaderErrorCode int

const (
	MV_CODEREADER_OK MvCodeReaderErrorCode = 0x00000000 // Succeeded.
	// General Error (from 0x80020000 to 0x800200FF)
	MV_CODEREADER_E_HANDLE         MvCodeReaderErrorCode = 0x80020000 // Error or invalid handle
	MV_CODEREADER_E_SUPPORT        MvCodeReaderErrorCode = 0x80020001 // The function is not supported
	MV_CODEREADER_E_BUFOVER        MvCodeReaderErrorCode = 0x80020002 // Buffer is full
	MV_CODEREADER_E_CALLORDER      MvCodeReaderErrorCode = 0x80020003 // Incorrect call order
	MV_CODEREADER_E_PARAMETER      MvCodeReaderErrorCode = 0x80020004 // Incorrect parameter
	MV_CODEREADER_E_RESOURCE       MvCodeReaderErrorCode = 0x80020005 // Resource request failed.
	MV_CODEREADER_E_NODATA         MvCodeReaderErrorCode = 0x80020006 // No data
	MV_CODEREADER_E_PRECONDITION   MvCodeReaderErrorCode = 0x80020007 // Incorrect precondition or running environment changed.
	MV_CODEREADER_E_VERSION        MvCodeReaderErrorCode = 0x80020008 // Version is mismatched
	MV_CODEREADER_E_NOENOUGH_BUF   MvCodeReaderErrorCode = 0x80020009 // Insufficient memory
	MV_CODEREADER_E_ABNORMAL_IMAGE MvCodeReaderErrorCode = 0x8002000A // Abnormal image. Incomplete image caused by packet loss.
	MV_CODEREADER_E_LOAD_LIBRARY   MvCodeReaderErrorCode = 0x8002000B // Importing DLL failed.
	MV_CODEREADER_E_NOOUTBUF       MvCodeReaderErrorCode = 0x8002000C // No output buffer
	MV_CODEREADER_E_FILE_PATH      MvCodeReaderErrorCode = 0x8002000F // Incorrect file path
	MV_CODEREADER_E_UNKNOW         MvCodeReaderErrorCode = 0x800200FF // Unknown error
	// GenICam Related Error (from 0x80020100 to 0x800201FF)Related                                                           MvCodeReaderErrorCode = // GenICam  // Error (from 0x80020100 to 0x800201FF)
	MV_CODEREADER_E_GC_GENERIC     MvCodeReaderErrorCode = 0x80020100 // Generic error
	MV_CODEREADER_E_GC_ARGUMENT    MvCodeReaderErrorCode = 0x80020101 // Invalid parameter
	MV_CODEREADER_E_GC_RANGE       MvCodeReaderErrorCode = 0x80020102 // The value is out of range.
	MV_CODEREADER_E_GC_PROPERTY    MvCodeReaderErrorCode = 0x80020103 // Attribute error
	MV_CODEREADER_E_GC_RUNTIME     MvCodeReaderErrorCode = 0x80020104 // Running environment error
	MV_CODEREADER_E_GC_LOGICAL     MvCodeReaderErrorCode = 0x80020105 // Incorrect logic
	MV_CODEREADER_E_GC_ACCESS      MvCodeReaderErrorCode = 0x80020106 // Node accessing condition error
	MV_CODEREADER_E_GC_TIMEOUT     MvCodeReaderErrorCode = 0x80020107 // Timed out
	MV_CODEREADER_E_GC_DYNAMICCAST MvCodeReaderErrorCode = 0x80020108 // Conversion exception
	MV_CODEREADER_E_GC_UNKNOW      MvCodeReaderErrorCode = 0x800201FF // GenICam unknown error
	// GigE_STATUS Reltaed Errors (from 0x80020200 to 0x800202FF)
	MV_CODEREADER_E_NOT_IMPLEMENTED MvCodeReaderErrorCode = 0x80020200 // Command is not supported by the device.
	MV_CODEREADER_E_INVALID_ADDRESS MvCodeReaderErrorCode = 0x80020201 // Target address does not exist.
	MV_CODEREADER_E_WRITE_PROTECT   MvCodeReaderErrorCode = 0x80020202 // Target address is not writable.
	MV_CODEREADER_E_ACCESS_DENIED   MvCodeReaderErrorCode = 0x80020203 // No access permission
	MV_CODEREADER_E_BUSY            MvCodeReaderErrorCode = 0x80020204 // Device is busy, or network is disconnected.
	MV_CODEREADER_E_PACKET          MvCodeReaderErrorCode = 0x80020205 // Network packet error
	MV_CODEREADER_E_NETER           MvCodeReaderErrorCode = 0x80020206 // Network error
	// GigE Cameras Related Error(s)
	MV_CODEREADER_E_IP_CONFLICT MvCodeReaderErrorCode = 0x80020221 // IP address conflicted
	// USB_STATUS Related Errors (from 0x80020300 to 0x800203FF)
	MV_CODEREADER_E_USB_READ      MvCodeReaderErrorCode = 0x80020300 // USB read error
	MV_CODEREADER_E_USB_WRITE     MvCodeReaderErrorCode = 0x80020301 // USB write error
	MV_CODEREADER_E_USB_DEVICE    MvCodeReaderErrorCode = 0x80020302 // Device exception
	MV_CODEREADER_E_USB_GENICAM   MvCodeReaderErrorCode = 0x80020303 // GenICam error
	MV_CODEREADER_E_USB_BANDWIDTH MvCodeReaderErrorCode = 0x80020304 // Insufficient bandwidth
	MV_CODEREADER_E_USB_DRIVER    MvCodeReaderErrorCode = 0x80020305 // Driver is mismatched, or the driver is not installed.
	MV_CODEREADER_E_USB_UNKNOW    MvCodeReaderErrorCode = 0x800203FF // USB unknown error
	// Upgrade Related Errors (from 0x80020400 to 0x800204FF)
	MV_CODEREADER_E_UPG_MIN_ERRCODE                  MvCodeReaderErrorCode = 0x80020400 // Minimum error code of upgrade module
	MV_CODEREADER_E_UPG_FILE_MISMATCH                MvCodeReaderErrorCode = 0x80020400 // Upgrade firmware is mismatched.
	MV_CODEREADER_E_UPG_LANGUSGE_MISMATCH            MvCodeReaderErrorCode = 0x80020401 // Upgrade firmware language is mismatched.
	MV_CODEREADER_E_UPG_CONFLICT                     MvCodeReaderErrorCode = 0x80020402 // Upgrading conflicted (repeated upgrading requests during device upg
	MV_CODEREADER_E_UPG_INNER_ERR                    MvCodeReaderErrorCode = 0x80020403 // Camera internal error during upgrade.
	MV_CODEREADER_E_UPG_REGRESH_TYPE_ERR             MvCodeReaderErrorCode = 0x80020404 // Acquiring camera model failed.
	MV_CODEREADER_E_UPG_COPY_FPGABIN_ERR             MvCodeReaderErrorCode = 0x80020405 // Copying FPGA file failed.
	MV_CODEREADER_E_UPG_ZIPEXTRACT_ERR               MvCodeReaderErrorCode = 0x80020406 // Extracting ZIP file failed.
	MV_CODEREADER_E_UPG_DAVEXTRACT_ERR               MvCodeReaderErrorCode = 0x80020407 // Extracting DAV file failed.
	MV_CODEREADER_E_UPG_DAVCOMPRESS_ERR              MvCodeReaderErrorCode = 0x80020408 // Compressing ZIP file failed.
	MV_CODEREADER_E_UPG_ZIPCOMPRESS_ERR              MvCodeReaderErrorCode = 0x80020409 // Compressing DAV file failed.
	MV_CODEREADER_E_UPG_GET_PROGRESS_TIMEOUT_ERR     MvCodeReaderErrorCode = 0x80020410 // Upgrade progress acquinsition timed out.
	MV_CODEREADER_E_UPG_SEND_QUERY_PROGRESS_ERR      MvCodeReaderErrorCode = 0x80020411 // Failed to send progress query instruction
	MV_CODEREADER_E_UPG_RECV_QUERY_PROGRESS_ERR      MvCodeReaderErrorCode = 0x80020412 // Failed to receive progress query instruction͘
	MV_CODEREADER_E_UPG_GET_QUERY_PROGRESS_ERR       MvCodeReaderErrorCode = 0x80020413 // Getting query progress failed.
	MV_CODEREADER_E_UPG_GET_MAX_QUERY_PROGRESS_ERR   MvCodeReaderErrorCode = 0x80020414 // Getting fastest progress failed.
	MV_CODEREADER_E_UPG_CHECKT_PACKET_FAILED         MvCodeReaderErrorCode = 0x80020465 // Verifying file failed.
	MV_CODEREADER_E_UPG_FPGA_PROGRAM_FAILED          MvCodeReaderErrorCode = 0x80020466 // FPGA program upgrade failed.
	MV_CODEREADER_E_UPG_WATCHDOG_FAILED              MvCodeReaderErrorCode = 0x80020467 // Watchdog upgrade failed.
	MV_CODEREADER_E_UPG_CAMERA_AND_BARE_FAILED       MvCodeReaderErrorCode = 0x80020468 // Bare camera upgrade failed.
	MV_CODEREADER_E_UPG_RETAIN_CONFIG_FAILED         MvCodeReaderErrorCode = 0x80020469 // Retaining configuration file failed.
	MV_CODEREADER_E_UPG_FPGA_DRIVER_FAILED           MvCodeReaderErrorCode = 0x8002046A // FPGA drive upgrade failed.
	MV_CODEREADER_E_UPG_SPI_DRIVER_FAILED            MvCodeReaderErrorCode = 0x8002046B // SPI drive upgrade failed.
	MV_CODEREADER_E_UPG_REBOOT_SYSTEM_FAILED         MvCodeReaderErrorCode = 0x8002046C // Restarting failed.
	MV_CODEREADER_E_UPG_UPSELF_FAILED                MvCodeReaderErrorCode = 0x8002046D // Service upgrade failed.
	MV_CODEREADER_E_UPG_STOP_RELATION_PROGRAM_FAILED MvCodeReaderErrorCode = 0x8002046E // Stop related service failed.
	MV_CODEREADER_E_UPG_DEVCIE_TYPE_INCONSISTENT     MvCodeReaderErrorCode = 0x8002046F // Inconsistent device type
	MV_CODEREADER_E_UPG_READ_ENCRYPT_INFO_FAILED     MvCodeReaderErrorCode = 0x80020470 // Read encryption message failed.
	MV_CODEREADER_E_UPG_PLAT_TYPE_INCONSISTENT       MvCodeReaderErrorCode = 0x80020471 // Incorrect device platform
	MV_CODEREADER_E_UPG_CAMERA_TYPE_INCONSISTENT     MvCodeReaderErrorCode = 0x80020472 // Incorrect camera model
	MV_CODEREADER_E_UPG_DEVICE_UPGRADING             MvCodeReaderErrorCode = 0x80020473 // Incorrect camera model
	MV_CODEREADER_E_UPG_UNZIP_FAILED                 MvCodeReaderErrorCode = 0x80020474 // Upgrade package decompression failed.
	MV_CODEREADER_E_UPG_BLE_DISCONNECT               MvCodeReaderErrorCode = 0x80020475 // PDA bluetooth is disconnected.
	MV_CODEREADER_E_UPG_BATTERY_NOTENOUGH            MvCodeReaderErrorCode = 0x80020476 // Low battery
	MV_CODEREADER_E_UPG_RTC_NOT_PRESENT              MvCodeReaderErrorCode = 0x80020477 // PDA is not on the base.
	MV_CODEREADER_E_UPG_APP_ERR                      MvCodeReaderErrorCode = 0x80020478 // Failed to upgrade the app.
	MV_CODEREADER_E_UPG_L3_ERR                       MvCodeReaderErrorCode = 0x80020479 // Failed to upgrade L3.
	MV_CODEREADER_E_UPG_MCU_ERR                      MvCodeReaderErrorCode = 0x8002047A // Failed to upgrade MCU.
	MV_CODEREADER_E_UPG_PLATFORM_DISMATCH            MvCodeReaderErrorCode = 0x8002047B // Platform does not match.
	MV_CODEREADER_E_UPG_TYPE_DISMATCH                MvCodeReaderErrorCode = 0x8002047C // Model does not match.
	MV_CODEREADER_E_UPG_SPACE_DISMATCH               MvCodeReaderErrorCode = 0x8002047D // Space does not match.
	MV_CODEREADER_E_UPG_MEM_DISMATCH                 MvCodeReaderErrorCode = 0x8002047E // Memory does not match.
	MV_CODEREADER_E_UPG_NET_TRANS_ERROR              MvCodeReaderErrorCode = 0x8002047F // Network transmission exception.Please upgrade again.
	MV_CODEREADER_E_UPG_UNKNOW                       MvCodeReaderErrorCode = 0x800204FF // Upgrade unknown error
	//Network Components Related Errors (from 0x80020500 to 0x800205FF)
	MV_CODEREADER_E_CREAT_SOCKET   MvCodeReaderErrorCode = 0x80020500 // Creating socket error
	MV_CODEREADER_E_BIND_SOCKET    MvCodeReaderErrorCode = 0x80020501 // Binding error
	MV_CODEREADER_E_CONNECT_SOCKET MvCodeReaderErrorCode = 0x80020502 // Connection error
	MV_CODEREADER_E_GET_HOSTNAME   MvCodeReaderErrorCode = 0x80020503 // Host name getting error
	MV_CODEREADER_E_NET_WRITE      MvCodeReaderErrorCode = 0x80020504 // Data write error
	MV_CODEREADER_E_NET_READ       MvCodeReaderErrorCode = 0x80020505 // Data read error
	MV_CODEREADER_E_NET_SELECT     MvCodeReaderErrorCode = 0x80020506 // Selection error
	MV_CODEREADER_E_NET_TIMEOUT    MvCodeReaderErrorCode = 0x80020507 // Timed out
	MV_CODEREADER_E_NET_ACCEPT     MvCodeReaderErrorCode = 0x80020508 // Receive error
	MV_CODEREADER_E_NET_UNKNOW     MvCodeReaderErrorCode = 0x800205FF // Network unknown error
)
