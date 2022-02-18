package bitso

type Details map[string]interface{}

func (d Details) GetTxHash() string {
	if id, ok := d["ripple_transaction_hash"]; ok {
		return id.(string)
	} else {
		return ""
	}
}

func (d Details) Get(key string) interface{} {
	if val, ok := d[key]; ok {
		return val
	} else {
		return nil
	}
}
