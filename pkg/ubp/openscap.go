package ubp

import "encoding/json"

func (ost OpenSCAPTailoring) SelectUnion() (TailoringProfiles, TailoringJSON, error) {
	var tp TailoringProfiles
	err := json.Unmarshal(ost.union, &tp)
	if err != nil {
		return TailoringProfiles{}, TailoringJSON{}, err
	}

	var tj TailoringJSON
	err = json.Unmarshal(ost.union, &tj)
	if err != nil {
		return TailoringProfiles{}, TailoringJSON{}, err
	}

	return tp, tj, nil
}

func OpenSCAPTailoringFromProfiles(node TailoringProfiles) *OpenSCAPTailoring {
	u, _ := json.Marshal(node)
	return &OpenSCAPTailoring{union: u}
}

func OpenSCAPTailoringFromJSON(node TailoringJSON) *OpenSCAPTailoring {
	u, _ := json.Marshal(node)
	return &OpenSCAPTailoring{union: u}
}
