package parse

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadYAMLWriteJSON(t *testing.T) {
	// XXX: this will fail with 1.24 which should clean up the output via omitzero
	in := "name: test\n"
	want := `{
  "accounts": {
    "groups": null,
    "users": null
  },
  "dnf": {},
  "fips": {},
  "ignition": null,
  "installer": {
    "anaconda": {},
    "coreos": {}
  },
  "kernel": {},
  "locale": {},
  "name": "test",
  "network": {
    "firewall": {}
  },
  "openscap": {
    "profile_id": ""
  },
  "registration": {
    "fdo": {
      "manufacturing_server_url": ""
    },
    "redhat": {
      "connector": {
        "enabled": false
      },
      "insights": {
        "enabled": false
      },
      "subscription_manager": {}
    }
  },
  "storage": {
    "partitions": null,
    "type": ""
  },
  "systemd": {},
  "timedate": {}
}`

	// XXX: however storing back YAML will not respect omitzero and this
	// needs to be fixed (relevant only for the convertor)
	want2 := `accounts:
  groups: null
  users: null
dnf: {}
fips: {}
ignition: null
installer:
  anaconda: {}
  coreos: {}
kernel: {}
locale: {}
name: test
network:
  firewall: {}
openscap:
  profile_id: ""
registration:
  fdo:
    manufacturing_server_url: ""
  redhat:
    connector:
      enabled: false
    insights:
      enabled: false
    subscription_manager: {}
storage:
  partitions: null
  type: ""
systemd: {}
timedate: {}
`
	b, err := ReadYAML(bytes.NewBufferString(in))
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("Expected data, got nil")
	}

	if b.Name != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name)
	}

	out, err := MarshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}

	out, err = MarshalYAML(b)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), want2); diff != "" {
		t.Fatalf("Unexpected YAML output: %s", diff)
	}
}

func TestReadJSONWriteYAML(t *testing.T) {
	// XXX: see above
	in := "{\n  \"name\": \"test\"\n}"
	want := `accounts:
  groups: null
  users: null
dnf: {}
fips: {}
ignition: null
installer:
  anaconda: {}
  coreos: {}
kernel: {}
locale: {}
name: test
network:
  firewall: {}
openscap:
  profile_id: ""
registration:
  fdo:
    manufacturing_server_url: ""
  redhat:
    connector:
      enabled: false
    insights:
      enabled: false
    subscription_manager: {}
storage:
  partitions: null
  type: ""
systemd: {}
timedate: {}
`
	want2 := `{
  "accounts": {
    "groups": null,
    "users": null
  },
  "dnf": {},
  "fips": {},
  "ignition": null,
  "installer": {
    "anaconda": {},
    "coreos": {}
  },
  "kernel": {},
  "locale": {},
  "name": "test",
  "network": {
    "firewall": {}
  },
  "openscap": {
    "profile_id": ""
  },
  "registration": {
    "fdo": {
      "manufacturing_server_url": ""
    },
    "redhat": {
      "connector": {
        "enabled": false
      },
      "insights": {
        "enabled": false
      },
      "subscription_manager": {}
    }
  },
  "storage": {
    "partitions": null,
    "type": ""
  },
  "systemd": {},
  "timedate": {}
}`

	b, err := ReadJSON(bytes.NewBufferString(in))
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("Expected data, got nil")
	}

	if b.Name != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name)
	}

	out, err := MarshalYAML(b)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), want); diff != "" {
		t.Fatalf("Unexpected YAML output: %s", diff)
	}

	out, err = MarshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), want2); diff != "" {
		t.Fatalf("Unexpected JSON output: %s", diff)
	}

	err = WriteJSON(b, bytes.NewBufferString(""), true)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), want2); diff != "" {
		t.Fatalf("Unexpected JSON output: %s", diff)
	}
}

func TestConvertJSONtoYAML(t *testing.T) {
	in := "{\n  \"name_test\": \"test\"\n}"
	want := "name_test: test\n"

	out, err := ConvertJSONtoYAML([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected YAML output: %s", out)
	}
}

func TestConvertYAMLtoJSON(t *testing.T) {
	in := "name_test: test\n"
	want := "{\"name_test\":\"test\"}"

	out, err := ConvertYAMLtoJSON([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}
}
