package config

func TargetsSpecificContainer() bool {
	return TargettedContainer() != ""
}

func TargettedContainer() string {
	return ""
}

func PodName() string {
	return ""
}
