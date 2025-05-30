package pulseaudio

type VolumeData struct {
	Value        int    `json:"value"`
	ValuePercent string `json:"value_percent"`
	Db           string `json:"db"`
}

type latency struct {
	Actual     float64 `json:"actual"`
	Configured float64 `json:"configured"`
}

type port struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Type              string `json:"type"`
	Priority          int    `json:"priority"`
	AvailabilityGroup string `json:"availability_group"`
	Availability      string `json:"availability"`
}

type audioDevice struct {
	Index               int                   `json:"index"`
	State               string                `json:"state"`
	Name                string                `json:"name"`
	Description         string                `json:"description"`
	Driver              string                `json:"driver"`
	SampleSpecification string                `json:"sample_specification"`
	ChannelMap          string                `json:"channel_map"`
	OwnerModule         int                   `json:"owner_module"`
	Mute                bool                  `json:"mute"`
	Volume              map[string]VolumeData `json:"volume"`
	BaseVolume          VolumeData            `json:"base_volume"`
	Balance             float64               `json:"balance"`
	Latency             latency               `json:"latency"`
	Flags               []string              `json:"flags"`
	Properties          map[string]string     `json:"properties"`
	Ports               []port                `json:"ports"`
	ActivePort          string                `json:"active_port"`
	Formats             []string              `json:"formats"`
}
