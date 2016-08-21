Package globs is used to get and set some global options such as last log ID. It has 2 functions to, accordingly, get and set last log ID and one constant *InitialPrev *which is the default "previous_log_id" value for the first log record in the database (for which there is not previous log id at all).

Endpoints:
func GetLastLogId() string
func SetLastLogId(id string)
const InitialPrev string

Usage:

	import "github.com/cryptobrains/CryptoLog/globs"

	lastLogId := globs.GetLastLogId()
	globs.SetLastLogId("theNewValueOfLastLogId")
	somevar := globs.InitialPrev

