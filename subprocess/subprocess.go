package subprocess

import "net/url"

type Subprocess struct {
	Active chan bool
	Kill   chan bool
	Url    *url.URL
}
