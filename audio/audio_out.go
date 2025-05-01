package audio

type AudioOutProvider interface {
	// IsReady is used to check if an audio provider is available for use.
	IsReady() error
	// PlayFile plays an audio file at the path provided.
	PlayFile(path string) error
}
