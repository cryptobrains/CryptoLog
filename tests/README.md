test example:

import tests "github.com/cryptobrains/CryptoLog/tests"

...
	secureLog := securelog.New(100, 2000)			// create new server
	err := tests.VerifySequence(secureLog, 1500, true);	// fill and read 1500 messages,
								//true means that there will be no incorrect messages purposely injected
								// so the test should be passed;
								//if false one incorrect message will be injected and test should show an error
...

VerifySequence test checks if all records in the DB are properly signed and satusfy the sequence 'id of the record must be previous_log_id of the following record'
