package sdkuci

type BrVlanPort struct {
	Device   string
	Tagged   bool
	Untagged bool
	Primary  bool
}

func (p *BrVlanPort) String() string {
	str := p.Device + ":"
	if p.Untagged {
		str += "u"
		if p.Primary {
			str += "*"
		}
	} else if p.Tagged {
		str += "t"
	} else {
		str += "t"
	}

	return str
}
