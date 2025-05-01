package audio

import "context"

type AudioDevice struct {
	ID                   string
	Name                 string
	CurrentVolumePercent float32
	minVolume            int32
	maxVolume            int32
	currentVolume        int32
	RawMap               map[string]string
}

type AudioDeviceManager interface {
	ListAudioDevices(ctx context.Context) ([]AudioDevice, error)
	// MuteDevice(id string) error
	// SetDeviceVolume(id string, volume uint) error
	// AdjustDeviceVolume(id string, relativeAmount int) error
}
