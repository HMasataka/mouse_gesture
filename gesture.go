package mousegesture

import orderedmap "github.com/wk8/go-ordered-map/v2"

type Gesture []Direction

var AllGestures *orderedmap.OrderedMap[string, Gesture]

func init() {
	AllGestures = orderedmap.New[string, Gesture](orderedmap.WithInitialData(
		// 四方向
		orderedmap.Pair[string, Gesture]{
			Key:   "上左下右",
			Value: []Direction{Up, Left, Down, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "上右下左",
			Value: []Direction{Up, Right, Down, Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "右上左下",
			Value: []Direction{Right, Up, Left, Down},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "右下左上",
			Value: []Direction{Right, Down, Left, Up},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下左上右",
			Value: []Direction{Down, Left, Up, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下右上左",
			Value: []Direction{Down, Right, Up, Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "左下上右",
			Value: []Direction{Left, Down, Up, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "左上右下",
			Value: []Direction{Left, Up, Right, Down},
		},
		// 三方向
		orderedmap.Pair[string, Gesture]{
			Key:   "上左下",
			Value: []Direction{Up, Left, Down},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "上右下",
			Value: []Direction{Up, Right, Down},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下右上",
			Value: []Direction{Down, Right, Up},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下左上",
			Value: []Direction{Down, Left, Up},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "左上右",
			Value: []Direction{Left, Up, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "左下右",
			Value: []Direction{Left, Down, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "右上左",
			Value: []Direction{Right, Up, Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "右下左",
			Value: []Direction{Right, Down, Left},
		},
		// 双方向
		orderedmap.Pair[string, Gesture]{
			Key:   "上下",
			Value: []Direction{Up, Down},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "左右",
			Value: []Direction{Left, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "上左",
			Value: []Direction{Up, Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "上右",
			Value: []Direction{Up, Right},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下左",
			Value: []Direction{Down, Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下右",
			Value: []Direction{Down, Right},
		},
		// 单方向
		orderedmap.Pair[string, Gesture]{
			Key:   "上",
			Value: []Direction{Up},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "下",
			Value: []Direction{Down},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "左",
			Value: []Direction{Left},
		},
		orderedmap.Pair[string, Gesture]{
			Key:   "右",
			Value: []Direction{Right},
		},
	))
}
