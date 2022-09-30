package InforJwtB3

type Sinhviens struct {
	Masv string `json:"masv" binding:"required"`
	Hotensv string `json:"hotensv" binding:"required"`
	Passwordsv string `json:"passwordsv" binding:"required"`
	Gmail string `json:"gmail" binding:"required"`
	Sodtsv string `json:"sodtsv" binding:"required"`
	Diachi string `json:"dia_chi" binding:"required"`
}




