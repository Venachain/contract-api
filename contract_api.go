package contract_api

import "math/big"

// Context 合约运行时的上下文
type Context interface {
	// GetState 通过 key 从 StateDB 中获取对应到 Value
	GetState(key []byte) []byte
	// SetState 设置 k-v 到 StateDB 中
	SetState(key []byte, value []byte)
	// GetBalance 获取给定地址的余额
	GetBalance(Address) *big.Int
	// Transfer 从 from 向 to 发送数量为 amount 的金额，发生错误返回 error
	// 本方法只会进行 from 余额是否足够，to 的金额是否会越界的校验
	Transfer(from, to Address, amount *big.Int) error
	// This 返回当前合约的地址
	This() Address
	// EmitEvent 触发自定义事件
	// topic 是必须的，是自定义事件的方法名，params 是它的参数
	EmitEvent(topic string, params ...string)
	// Log 记录日志
	Log(msg string)

	// BlockHash 返回指定区块的哈希
	BlockHash(blockNumber uint64) string
	// GasLimit 返回当前区块 gas 限额
	GasLimit() uint64
	// Number 返回当前区块号
	Number() *big.Int
	// Timestamp 返回当前区块以秒计的时间戳
	Timestamp() uint64
	// GasLeft 返回剩余 gas
	GasLeft() uint64

	// Data 返回当前调用完整的 calldata
	Data() []byte
	// Sender 返回当前调用的消息发送者
	Sender() Address
	// Sig 返回当前调用的函数名
	Sig() []byte
	// Value 返回随消息发送的金额
	Value() *big.Int

	// Origin 返回交易的发起者（完全的调用链）
	Origin() Address

	// Call 直接调用
	// 调用后 Context.Data、Context.Sender、Context.Sig、Context.Value 的值会修改为调用者，执行环境为被调用者的运行环境(合约的 storage)
	// addr 目标合约地址
	// param 调用目标合约的参数
	// 调用成功返回被调用合约方法的返回值，失败返回 error
	Call(addr string, param []byte) ([]byte, error)

	// DelegateCall 代理调用
	// 调用后 Context.Data、Context.Sender、Context.Sig、Context.Value 的值不会修改为调用者（注：会改为以太坊账户的地址）， 但执行环境为调用者的运行环境。
	// addr 目标合约地址
	// param 调用目标合约的参数
	// 调用成功返回被调用合约方法的返回值，失败返回 error
	DelegateCall(addr string, param []byte) ([]byte, error)

	// HexToAddress 将字符地址转为 Address
	HexToAddress(addr string) Address

	// BytesToAddress 将字节数组转为 Address
	BytesToAddress(addr []byte) Address
}

type Address interface {
	// Bytes 返回该地址对应的字节数组
	Bytes() []byte
	// Hex 返回该地址的字符串
	Hex() string
}

//type ErrHandle interface {
//	// Assert 如果条件不满足，则使当前交易没有效果 — 用于检查内部错误。
//	Assert(condition bool)
//	// Require 如果条件不满足则撤销状态更改 - 用于检查由输入或者外部组件引起的错误。
//	Require(condition bool)
//	// Require2 如果条件不满足则撤销状态更改 - 用于检查由输入或者外部组件引起的错误，可以同时提供一个错误消息。
//	Require2(condition bool, message string)
//	// Revert 终止运行并撤销状态更改
//	Revert()
//	// Revert2 终止运行并撤销状态更改，可以同时提供一个解释性的字符串。
//	Revert2(reason string)
//}
