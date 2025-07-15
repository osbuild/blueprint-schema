package ubp

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPopulateDefaults(t *testing.T) {
	ubpDefaults := &Blueprint{
		Containers: []Container{
			{
				Name:   "container",
				Source: "source",
			},
		},
		DNF: &DNF{
			Repositories: []DNFRepository{{ID: "repo"}},
		},
		FSNodes: []FSNode{
			{
				Path: "file",
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
					{
						union: []byte(`{"port": 22}`),
					},
					{
						union: []byte(`{"from": 200, "to": 300}`),
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
			name: "container-tls-verify",
			val: func(ubp *Blueprint) any {
				return *ubp.Containers[0].TLSVerify
			},
			want: true,
		},
		{
			name: "dnf-repo-tls-verify",
			val: func(ubp *Blueprint) any {
				return *ubp.DNF.Repositories[0].TLSVerify
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
			name: "network-firewall-port-protocol",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[1].AsFirewallPort()
				if err != nil {
					t.Fatalf("firewall as call failed: %v", err)
				}
				return s.Protocol
			},
			want: ProtocolAny,
		},
		{
			name: "network-firewall-service-protocol",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[0].AsFirewallService()
				if err != nil {
					t.Fatalf("firewall as call failed: %v", err)
				}
				return s.Protocol
			},
			want: ProtocolAny,
		},
		{
			name: "network-firewall-from-to-protocol",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[2].AsFirewallFromTo()
				if err != nil {
					t.Fatalf("firewall as call failed: %v", err)
				}
				return s.Protocol
			},
			want: ProtocolAny,
		},
		{
			name: "network-firewall-port-enabled",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[1].AsFirewallPort()
				if err != nil {
					t.Fatalf("firewall as call failed: %v", err)
				}
				return *s.Enabled
			},
			want: true,
		},
		{
			name: "network-firewall-service-enabled",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[0].AsFirewallService()
				if err != nil {
					t.Fatalf("firewall as call failed: %v", err)
				}
				return *s.Enabled
			},
			want: true,
		},
		{
			name: "network-firewall-from-to-enabled",
			val: func(ubp *Blueprint) any {
				s, err := ubp.Network.Firewall.Services[2].AsFirewallFromTo()
				if err != nil {
					t.Fatalf("firewall as call failed: %v", err)
				}
				return *s.Enabled
			},
			want: true,
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
	"containers": [
		{
			"name": "container",
			"source": "source"
		}
	],
	"dnf": {
		"repositories": [
			{
				"id": "repo"
			}
		]
	},
	"fsnodes": [
		{
			"path": "file"
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
				},
				{
					"port": 22
				},
				{
					"from": 200,
					"to": 300
				}
			]
		}
	}
}`

	if diff := cmp.Diff(want, string(data)); diff != "" {
		t.Errorf("MarshalJSON defaults mismatch (-want +got):\n%s", diff)
	}
}
