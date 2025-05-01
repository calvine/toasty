package pulseaudio

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/calvine/toasty/audio"
)

type pulseAudioDeviceManager struct {
}

func NewPulseAudioDeviceManager() audio.AudioDeviceManager {
	return &pulseAudioDeviceManager{}
}

func parsePACtlListSourceSinkOutput(cmdOutput []byte) ([]audio.AudioDevice, error) {
	// `\n\n`gm
	emptyLineRegex := regexp.MustCompile(`(?m)^\s*$`)
	elResults := emptyLineRegex.FindAllIndex(cmdOutput, -1)
	fmt.Print(elResults)
	return nil, errors.New("TODO, Implement me")
}

func (p *pulseAudioDeviceManager) ListAudioDevices(ctx context.Context) ([]audio.AudioDevice, error) {
	// TODO: implement an interface to check availability of underlying utility in this case like pactl.
	cmd := exec.CommandContext(ctx, "pactl", "list", "sinks")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list sinks: %w", err)
	}
	devices, err := parsePACtlListSourceSinkOutput(output)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pactl list sinks output: %w", err)
	}
	fmt.Println(output)
	return devices, nil
}
