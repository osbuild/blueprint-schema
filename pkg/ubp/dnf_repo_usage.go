package ubp

// IsZero checks if the DNFRepoUsage is empty.
func (dr DNFRepoUsage) IsZero() bool {
	return (dr.Configure == nil || *dr.Configure) && (dr.Install == nil || *dr.Install)
}
