package pulseaudio

import (
	"context"
	"fmt"
	"os"
	"testing"
)

type mockPulseAudioDeviceManager struct {

}

func TestParseStuff(t *testing.T) {
	cmdOutput, err := os.ReadFile("../../sample_command_output/pactl_list_sources_output.json")
	if err != nil {
		t.Errorf("failed to read command output file: %s", err)
	}
	pad := NewPulseAudioDeviceManager()
	ctx := context.TODO()
	pad.ListAudioDevices(ctx)
	fmt.Print("t: This is a test line\n")
}
