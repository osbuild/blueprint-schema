package ubp

import "encoding/json"

func (ost Ignition) SelectUnion() (IgnitionURL, IgnitionText, error) {
	var iu IgnitionURL

	if len(ost.union) == 0 {
		return IgnitionURL{}, IgnitionText{}, nil
	}

	err := json.Unmarshal(ost.union, &iu)
	if err != nil {
		return IgnitionURL{}, IgnitionText{}, err
	}

	var it IgnitionText
	err = json.Unmarshal(ost.union, &it)
	if err != nil {
		return IgnitionURL{}, IgnitionText{}, err
	}

	return iu, it, nil
}

func IgnitionFromURL(url IgnitionURL) Ignition {
	u, _ := json.Marshal(url)

	return Ignition{
		union: u,
	}
}

func IgnitionFromText(text IgnitionText) Ignition {
	t, _ := json.Marshal(text)

	return Ignition{
		union: t,
	}
}
