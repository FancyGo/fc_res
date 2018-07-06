package fc_res

type Parser interface {
	DoParse(res interface{}) (interface{}, error)
	GenKey(ptt interface{}) int
	//DoCheck(ptt interface{}) error
	//DoAssemble(ptt interface{}) error
}
