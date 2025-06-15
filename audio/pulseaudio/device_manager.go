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

var volumePrecentRegex = regexp.MustCompile(`\d{1,3}`)

type pulseAudioDeviceManager struct {
	outputDevices []audio.AudioDevice
	// TODO: input devices?
}

func NewPulseAudioDeviceManager() audio.AudioDeviceManager {
	outputDevices := make([]audio.AudioDevice, 0, 5)
	pad := &pulseAudioDeviceManager{
		outputDevices: outputDevices,
	}
	return pad
}

func pulseAudioDeviceToDTO(device audioDevice) (audio.AudioDevice, error) {
	dto := audio.AudioDevice{}
	dto.ID = strconv.Itoa(device.Index)
	dto.Description = device.Properties["device.description"]
	cvp, err := strconv.Atoi(volumePrecentRegex.FindString(device.BaseVolume.ValuePercent))
	if err != nil {
		return dto, fmt.Errorf("failed to get current volume percentage for device %s: %w", dto.ID, err)
	}
	dto.CurrentVolumePercent = int8(cvp)
	switch {
	case device.Name != "":
		dto.Name = device.Name
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

func parsePactlOutputForListSinks(output []byte) ([]audioDevice, error) {
	devices := make([]audioDevice, 0, 3)
	err := json.Unmarshal(output, &devices)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pactl list sinks output: %w", err)
	}
	return devices, nil
}

func (p *pulseAudioDeviceManager) getOutputDevices(ctx context.Context) ([]audioDevice, []audio.AudioDevice, error) {
	// TODO: implement an interface to check availability of underlying utility in this case like pactl.
	cmd := exec.CommandContext(ctx, "pactl", "list", "sinks")

	output, err := cmd.Output()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list sinks: %w", err)
	}
	devices, err := parsePactlOutputForListSinks(output)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal pactl output: %w", err)
	}
	dtoDevices := make([]audio.AudioDevice, 0, len(devices))
	for _, device := range devices {
		dtoDevice, err := pulseAudioDeviceToDTO(device)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert pulse audio device (%d) to DTO version: %w", device.Index, err)
		}
		dtoDevices = append(dtoDevices, dtoDevice)
	}
	p.outputDevices = dtoDevices
	return devices, dtoDevices, nil
}

func (p *pulseAudioDeviceManager) SetDefaultOutputDevice(ctx context.Context, id string) error {
	cmd := exec.CommandContext(ctx, "pactl", "set-default-sink", id)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to set default output device: (%s) %s - %w", output, id, err)
	}
	return nil
}

func (p *pulseAudioDeviceManager) ListOutputDevices(ctx context.Context) ([]audio.AudioDevice, error) {
	_, dtoDevices, err := p.getOutputDevices(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list output devices: %w", err)
	}
	return dtoDevices, nil
}

func (p *pulseAudioDeviceManager) SetOutputMute(ctx context.Context, id string, toggle, state bool) error {
	var cmd *exec.Cmd
	if toggle {
		cmd = exec.CommandContext(ctx, "pactl", "set-sink-mute", id, "toggle")
	} else {
		var muteState = "0"
		if state {
			muteState = "1"
		}
		cmd = exec.CommandContext(ctx, "pactl", "set-sink-mute", id, muteState)
	}
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to set output device mute: %s - %w", id, err)
	}
	return nil
}

func (p *pulseAudioDeviceManager) SetOutputVolume(ctx context.Context, id string, volume uint) error {
	cmd := exec.CommandContext(ctx, "pactl", "set-sink-volume", id, fmt.Sprintf("%d%%", volume))
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to set output device volume: %s - %w", id, err)
	}
	return nil
}

func (p *pulseAudioDeviceManager) AdjustOutputVolume(ctx context.Context, id string, relativeAmount int) error {
	_, dtoDevices, err := p.getOutputDevices(ctx)
	if err != nil {
		return fmt.Errorf("failed to get devices to adjust volume: %s - %w", id, err)
	}
	var targetDevice *audio.AudioDevice
	for i := range dtoDevices {
		if dtoDevices[i].ID == id {
			targetDevice = &dtoDevices[i]
			break
		}
	}
	if targetDevice == nil {
		return fmt.Errorf("could not find target device: %s", id)
	}
	currentVolume := targetDevice.CurrentVolumePercent
	newVolume := currentVolume + int8(relativeAmount)
	newVolume = util.Max(util.Min(newVolume, 100), 0)
	err = p.SetOutputVolume(ctx, id, uint(newVolume))
	// NOTE: no error created here because the error from the SetOutputVolume function is sufficient
	return err
}
