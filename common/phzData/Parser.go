package phzData

type PHZParser interface {
	CanTi()
	CanPeng()
	CanChi()
	CanHu()
	XiPai()
}

type ParserCore struct {
}

func (p *ParserCore) CountHandPais(pais []*PHZPoker) []int {
	//todo
	/*counts := make([]int, 20)
	for _, p := range pais {
	}*/
	return nil
}