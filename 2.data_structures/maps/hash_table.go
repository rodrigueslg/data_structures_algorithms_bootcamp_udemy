package maps

type HashMapItem struct {
	Key    string
	Values []any
}

type HashMap struct {
	Items []HashMapItem
}

func NewMap(size int) *HashMap {
	return &HashMap{
		Items: make([]HashMapItem, size),
	}
}

func (h *HashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		char := rune(key[i])
		hash = (hash + int(char)*i) % len(h.Items)
	}
	return hash
}

func (h *HashMap) Set(key string, value any) {
	address := h.hash(key)
	h.Items[address].Key = key
	h.Items[address].Values = append(h.Items[address].Values, value)
}

func (h *HashMap) Get(key string) []any {
	address := h.hash(key)
	return h.Items[address].Values
}

func (h *HashMap) Keys() []string {
	keys := make([]string, len(h.Items))
	for _, v := range h.Items {
		keys = append(keys, v.Key)
	}
	return keys
}
