package blueprint

import "encoding/json"

func (np NetworkProtocol) String() string {
	return string(np)
}

func (t NetworkService) SelectUnion() (FirewallService, FirewallPort, FirewallFromTo, error) {
	var fs FirewallService
	err := json.Unmarshal(t.union, &fs)
	if err != nil {
		return FirewallService{}, FirewallPort{}, FirewallFromTo{}, err
	}

	var fp FirewallPort
	err = json.Unmarshal(t.union, &fp)
	if err != nil {
		return FirewallService{}, FirewallPort{}, FirewallFromTo{}, err
	}

	var fft FirewallFromTo
	err = json.Unmarshal(t.union, &fft)
	if err != nil {
		return FirewallService{}, FirewallPort{}, FirewallFromTo{}, err
	}

	return fs, fp, fft, nil
}
