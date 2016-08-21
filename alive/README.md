After starting, alive package first check all existing named pipes in the specified folder sending them "ALIVE_QUERY" message. If one of them answers "TRUE", it panics and exits.

If no, it creates a named pipe in which receives "ALIVE_QUERY" queries of other alive processes and responds "TRUE" to them.

So only the one process can exist at a moment of time.

Usage:

	import "github.com/cryptobrains/CryptoLog/alive"

	//... somewhere at the beginning ...
	alive.Serve()

That's all.
