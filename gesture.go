package mousegesture

import orderedmap "github.com/wk8/go-ordered-map/v2"

type Gesture []Direction

var AllGestures *orderedmap.OrderedMap[string, Gesture]

func init() {
	AllGestures = orderedmap.New[string, Gesture](orderedmap.WithInitialData(
		orderedmap.Pair[string, Gesture]{
			Key:   "Up",
			Value: []Direction{Up},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "Down",
			Value: []Direction{Down},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "Left",
			Value: []Direction{Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "Right",
			Value: []Direction{Right},
		},
	))
}
