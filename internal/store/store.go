package store

// Store is an in-memory key-value store.
// Week 1: single-threaded, no locking. Week 2 is where concurrency lands.
type Store struct {
	// TODO: pick a backing map. What key type? What value type — string, []byte, any?
	//       Read week 1's prompts in README.md before choosing.
}

// New returns an initialized Store.
func New() *Store {
	// TODO: a map declared but not constructed is nil — reads are fine but writes panic.
	//       Construct the map here.
	return &Store{}
}

// Get returns the value for key and whether it was present.
func (s *Store) Get(key string) ([]byte, bool) {
	// TODO: use the two-value map lookup idiom: v, ok := m[k]
	return nil, false
}

// Put stores value under key, overwriting any prior value.
func (s *Store) Put(key string, value []byte) {
	// TODO: store the value. Question to sit with: should you copy the slice before storing?
	//       What happens if the caller mutates it afterwards?
}

// Delete removes key from the store. No-op if absent.
func (s *Store) Delete(key string) {
	// TODO: the builtin `delete(m, k)` is a no-op if k isn't there — is that what you want?
}
