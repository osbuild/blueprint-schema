package ubp

import (
	"encoding/json"
)

func (s *DNFSource) SelectUnion() (DNFSourceBaseURLs, DNFSourceMetalink, DNFSourceMirrorlist, error) {
	if s == nil {
		return DNFSourceBaseURLs{}, DNFSourceMetalink{}, DNFSourceMirrorlist{}, nil
	}

	var burl DNFSourceBaseURLs
	err := json.Unmarshal(s.union, &burl)
	if err != nil {
		return DNFSourceBaseURLs{}, DNFSourceMetalink{}, DNFSourceMirrorlist{}, err
	}

	var bmeta DNFSourceMetalink
	err = json.Unmarshal(s.union, &bmeta)
	if err != nil {
		return DNFSourceBaseURLs{}, DNFSourceMetalink{}, DNFSourceMirrorlist{}, err
	}

	var bmirror DNFSourceMirrorlist
	err = json.Unmarshal(s.union, &bmirror)
	if err != nil {
		return DNFSourceBaseURLs{}, DNFSourceMetalink{}, DNFSourceMirrorlist{}, err
	}

	return burl, bmeta, bmirror, nil
}

func DNFSourceFromBaseURLs(node DNFSourceBaseURLs) *DNFSource {
	u, _ := json.Marshal(node)
	return &DNFSource{union: u}
}

func DNFSourceFromMetalink(node DNFSourceMetalink) *DNFSource {
	u, _ := json.Marshal(node)
	return &DNFSource{union: u}
}

func DNFSourceFromMirrorlist(node DNFSourceMirrorlist) *DNFSource {
	u, _ := json.Marshal(node)
	return &DNFSource{union: u}
}
