package wechat

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_GetUUID(t *testing.T) {
	wxlogger := log.New(os.Stdout, "[wechat]", log.LstdFlags)
	we := NewWechat(wxlogger)
	we.GetUUID()
	fmt.Println(we.Uuid)
	t.Logf("%v", we.Uuid)
}
