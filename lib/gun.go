package lib

import (
	"encoding/json"
	"math"
)

// Gun is
type Gun struct {
	Owner     *Object `json:"owner"`
	Bullet    Object  `json:"bullet"`
	Sx        float64 `json:"sx"`
	Sy        float64 `json:"sy"`
	Dir       float64 `json:"dir"`
	Rdir      float64 `json:"rdir"`
	Bound     float64 `json:"bound"`
	Reload    float64 `json:"reload"`    // default delay
	WaitTime  float64 `json:"waittime"`  // first wait time
	DelayTime float64 `json:"delaytime"` // when shot
	ShotTime  float64 `json:"shottime"`  // click time to delay
	Limit     int     `json:"limit"`
	AutoShot  bool    `json:"autoshot"`
}

func (g Gun) Shot() *Object {

	return nil
}

// New Gun1!!!!!!!!!!11!!!
func NewGun(value map[string]interface{}) *Gun {
	m := map[string]interface{}{
		"sx":        0,
		"sy":        1.88,
		"dir":       0,
		"rdir":      math.Pi / 36,
		"bound":     1,
		"reload":    1,
		"waittime":  0,
		"delaytime": 0,
		"shottime":  0,
		"limit":     -1,
		"autoshot":  false,
	}
	objID++

	for key, val := range value {
		m[key] = val
	}

	jsonString, _ := json.Marshal(m)

	s := Gun{}
	json.Unmarshal(jsonString, &s)
	return &s
}
