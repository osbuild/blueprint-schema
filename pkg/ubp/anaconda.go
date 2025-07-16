package ubp

func ParseAnacondaModule(m string) AnacondaModules {
	switch m {
	case "org.fedoraproject.Anaconda.Modules.Localization":
		return AnacondaLocalization
	case "org.fedoraproject.Anaconda.Modules.Network":
		return AnacondaNetwork
	case "org.fedoraproject.Anaconda.Modules.Payloads":
		return AnacondaPayloads
	case "org.fedoraproject.Anaconda.Modules.Runtime":
		return AnacondaRuntime
	case "org.fedoraproject.Anaconda.Modules.Security":
		return AnacondaSecurity
	case "org.fedoraproject.Anaconda.Modules.Services":
		return AnacondaServices
	case "org.fedoraproject.Anaconda.Modules.Storage":
		return AnacondaStorage
	case "org.fedoraproject.Anaconda.Modules.Subscription":
		return AnacondaSubscription
	case "org.fedoraproject.Anaconda.Modules.Timezone":
		return AnacondaTimezone
	case "org.fedoraproject.Anaconda.Modules.Users":
		return AnacondaUsers
	default:
		return ""
	}
}
