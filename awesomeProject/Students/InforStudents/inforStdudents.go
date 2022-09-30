package InforStudents

type Students struct {
	Id        string `json:"id"binding:"required" form:"id" gorm : "primary_key"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Grade     string `json:"grade" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Gmail     string `json:"gmail" binding:"required"`
}
