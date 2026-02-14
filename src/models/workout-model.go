package models

type Weight_Exercise struct {
	Name   string `json:"name" bson:"name"`
	Sets   int    `json:"sets" bson:"sets"`
	Reps   int    `json:"reps" bson:"reps"`
	Weight int    `json:"weight" bson:"weight"`
}

type Cardio_Exercise struct {
	Name     string `json:"name" bson:"name"`
	Time     int    `json:"time" bson:"time"`
	Sets     int    `json:"sets" bson:"sets"`
	Reps     int    `json:"reps" bson:"reps"`
	Distance int    `json:"distance" bson:"distance"`
}

type Workout struct {
	BaseModel
	UID       string            `json:"uid" bson:"uid"`
	Name      string            `json:"name" bson:"name"`
	Exercises []Weight_Exercise `json:"exercises" bson:"exercises"`
	Cardio    []Cardio_Exercise `json:"cardio" bson:"cardio"`
}
