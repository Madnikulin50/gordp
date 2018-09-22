package gordp

import "testing"

func Test_RdpClient(t *testing.T) {
	params := RdpConnectionParams{
		"test",
		"user",
		"123"}
	rdp := NewRdpConnectionBase()
	err := rdp.Connect(&params)
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
}
