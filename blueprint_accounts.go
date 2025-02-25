package blueprint

import "time"

// accounts:
//
//	users:
//	  - name: "lzap"
//	    description: ""
//	    password: ""
//	    ssh_keys:
//	      - "ssh-key 1234"
//	    home: "/home/lzap"
//	    shell: "/usr/bin/bash"
//	    uid: 1001
//	    gid: 1001
//	    groups: ["wheel", "operators"]
//	    expires: 2050-05-13
//	groups:
//	  - name: "operators"
//	    gid: 1042
type Accounts struct {
	// Operating system user accounts to be created on the image.
	Users []UserAccount `json:"users,omitempty" jsonschema:"nullable"`

	// Operating system group accounts to be created on the image.
	Groups []GroupAccount `json:"groups,omitempty" jsonschema:"nullable"`
}

type UserAccount struct {
	// Account name. Accepted characters: lowercase letters, digits, underscores, dollars, and hyphens.
	// Name must not start with a hyphen. Maximum length is 256 characters. The validation pattern is
	// a relaxed version of https://github.com/shadow-maint/shadow/blob/master/lib/chkname.c
	Name string `json:"name" jsonschema:"required,maxLength=256,pattern=^[a-zA-Z0-9_.][a-zA-Z0-9_.$-]*$"`

	// A longer description of the account.
	Description string `json:"description,omitempty" jsonschema:"maxLength=4096"`

	// Password either in plain text or encrypted form. If the password is not provided, the account will be
	// locked and the user will not be able to log in with a password. The password can be encrypted using
	// the crypt(3) function. The format of the encrypted password is $id$salt$hashed,
	// where $id is the algorithm used (1, 5, 6, or 2a).
	Password string `json:"password,omitempty"`

	// SSH keys to be added to the account's authorized_keys file.
	SshKeys []string `json:"ssh_keys,omitempty"`

	// The home directory of the user.
	Home string `json:"home,omitempty"`

	// The shell of the user.
	Shell string `json:"shell,omitempty"`

	// The user ID (UID) of the user.
	UID int `json:"uid,omitempty" jsonschema:"minimum=1"`

	// The primary group ID (GID) of the user.
	GID int `json:"gid,omitempty" jsonschema:"minimum=1"`

	// Additional groups that the user should be a member of.
	Groups []string `json:"groups,omitempty"`

	// The expiration date of the account in the format YYYY-MM-DD.
	Expires time.Time `json:"expires,omitempty" jsonschema:"format=date"`
}

type GroupAccount struct {
	// Group name. Accepted characters: lowercase letters, digits, underscores, dollars, and hyphens.
	// Name must not start with a hyphen. Maximum length is 256 characters. The validation pattern is
	// a relaxed version of https://github.com/shadow-maint/shadow/blob/master/lib/chkname.c
	Name string `json:"name" jsonschema:"required,maxLength=256,pattern=^[a-zA-Z0-9_.][a-zA-Z0-9_.$-]*$"`

	// The group ID (GID) of the group.
	GID int `json:"gid,omitempty" jsonschema:"minimum=1"`
}
