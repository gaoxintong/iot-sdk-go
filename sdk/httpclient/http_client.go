package httpclient

import (
	"net/http"
	"time"
)

// DefaultClient 默认的 http 客户端
var DefaultClient = http.Client{
	Timeout: 5 * time.Second,
}
