package rwfile

import (
	"DataCompress/getapi"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func Opt(file string){
	fileTmp,err:= os.Open(file)
	defer func(){fileTmp.Close()}()
	if err!=nil && os.IsNotExist(err){
		log.Printf("Not Find File!")
		return
	}

	fileNew := file[0:len(file)-4] + "_1.txt"
	//fileTmp2,err:= os.Open(fileNew)
	//defer func(){fileTmp2.Close()}()
	//if err!=nil && os.IsNotExist(err){
		os.Create(fileNew)
	//}

	strData := ReadFile(file)

	start := time.Now()

	indexP := 0
	//区块头 80
	head := strData[indexP : indexP + 160]
	//log.Printf("head %s ", head)
	WriteFile(fileNew,head)
	indexP += 160

	//交易数 变长
	txCount,selfLength := CompactSize(strData[indexP : indexP + 18])
	//log.Printf("txCount %s selfLength %s", txCount, selfLength)
	txCountStr := strData[indexP : indexP + selfLength]
	WriteFile(fileNew,txCountStr)
	indexP += selfLength

	//coinbase交易版本号 4
	ver := strData[indexP : indexP + 8]
	//log.Printf("ver %s ", ver)
	WriteFile(fileNew, ver)
	indexP += 8

	//若有witness，则此处一定为0001
	witnessAppend := strData[indexP : indexP + 4]
	if witnessAppend == "0001" {
		WriteFile(fileNew, witnessAppend)
		indexP += 4
	}

	//coinbase交易输入个数  变长
	txInCount := strData[indexP : indexP + 2]
	log.Printf("txInCount %s ", txInCount)
	indexP += 2

	//coinbase交易哈希值 32
	txHash := strData[indexP : indexP + 64]
	log.Printf("txHash %s ", txHash)
	indexP += 64

	//coinbase交易输出索引 4
	outIndex := strData[indexP : indexP + 8]
	log.Printf("outIndex %s ", outIndex)
	indexP += 8

	//coinbase脚本长度  变长
	scriptLength,selfLength2 := CompactSize(strData[indexP : indexP + 18])
	//log.Printf("scriptLength %s %s ", scriptLength, selfLength2)
	scriptLengthStr := strData[indexP : indexP + selfLength2]
	WriteFile(fileNew,scriptLengthStr)
	indexP += selfLength2

	//coinbase脚本 不定长
	script := strData[indexP : indexP + Convert2Int(scriptLength)*2]
	//log.Printf("script %s ", script)
	WriteFile(fileNew,script)
	indexP += Convert2Int(scriptLength)*2

	//序列号　４
	seq := strData[indexP : indexP + 8]
	//log.Printf("seq %s ", seq)
	WriteFile(fileNew,seq)
	indexP += 8

	//交易输出个数  变长
	txOutCount := strData[indexP : indexP + 2]
	txOutCountInt,_ := strconv.ParseInt(txOutCount,16,64)
	//log.Printf("txOutCount %s ", txOutCount)
	WriteFile(fileNew, txOutCount)
	indexP += 2

	for j:=0;j<Convert2Int(txOutCountInt);j++ {
		//奖励币数量 8
		amcount := strData[indexP : indexP+16]
		//log.Printf("amcount %s ", amcount)
		WriteFile(fileNew, amcount)
		indexP += 16

		//锁定脚本长度 变长
		signLength, selfLength3 := CompactSize(strData[indexP : indexP+18])
		//log.Printf("signLength %s selfLength3 %s ", signLength, selfLength3)
		signLengthStr := strData[indexP : indexP+selfLength3]
		WriteFile(fileNew, signLengthStr)
		indexP += selfLength3

		//锁定脚本
		sign := strData[indexP : indexP+Convert2Int(signLength)*2]
		//log.Printf("sign %s ", sign)
		WriteFile(fileNew, sign)
		indexP += Convert2Int(signLength) * 2
	}

	if witnessAppend == "0001" {
		strsCount, selfLength9 := CompactSize(strData[indexP : indexP+18])
		strsCountStr := strData[indexP : indexP + selfLength9]
		WriteFile(fileNew, strsCountStr)
		indexP += selfLength9

		for p:=0; p<Convert2Int(strsCount); p++ {
			witnessLength, selfLength8 := CompactSize(strData[indexP : indexP+18])
			witnessLengthStr := strData[indexP : indexP + selfLength8]
			WriteFile(fileNew, witnessLengthStr)
			indexP += selfLength8

			witnessStr := strData[indexP : indexP+Convert2Int(witnessLength)*2]
			WriteFile(fileNew, witnessStr)
			indexP += Convert2Int(witnessLength)*2
		}
	}

	//锁定时间　４
	signTime := strData[indexP : indexP+8]
	if signTime != "00000000" {
		log.Printf("signTime %s ", signTime)
	}
	WriteFile(fileNew, signTime)
	indexP += 8

	for i:=1;i<Convert2Int(txCount);i++ {
		//if i==434{
		//	log.Printf("txCount Index %s %s", i, txCount)
		//}
		//log.Printf("txCount Index %s %s", i, txCount)

		//交易版本号 4
		ver := strData[indexP : indexP+8]
		//	log.Printf("ver %s ", ver)
		WriteFile(fileNew, ver)
		indexP += 8

		witnessAppendTmp := strData[indexP : indexP + 4]
		if witnessAppendTmp == "0001" {
			WriteFile(fileNew, witnessAppend)
			indexP += 4
		}

		//交易输入个数  变长
		txInCount,selfLength := CompactSize(strData[indexP : indexP + 18])
		//log.Printf("txInCount %s selfLength %s", txInCount, selfLength)
		WriteFile(fileNew,strData[indexP : indexP + selfLength])
		indexP += selfLength

		//交易输入
		for j:=0;j<Convert2Int(txInCount);j++ {
			//log.Printf("txInCount Index %s %s", j, txInCount)
			//交易输入哈希值 32
			txHash := ConvertHash(strData[indexP : indexP + 64])
			//log.Printf("txHash %s ", txHash)

			//optStr := GetTxSimple(txHash)

			optStr := GetTxSimple2(txHash)

			if strings.HasPrefix(optStr, "error") {
				WriteFile(fileNew, optStr)
			} else {
				WriteFile(fileNew, optStr)
			}
			indexP += 64

			//交易输出索引 4
			//outIndex := strData[indexP : indexP + 8]
			//log.Printf("outIndex %s ", outIndex)
			tmpInt,_ := strconv.ParseInt(strData[indexP + 6 : indexP + 8] + strData[indexP + 4 : indexP + 6] + strData[indexP + 2 : indexP + 4] + strData[indexP : indexP + 2],16,64)
			strTmp := Convert2CompactSize(Convert2Int(tmpInt))
			log.Printf("outIndex2 %s ", strTmp)
			WriteFile(fileNew,strTmp)
			indexP += 8

			//解锁脚本长度  变长
			scriptLength,selfLength2 := CompactSize(strData[indexP : indexP + 18])
			//log.Printf("scriptLength %s selfLength2 %s ", scriptLength, selfLength2)
			scriptLengthStr := strData[indexP : indexP + selfLength2]
			WriteFile(fileNew,scriptLengthStr)
			indexP += selfLength2

			//解锁脚本 不定长
			script := strData[indexP : indexP + Convert2Int(scriptLength)*2]
			//log.Printf("script %s ", script)
			WriteFile(fileNew,script)
			indexP += Convert2Int(scriptLength)*2

			//序列号　４
			seq := strData[indexP : indexP + 8]
			if seq != "ffffffff" {
				log.Printf("seq %s ", seq)
			}
			WriteFile(fileNew,seq)
			indexP += 8
		}

		//交易输出个数  变长
		txOutCount,selfLength2 := CompactSize(strData[indexP : indexP + 18])
		//log.Printf("txOutCount %s selfLength2 %s", txOutCount, selfLength2)
		WriteFile(fileNew,strData[indexP : indexP + selfLength2])
		indexP += selfLength2

		//交易输出
		for j:=0;j<Convert2Int(txOutCount);j++ {
			//log.Printf("txOutCount Index %s %s", j, txOutCount)
			//交易额 8
			amcount := strData[indexP : indexP + 16]
			//log.Printf("amcount %s ", amcount)
			WriteFile(fileNew, amcount)
			indexP += 16

			//锁定脚本长度 变长
			signLength, selfLength3 := CompactSize(strData[indexP : indexP + 18])
			//log.Printf("signLength %s selfLength3 %s ", signLength, selfLength3)
			signLengthStr := strData[indexP : indexP + selfLength3]
			WriteFile(fileNew, signLengthStr)
			indexP += selfLength3

			//锁定脚本
			sign := strData[indexP : indexP + Convert2Int(signLength)*2]
			//log.Printf("sign %s ", sign)
			WriteFile(fileNew, sign)
			indexP += Convert2Int(signLength)*2
		}

		if witnessAppendTmp == "0001" {
			for j:=0;j<Convert2Int(txInCount);j++ { //交易输入
				strsCount, selfLength9 := CompactSize(strData[indexP : indexP+18])
				strsCountStr := strData[indexP : indexP+selfLength9]
				WriteFile(fileNew, strsCountStr)
				indexP += selfLength9

				for p := 0; p < Convert2Int(strsCount); p++ {
					witnessLength, selfLength8 := CompactSize(strData[indexP : indexP+18])
					witnessLengthStr := strData[indexP : indexP+selfLength8]
					WriteFile(fileNew, witnessLengthStr)
					indexP += selfLength8

					witnessStr := strData[indexP : indexP+Convert2Int(witnessLength)*2]
					WriteFile(fileNew, witnessStr)
					indexP += Convert2Int(witnessLength) * 2
				}
			}
		}

		//锁定时间　４
		signTime := strData[indexP : indexP + 8]
		if signTime != "00000000" {
			log.Printf("signTime %s ", signTime)
		}
		WriteFile(fileNew, signTime)
		indexP += 8
	}

	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}

func UnOpt(file string){
	fileTmp,err:= os.Open(file)
	defer func(){fileTmp.Close()}()
	if err!=nil && os.IsNotExist(err){
		log.Printf("Not Find File!")
		return
	}

	fileNew := file[0:len(file)-4]+"_2.txt"
	fileTmp2,err:= os.Open(fileNew)
	defer func(){fileTmp2.Close()}()
	if err!=nil && os.IsNotExist(err){
		os.Create(fileNew)
	}

	strData := ReadFile(file)

	indexP := 0
	//区块头 80
	head := strData[indexP : indexP + 160]
	log.Printf("head %s ", head)
	WriteFile(fileNew,head)
	indexP += 160

	//交易数 变长
	txCount,selfLength := CompactSize(strData[indexP : indexP + 18])
	log.Printf("txCount %s selfLength %s", txCount, selfLength)
	txCountStr := strData[indexP : indexP + selfLength]
	WriteFile(fileNew,txCountStr)
	indexP += selfLength

	//coinbase交易版本号 4
	ver := strData[indexP : indexP + 8]
	log.Printf("ver %s ", ver)
	WriteFile(fileNew,ver)
	indexP += 8

	//若有witness，则此处一定为0001
	witnessAppend := strData[indexP : indexP + 4]
	if witnessAppend == "0001" {
		WriteFile(fileNew, witnessAppend)
		indexP += 4
	}

	//coinbase交易输入个数  变长
	WriteFile(fileNew,"01")

	//coinbase交易哈希值 32
	WriteFile(fileNew,"0000000000000000000000000000000000000000000000000000000000000000")

	//coinbase交易输出索引 4
	WriteFile(fileNew,"ffffffff")

	//coinbase脚本长度  变长
	scriptLength,selfLength2 := CompactSize(strData[indexP : indexP + 18])
	log.Printf("scriptLength %s %s ", scriptLength, selfLength2)
	scriptLengthStr := strData[indexP : indexP + selfLength2]
	WriteFile(fileNew,scriptLengthStr)
	indexP += selfLength2

	//coinbase脚本 不定长
	script := strData[indexP : indexP + Convert2Int(scriptLength)*2]
	log.Printf("script %s ", script)
	WriteFile(fileNew,script)
	indexP += Convert2Int(scriptLength)*2

	//序列号　４
	seq := strData[indexP : indexP + 8]
	//log.Printf("seq %s ", seq)
	WriteFile(fileNew,seq)
	indexP += 8

	//交易输出个数  变长
	txOutCount := strData[indexP : indexP + 2]
	txOutCountInt,_ := strconv.ParseInt(txOutCount,16,64)
	log.Printf("txOutCount %s ", txOutCount)
	WriteFile(fileNew, txOutCount)
	indexP += 2

	for j:=0;j<Convert2Int(txOutCountInt);j++ {
		//奖励币数量 8
		amcount := strData[indexP : indexP+16]
		log.Printf("amcount %s ", amcount)
		WriteFile(fileNew, amcount)
		indexP += 16

		//锁定脚本长度 变长
		signLength, selfLength3 := CompactSize(strData[indexP : indexP+18])
		log.Printf("signLength %s selfLength3 %s ", signLength, selfLength3)
		signLengthStr := strData[indexP : indexP+selfLength3]
		WriteFile(fileNew, signLengthStr)
		indexP += selfLength3

		//锁定脚本
		sign := strData[indexP : indexP+Convert2Int(signLength)*2]
		log.Printf("sign %s ", sign)
		WriteFile(fileNew, sign)
		indexP += Convert2Int(signLength) * 2
	}

	if witnessAppend == "0001" {
		strsCount, selfLength9 := CompactSize(strData[indexP : indexP+18])
		strsCountStr := strData[indexP : indexP + selfLength9]
		WriteFile(fileNew, strsCountStr)
		indexP += selfLength9

		for p:=0; p<Convert2Int(strsCount); p++ {
			witnessLength, selfLength8 := CompactSize(strData[indexP : indexP+18])
			witnessLengthStr := strData[indexP : indexP + selfLength8]
			WriteFile(fileNew, witnessLengthStr)
			indexP += selfLength8

			witnessStr := strData[indexP : indexP+Convert2Int(witnessLength)*2]
			WriteFile(fileNew, witnessStr)
			indexP += Convert2Int(witnessLength)*2
		}
	}

	//锁定时间　４
	signTime := strData[indexP: indexP + 8]
	log.Printf("signTime %s ", signTime)
	WriteFile(fileNew, signTime)
	indexP += 8

	for i:=1;i<Convert2Int(txCount);i++ {
		//交易版本号 4
		ver := strData[indexP : indexP + 8]
		log.Printf("ver %s ", ver)
		WriteFile(fileNew,ver)
		indexP += 8

		witnessAppendTmp := strData[indexP : indexP + 4]
		if witnessAppendTmp == "0001" {
			WriteFile(fileNew, witnessAppend)
			indexP += 4
		}

		//交易输入个数  变长
		txInCount,selfLength := CompactSize(strData[indexP : indexP + 18])
		log.Printf("txCount %s selfLength %s", txInCount, selfLength)
		WriteFile(fileNew,strData[indexP : indexP + selfLength])
		indexP += selfLength

		//交易输入
		for j:=0;j<Convert2Int(txInCount);j++ {
			//交易输入哈希值 32
			txHash,selfLength := GetTxHash(strData[indexP : indexP + 18*2])
			log.Printf("txHash %s selfLength %s", txHash, selfLength)
			WriteFile(fileNew,ConvertHash(txHash))
			indexP += selfLength

			//交易输出索引 4
			outIndex,selfLength := Convert2Size(strData[indexP : indexP + 18])
			WriteFile(fileNew,outIndex)
			indexP += selfLength

			//解锁脚本长度  变长!!!
			scriptLength,selfLength2 := CompactSize(strData[indexP : indexP + 18])
			log.Printf("scriptLength %s selfLength2 %s ", scriptLength, selfLength2)
			scriptLengthStr := strData[indexP : indexP + selfLength2]
			WriteFile(fileNew,scriptLengthStr)
			indexP += selfLength2

			//解锁脚本 不定长
			script := strData[indexP : indexP + Convert2Int(scriptLength)*2]
			log.Printf("script %s ", script)
			WriteFile(fileNew,script)
			indexP += Convert2Int(scriptLength)*2

			//序列号　４
			seq := strData[indexP : indexP + 8]
			log.Printf("seq %s ", seq)
			WriteFile(fileNew,seq)
			indexP += 8
		}

		//交易输出个数  变长
		txOutCount,selfLength2 := CompactSize(strData[indexP : indexP + 18])
		log.Printf("txOutCount %s selfLength2 %s", txOutCount, selfLength2)
		WriteFile(fileNew,strData[indexP : indexP + selfLength2])
		indexP += selfLength2

		//交易输出
		for j:=0;j<Convert2Int(txOutCount);j++ {
			//交易额 8
			amcount := strData[indexP : indexP + 16]
			log.Printf("amcount %s ", amcount)
			WriteFile(fileNew, amcount)
			indexP += 16

			//锁定脚本长度 变长
			signLength, selfLength3 := CompactSize(strData[indexP : indexP + 18])
			log.Printf("signLength %s selfLength3 %s ", signLength, selfLength3)
			signLengthStr := strData[indexP : indexP + selfLength3]
			WriteFile(fileNew, signLengthStr)
			indexP += selfLength3

			//锁定脚本
			sign := strData[indexP : indexP + Convert2Int(signLength)*2]
			log.Printf("sign %s ", sign)
			WriteFile(fileNew, sign)
			indexP += Convert2Int(signLength)*2
		}

		if witnessAppendTmp == "0001" {
			for j:=0;j<Convert2Int(txInCount);j++ { //交易输入
				strsCount, selfLength9 := CompactSize(strData[indexP : indexP+18])
				strsCountStr := strData[indexP : indexP+selfLength9]
				WriteFile(fileNew, strsCountStr)
				indexP += selfLength9

				for p := 0; p < Convert2Int(strsCount); p++ {
					witnessLength, selfLength8 := CompactSize(strData[indexP : indexP+18])
					witnessLengthStr := strData[indexP : indexP+selfLength8]
					WriteFile(fileNew, witnessLengthStr)
					indexP += selfLength8

					witnessStr := strData[indexP : indexP+Convert2Int(witnessLength)*2]
					WriteFile(fileNew, witnessStr)
					indexP += Convert2Int(witnessLength) * 2
				}
			}
		}

		//锁定时间　４
		signTime := strData[indexP : indexP + 8]
		log.Printf("signTime %s ", signTime)
		WriteFile(fileNew, signTime)
		indexP += 8
	}
}

func CompactSize(str string)(int64,int){
	//strs :=[]rune(str)
	switch str[0:2] {
	case "fd":
		tmp := str[4:6] + str[2:4]
		tmpInt,_ := strconv.ParseInt(tmp,16,64)
		return tmpInt,6
	case "fe":
		tmp := str[8:10] + str[6:8]+ str[4:6] + str[2:4]
		tmpInt,_ := strconv.ParseInt(tmp,16,64)
		return tmpInt,10
	case "ff":
		tmp := str[16:18] + str[14:16] + str[12:14] + str[10:12] + str[8:10] + str[6:8]+ str[4:6] + str[2:4]
		tmpInt,_ := strconv.ParseInt(tmp,16,64)
		return tmpInt,18
	default:
		tmp := str[0:2]
		tmpInt,_ := strconv.ParseInt(tmp,16,64)
		return tmpInt,2
	}
}

func Convert2CompactSize(num int)(string){
	if num <= 252 {
		str := strconv.FormatInt(int64(num), 16)
		if len(str) == 1 {
			str = "0" + str
		}
		return str
	} else if num >= 253 && num <= 65535 {
		str := strconv.FormatInt(int64(num), 16)
		if len(str) == 2 {
			str = "00" + str
		} else if len(str) == 3 {
			str = "0" + str
		}
		return "fd" + str[2:4] + str[0:2]
	} else if num >= 65536 && num <= 4294967295 {
		str := strconv.FormatInt(int64(num), 16)
		if len(str) == 5 {
			str = "000" + str
		} else if len(str) == 6 {
			str = "00" + str
		}else if len(str) == 7 {
			str = "0" + str
		}
		return "fe" + str[6:8]+ str[4:6] + str[2:4] + str[0:2]
	} else if num >= 4294967296 {
		str := strconv.FormatInt(int64(num), 16)
		if len(str) == 9 {
			str = "0000000" + str
		} else if len(str) == 10 {
			str = "000000" + str
		}else if len(str) == 11 {
			str = "00000" + str
		}else if len(str) == 12 {
			str = "0000" + str
		}else if len(str) == 13 {
			str = "000" + str
		}else if len(str) == 14 {
			str = "00" + str
		}else if len(str) == 15 {
			str = "0" + str
		}
		return "ff" + str[14:16] + str[12:14] + str[10:12] + str[8:10] + str[6:8]+ str[4:6] + str[2:4] + str[0:2]
	}
	return ""
}

func Convert2Size(str string) (string,int){
	switch str[0:2] {
	case "fd":
		return str[4:6] + str[2:4] + "0000",6
	case "fe":
		return str[8:10] + str[6:8]+ str[4:6] + str[2:4],10
	default:
		tmp := str[0:2]
		return tmp + "000000",2
	}
}

func Convert2Int(num int64) (int){
	// 设置一个 int64 的数据
	//int64_num := int64(6)
	// 将 int64 转化为 int
	int_num := *(*int)(unsafe.Pointer(&num))
	return int_num
}

func GetTxSimple(hash string) (string) {
	start := time.Now()

	//交易索引不准确
	//data := getapi.GetData("https://blockchain.info/rawtx/"+hash)
	//查询次数限制
	data := getapi.GetData("https://api.blockcypher.com/v1/btc/main/txs/"+hash)

	//log.Printf("data %s ", data)
	if data == "error" {
		return data + "  " + hash
	}

	var txData TxStruct
	if err := json.Unmarshal([]byte(data), &txData); err == nil {
		//fmt.Println(txData)
	} else {
		log.Printf("GetTxSimple Error: %s ", err)
	}

	cost := time.Since(start)
	fmt.Printf("cost2=[%s],",cost)

	return Convert2CompactSize(txData.BlockHeight) + Convert2CompactSize(txData.TxIndex)
}

func GetTxSimple2(hash string) (string) {
	start := time.Now()

	data0 := getapi.GetData("https://blockchain.info/rawtx/"+hash)
	//log.Printf("data %s ", data)
	if data0 == "error" {
		//return "  " + data0 + "  " + hash + "  "

		return GetTxSimple(hash)
	}
	var txData TxStruct
	if err := json.Unmarshal([]byte(data0), &txData); err == nil {
		//fmt.Println(txData)
	} else {
		log.Printf("GetTxSimple Error: %s ", err)
	}

	heightStr := strconv.Itoa(txData.BlockHeight)

	data := getapi.GetData("https://blockchain.info/block-height/"+heightStr+"?format=json")
	//log.Printf("data %s ", data)

	if data == "error" || data == "Block Not Found"{
		//return " error " + hash + " "

		return GetTxSimple(hash)
	}

	var blockData blockStruct
	if err := json.Unmarshal([]byte(data), &blockData); err == nil {
		//fmt.Println(blockData)
	} else {
		log.Printf("GetTxHash Error: %s ", err)
	}

	tx_index := 0
	for p:=0; p< len(blockData.Blocks[0].Tx); p++ {
		if blockData.Blocks[0].Tx[p].TxHash == hash {
			tx_index = p
			break
		}
	}

	cost := time.Since(start)
	fmt.Printf("cost1=[%s],",cost)

	//heightInt,_ := strconv.Atoi(height)
	return Convert2CompactSize(txData.BlockHeight) + Convert2CompactSize(tx_index)


	/*
	//data := getapi.GetData2("https://sochain.com/api/v2/get_block/BTC/"+heightStr)
	data:= getapi.GetData("https://sochain.com/api/v2/get_block/BTC/"+heightStr)
	if data == "error" {
		return "  " + data + "  " + hash + "  "
	}

	var blockData blockStruct2
	if err := json.Unmarshal([]byte(data), &blockData); err == nil {
		//fmt.Println(blockData)
	} else {
		log.Printf("GetTxHash Error: %s ", err)

	}

	tx_index := 0
	for p:=0; p< len(blockData.Blocks.Tx); p++ {
		if blockData.Blocks.Tx[p] == hash {
			tx_index = p
			break
		}
	}

	cost := time.Since(start)
	fmt.Printf("cost=[%s],",cost)

	//heightInt,_ := strconv.Atoi(height)
	return Convert2CompactSize(txData.BlockHeight) + Convert2CompactSize(tx_index)

	 */
}

func GetTxHash(str string) (string,int){
	valueStr, selfLength := CompactSize(str[0:18])
	height := strconv.Itoa(Convert2Int(valueStr))

	valueStr2,selfLength2 := CompactSize(str[selfLength:selfLength+18])
	index := Convert2Int(valueStr2)

	data := getapi.GetData("https://blockchain.info/block-height/"+height+"?format=json")
	//log.Printf("data %s ", data)

	var blockData blockStruct
	if err := json.Unmarshal([]byte(data), &blockData); err == nil {
		//fmt.Println(blockData)
	} else {
		log.Printf("GetTxHash Error: %s ", err)
	}
	return blockData.Blocks[0].Tx[index].TxHash,selfLength + selfLength2
}

type TxStruct struct {
	BlockHeight int `json:"block_height"`
	//TxIndex int `json:"tx_index"`
	TxIndex int `json:"block_index"`
}

type blockStruct struct {
	Blocks []TxStruct2 `json:"blocks"`
}

type TxStruct2 struct {
	Tx []TxHashStruct `json:"tx"`
}

type TxHashStruct struct {
	TxHash string `json:"hash"`
}

type blockStruct2 struct {
	Blocks TxStruct22 `json:"data"`
}

type TxStruct22 struct {
	Tx []string `json:"txs"`
}

//原生数据是小端格式，json格式数据是大端格式
func ConvertHash(str string) (string){
	return str[62:64]+str[60:62]+str[58:60]+str[56:58]+str[54:56]+str[52:54]+str[50:52]+str[48:50]+str[46:48]+str[44:46]+str[42:44]+str[40:42]+str[38:40]+str[36:38]+str[34:36]+str[32:34] + str[30:32]+str[28:30]+str[26:28]+str[24:26]+str[22:24]+str[20:22]+str[18:20]+str[16:18]+str[14:16]+str[12:14]+str[10:12]+str[8:10]+str[6:8]+str[4:6]+str[2:4]+str[0:2]
}
