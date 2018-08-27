package license

import (
	"testing"
	"fmt"
)

func Test_License(t *testing.T){
	 GenLicense(10)
	 end,flag := VerifyLicense()
	 fmt.Println(end,flag)
}
