package MvCodeReaderSDK

import (
	"errors"
	"fmt"
)

var ErrUnknown error = errors.New("unknown error")

type MvCodeReaderError struct {
	Code MvCodeReaderErrorCode
	Desc string
}

func (e MvCodeReaderError) Error() string {
	return fmt.Sprintf("error %s (%v)", e.Desc, e.Code)
}

func Err(errCode int32) error {

	return errorByCode(MvCodeReaderErrorCode(errCode))

}

func errorByCode(code MvCodeReaderErrorCode) error {

	switch code {
	case MV_CODEREADER_OK:
		return nil
	case MV_CODEREADER_E_HANDLE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_HANDLE, Desc: "Error or invalid handle"}
	case MV_CODEREADER_E_SUPPORT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_SUPPORT, Desc: "The function is not supported"}
	case MV_CODEREADER_E_BUFOVER:
		return MvCodeReaderError{Code: MV_CODEREADER_E_BUFOVER, Desc: "Buffer is full"}
	case MV_CODEREADER_E_CALLORDER:
		return MvCodeReaderError{Code: MV_CODEREADER_E_CALLORDER, Desc: "Incorrect call order"}
	case MV_CODEREADER_E_PARAMETER:
		return MvCodeReaderError{Code: MV_CODEREADER_E_PARAMETER, Desc: "Incorrect parameter"}
	case MV_CODEREADER_E_RESOURCE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_RESOURCE, Desc: "Resource request failed."}
	case MV_CODEREADER_E_NODATA:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NODATA, Desc: "No data"}
	case MV_CODEREADER_E_PRECONDITION:
		return MvCodeReaderError{Code: MV_CODEREADER_E_PRECONDITION, Desc: "Incorrect precondition or running environment changed."}
	case MV_CODEREADER_E_VERSION:
		return MvCodeReaderError{Code: MV_CODEREADER_E_VERSION, Desc: "Version is mismatched"}
	case MV_CODEREADER_E_NOENOUGH_BUF:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NOENOUGH_BUF, Desc: "Insufficient memory"}
	case MV_CODEREADER_E_ABNORMAL_IMAGE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_ABNORMAL_IMAGE, Desc: "Abnormal image. Incomplete image caused by packet loss."}
	case MV_CODEREADER_E_LOAD_LIBRARY:
		return MvCodeReaderError{Code: MV_CODEREADER_E_LOAD_LIBRARY, Desc: "Importing DLL failed."}
	case MV_CODEREADER_E_NOOUTBUF:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NOOUTBUF, Desc: "No output buffer"}
	case MV_CODEREADER_E_FILE_PATH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_FILE_PATH, Desc: "Incorrect file path"}
	case MV_CODEREADER_E_UNKNOW:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UNKNOW, Desc: "Unknown error"}
	case MV_CODEREADER_E_GC_GENERIC:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_GENERIC, Desc: "Generic error"}
	case MV_CODEREADER_E_GC_ARGUMENT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_ARGUMENT, Desc: "Invalid parameter"}
	case MV_CODEREADER_E_GC_RANGE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_RANGE, Desc: "The value is out of range."}
	case MV_CODEREADER_E_GC_PROPERTY:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_PROPERTY, Desc: "Attribute error"}
	case MV_CODEREADER_E_GC_RUNTIME:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_RUNTIME, Desc: "Running environment error"}
	case MV_CODEREADER_E_GC_LOGICAL:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_LOGICAL, Desc: "Incorrect logic"}
	case MV_CODEREADER_E_GC_ACCESS:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_ACCESS, Desc: "Node accessing condition error"}
	case MV_CODEREADER_E_GC_TIMEOUT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_TIMEOUT, Desc: "Timed out"}
	case MV_CODEREADER_E_GC_DYNAMICCAST:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_DYNAMICCAST, Desc: "Conversion exception"}
	case MV_CODEREADER_E_GC_UNKNOW:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GC_UNKNOW, Desc: "GenICam unknown error"}
	case MV_CODEREADER_E_NOT_IMPLEMENTED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NOT_IMPLEMENTED, Desc: "Command is not supported by the device."}
	case MV_CODEREADER_E_INVALID_ADDRESS:
		return MvCodeReaderError{Code: MV_CODEREADER_E_INVALID_ADDRESS, Desc: "Target address does not exist."}
	case MV_CODEREADER_E_WRITE_PROTECT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_WRITE_PROTECT, Desc: "Target address is not writable."}
	case MV_CODEREADER_E_ACCESS_DENIED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_ACCESS_DENIED, Desc: "No access permission"}
	case MV_CODEREADER_E_BUSY:
		return MvCodeReaderError{Code: MV_CODEREADER_E_BUSY, Desc: "Device is busy, or network is disconnected."}
	case MV_CODEREADER_E_PACKET:
		return MvCodeReaderError{Code: MV_CODEREADER_E_PACKET, Desc: "Network packet error"}
	case MV_CODEREADER_E_NETER:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NETER, Desc: "Network error"}
	case MV_CODEREADER_E_IP_CONFLICT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_IP_CONFLICT, Desc: "IP address conflicted"}
	case MV_CODEREADER_E_USB_READ:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_READ, Desc: "USB read error"}
	case MV_CODEREADER_E_USB_WRITE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_WRITE, Desc: "USB write error"}
	case MV_CODEREADER_E_USB_DEVICE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_DEVICE, Desc: "Device exception"}
	case MV_CODEREADER_E_USB_GENICAM:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_GENICAM, Desc: "GenICam error"}
	case MV_CODEREADER_E_USB_BANDWIDTH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_BANDWIDTH, Desc: "Insufficient bandwidth"}
	case MV_CODEREADER_E_USB_DRIVER:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_DRIVER, Desc: "Driver is mismatched, or the driver is not installed."}
	case MV_CODEREADER_E_USB_UNKNOW:
		return MvCodeReaderError{Code: MV_CODEREADER_E_USB_UNKNOW, Desc: "USB unknown error"}
	//case MV_CODEREADER_E_UPG_MIN_ERRCODE: return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_MIN_ERRCODE, Desc: "Minimum error code of upgrade module"}
	case MV_CODEREADER_E_UPG_FILE_MISMATCH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_FILE_MISMATCH, Desc: "Upgrade firmware is mismatched."}
	case MV_CODEREADER_E_UPG_LANGUSGE_MISMATCH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_LANGUSGE_MISMATCH, Desc: "Upgrade firmware language is mismatched."}
	case MV_CODEREADER_E_UPG_CONFLICT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_CONFLICT, Desc: "Upgrading conflicted (repeated upgrading requests during device upgrade)."}
	case MV_CODEREADER_E_UPG_INNER_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_INNER_ERR, Desc: "Camera internal error during upgrade."}
	case MV_CODEREADER_E_UPG_REGRESH_TYPE_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_REGRESH_TYPE_ERR, Desc: "Acquiring camera model failed."}
	case MV_CODEREADER_E_UPG_COPY_FPGABIN_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_COPY_FPGABIN_ERR, Desc: "Copying FPGA file failed."}
	case MV_CODEREADER_E_UPG_ZIPEXTRACT_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_ZIPEXTRACT_ERR, Desc: "Extracting ZIP file failed."}
	case MV_CODEREADER_E_UPG_DAVEXTRACT_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_DAVEXTRACT_ERR, Desc: "Extracting DAV file failed."}
	case MV_CODEREADER_E_UPG_DAVCOMPRESS_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_DAVCOMPRESS_ERR, Desc: "Compressing ZIP file failed."}
	case MV_CODEREADER_E_UPG_ZIPCOMPRESS_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_ZIPCOMPRESS_ERR, Desc: "Compressing DAV file failed."}
	case MV_CODEREADER_E_UPG_GET_PROGRESS_TIMEOUT_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_GET_PROGRESS_TIMEOUT_ERR, Desc: "Upgrade progress acquinsition timed out."}
	case MV_CODEREADER_E_UPG_SEND_QUERY_PROGRESS_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_SEND_QUERY_PROGRESS_ERR, Desc: "Failed to send progress query instruction"}
	case MV_CODEREADER_E_UPG_RECV_QUERY_PROGRESS_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_RECV_QUERY_PROGRESS_ERR, Desc: "Failed to receive progress query instruction͘"}
	case MV_CODEREADER_E_UPG_GET_QUERY_PROGRESS_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_GET_QUERY_PROGRESS_ERR, Desc: "Getting query progress failed."}
	case MV_CODEREADER_E_UPG_GET_MAX_QUERY_PROGRESS_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_GET_MAX_QUERY_PROGRESS_ERR, Desc: "Getting fastest progress failed."}
	case MV_CODEREADER_E_UPG_CHECKT_PACKET_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_CHECKT_PACKET_FAILED, Desc: "Verifying file failed."}
	case MV_CODEREADER_E_UPG_FPGA_PROGRAM_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_FPGA_PROGRAM_FAILED, Desc: "FPGA program upgrade failed."}
	case MV_CODEREADER_E_UPG_WATCHDOG_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_WATCHDOG_FAILED, Desc: "Watchdog upgrade failed."}
	case MV_CODEREADER_E_UPG_CAMERA_AND_BARE_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_CAMERA_AND_BARE_FAILED, Desc: "Bare camera upgrade failed."}
	case MV_CODEREADER_E_UPG_RETAIN_CONFIG_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_RETAIN_CONFIG_FAILED, Desc: "Retaining configuration file failed."}
	case MV_CODEREADER_E_UPG_FPGA_DRIVER_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_FPGA_DRIVER_FAILED, Desc: "FPGA drive upgrade failed."}
	case MV_CODEREADER_E_UPG_SPI_DRIVER_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_SPI_DRIVER_FAILED, Desc: "SPI drive upgrade failed."}
	case MV_CODEREADER_E_UPG_REBOOT_SYSTEM_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_REBOOT_SYSTEM_FAILED, Desc: "Restarting failed."}
	case MV_CODEREADER_E_UPG_UPSELF_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_UPSELF_FAILED, Desc: "Service upgrade failed."}
	case MV_CODEREADER_E_UPG_STOP_RELATION_PROGRAM_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_STOP_RELATION_PROGRAM_FAILED, Desc: "Stop related service failed."}
	case MV_CODEREADER_E_UPG_DEVCIE_TYPE_INCONSISTENT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_DEVCIE_TYPE_INCONSISTENT, Desc: "Inconsistent device type"}
	case MV_CODEREADER_E_UPG_READ_ENCRYPT_INFO_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_READ_ENCRYPT_INFO_FAILED, Desc: "Read encryption message failed."}
	case MV_CODEREADER_E_UPG_PLAT_TYPE_INCONSISTENT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_PLAT_TYPE_INCONSISTENT, Desc: "Incorrect device platform"}
	case MV_CODEREADER_E_UPG_CAMERA_TYPE_INCONSISTENT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_CAMERA_TYPE_INCONSISTENT, Desc: "Incorrect camera model"}
	case MV_CODEREADER_E_UPG_DEVICE_UPGRADING:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_DEVICE_UPGRADING, Desc: "Incorrect camera model"}
	case MV_CODEREADER_E_UPG_UNZIP_FAILED:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_UNZIP_FAILED, Desc: "Upgrade package decompression failed."}
	case MV_CODEREADER_E_UPG_BLE_DISCONNECT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_BLE_DISCONNECT, Desc: "PDA bluetooth is disconnected."}
	case MV_CODEREADER_E_UPG_BATTERY_NOTENOUGH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_BATTERY_NOTENOUGH, Desc: "Low battery"}
	case MV_CODEREADER_E_UPG_RTC_NOT_PRESENT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_RTC_NOT_PRESENT, Desc: "PDA is not on the base."}
	case MV_CODEREADER_E_UPG_APP_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_APP_ERR, Desc: "Failed to upgrade the app."}
	case MV_CODEREADER_E_UPG_L3_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_L3_ERR, Desc: "Failed to upgrade L3."}
	case MV_CODEREADER_E_UPG_MCU_ERR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_MCU_ERR, Desc: "Failed to upgrade MCU."}
	case MV_CODEREADER_E_UPG_PLATFORM_DISMATCH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_PLATFORM_DISMATCH, Desc: "Platform does not match."}
	case MV_CODEREADER_E_UPG_TYPE_DISMATCH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_TYPE_DISMATCH, Desc: "Model does not match."}
	case MV_CODEREADER_E_UPG_SPACE_DISMATCH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_SPACE_DISMATCH, Desc: "Space does not match."}
	case MV_CODEREADER_E_UPG_MEM_DISMATCH:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_MEM_DISMATCH, Desc: "Memory does not match."}
	case MV_CODEREADER_E_UPG_NET_TRANS_ERROR:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_NET_TRANS_ERROR, Desc: "Network transmission exception.Please upgrade again."}
	case MV_CODEREADER_E_UPG_UNKNOW:
		return MvCodeReaderError{Code: MV_CODEREADER_E_UPG_UNKNOW, Desc: "Upgrade unknown error"}
	case MV_CODEREADER_E_CREAT_SOCKET:
		return MvCodeReaderError{Code: MV_CODEREADER_E_CREAT_SOCKET, Desc: "Creating socket error"}
	case MV_CODEREADER_E_BIND_SOCKET:
		return MvCodeReaderError{Code: MV_CODEREADER_E_BIND_SOCKET, Desc: "Binding error"}
	case MV_CODEREADER_E_CONNECT_SOCKET:
		return MvCodeReaderError{Code: MV_CODEREADER_E_CONNECT_SOCKET, Desc: "Connection error"}
	case MV_CODEREADER_E_GET_HOSTNAME:
		return MvCodeReaderError{Code: MV_CODEREADER_E_GET_HOSTNAME, Desc: "Host name getting error"}
	case MV_CODEREADER_E_NET_WRITE:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NET_WRITE, Desc: "Data write error"}
	case MV_CODEREADER_E_NET_READ:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NET_READ, Desc: "Data read error"}
	case MV_CODEREADER_E_NET_SELECT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NET_SELECT, Desc: "Selection error"}
	case MV_CODEREADER_E_NET_TIMEOUT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NET_TIMEOUT, Desc: "Timed out"}
	case MV_CODEREADER_E_NET_ACCEPT:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NET_ACCEPT, Desc: "Receive error"}
	case MV_CODEREADER_E_NET_UNKNOW:
		return MvCodeReaderError{Code: MV_CODEREADER_E_NET_UNKNOW, Desc: "Network unknown erro"}
	}

	return ErrUnknown

}
