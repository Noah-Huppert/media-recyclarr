package main

// ByUserEmbyIDMap is a map where keys are user Emby IDs
type ByUserEmbyIDMap[Item interface{}] struct {
	data map[string]Item
}

// SetByUserEmbyID stores an item based on its Emby user ID
func (m *ByUserEmbyIDMap[Item]) SetByUserEmbyID(userEmbyID string, item Item) {
	m.data[userEmbyID] = item
}

// GetByUserEmbyID retrieves an item based on its Emby user ID
func (m *ByUserEmbyIDMap[Item]) GetByUserEmbyID(userEmbyID string) (Item, bool) {
	item, ok := m.data[userEmbyID]
	return item, ok
}

// Iterate returns a map which can be used to iterate over the keys and values, do not modify this
func (m *ByUserEmbyIDMap[Item]) Iterate() map[string]Item {
	return m.data
}

// ByMediaEmbyIDMap is a map where keys are media Emby IDs
