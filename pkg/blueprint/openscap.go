package blueprint

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
