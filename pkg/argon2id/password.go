package argon2id

// Cost for argon2id
type Cost struct {
	Time    uint32 `json:"time"`
	Memory  uint32 `json:"memory"`
	Threads uint8  `json:"threads"`
	KeyLen  uint32 `json:"keyLen"`
	SaltLen int32  `json:"saltLen"`
}

// FillEmpty override struct defaults with DefaultCost
func (c *Cost) FillEmpty() {
	if c.Time == 0 {
		c.Time = DefaultCost.Time
	}
	if c.Memory == 0 {
		c.Memory = DefaultCost.Memory
	}
	if c.Threads == 0 {
		c.Threads = DefaultCost.Threads
	}
	if c.KeyLen == 0 {
		c.KeyLen = DefaultCost.KeyLen
	}
	if c.SaltLen <= 0 {
		c.SaltLen = DefaultCost.SaltLen
	}
}

// Password Argon2ID password struct
type Password struct {
	Version int
	Cost    Cost
	Salt    []byte
	Key     []byte
}

// NeedRehash checks hash values to
func (p *Password) NeedRehash(cost Cost) bool {
	if cost.Time > p.Cost.Time || cost.Memory > p.Cost.Memory || cost.Threads > p.Cost.Threads {
		return true
	}
	if len(p.Salt) < int(cost.KeyLen) || len(p.Key) < int(cost.KeyLen) {
		return true
	}
	return false
}
