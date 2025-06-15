package pulseaudio

import (
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
	devices, err := parsePactlOutputForListSinks(cmdOutput)
	if err != nil {
		t.Errorf("failed to parse example output: %s", err)
	}
	fmt.Printf("devices: %+v", devices)
	for i := range devices {
		dtoDevice, err := pulseAudioDeviceToDTO(devices[i])
		if err != nil {
			t.Errorf("failed to convert pulse audio device to dto device: %s", err)
		}
		fmt.Printf("dto device %d: %+v", i, dtoDevice)
	}
}
