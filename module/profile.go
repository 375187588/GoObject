package module

import "encoding/json"

//Profile struct
type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int    //身高
	Weight     int    //体重
	Income     string //收入
	Marriage   string
	Education  string
	Occupation string
	Hokou      string //户口
	Xinzuo     string
	House      string
	Car        string
	WorkAddr   string
	WorkCity   string
}

// FromJsonObject ....
func FromJsonObject(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
