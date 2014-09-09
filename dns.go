package dns

import "strings"
import "time"
import "net"
import "fmt"

// LookupHostTimeout looks up host `addr` and will retry until the timeout is reached.
func LookupHostTimeout(addr string, timeout time.Duration) ([]string, error) {
	t := time.After(timeout)

	for {
		select {
		case <-t:
			return nil, fmt.Errorf("host lookup timed out")
		default:
			addrs, err := net.LookupHost(addr)

			if err == nil {
				return addrs, nil
			}

			if strings.Contains(err.Error(), "no such host") {
				time.Sleep(time.Second)
				continue
			}

			return nil, err
		}
	}
}
