package parse

type OutPut string

const (
	OutPutOnly     OutPut = "OUTPUT_ONLY"
	OutPutOptional OutPut = "OUTPUT_OPTIONAL"
	OutPutRequired OutPut = "REQUIRED"
	OutPutInhibit  OutPut = "INHIBIT"
)
