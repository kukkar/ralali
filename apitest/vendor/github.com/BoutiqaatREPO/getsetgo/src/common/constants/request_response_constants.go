package constants

const (

	// Request
	Request = "REQUEST"
	// RequestContext
	RequestContext = "REQUEST_CONTEXT"
	// RequestBodyParam
	RequestBodyParam = "REQUEST_BODY_PARAM"
	// RequestPathParam
	RequestPathParam = "REQUEST_PATH_PARAMETER"

	// Response
	Response = "RESPONSE"

	//Response Type
	ResponseType = "RESPONSE_TYPE"

	// ResponseMetaData
	ResponseMetaData = "RESPONSE_META_DATA"
	// ResponseData
	ResponseData          = "RESPONSE_DATA"
	ResponseStatus        = "RESPONSE_STATUS"
	ResponseHeadersConfig = "RESPONSE_HEADERS_CONFIG"
	APIResponse           = "API_RESPONSE"

	APPError = "APPERROR"

	HTTPVerb = "HTTPVERB"
	URI      = "URI"

	BucketID = "BUCKETID"

	Resource    = "RESOURCE"
	Version     = "VERSION"
	Action      = "ACTION"
	PathParams  = "PATH_PARAMS"
	QueryString = "URL_QUERY_STRING"

	NewRelicTransaction = "NewRelicTransaction"

	Result = "RESULT"

	UserAgent    = "USER_AGENT"
	HTTPReferrer = "HTTP_REFERRER"
	BucketsList  = "BUCKETSLIST"

	FieldSeperator    = ","
	KeyValueSeperator = ":"

	HealthCheckAPI  = "HEALTHCHECK"
	HealthCheckList = "HEALTH_CHECK_LIST"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

//constant for buckets and their values
const (
	OrchestratorBucketKey          = "Algo"
	OrchestratorBucketDefaultValue = "Old"
	OrchestratorBucketNewAlgo      = "New"
)

// constant for monitor
const (
	MonitorCustomMetric = "monitor_custom_metric"
)

const (
	RESPONSE_TYPE_JSON    = "json"
	RESPONSE_TYPE_CSV     = "csv"
	RESPONSE_TYPE_IMG_JPG = "image-jpeg"
	RESPONSE_TYPE_IMG_GIF = "image-gif"
	RESPONSE_TYPE_IMG_PNG = "image-png"
)
