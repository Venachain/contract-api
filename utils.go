package contract_api

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Venachain/contract-api/rlp"
	"math/big"
)

// tx-type
const (
	CallContractFlag                 = 9
	TxTypeCallSollGovmCompatibleWasm = 14
	TxTypeCallSollWasmCompatibleGovm = 15
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
	return Int64ToBytes(int64(n))
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
	return Uint64ToBytes(uint64(n))
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
	n := len(b)
	switch n {
	case 0:
		return 0
	case 1:
		return int(BytesToInt8(b))
	case 2:
		return int(BytesToInt16(b))
	case 4:
		return int(BytesToInt32(b))
	default:
		return int(BytesToInt64(b))
	}
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
	n := len(b)
	switch n {
	case 0:
		return 0
	case 1:
		return uint(BytesToUint8(b))
	case 2:
		return uint(BytesToUint16(b))
	case 4:
		return uint(BytesToUint32(b))
	default:
		return uint(BytesToUint64(b))
	}
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

type SolInput struct {
	FuncName   string   `json:"func_name"`
	FuncParams []string `json:"func_params"`
}

type WasmInput struct {
	TxType     int      `json:"-"`
	FuncName   string   `json:"func_name"`
	FuncParams []string `json:"func_params"`
}

type GovmInput struct {
	TxType     int      `json:"-"`
	FuncName   string   `json:"func_name"`
	FuncParams []string `json:"func_params"`
}

// BuildGovmCallData 构建 govm 调用 govm 合约的 data/input
// 调用例子：BuildWasmCallData(2, "test", IntToBytes(1), []byte("test1")])
// 参数说明：
//		2			: 交易类型
//		"test"		: 被调用合约方法名
//	IntToBytes(1)	: 被调用合约方法的第一个参数，参数类型为 int，实参的值为 1，传入时需要先转为 []byte
//	[]byte("test1")	: 被调用合约方法的第二个参数，参数类型为 string，实参的值为 test1，传入时需要先转为 []byte
func BuildGovmCallData(txType int, methodName string, params ...[]byte) ([]byte, error) {
	paramArr := [][]byte{
		IntToBytes(txType),
		[]byte(methodName),
	}
	for _, v := range params {
		paramArr = append(paramArr, v)
	}
	return rlp.EncodeToBytes(paramArr)
}

// BuildWasmCallData 构建 govm 调用 wasm 合约的 data/input
// 调用例子：BuildWasmCallData(2, "test", "int(1)", "string(test1)")
// 参数说明：
//		2			: 交易类型
//		"test"		: 被调用合约方法名
//		"int(1)"	: 被调用合约方法的第一个参数，参数类型为 int，实参的值为 1
//	"string(test1)"	: 被调用合约方法的第二个参数，参数类型为 string，实参的值为 test1
func BuildWasmCallData(txType int, methodName string, params ...string) ([]byte, error) {
	input := WasmInput{
		TxType:     txType,
		FuncName:   methodName,
		FuncParams: nil,
	}
	var fnParams []string
	for _, v := range params {
		fnParams = append(fnParams, v)
	}
	input.FuncParams = fnParams
	return json.Marshal(input)
}

// BuildEvmCallData 构建 govm 调用 solidity 合约的 data/input
// 调用例子：BuildEvmCallData("test", "int(1)", "string(test1)")
// 参数说明：
//		"test"		: 被调用合约方法名
//		"int(1)"	: 被调用合约方法的第一个参数，参数类型为 int，实参的值为 1
//	"string(test1)"	: 被调用合约方法的第二个参数，参数类型为 string，实参的值为 test1
func BuildEvmCallData(methodName string, params ...string) ([]byte, error) {
	paramArr := [][]byte{
		IntToBytes(TxTypeCallSollWasmCompatibleGovm),
		[]byte(methodName),
	}
	for _, v := range params {
		paramArr = append(paramArr, []byte(v))
	}
	return rlp.EncodeToBytes(paramArr)
}
