package audio

import "context"

type AudioDevice struct {
	ID          string
	Name        string
	Description string
	// This will be an integer between 0 and 100
	CurrentVolumePercent int8
	minVolume            int32
	maxVolume            int32
	currentVolume        int32
	DeviceProperties     map[string]string
	RawDevice            any
}

type AudioDeviceManager interface {
	ListAudioDevices(ctx context.Context) ([]AudioDevice, error)
	// MuteDevice(id string) error
	// SetDeviceVolume(id string, volume uint) error
	// AdjustDeviceVolume(id string, relativeAmount int) error
}
