package ubp

import (
	"encoding/json"
	"testing"
)

func TestPopulateDefaults(t *testing.T) {
	ubpDefaults := &Blueprint{
		DNF: &DNF{
			Repositories: []DNFRepository{{ID: "repo"}},
		},
		FSNodes: []FSNode{
			{
				Path: "file",
				Type: "file",
			},
			{
				Path: "dir",
				Type: "dir",
			},
		},
		Network: &Network{
			Firewall: &NetworkFirewall{
				Services: []NetworkService{
					{
						union: []byte(`{"name": "ssh"}`),
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		val  func(*Blueprint) any
		want any
	}{
		{
			name: "dnf-repo-ssl-verify",
			val: func(ubp *Blueprint) any {
				return *ubp.DNF.Repositories[0].SSLVerify
			},
			want: true,
		},
		{
			name: "dnf-repo-usage-configure",
			val: func(ubp *Blueprint) any {
				return *ubp.DNF.Repositories[0].Usage.Configure
			},
			want: true,
		},
		{
			name: "dnf-repo-usage-install",
			val: func(ubp *Blueprint) any {
				return *ubp.DNF.Repositories[0].Usage.Install
			},
			want: true,
		},
		{
			name: "fsnode-mode-file",
			val: func(ubp *Blueprint) any {
				return ubp.FSNodes[0].Mode
			},
			want: DefaultFileFSNodeMode,
		},
		{
			name: "fsnode-mode-dir",
			val: func(ubp *Blueprint) any {
				return ubp.FSNodes[1].Mode
			},
			want: DefaultDirFSNodeMode,
		},
		{
			name: "network-firewall-service-name",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[0].AsFirewallPort()
				if err != nil {
					t.Fatalf("AsFirewallPort failed: %v", err)
				}
				return s.Protocol
			},
			want: ProtocolAny,
		},
	}

	if err := PopulateDefaults(nil); err != nil {
		t.Errorf("PopulateDefaults should not return an error for nil input: %v", err)
	}

	if err := PopulateDefaults(ubpDefaults); err != nil {
		t.Errorf("PopulateDefaults returned an error: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val(ubpDefaults); got != tt.want {
				t.Errorf("PopulateDefaults() = %v, want %v", got, tt.want)
			}
		})
	}

	data, err := json.MarshalIndent(ubpDefaults, "", "\t")
	if err != nil {
		t.Fatalf("MarshalJSON defaults failed: %v", err)
	}

	want := `{
	"dnf": {
		"repositories": [
			{
				"id": "repo"
			}
		]
	},
	"fsnodes": [
		{
			"path": "file",
			"type": "file"
		},
		{
			"path": "dir",
			"type": "dir"
		}
	],
	"network": {
		"firewall": {
			"services": [
				{
					"name": "ssh"
				}
			]
		}
	}
}`

	if string(data) != want {
		t.Errorf("MarshalJSON defaults = %s, want %s", data, want)
	}
}
