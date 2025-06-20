package ubp

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// UnmarshalJSON handles default values
func (node *FSNode) UnmarshalJSON(data []byte) error {
	type tmpType FSNode
	tmp := tmpType{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.Mode == 0 {
		if tmp.Type.IsDir() {
			tmp.Mode = 0755
		} else {
			tmp.Mode = 0644
		}
	}

	*node = FSNode(tmp)
	return nil
}

func (node FSNodeContents) SelectUnion() (FSNodeContentsText, FSNodeContentsBase64, error) {
	var ft FSNodeContentsText
	err := json.Unmarshal(node.union, &ft)
	if err != nil {
		return FSNodeContentsText{}, FSNodeContentsBase64{}, err
	}

	var fb FSNodeContentsBase64
	err = json.Unmarshal(node.union, &fb)
	if err != nil {
		return FSNodeContentsText{}, FSNodeContentsBase64{}, err
	}

	return ft, fb, nil
}

func (n FSNodeState) String() string {
	return string(n)
}

func (n FSNodeState) IsPresent() bool {
	return n == FSStatePresent
}

func (n FSNodeState) IsAbsent() bool {
	return n == FSStateAbsent
}

func (c FSNodeContentsText) String() (string, error) {
	return c.Text, nil
}

func (c FSNodeContentsBase64) String() (string, error) {
	buf, err := base64.StdEncoding.DecodeString(c.Base64)
	return string(buf), err
}

func (node FSNodeContents) String() (string, error) {
	text, base64, err := node.SelectUnion()
	if err != nil {
		return "", err
	}

	if text.Text != "" {
		return text.String()
	} else if base64.Base64 != "" {
		return base64.String()
	}

	return "", nil
}

func FSNodeContentsFromText(node FSNodeContentsText) *FSNodeContents {
	u, _ := json.Marshal(node)
	return &FSNodeContents{union: u}
}

func FSNodeContentsFromBase64(node FSNodeContentsBase64) *FSNodeContents {
	u, _ := json.Marshal(node)
	return &FSNodeContents{union: u}
}

func (nt FSNodeType) String() string {
	return string(nt)
}

func (nt FSNodeType) IsDir() bool {
	return strings.EqualFold(nt.String(), FSNodeDir.String())
}

func (nt FSNodeType) IsFile() bool {
	return strings.EqualFold(nt.String(), FSNodeFile.String())
}
