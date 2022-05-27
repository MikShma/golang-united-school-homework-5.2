package cache

import "time"

type MyStruct struct {
	Value    string
	Deadline time.Time
}

type Cache struct {
	KeyValue map[string]MyStruct
}

func NewCache() Cache {
	return Cache{KeyValue: make(map[string]MyStruct)}
}

func (c Cache) ClearExpired() {
	for i := range c.KeyValue {
		if c.KeyValue[i].Deadline.Before(time.Now()) && !c.KeyValue[i].Deadline.IsZero() {
			delete(c.KeyValue, i)
		}
	}
}

func (c Cache) Get(key string) (string, bool) {
	v, ok := c.KeyValue[key]
	c.ClearExpired()
	if !ok {
		return "", false

	}
	return v.Value, ok

}

func (c Cache) Put(key, value string) {
	c.KeyValue[key] = MyStruct{Value: value}
}

func (c Cache) Keys() (s []string) {
	c.ClearExpired()
	for i := range c.KeyValue {
		s = append(s, i)
	}
	return s
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.KeyValue[key] = MyStruct{Value: value, Deadline: deadline}
}
