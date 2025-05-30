package pulseaudio

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/calvine/toasty/audio"
	"github.com/calvine/toasty/util"
)

type pulseAudioDeviceManager struct {
	volumePrecentRegex *regexp.Regexp
}

func NewPulseAudioDeviceManager() audio.AudioDeviceManager {
	pad := &pulseAudioDeviceManager{}
	pad.volumePrecentRegex = regexp.MustCompile(`\d{1,3}`)
	return pad
}

func (p *pulseAudioDeviceManager) pulseAudioDeviceToDTO(device audioDevice) (audio.AudioDevice, error) {
	dto := audio.AudioDevice{}
	dto.ID = strconv.Itoa(device.Index)
	dto.Description = device.Properties["device.description"]
	cvp, err := strconv.Atoi(p.volumePrecentRegex.FindString(device.BaseVolume.ValuePercent))
	if err != nil {
		return dto, fmt.Errorf("failed to get current volume percentage for device %d: %w", dto.ID, err)
	}
	dto.CurrentVolumePercent = int8(cvp)
	switch {
	case util.MapContains(device.Properties, "device.product.name"):
		dto.Name = device.Properties["device.product.name"]
	case util.MapContains(device.Properties, "alsa.card_name"):
		dto.Name = device.Properties["alsa.card_name"]
	case util.MapContains(device.Properties, "alsa.long_card_name"):
		dto.Name = device.Properties["alsa.long_card_name"]
	default:
		dto.Name = "Unknown Device"
	}
	dto.DeviceProperties = device.Properties
	dto.RawDevice = device
	return dto, nil
}

func (p *pulseAudioDeviceManager) ListAudioDevices(ctx context.Context) ([]audio.AudioDevice, error) {
	// TODO: implement an interface to check availability of underlying utility in this case like pactl.
	cmd := exec.CommandContext(ctx, "pactl", "list", "sinks")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list sinks: %w", err)
	}
	devices := make([]audioDevice, 0, 3)
	err = json.Unmarshal(output, &devices)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pactl list sinks output: %w", err)
	}
	fmt.Println(output)
	dtoDevices := make([]audio.AudioDevice, 0, len(devices))
	for _, device := range devices {
		dtoDevice, err := p.pulseAudioDeviceToDTO(device)
		if err != nil {
			return nil, fmt.Errorf("failed to convert pulse audio device (%d) to DTO version: %w", device.Index, err)
		}
		dtoDevices = append(dtoDevices, dtoDevice)
	}
	return dtoDevices, nil
}
