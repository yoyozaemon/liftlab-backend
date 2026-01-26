package models

type Exercise struct {
	Name string `json:"name" bson:"name"`
	Sets int    `json:"sets" bson:"sets"`
	Reps int    `json:"reps" bson:"reps"`
}

type Workout struct {
	BaseModel
	Name      string     `json:"name" bson:"name"`
	Exercises []Exercise `json:"exercises" bson:"exercises"`
}
