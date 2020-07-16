package universe

import (
	"sort"
	"sync"

	"github.com/shadracnicholas/home-automation/service.dmx/domain"
)

// Universe represents a set of fixtures in a 512 channel space
type Universe struct {
	Number uint8

	fixtures map[string]domain.Fixture
	mux      *sync.RWMutex
}

// New returns an initialised universe
func New(number uint8) *Universe {
	return &Universe{
		Number:   number,
		fixtures: make(map[string]domain.Fixture),
		mux:      &sync.RWMutex{},
	}
}

// AddFixture adds the given fixture to
// the universe if it does not already exist
func (u *Universe) AddFixture(f domain.Fixture) {
	u.mux.Lock()
	defer u.mux.Unlock()

	if _, ok := u.fixtures[f.ID()]; ok {
		return
	}

	u.fixtures[f.ID()] = f
}

// Find returns the fixture with the given ID
func (u *Universe) Find(id string) domain.Fixture {
	u.mux.RLock()
	defer u.mux.RUnlock()

	f, ok := u.fixtures[id]
	if !ok {
		return nil
	}

	return f.Copy()
}

// Save adds the given fixture to the universe
// replacing any existing fixture with the same ID
func (u *Universe) Save(f domain.Fixture) {
	u.mux.Lock()
	defer u.mux.Unlock()

	u.fixtures[f.ID()] = f
}

// Valid returns false if any fixtures have overlapping channel ranges
func (u *Universe) Valid() bool {
	u.mux.RLock()
	defer u.mux.RUnlock()

	var f []domain.Fixture
	for _, fixture := range u.fixtures {
		f = append(f, fixture)
	}

	// Sort the fixtures by offset
	sort.Slice(f, func(i, j int) bool {
		return f[i].Offset() < f[j].Offset()
	})

	// Make sure each fixture ends before the next one begins
	for i := 0; i < len(f)-1; i++ {
		if f[i].Offset()+len(f[i].DMXValues()) > f[i+1].Offset() {
			return false
		}
	}

	return true
}

// DMXValues returns the value of all channels in the universe
// The given fixture will override the locally held version.
func (u *Universe) DMXValues(override domain.Fixture) [512]byte {
	u.mux.RLock()
	defer u.mux.RUnlock()

	var v [512]byte
	for id, f := range u.fixtures {
		if override.ID() == id {
			f = override
		}
		copy(v[f.Offset():], f.DMXValues())
	}
	return v
}
