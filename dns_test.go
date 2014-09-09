package dns

import "testing"
import "time"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestLookupHostTimeout(t *testing.T) {
	{
		_, err := LookupHostTimeout("foo.segment.io", 5*time.Second)
		if err.Error() != "host lookup timed out" {
			t.Fail()
		}
	}

	{
		hosts, err := LookupHostTimeout("segment.io", 5*time.Second)

		if err != nil {
			t.Fail()
		}

		if len(hosts) == 0 {
			t.Fail()
		}
	}
}
