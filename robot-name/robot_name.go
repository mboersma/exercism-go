// Package robotname manages random and unique robot names.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var used = map[string]bool{}

const maxRobotNames = 26 * 26 * 10 * 10 * 10

// Robot has factory settings: specifically a unique, random name.
type Robot struct {
	name string
}

// Name returns the unique, random name of this robot, such as  "RX837" or "BC811".
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(used) >= maxRobotNames {
		return "", errors.New("No names available")
	}
	for {
		name := fmt.Sprintf("%c%c%03d",
			rnd.Intn(26)+'A', rnd.Intn(26)+'A', rnd.Intn(1000))
		if _, ok := used[name]; !ok {
			used[name] = true
			r.name = name
			break
		}
	}
	return r.name, nil
}

// Reset returns a Robot to its factory settings, giving it a new name.
func (r *Robot) Reset() {
	r.name = ""
}
