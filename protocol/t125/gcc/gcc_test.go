package gcc_test

import (
	"bytes"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"github.com/icodeface/grdp/protocol/t125"
	"github.com/icodeface/grdp/protocol/t125/gcc"
	"testing"
)

func TestMakeConferenceCreateRequest(t *testing.T) {
	//userData, _ := hex.DecodeString("01c0d800040008000005200301ca03aa09040000ce0e0000720064007000790000000000000000000000000000000000000000000000000004000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001ca01000000000018000f0001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003c008000000000002c00c000b00000000000000")
	//
	//clientCoreData := gcc.NewClientCoreData()
	////data := gcc.NewClientSecurityData()
	////clientNetData := gcc.NewClientNetworkData()
	//
	//fmt.Println(hex.EncodeToString(clientCoreData.Block()))
	//
	//ccReq := gcc.MakeConferenceCreateRequest(userData)
	//fmt.Println(hex.EncodeToString(ccReq))

	userDataBuff := bytes.Buffer{}
	userDataBuff.Write(gcc.NewClientCoreData().Block())
	userDataBuff.Write(gcc.NewClientNetworkData().Block())
	userDataBuff.Write(gcc.NewClientSecurityData().Block())
	ccReq := gcc.MakeConferenceCreateRequest(userDataBuff.Bytes())

	connectInitial := t125.NewConnectInitial(ccReq)
	connectInitialBerEncoded, err := asn1.Marshal(connectInitial)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hex.EncodeToString(connectInitialBerEncoded))

	d := t125.NewDomainParameters(34, 2, 0, 1, 0, 1, 0xffff, 2)
	dd, err := asn1.Marshal(*d)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hex.EncodeToString(dd))
}

// rdpy 3019 020122020102020100020101020100020101 0202 ffff 020102
// me 	301a 020122020102020100020101020100020101 0203 00ffff 020102

// rdpy
// 040101
// 040101
// 0101
// ff30190201220201020201000201010201000201010202ffff020102301902010102010102010102010102010002010102020420020102301c0202ffff0202fc170202ffff0201010201000201010202ffff02010204820103000500147
// c000180fa000800100001c0004475636180ec01c0d800040008000005200301ca03aa09040000ce0e0000720064007000790000000000000000000000000000000000000000000000000004000000000000000c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
// 01ca01000000000018000f0001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003c008000000000002c00c000b00000000000000

// me 30820164
// 040101
// 040101
// 0101
// ff301a020122020102020100020101020100020101020300ffff0201023019020101020101020101020101020100020101020204200201023020020300ffff020300fc17020300ffff020101020100020101020300ffff0201020481ff000500147
// c000180f6000800100001c0004475636180e801c0d800040008000005200301ca03aa09040000ce0e00006d73747363000000000000000000000000000000000000000000000000000000040000000c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
// 01ca01000000000018000f0001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003c008000000000002c00c000b00000000000000

//00
// 0500147c0001
// 0e // 长度+14
// 00
// 08
// 00 10
// 00
// 01
// c0  per.WriteChoice(0xc0, buff)
// 00 44:75:63:61
// 00

//00:05:00:14:7c:00:01:
// 80:fa
// 00:
// 08:
// 00: 10:
// 00:
// 01:
// c0: per.WriteChoice(0xc0, buff)
// 00 44:75:63:61:

// 80:ec:
// 01:c0:d8:00:04:00:08:00:00:05:20:03:01:ca:03:aa:09:04:00:00:ce:0e:00:00:72:00:64:00:70:00:79:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:04:00:00:00:00:00:00:00:0c:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:01:ca:01:00:00:00:00:00:18:00:0f:00:01:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:03:c0:08:00:00:00:00:00:02:c0:0c:00:0b:00:00:00:00:00:00:00

// 01c0
// d800
// 040008000005200301ca03aa09040000ce0e0000676f726470
// 000000000000000000000000000000000000000000000000000000040000000
// c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001ca01000000000018000f00010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

// user data
// 01c0
// d800
// 040008000005200301ca03aa09040000ce0e00007200640070
// 00790000000000000000000000000000000000000000000000000004000000000000000
// c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001ca01000000000018000f00010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

// client network data
// 03c0
// 08 00
// 00 00 00 00

// client sec data
// 02c0
// 0c00
// 0b00000000000000
//sendConnectInitial

// 3081c58401018401010101ff301a820122820102820100820101820100820101820300ffff8201023019820101820101820101820101820100820101820204208201023020820300ffff820300fc17820300ffff820101820100820101820300ffff8201024061000500147c00015a000800100001c000447563614c
// 01c0d800040008000005200301ca03aa09040000ce0e00006d73747363000000000000000000000000000000000000000000000000000000
// 03c0080000000000
// 02c00c000b00000000000000
