package main

import (
	"github.com/FancyGo/fc_log"
	"github.com/FancyGo/fc_res"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

type TestDataRes struct {
	a int
	b int
	c int
}
type TestRes struct {
	tid  int
	name string
	data []TestDataRes
}

type TestParser struct {
}

func (t *TestParser) DoParse(data interface{}) (interface{}, error) {
	cells := data.([]*xlsx.Cell)
	testres := &TestRes{
		data: make([]TestDataRes, 0, 0),
	}
	testres.tid, _ = strconv.Atoi(cells[0].Value)
	testres.name = cells[1].Value

	s := strings.Split(cells[2].Value, ";")
	for _, v := range s {
		s1 := strings.Split(v, ",")
		i1, _ := strconv.Atoi(s1[0])
		i2, _ := strconv.Atoi(s1[1])
		i3, _ := strconv.Atoi(s1[2])
		testdata := TestDataRes{
			a: i1,
			b: i2,
			c: i3,
		}
		testres.data = append(testres.data, testdata)
	}
	return testres, nil
}

func (t *TestParser) GenKey(ptt interface{}) int {
	p := ptt.(*TestRes)
	return p.tid
}

func main() {
	resload := fc_res.NewResload()
	testpaser := new(TestParser)
	resload.Register("test.xlsx", testpaser)
	if err := resload.LoadAddRes(); err != nil {
		fc_log.Sys("err = %v\n", err)
	}
	fc_log.Sys("main key = 1, val = %+v\n", resload.GetPttByKey("test.xlsx", 1))
	fc_log.Sys("main key = 2, val = %+v\n", resload.GetPttByKey("test.xlsx", 2))
	fc_log.Sys("main key = 3, val = %+v\n", resload.GetPttByKey("test.xlsx", 3))
	fc_log.Sys("main key = 4, val = %+v\n", resload.GetPttByKey("test.xlsx", 4))
	fc_log.Sys("main key = 5, val = %+v\n", resload.GetPttByKey("test.xlsx", 5))
	fc_log.Sysln("fancygo")
}
