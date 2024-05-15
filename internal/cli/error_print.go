package cli

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func detailedErrorInformation(resp *resty.Response) string {
	var info string
	// Explore response object
	info += "Response Info:\n"
	info += fmt.Sprintf("  Status Code: %v\n", resp.StatusCode())
	info += fmt.Sprintf("  Status     : %v\n", resp.Status())
	info += fmt.Sprintf("  Proto      : %v\n", resp.Proto())
	info += fmt.Sprintf("  Time       : %v\n", resp.Time())
	info += fmt.Sprintf("  Received At: %v\n", resp.ReceivedAt())
	info += fmt.Sprintf("  Body       : %v\n", resp)

	// Explore trace info
	info += "Request Trace Info:"
	ti := resp.Request.TraceInfo()
	info += fmt.Sprintf("  DNSLookup     : %v\n", ti.DNSLookup)
	info += fmt.Sprintf("  ConnTime      : %v\n", ti.ConnTime)
	info += fmt.Sprintf("  TCPConnTime   : %v\n", ti.TCPConnTime)
	info += fmt.Sprintf("  TLSHandshake  : %v\n", ti.TLSHandshake)
	info += fmt.Sprintf("  ServerTime    : %v\n", ti.ServerTime)
	info += fmt.Sprintf("  ResponseTime  : %v\n", ti.ResponseTime)
	info += fmt.Sprintf("  TotalTime     : %v\n", ti.TotalTime)
	info += fmt.Sprintf("  IsConnReused  : %v\n", ti.IsConnReused)
	info += fmt.Sprintf("  IsConnWasIdle : %v\n", ti.IsConnWasIdle)
	info += fmt.Sprintf("  ConnIdleTime  : %v\n", ti.ConnIdleTime)
	info += fmt.Sprintf("  RequestAttempt: %v\n", ti.RequestAttempt)
	// fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
	return info
}
