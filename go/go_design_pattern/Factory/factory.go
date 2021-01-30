package factory

// IRuleConfigParser struct
type IRuleConfigParser interface{}
type yamlRuleCinfigParser struct{}
type jsonRuleConfigParser struct{}

// IRuleConfigParserFactory interface
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleCinfigParserFactory factory class
type yamlRuleCinfigParserFactory struct{}

func (y yamlRuleCinfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleCinfigParser{}
}

type jsonRuleConfigParserFactory struct{}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory create an NewIRuleConfigParserFactory factory function
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleCinfigParserFactory{}
	}

	return nil
}
