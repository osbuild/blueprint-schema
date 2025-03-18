package blueprint

type TimeDate struct {
	// System time zone. Defaults to UTC.
	// To list available time zones run: timedatectl list-timezones
	Timezone string `json:"timezone,omitempty" jsonschema:"required,default=UTC"`

	// An optional list of strings containing NTP servers to use. If not provided the distribution defaults are used
	NTPServers []string `json:"ntp_servers,omitempty"`
}
