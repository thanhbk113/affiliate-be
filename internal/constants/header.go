package constants

// Header keys
const (
	HeaderOrigin         = "Origin"
	HeaderContentLength  = "Content-Length"
	HeaderContentType    = "Content-Type"
	HeaderAuthorization  = "Authorization"
	HeaderAcceptLanguage = "Accept-Language"
	ResponseType         = "responseType"
	HeaderVersion        = "version"
	HeaderApiKey         = "Api-Key"
	HeaderAppName        = "AppName"
	HeaderOSName         = "OS-NAME"
	HeaderOSVersion      = "OS-VERSION"
	HeaderPlatform       = "PLATFORM"
	HeaderDeviceType     = "DEVICE-TYPE"
	HeaderBrowserName    = "BROWSER-NAME"
	HeaderBrowserVersion = "BROWSER-VERSION"
	HeaderAppVersion     = "App-Version"
	HeaderAppVersionCode = "App-Version-Code"
	HeaderUserAgent      = "User-Agent"
	HeaderFCMToken       = "Fcm-Token"
	HeaderModel          = "Model"
	HeaderDeviceID       = "Device-ID"
	HeaderIP             = "IP"
	HeaderSource         = "Source"
)

var (
	ListHeaderAllow = []string{
		HeaderOrigin,
		HeaderContentLength,
		HeaderContentType,
		HeaderAuthorization,
		HeaderAcceptLanguage,
		ResponseType,
		HeaderVersion,
		HeaderApiKey,
		HeaderOSName,
		HeaderOSVersion,
		HeaderPlatform,
		HeaderDeviceType,
		HeaderBrowserName,
		HeaderBrowserVersion,
		HeaderAppVersion,
		HeaderAppVersionCode,
		HeaderUserAgent,
		HeaderFCMToken,
		HeaderModel,
		HeaderDeviceID,
		HeaderIP,
		HeaderSource,
		HeaderAppName,
	}
)
