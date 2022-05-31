package lib

import (
	"time"

	"github.com/xlzd/gotp"
)

func OTPGenerator() string {
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	return totp.At(int(time.Now().Unix()))
}
