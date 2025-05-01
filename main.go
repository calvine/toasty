package main

import (
	"context"
	"fmt"

	"github.com/calvine/toasty/audio/pulseaudio"
)

func main() {
	adm := pulseaudio.NewPulseAudioDeviceManager()
	ctx := context.Background()
	_, err := adm.ListAudioDevices(ctx)
	if err != nil {
		fmt.Printf("Failed to list audio devices: %s", err)
	}
}
