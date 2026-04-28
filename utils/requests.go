package utils

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetIpFromRequest Gets the client's proper ip address from a request.
func GetIpFromRequest(c *gin.Context) string {
	headers := []string{
		"CF-Connecting-IP",
		"X-Real-IP",
	}

	for _, h := range headers {
		if ip := strings.TrimSpace(c.GetHeader(h)); ip != "" {
			if parsed := net.ParseIP(ip); parsed != nil {
				return parsed.String()
			}
		}
	}

	if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
		ip := strings.TrimSpace(strings.Split(xff, ",")[0])
		if parsed := net.ParseIP(ip); parsed != nil {
			return parsed.String()
		}
	}

	if host, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
		if parsed := net.ParseIP(host); parsed != nil {
			return parsed.String()
		}
	}

	if parsed := net.ParseIP(strings.TrimSpace(c.Request.RemoteAddr)); parsed != nil {
		return parsed.String()
	}

	return "::1"
}
