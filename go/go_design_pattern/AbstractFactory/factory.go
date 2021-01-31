package factory

// IRuleConfigParser interface
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser struct
type jsonRuleConfigParser struct{}

// Parse func
func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// ISystemConfigParser interface
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

// jsonSystemConfigParser struct
type jsonSystemConfigParser struct{}

// ParseSystem func
func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory factory functions
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
