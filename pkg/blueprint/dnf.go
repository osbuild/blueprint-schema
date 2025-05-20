package blueprint

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
