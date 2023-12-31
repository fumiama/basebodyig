package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/fumiama/basebodyig"
)

const header = `// Code generated by gen/define. DO NOT EDIT.

package basebodyig

// BodYig基字映射表 9位 映射 BodYig基字超集 到编码
var BodYig基字映射表 = [512][2]uint16{
	`

const 基字逆映射表 = `

// BodYig基字逆映射表 还原基字到编码
var BodYig基字逆映射表 = map[string]uint16{
	`

const 小逆映射表模版 = `

// BodYig%s逆映射表 还原%s到编码
var BodYig%s逆映射表 = map[string]uint8{
	`

const 附标逆映射表 = `

// BodYig附标逆映射表 还原附标到编码
var BodYig附标逆映射表 = map[string]uint8{
	`

func 写逆映射表of(typ string, w io.Writer) {
	fmt.Fprintf(w, 小逆映射表模版, typ, typ, typ)
}

func main() {
	f, err := os.Create("define_gen.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ptr := 0
	cnt := 0
	writebatch := func(charlen int, str string) {
		if cnt >= 512 {
			return
		}
		chls := strconv.Itoa(charlen)
		for i := 0; i < len(str)/charlen; i++ {
			if cnt > 0 {
				if cnt%16 == 0 {
					f.WriteString("},\n\t")
				} else {
					f.WriteString("}, ")
				}
			}
			f.WriteString("{")
			f.WriteString(strconv.Itoa(ptr))
			f.WriteString(", ")
			f.WriteString(chls)
			cnt++
			if cnt >= 512 {
				f.WriteString("},\n}")
				return
			}
			ptr += charlen
		}
	}
	f.WriteString(header)
	writebatch(3, basebodyig.BodYig单符基字)
	writebatch(6, basebodyig.BodYig双符合字)
	writebatch(9, basebodyig.BodYig三符合字)
	writebatch(12, basebodyig.BodYig四符合字)
	cnt = 0
	writebatch = func(charlen int, str string) {
		if cnt >= 512 {
			return
		}
		ptr := 0
		for i := 0; i < len(str)/charlen; i++ {
			if cnt > 0 {
				if cnt%16 == 0 {
					f.WriteString(",\n\t")
				} else {
					f.WriteString(", ")
				}
			}
			f.WriteString(`"`)
			f.WriteString(str[ptr : ptr+charlen])
			f.WriteString(`": `)
			f.WriteString(strconv.Itoa(cnt))
			cnt++
			if cnt >= 512 {
				f.WriteString(",\n}")
				return
			}
			ptr += charlen
		}
	}
	f.WriteString(基字逆映射表)
	writebatch(3, basebodyig.BodYig单符基字)
	writebatch(6, basebodyig.BodYig双符合字)
	writebatch(9, basebodyig.BodYig三符合字)
	writebatch(12, basebodyig.BodYig四符合字)
	cnt = 0
	写逆映射表of("前加字", f)
	writebatch(3, basebodyig.BodYig前加字)
	f.WriteString(",\n}")
	cnt = 0
	写逆映射表of("后加字", f)
	writebatch(3, basebodyig.BodYig后加字)
	f.WriteString(",\n}")
	cnt = 0
	写逆映射表of("再后加字", f)
	writebatch(3, basebodyig.BodYig再后加字)
	f.WriteString(",\n}")
	writebatch = func(charlen int, str string) {
		cnt := 1
		f.WriteString(`"": 0`)
		for i := 1; i < 8; i++ {
			if cnt > 0 {
				if cnt%16 == 0 {
					f.WriteString(",\n\t")
				} else {
					f.WriteString(", ")
				}
			}
			a := 3 * (2*i - 1)
			b := a + 3
			f.WriteString(`"`)
			f.WriteString(str[a:b])
			f.WriteString(`": `)
			f.WriteString(strconv.Itoa(cnt))
			cnt++
			if cnt >= 8 {
				f.WriteString(",\n}")
				return
			}
		}
	}
	f.WriteString(附标逆映射表)
	writebatch(3, basebodyig.BodYig单符附标)
	f.WriteString("\n")
}
