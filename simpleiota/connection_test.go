package simpleiota

import (
	"encoding/json"
	"testing"
	"time"
)

func TestConnectionSend(t *testing.T) {
	c, err := NewConnection("http://node02.iotatoken.nl:14265", "SIERTBRUINSISBEZIGOMEENRONDJESAMENMETWIMAMENTTEMAKENOMZODESUBSIDIERONDTEKRIJGENH9")

	if err != nil {
		t.Error(err)
	}

	var someJSON struct {
		Id        int
		Message   string
		Timestamp time.Time
	}

	someJSON.Id = 12345
	someJSON.Message = "Hello world this is a JSON"
	someJSON.Timestamp = time.Now()

	stringifiedJSON, err := json.Marshal(someJSON)
	if err != nil {
		t.Error(err)
	}

	id, err := Send("RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY", 0, string(stringifiedJSON), c)
	if err != nil {
		t.Error(err)
	}

	t.Logf("TransactionId: %v\n", id)
}
