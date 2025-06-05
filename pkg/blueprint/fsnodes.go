package blueprint

import (
	"encoding/base64"
	"encoding/json"
)

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
