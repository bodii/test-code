package chain

// SensitiveWordFilter Filter word sensitive
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain Filter word sensitive chain
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// AddFilter add filter
func (s *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	s.filters = append(s.filters, filter)
}

// Filter Filter word sensitive chain function
func (s *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range s.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter struct
type AdSensitiveWordFilter struct{}

// Filter AdSensitiveWordFilter filter
func (a *AdSensitiveWordFilter) Filter(content string) bool {
	return false
}

// PoliticalWordFilter filter word politial
type PoliticalWordFilter struct{}

// Filter PoliticalWordFilter implement
func (p *PoliticalWordFilter) Filter(content string) bool {
	// TODO: implement content
	return true
}
