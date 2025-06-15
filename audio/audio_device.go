package audio

import "context"

// TODO: Add support for granular volume adjustment if possible. I am just figuring this stuff out, but device that have multiple sound outputs like stereo have more thatn one volume (i.e. left and right) I assume more sophisitcated audio devices can have even more. This is not a concern now, but in the future it might be?

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
	ListOutputDevices(ctx context.Context) ([]AudioDevice, error)
	SetDefaultOutputDevice(ctx context.Context, id string) error
	// SetOutputMute takes 3 parameters. The first is the id of the target device. Next is `toggle` which if true will ignore the thrid parameter and invert the current state of the target devices mute. Finally `state` if true unmutes the device and unmutes if false. The `state` parameter is ignored if `toggle` is true.
	SetOutputMute(ctx context.Context, id string, toggle, state bool) error
	SetOutputVolume(ctx context.Context, id string, volume uint) error
	AdjustOutputVolume(ctx context.Context, id string, relativeAmount int) error
}
