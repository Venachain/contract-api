package contract_api

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

// ValueToBytes 其他类型转 []byte
func ValueToBytes(source interface{}) ([]byte, error) {
	switch dest := source.(type) {
	case string:
		return StringToBytes(dest), nil
	case int:
		return IntToBytes(dest), nil
	case int8:
		return Int8ToBytes(dest), nil
	case int16:
		return Int16ToBytes(dest), nil
	case int32:
		return Int32ToBytes(dest), nil
	case int64:
		return Int64ToBytes(dest), nil
	case uint:
		return UintToBytes(dest), nil
	case uint8:
		return Uint8ToBytes(dest), nil
	case uint16:
		return Uint16ToBytes(dest), nil
	case uint32:
		return Uint32ToBytes(dest), nil
	case uint64:
		return Uint64ToBytes(dest), nil
	case bool:
		return BoolToBytes(dest), nil
	case *big.Int:
		return BigIntToBytes(dest), nil
	case Address:
		return dest.Bytes(), nil
	case []byte:
		return dest, nil
	}
	return nil, errors.New(fmt.Sprintf("ToBytes function not support %v", source))
}

func StringToBytes(s string) []byte {
	return []byte(s)
}

func BytesToString(b []byte) string {
	return string(b)
}

func IntToBytes(n int) []byte {
	tmp := int(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Int8ToBytes(n int8) []byte {
	tmp := int8(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Int16ToBytes(n int16) []byte {
	tmp := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Int32ToBytes(n int32) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Int64ToBytes(n int64) []byte {
	tmp := int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func UintToBytes(n uint) []byte {
	tmp := uint(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Uint8ToBytes(n uint8) []byte {
	tmp := uint8(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Uint16ToBytes(n uint16) []byte {
	tmp := uint16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Uint32ToBytes(n uint32) []byte {
	tmp := uint32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Uint64ToBytes(n uint64) []byte {
	tmp := uint64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func BigIntToBytes(n *big.Int) []byte {
	if n == nil {
		return nil
	}
	return n.Bytes()
}

func BytesToInt(b []byte) int {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return int(n.Int64())
	}
	if len(b) < 4 {
		b = append(make([]byte, 4-len(b)), b...)
	} else {
		b = b[len(b)-4:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToInt8(b []byte) int8 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return int8(n.Int64())
	}
	if len(b) < 1 {
		b = append(make([]byte, 1-len(b)), b...)
	} else {
		b = b[len(b)-1:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int8
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToInt16(b []byte) int16 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return int16(n.Int64())
	}
	if len(b) < 2 {
		b = append(make([]byte, 2-len(b)), b...)
	} else {
		b = b[len(b)-2:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToInt32(b []byte) int32 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return int32(n.Int64())
	}
	if len(b) < 4 {
		b = append(make([]byte, 4-len(b)), b...)
	} else {
		b = b[len(b)-4:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int32(tmp)
}

func BytesToInt64(b []byte) int64 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return n.Int64()
	}
	if len(b) < 8 {
		b = append(make([]byte, 8-len(b)), b...)
	} else {
		b = b[len(b)-8:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToUint(b []byte) uint {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return uint(n.Uint64())
	}
	if len(b) < 4 {
		b = append(make([]byte, 4-len(b)), b...)
	} else {
		b = b[len(b)-4:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToUint8(b []byte) uint8 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return uint8(n.Uint64())
	}
	if len(b) < 1 {
		b = append(make([]byte, 1-len(b)), b...)
	} else {
		b = b[len(b)-1:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint8
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToUint16(b []byte) uint16 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return uint16(n.Uint64())
	}
	if len(b) < 2 {
		b = append(make([]byte, 2-len(b)), b...)
	} else {
		b = b[len(b)-2:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToUint32(b []byte) uint32 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return uint32(n.Uint64())
	}
	if len(b) < 4 {
		b = append(make([]byte, 4-len(b)), b...)
	} else {
		b = b[len(b)-4:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToUint64(b []byte) uint64 {
	if len(b) >= 32 {
		n := big.NewInt(0)
		n.SetBytes(b)
		return n.Uint64()
	}
	if len(b) < 8 {
		b = append(make([]byte, 8-len(b)), b...)
	} else {
		b = b[len(b)-8:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func BytesToBigInt(b []byte) *big.Int {
	return big.NewInt(0).SetBytes(b)
}

func BoolToBytes(b bool) []byte {
	tmp := bool(b)
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

func BytesToBool(b []byte) bool {
	if len(b) < 1 {
		b = append(make([]byte, 1-len(b)), b...)
	} else {
		b = b[len(b)-1:]
	}
	bytesBuffer := bytes.NewBuffer(b)
	var tmp bool
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

func InterfaceToBytes(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
