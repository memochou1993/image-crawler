package formatter

// Payload struct
type Payload struct {
	Data interface{} `json:"data"`
}

// Set sets the formatter.
func (p *Payload) Set(data interface{}) {
	p.Data = data
}
