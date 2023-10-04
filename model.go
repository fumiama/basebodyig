package basebodyig

import (
	"encoding/binary"
	"errors"
	"strings"
)

var (
	ErrInvalidBodYigCharString = errors.New("invalid bodyig char string")
	ErrInvalidBodYig前加字        = errors.New("invalid bodyig 前加字")
	ErrInvalidBodYig基字         = errors.New("invalid bodyig 基字")
	ErrInvalidBodYig附标         = errors.New("invalid bodyig 附标")
	ErrInvalidBodYig后加字        = errors.New("invalid bodyig 后加字")
	ErrInvalidBodYig再后加字       = errors.New("invalid bodyig 再后加字")
)

// BodYigChar 线程不安全
type BodYigChar [3]byte

// NewBodYigChar 从合法 string 生成 BodYigChar
func NewBodYigChar(charstr string) (c BodYigChar, err error) {
	switch len(charstr) {
	case 3: // 只有基字
		err = c.Set基字(charstr)
		return
	case 6: // 单符基字+附标 / 单符基字+后加字 / 单符基字+再后加字 / 双符合字
		err = c.Set基字(charstr)
		if err == nil { // 是 双符合字
			return
		}
		err = c.Set基字(charstr[:3])
		if err != nil {
			return
		}
		err = c.Set附标(charstr[3:])
		if err == nil { // 是 单符基字+附标
			return
		}
		err = c.Set后加字(charstr[3:])
		if err == nil { // 是 单符基字+后加字
			return
		}
		// 是 单符基字+再后加字
		err = c.Set再后加字(charstr[3:])
		return
	case 9: // 前加字+单符基字+区分符/后加字 / 单符基字+附标+后加字 / 单符基字+附标+再后加字 / 单符基字+后加字+再后加字 / 前加字+双符合字 / 双符合字+附标 / 双符合字+后加字 / 双符合字+再后加字 / 三符合字
		err = c.Set基字(charstr)
		if err == nil { // 是 三符合字
			return
		}
		err = c.Set基字(charstr[:6])
		if err == nil { // 是 双符合字+...
			err = c.Set附标(charstr[6:])
			if err == nil { // 是 +附标
				return
			}
			err = c.Set后加字(charstr[6:])
			if err == nil { // 是 +后加字
				return
			}
			// 是 +再后加字
			err = c.Set再后加字(charstr[6:])
			return
		}
		err = c.Set基字(charstr[3:9])
		if err == nil { // 是 前加字+双符合字
			err = c.Set前加字(charstr[:3])
			return
		}
		err = c.Set附标(charstr[3:6])
		if err == nil { // 是 单符基字+附标+后加字 / 单符基字+附标+再后加字
			err = c.Set基字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[6:9])
			if err == nil { // 单符基字+附标+再后加字
				return
			}
			// 单符基字+附标+后加字
			err = c.Set后加字(charstr[6:9])
			return
		}
		err = c.Set再后加字(charstr[6:9])
		if err == nil { // 是 单符基字+后加字+再后加字
			err = c.Set基字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set后加字(charstr[3:6])
			return
		}
		// 余下一种可能 前加字+单符基字+区分符/后加字
		err = c.Set前加字(charstr[:3])
		if err != nil {
			return
		}
		err = c.Set基字(charstr[3:6])
		if err != nil {
			return
		}
		err = c.Set后加字(charstr[6:9])
		return
	case 12: // 前加字+单符基字+后加字/区分符+再后加字 / 前加字+单符基字+附标+后加字 / 前加字+单符基字+附标+再后加字 / 单符基字+附标+后加字+再后加字 / 前加字+双符合字+后加字 / 前加字+双符合字+再后加字 / 双符合字+附标+后加字 / 双符合字+附标+再后加字 / 双符合字+后加字+再后加字 / 前加字+三符合字 / 三符合字+附标 / 三符合字+后加字 / 三符合字+再后加字
		err = c.Set基字(charstr[:9])
		if err == nil { // 是 三符合字+...
			err = c.Set附标(charstr[9:12])
			if err == nil { // 是 +附标
				return
			}
			err = c.Set再后加字(charstr[9:12])
			if err == nil { // 是 +再后加字
				return
			}
			// 是 +后加字
			err = c.Set后加字(charstr[9:12])
			return
		}
		err = c.Set基字(charstr[3:12])
		if err == nil { // 是 前加字+三符合字
			err = c.Set前加字(charstr[:3])
			return
		}
		err = c.Set基字(charstr[:6])
		if err == nil { // 是 双符合字+附标+后加字 / 双符合字+附标+再后加字 / 双符合字+后加字+再后加字
			err = c.Set附标(charstr[6:9])
			if err == nil { // 是 +附标+...
				err = c.Set再后加字(charstr[9:12])
				if err == nil { // ...+再后加字
					return
				}
				// ...+后加字
				err = c.Set后加字(charstr[9:12])
				return
			}
			// 是 双符合字+后加字+再后加字
			err = c.Set后加字(charstr[6:9])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[9:12])
			return
		}
		err = c.Set基字(charstr[3:9])
		if err == nil { // 是 前加字+双符合字+后加字 / 前加字+双符合字+再后加字
			err = c.Set前加字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[9:12])
			if err == nil {
				return
			}
			err = c.Set后加字(charstr[9:12])
			return
		}
		err = c.Set附标(charstr[3:6])
		if err == nil { // 是 单符基字+附标+后加字+再后加字
			err = c.Set基字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set后加字(charstr[6:9])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[9:12])
			return
		}
		err = c.Set附标(charstr[6:9])
		if err == nil { // 是 前加字+单符基字+附标+后加字 / 前加字+单符基字+附标+再后加字
			err = c.Set前加字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set基字(charstr[3:6])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[9:12])
			if err == nil {
				return
			}
			err = c.Set后加字(charstr[9:12])
			return
		}
		// 余下一种可能 前加字+单符基字+后加字/区分符+再后加字
		err = c.Set前加字(charstr[:3])
		if err != nil {
			return
		}
		err = c.Set基字(charstr[3:6])
		if err != nil {
			return
		}
		err = c.Set后加字(charstr[6:9])
		if err != nil {
			return
		}
		err = c.Set再后加字(charstr[9:12])
		return
	case 15: // 前加字+单符基字+附标+后加字+再后加字 / 前加字+双符合字+后加字+再后加字 / 前加字+双符合字+附标+后加字 / 双符合字+附标+后加字+再后加字 / 前加字+三符合字+后加字 / 三符合字+附标+后加字 / 三符合字+附标+再后加字 / 三符合字+后加字+再后加字
		err = c.Set基字(charstr[:9])
		if err == nil { // 是 三符合字+...
			err = c.Set附标(charstr[9:12])
			if err == nil { // ...+附标+后加字
				err = c.Set后加字(charstr[12:15])
				return
			}
			// ...+后加字+再后加字
			err = c.Set后加字(charstr[9:12])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[12:15])
			return
		}
		err = c.Set基字(charstr[3:12])
		if err == nil { // 是 前加字+三符合字+后加字
			err = c.Set前加字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set后加字(charstr[12:15])
			return
		}
		err = c.Set基字(charstr[:6])
		if err == nil { // 是 双符合字+附标+后加字+再后加字
			err = c.Set附标(charstr[6:9])
			if err != nil {
				return
			}
			err = c.Set后加字(charstr[9:12])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[12:15])
			return
		}
		err = c.Set基字(charstr[3:9])
		if err == nil { // 是 前加字+双符合字+...
			err = c.Set前加字(charstr[:3])
			if err != nil {
				return
			}
			err = c.Set附标(charstr[9:12])
			if err == nil { // ...+附标+后加字
				err = c.Set后加字(charstr[12:15])
				return
			}
			// ...+后加字+再后加字
			err = c.Set后加字(charstr[9:12])
			if err != nil {
				return
			}
			err = c.Set再后加字(charstr[12:15])
			return
		}
		// 余下一种可能 前加字+单符基字+附标+后加字+再后加字
		err = c.Set前加字(charstr[:3])
		if err != nil {
			return
		}
		err = c.Set基字(charstr[3:6])
		if err != nil {
			return
		}
		err = c.Set附标(charstr[6:9])
		if err != nil {
			return
		}
		err = c.Set后加字(charstr[9:12])
		if err != nil {
			return
		}
		err = c.Set再后加字(charstr[12:15])
		return
	case 18: // 前加字+双符合字+附标+后加字+再后加字 / 前加字+三符合字+后加字+再后加字 / 前加字+三符合字+附标+后加字 / 三符合字+附标+后加字+再后加字
		return
	default:
		err = ErrInvalidBodYigCharString
		return
	}
}

func (c BodYigChar) String() string {
	sb := strings.Builder{}
	前加字 := c.Get前加字()
	基字 := c.Get基字()
	附标 := c.Get附标()
	后加字 := c.Get后加字()
	再后加字 := c.Get再后加字()
	if 前加字 != "" && len(基字) == 3 && 附标 == "" && 后加字 == "" {
		后加字 = BodYig区分符
	}
	sb.WriteString(前加字)
	sb.WriteString(基字)
	sb.WriteString(附标)
	sb.WriteString(后加字)
	sb.WriteString(再后加字)
	return sb.String()
}

// Get前加字 11110000 00000000 00000000
func (c BodYigChar) Get前加字() string {
	p := ((([3]byte)(c))[0] >> 4) & 0x0f
	前加字 := BodYig前加字[p*3 : (p+1)*3]
	if 前加字 == BodYig空字符 {
		return ""
	}
	return 前加字
}

// Set前加字 11110000 00000000 00000000
func (c BodYigChar) Set前加字(ch string) error {
	if ch == "" {
		ch = BodYig空字符
	}
	n, ok := BodYig前加字逆映射表[ch]
	if !ok {
		return ErrInvalidBodYig前加字
	}
	x := ([3]byte)(c)
	x[0] |= byte((n << 4) & 0xf0)
	return nil
}

// Get基字 00001111 11111000 00000000
func (c BodYigChar) Get基字() string {
	p := (binary.BigEndian.Uint16(c[0:2]) >> 3) & 0x01ff
	基字址 := BodYig基字映射表[p]
	return BodYig基字超集[基字址[0] : 基字址[0]+基字址[1]]
}

// Set基字 00001111 11111000 00000000
func (c BodYigChar) Set基字(ch string) error {
	n, ok := BodYig基字逆映射表[ch]
	if !ok {
		return ErrInvalidBodYig基字
	}
	x := binary.BigEndian.Uint16(c[0:2])
	x |= (n << 3) & 0x0ff8
	binary.BigEndian.PutUint16(c[0:2], x)
	return nil
}

// Get附标 00000000 00000111 00000000
func (c BodYigChar) Get附标() string {
	p := (([3]byte)(c))[1] & 0x07
	if p == 0 {
		return ""
	}
	a := 3 * (2*p - 1)
	b := a + 3
	return BodYig单符附标[a:b]
}

// Set附标 00000000 00000111 00000000
func (c BodYigChar) Set附标(ch string) error {
	n, ok := BodYig附标逆映射表[ch]
	if !ok {
		return ErrInvalidBodYig附标
	}
	x := ([3]byte)(c)
	x[1] |= n & 0x07
	return nil
}

// Get后加字 00000000 00000000 11110000
func (c BodYigChar) Get后加字() string {
	p := ((([3]byte)(c))[2] >> 4) & 0x0f
	后加字 := BodYig后加字[p*3 : (p+1)*3]
	if 后加字 == BodYig空字符 {
		return ""
	}
	return 后加字
}

// Set后加字 00000000 00000000 11110000
func (c BodYigChar) Set后加字(ch string) error {
	if ch == "" || ch == BodYig区分符 {
		ch = BodYig空字符
	}
	n, ok := BodYig后加字逆映射表[ch]
	if !ok {
		return ErrInvalidBodYig后加字
	}
	x := ([3]byte)(c)
	x[2] |= (n << 4) & 0xf0
	return nil
}

// Get再后加字 00000000 00000000 00001111
func (c BodYigChar) Get再后加字() string {
	p := (([3]byte)(c))[2] & 0x0f
	再后加字 := BodYig再后加字[p*3 : (p+1)*3]
	if 再后加字 == BodYig空字符 {
		return ""
	}
	return 再后加字
}

// Set再后加字 00000000 00000000 00001111
func (c BodYigChar) Set再后加字(ch string) error {
	if ch == "" {
		ch = BodYig空字符
	}
	n, ok := BodYig再后加字逆映射表[ch]
	if !ok {
		return ErrInvalidBodYig再后加字
	}
	x := ([3]byte)(c)
	x[2] |= n & 0x0f
	return nil
}
