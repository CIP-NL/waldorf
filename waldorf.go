package waldorf

import (
	"fmt"
	"sync"
)

// Complaints is an alias for a []string to allow for custom formatting. It contains
// the resultset of an observation by Waldorf (which of course contains only everything
// wrong and bad.
type Complaints struct {
	comps      []string
	formatting func(Complaints) string
}

func (c Complaints) String() string {
	return c.formatting(c)
}

// Observation is the single instance of Waldorf. For each series of checks, create a new observation, which ends with
// a Ridicule
type Observation struct {
	complaintsMu *sync.RWMutex
	complaints   *Complaints

	RidiculedTwiceFunc RidiculedTwiceFunc

	ridiMu    *sync.Mutex
	ridiculed bool
}

// Observe is the constructor for an Observation (Equivalent of New, but more in tune with the library)
func Observe() *Observation {
	return &Observation{
		complaintsMu: &sync.RWMutex{},
		complaints: &Complaints{
			comps:      []string{},
			formatting: defaultFormatter,
		},
		RidiculedTwiceFunc: func(observation *Observation, ridiculedAlready bool) {
			if ridiculedAlready {
				panic("ridiculed already!")
			}
		},
		ridiMu:    &sync.Mutex{},
		ridiculed: false,
	}
}

// RidiculedTwiceFunc is called when Waldorf has already Ridiculed an Observation. The default function it panics (since reusing
// the same observation leads to a possible memory leak. A good alternative would be to reset the Complaints.
// The paramaters are the Observation it is called in, and if it has been Ridiculed already.
type RidiculedTwiceFunc func(observation *Observation, ridiculedAlready bool)

// Ridicule returns the resultset of an observation (the observation is actually validated
// during each call (an alternative would be to alter Waldorf to accept a func(arguments) bool as each argument in
// assertions.go to allow for evaluating the complaints during a Ridicule walk.
//
// TODO when we need to implement context.Context in Walter, check and update this strategy
// TODO a possible strategy would be to not implement ctx in each argument in the functions in assertion.go,
// TODO but to have Ridicule accept a context.Context object.
func (r *Observation) Ridicule() *Complaints {
	r.ridiMu.Lock()
	r.RidiculedTwiceFunc(r, r.ridiculed)
	r.ridiMu.Unlock()

	r.complaintsMu.RLock()
	defer r.complaintsMu.RUnlock()
	if len(r.complaints.comps) == 0 {
		return nil
	}
	return r.complaints
}

func generateMSG(msg string, formatting ...interface{}) string {
	if len(formatting) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, formatting)
}

func defaultFormatter(complaints Complaints) string {
	var res string
	for _, comp := range complaints.comps {
		res = res + comp + "\n"
	}
	return res
}
