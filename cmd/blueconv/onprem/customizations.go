package blueprint

type Customizations struct {
	Hostname           *string                        `json:"hostname,omitempty" toml:"hostname,omitempty"`
	Kernel             *KernelCustomization           `json:"kernel,omitempty" toml:"kernel,omitempty"`
	User               []UserCustomization            `json:"user,omitempty" toml:"user,omitempty"`
	Group              []GroupCustomization           `json:"group,omitempty" toml:"group,omitempty"`
	Timezone           *TimezoneCustomization         `json:"timezone,omitempty" toml:"timezone,omitempty"`
	Locale             *LocaleCustomization           `json:"locale,omitempty" toml:"locale,omitempty"`
	Firewall           *FirewallCustomization         `json:"firewall,omitempty" toml:"firewall,omitempty"`
	Services           *ServicesCustomization         `json:"services,omitempty" toml:"services,omitempty"`
	Filesystem         []FilesystemCustomization      `json:"filesystem,omitempty" toml:"filesystem,omitempty"`
	Disk               *DiskCustomization             `json:"disk,omitempty" toml:"disk,omitempty"`
	InstallationDevice string                         `json:"installation_device,omitempty" toml:"installation_device,omitempty"`
	FDO                *FDOCustomization              `json:"fdo,omitempty" toml:"fdo,omitempty"`
	OpenSCAP           *OpenSCAPCustomization         `json:"openscap,omitempty" toml:"openscap,omitempty"`
	Ignition           *IgnitionCustomization         `json:"ignition,omitempty" toml:"ignition,omitempty"`
	Directories        []DirectoryCustomization       `json:"directories,omitempty" toml:"directories,omitempty"`
	Files              []FileCustomization            `json:"files,omitempty" toml:"files,omitempty"`
	Repositories       []RepositoryCustomization      `json:"repositories,omitempty" toml:"repositories,omitempty"`
	FIPS               *bool                          `json:"fips,omitempty" toml:"fips,omitempty"`
	ContainersStorage  *ContainerStorageCustomization `json:"containers-storage,omitempty" toml:"containers-storage,omitempty"`
	Installer          *InstallerCustomization        `json:"installer,omitempty" toml:"installer,omitempty"`
	RPM                *RPMCustomization              `json:"rpm,omitempty" toml:"rpm,omitempty"`
	RHSM               *RHSMCustomization             `json:"rhsm,omitempty" toml:"rhsm,omitempty"`
	CACerts            *CACustomization               `json:"cacerts,omitempty" toml:"cacerts,omitempty"`
}

type IgnitionCustomization struct {
	Embedded  *EmbeddedIgnitionCustomization  `json:"embedded,omitempty" toml:"embedded,omitempty"`
	FirstBoot *FirstBootIgnitionCustomization `json:"firstboot,omitempty" toml:"firstboot,omitempty"`
}

type EmbeddedIgnitionCustomization struct {
	Config string `json:"config,omitempty" toml:"config,omitempty"`
}

type FirstBootIgnitionCustomization struct {
	ProvisioningURL string `json:"url,omitempty" toml:"url,omitempty"`
}

type FDOCustomization struct {
	ManufacturingServerURL string `json:"manufacturing_server_url,omitempty" toml:"manufacturing_server_url,omitempty"`
	DiunPubKeyInsecure     string `json:"diun_pub_key_insecure,omitempty" toml:"diun_pub_key_insecure,omitempty"`
	// This is the output of:
	// echo "sha256:$(openssl x509 -fingerprint -sha256 -noout -in diun_cert.pem | cut -d"=" -f2 | sed 's/://g')"
	DiunPubKeyHash          string `json:"diun_pub_key_hash,omitempty" toml:"diun_pub_key_hash,omitempty"`
	DiunPubKeyRootCerts     string `json:"diun_pub_key_root_certs,omitempty" toml:"diun_pub_key_root_certs,omitempty"`
	DiMfgStringTypeMacIface string `json:"di_mfg_string_type_mac_iface,omitempty" toml:"di_mfg_string_type_mac_iface,omitempty"`
}

type KernelCustomization struct {
	Name   string `json:"name,omitempty" toml:"name,omitempty"`
	Append string `json:"append" toml:"append"`
}

type SSHKeyCustomization struct {
	User string `json:"user" toml:"user"`
	Key  string `json:"key" toml:"key"`
}

type UserCustomization struct {
	Name               string   `json:"name" toml:"name"`
	Description        *string  `json:"description,omitempty" toml:"description,omitempty"`
	Password           *string  `json:"password,omitempty" toml:"password,omitempty"`
	Key                *string  `json:"key,omitempty" toml:"key,omitempty"`
	Home               *string  `json:"home,omitempty" toml:"home,omitempty"`
	Shell              *string  `json:"shell,omitempty" toml:"shell,omitempty"`
	Groups             []string `json:"groups,omitempty" toml:"groups,omitempty"`
	UID                *int     `json:"uid,omitempty" toml:"uid,omitempty"`
	GID                *int     `json:"gid,omitempty" toml:"gid,omitempty"`
	ExpireDate         *int     `json:"expiredate,omitempty" toml:"expiredate,omitempty"`
	ForcePasswordReset *bool    `json:"force_password_reset,omitempty" toml:"force_password_reset,omitempty"`
}

type GroupCustomization struct {
	Name string `json:"name" toml:"name"`
	GID  *int   `json:"gid,omitempty" toml:"gid,omitempty"`
}

type TimezoneCustomization struct {
	Timezone   *string  `json:"timezone,omitempty" toml:"timezone,omitempty"`
	NTPServers []string `json:"ntpservers,omitempty" toml:"ntpservers,omitempty"`
}

type LocaleCustomization struct {
	Languages []string `json:"languages,omitempty" toml:"languages,omitempty"`
	Keyboard  *string  `json:"keyboard,omitempty" toml:"keyboard,omitempty"`
}

type FirewallCustomization struct {
	Ports    []string                       `json:"ports,omitempty" toml:"ports,omitempty"`
	Services *FirewallServicesCustomization `json:"services,omitempty" toml:"services,omitempty"`
	Zones    []FirewallZoneCustomization    `json:"zones,omitempty" toml:"zones,omitempty"`
}

type FirewallZoneCustomization struct {
	Name    *string  `json:"name,omitempty" toml:"name,omitempty"`
	Sources []string `json:"sources,omitempty" toml:"sources,omitempty"`
}

type FirewallServicesCustomization struct {
	Enabled  []string `json:"enabled,omitempty" toml:"enabled,omitempty"`
	Disabled []string `json:"disabled,omitempty" toml:"disabled,omitempty"`
}

type ServicesCustomization struct {
	Enabled  []string `json:"enabled,omitempty" toml:"enabled,omitempty"`
	Disabled []string `json:"disabled,omitempty" toml:"disabled,omitempty"`
	Masked   []string `json:"masked,omitempty" toml:"masked,omitempty"`
}

type OpenSCAPCustomization struct {
	DataStream    string                               `json:"datastream,omitempty" toml:"datastream,omitempty"`
	ProfileID     string                               `json:"profile_id,omitempty" toml:"profile_id,omitempty"`
	Tailoring     *OpenSCAPTailoringCustomizations     `json:"tailoring,omitempty" toml:"tailoring,omitempty"`
	JSONTailoring *OpenSCAPJSONTailoringCustomizations `json:"json_tailoring,omitempty" toml:"json_tailoring,omitempty"`
}

type OpenSCAPTailoringCustomizations struct {
	Selected   []string `json:"selected,omitempty" toml:"selected,omitempty"`
	Unselected []string `json:"unselected,omitempty" toml:"unselected,omitempty"`
}

type OpenSCAPJSONTailoringCustomizations struct {
	ProfileID string `json:"profile_id,omitempty" toml:"profile_id,omitempty"`
	Filepath  string `json:"filepath,omitempty" toml:"filepath,omitempty"`
}

// Configure the container storage separately from containers, since we most likely would
// like to use the same storage path for all of the containers.
type ContainerStorageCustomization struct {
	// destination is always `containers-storage`, so we won't expose this
	StoragePath *string `json:"destination-path,omitempty" toml:"destination-path,omitempty"`
}

type CACustomization struct {
	PEMCerts []string `json:"pem_certs,omitempty" toml:"pem_certs,omitempty"`
}
