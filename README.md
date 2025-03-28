# Toasty

The goal of this project is to create a simple HTTP server that will play audio
locally on the machine the server is running on.

Initially with audio files uploaded to the server, then using AI to TTS text in
request payloads and play it on the device.

Eventially I want to add audio input functionality to enable a mesh of these
servers to work as in intercomm system.

## Dependencies

### Install Dependencies on Ubuntu

```bash
sudo apt update;
sudo apt install pulseaudio libpulse-dev pulseaudio-utils;

```
