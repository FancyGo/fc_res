package fc_res

import (
	"errors"
	"github.com/fancygo/fc_log"
	"github.com/fancygo/fc_sys"
	"github.com/tealeg/xlsx"
	"path"
)

type Ptt interface{}

type ResMgr struct {
	resname string
	pttdata []Ptt
	idxmap  map[int]int
	parser  Parser
}

type Resload struct {
	Mgrs map[string]*ResMgr
}

func NewResload() *Resload {
	return &Resload{
		Mgrs: make(map[string]*ResMgr),
	}
}

func (r *Resload) Register(resname string, parser Parser) error {
	mgr := &ResMgr{
		resname: resname,
		pttdata: make([]Ptt, 0),
		idxmap:  make(map[int]int),
		parser:  parser,
	}
	r.Mgrs[resname] = mgr
	fc_log.Sys("Resload Register resname = %s\n", resname)

	return nil
}

func (r *Resload) LoadAddRes() error {
	for resname, _ := range r.Mgrs {
		err := r.parseOne(resname)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Resload) GetPttByIdx(resname string, idx int) interface{} {
	mgr, ok := r.Mgrs[resname]
	if !ok {
		return nil
	}
	if idx >= len(mgr.pttdata) {
		return nil
	}
	return mgr.pttdata[idx]
}

func (r *Resload) GetPttByKey(resname string, key int) interface{} {
	mgr, ok := r.Mgrs[resname]
	if !ok {
		return nil
	}
	idx, ok := mgr.idxmap[key]
	if !ok {
		return nil
	}
	return r.GetPttByIdx(resname, idx)
}

func (r *Resload) parseOne(resname string) error {
	mgr := r.Mgrs[resname]
	respath := fc_sys.GetResDir()
	path := path.Join(respath, resname)
	xlsfile, err := xlsx.OpenFile(path)
	if err != nil {
		return errors.New("Resload parseOne res data empty")
	}
	if len(xlsfile.Sheets[0].Rows) <= 1 {
		return err
	}
	for i := 1; i < len(xlsfile.Sheets[0].Rows); i++ {
		cells := xlsfile.Sheets[0].Rows[i].Cells
		ptt, err := mgr.parser.DoParse(cells)
		if err != nil {
			fc_log.Sys("Resload parseOne err")
			return err
		}
		mgr.pttdata = append(mgr.pttdata, ptt)
		key := mgr.parser.GenKey(ptt)
		mgr.idxmap[key] = i - 1
	}
	return nil
}
