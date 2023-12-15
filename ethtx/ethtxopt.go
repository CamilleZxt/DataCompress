package ethtx

import (
	"DataCompress/getapi"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func OptTest(strData string) {
	fileNew := "E:\\1_1.txt"
	os.Create(fileNew)

	//数据总长度的长度
	indexP := 0
	lengthP := 1 * 2
	strTmp := strData[indexP : indexP+lengthP]
	length := GetLength1(strTmp)
	//WriteFile(fileNew,strTmp)
	indexP += lengthP
	//数据总长度
	lengthP = length * 2
	strTmp = strData[indexP : indexP+lengthP]
	//WriteFile(fileNew,strTmp)
	indexP += lengthP

	// AccountNonce长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		WriteFile(fileNew, strTmp)
		indexP += lengthP
		//AccountNonce
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		WriteFile(fileNew, strTmp)
	}
	indexP += lengthP

	// Fee wei长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		lenTmp := strTmp
		indexP += lengthP
		//Fee wei
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		//WriteFile(fileNew,strTmp)
		newStr := GetRLPZero(lenTmp + strTmp)
		WriteFile(fileNew, newStr)
	}
	indexP += lengthP

	//Gas Limit长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		lenTmp := strTmp
		indexP += lengthP
		//Gas Limit
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		//WriteFile(fileNew,strTmp)
		newStr := GetRLPZero(lenTmp + strTmp)
		WriteFile(fileNew, newStr)
	}
	indexP += lengthP

	//Recipient Address长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	//WriteFile(fileNew,strTmp)
	indexP += lengthP
	//Recipient Address
	lengthP = length * 2
	strTmp = strData[indexP : indexP+lengthP]
	txHash := GetTx(strTmp)
	heightIndex := GetHeightIndex(txHash)
	WriteFile(fileNew, heightIndex)
	indexP += lengthP
}

func Opt(file string) {

	fileTmp, err := os.Open(file)
	defer func() { fileTmp.Close() }()
	if err != nil && os.IsNotExist(err) {
		log.Printf("Not Find File!")
		return
	}

	fileNew := file + "_1.txt"
	os.Create(fileNew)

	strData := ReadFile(file)

	start := time.Now()

	//数据总长度的长度
	indexP := 0
	lengthP := 1 * 2
	strTmp := strData[indexP : indexP+lengthP]
	length := GetLength1(strTmp)
	//WriteFile(fileNew,strTmp)
	indexP += lengthP
	//数据总长度
	lengthP = length * 2
	strTmp = strData[indexP : indexP+lengthP]
	//WriteFile(fileNew,strTmp)
	indexP += lengthP

	// AccountNonce长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		WriteFile(fileNew, strTmp)
		indexP += lengthP
		//AccountNonce
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		WriteFile(fileNew, strTmp)
	}
	indexP += lengthP

	// Fee wei长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		lenTmp := strTmp
		indexP += lengthP
		//Fee wei
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		//WriteFile(fileNew,strTmp)
		newStr := GetRLPZero(lenTmp + strTmp)
		WriteFile(fileNew, newStr)
	}
	indexP += lengthP

	//Gas Limit长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		lenTmp := strTmp
		indexP += lengthP
		//Gas Limit
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		//WriteFile(fileNew,strTmp)
		newStr := GetRLPZero(lenTmp + strTmp)
		WriteFile(fileNew, newStr)
	}
	indexP += lengthP

	//Recipient Address长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	//WriteFile(fileNew,strTmp)
	indexP += lengthP
	//Recipient Address
	lengthP = length * 2
	strTmp = strData[indexP : indexP+lengthP]
	txHash := GetTx(strTmp)
	heightIndex := GetHeightIndex(txHash)
	WriteFile(fileNew, heightIndex)
	indexP += lengthP

	//Amount长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	if length == -1 {
		WriteFile(fileNew, strTmp)
	} else {
		lenTmp := strTmp
		indexP += lengthP
		//Amount
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		//WriteFile(fileNew,strTmp)
		newStr := GetRLPZero(lenTmp + strTmp)
		WriteFile(fileNew, newStr)
	}
	indexP += lengthP

	//Payload长度
	lengthP = 1 * 2
	strTmp = strData[indexP : indexP+lengthP]
	length = GetLength2(strTmp)
	//WriteFile(fileNew,strTmp)
	indexP += lengthP
	strTmp6 := ""
	if strTmp > "b7" {
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		length = GetLength0(strTmp)
		indexP += lengthP

		//Payload
		lengthP = length * 2
		strTmp = strData[indexP : indexP+lengthP]
		//智能合约MethodID
		strTmp6 += strTmp[0:8]
		// 智能合约参数
		strTmp6 += GetZeroCode(strTmp[8:])
		WriteFile(fileNew, strTmp6)
	} else {
		WriteFile(fileNew, strTmp)
	}
	indexP += lengthP

	//VRS
	strTmp = strData[indexP:]
	WriteFile(fileNew, strTmp)

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}

func ReadFile(file string) string {
	//从命令行标记参数中获取文件路径
	fptr := flag.String("fpath", file, "the file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error: ", err)
	}
	strTmp := string(data)
	fmt.Println(strTmp)
	return strTmp
}

func WriteFile(file, str string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err = fmt.Fprint(f, str)
	//_, err = fmt.Fprintln(f, str)
	if err != nil {
		log.Fatal(err)
	}
}

func Convert2Int(num int64) int {
	// 设置一个 int64 的数据
	//int64_num := int64(6)
	// 将 int64 转化为 int
	int_num := *(*int)(unsafe.Pointer(&num))
	return int_num
}

func GetLength0(hexStr string) int {
	tmpInt, _ := strconv.ParseInt(hexStr, 16, 64)
	return Convert2Int(tmpInt)
}

//RLP总长度：总长度小于 55，(0xc0+长度的长度)+总长度；总长度大于 55，(0xf7+长度的长度)+总长度；
func GetLength1(hexStr string) int {
	tmpInt, _ := strconv.ParseInt(hexStr, 16, 64)
	if hexStr > "f7" { //表示取hexStr数值后返回值的位数，其值为需要取值的长度
		return Convert2Int(tmpInt - 247)
	} else { //表示取hexStr数值后返回值的位数
		return Convert2Int(tmpInt - 192)
	}
}

//([0x00, 0x7f]）之间编码是其本身;长度为小于 55，0x80+长度 ;长度为大于 55，(0xb7+长度的长度)+长度 ;
func GetLength2(hexStr string) int {
	tmpInt, _ := strconv.ParseInt(hexStr, 16, 64)
	if hexStr <= "80" { //表示取hexStr数值本身
		return -1
	} else if hexStr < "b7" { //表示取hexStr数值后返回值的位数
		return Convert2Int(tmpInt - 128)
	} else { //表示取hexStr数值后返回值的位数，其值为需要取值的长度
		return Convert2Int(tmpInt - 183)
	}
}

func GetRLPCode(k int64) string {
	hexStr := strconv.FormatInt(k, 16)
	length := len(hexStr)
	if length%2 == 1 {
		length = length + 1
		hexStr = "0" + hexStr
	}
	if k < 128 { //表示取hexStr数值本身
		return hexStr
	} else {
		length = length / 2
		if length < 55 { //表示取hexStr数值后返回值的位数
			hexStrLength := strconv.FormatInt(int64(128+length), 16)
			return hexStrLength + hexStr
		} else { //表示取hexStr数值后返回值的位数，其值为需要取值的长度
			lengthHex := strconv.FormatInt(int64(length), 16)
			length2 := len(lengthHex)
			if length2%2 == 1 {
				length2 = length2 + 1
				lengthHex = "0" + lengthHex
			}
			length2 = length2 / 2
			hexStrLength1 := strconv.FormatInt(int64(183+length2), 16)
			return hexStrLength1 + lengthHex + hexStr
		}
	}
}

func GetRLPStr(str string) string {
	length := len(str) / 2
	if length < 128 { //表示取hexStr数值本身
		indexStr := strconv.FormatInt(int64(length+128), 16) //strconv.Itoa(length+128)
		if len(indexStr)%2 != 0 {
			indexStr = "0" + indexStr
		}
		return indexStr + str
	}
	return str
}

//将包含很多0的数字转化为较短的RLP编码 110000为 a1FF04
func GetRLPZero(str string) string {
	strTmp := str[0:2]
	if strTmp < "80" {
		return str
	} else {
		tmpInt, _ := strconv.ParseInt(strTmp, 16, 64)
		if strTmp < "b7" {
			length := Convert2Int(tmpInt - 128)
			strTmp1 := str[2 : 2*length+2]
			tmpInt1, _ := strconv.ParseInt(strTmp1, 16,
				64)
			countZero := GetLastZeroCount(tmpInt1, 0)

			strTmp2 := strconv.FormatInt(tmpInt1, 10)
			strTmp2 = strTmp2[0 : len(strTmp2)-int(countZero)]
			intTmp2, _ := strconv.ParseInt(strTmp2, 10, 64)

			res := GetRLPCode(intTmp2) + "FF" + GetRLPCode(countZero)
			if len(res) < len(str) {
				return res
			} else {
				return str
			}
		} else {

		}
	}
	return str
}

func GetLastZeroCount(k, j int64) int64 {
	if k%10 == 0 {
		j = GetLastZeroCount(k/10, j+1)
		return j
	}
	return j
}

func GetZeroCode(str string) string {
	strRet := ""
	for i := 0; i < len(str)/64; i++ {
		index := 64 * i
		strTmp := str[index : index+64]
		indexZero := 0
		for j := 0; j < 64; j = j + 2 {
			if strTmp[j:j+2] != "00" {
				break
			}
			indexZero++
		}
		//indexStr :=strconv.Itoa(indexZero)
		//indexStr := strconv.FormatInt(int64(indexZero), 16)
		//if(len(indexStr)%2 !=0){
		//	indexStr = "0"+indexStr
		//}
		//strRet = strRet + "FF" + GetRLPCode(int64(indexZero*2)) + GetRLPStr(strTmp[indexZero*2 : 64])
		strRet = strRet + GetRLPStr(strTmp[indexZero*2:64])
	}
	if len(strRet) < len(str) {
		return strRet
	} else {
		return str
	}
}

func GetTx(address string) string {
	start := time.Now()

	data := getapi.GetData("https://api.blockchair.com/ethereum/dashboards/address/0x" + address + "?limit=10&offset=0")
	log.Printf("data %s ", data)
	if data == "error" {
		return data + "  " + address
	}
	dataTmp := gojsonq.New().FromString(data).From("data.0x"+address+".calls").Select("transaction_hash", "recipient")
	lst, ok := dataTmp.Get().([]interface{})
	if !ok {
		fmt.Println("Convert Tx error")
	} else {
		for _, item := range lst {
			items, ok := item.(map[string]interface{})
			if !ok {
				fmt.Println("Convert Tx error")
			}
			if items["recipient"] == "0x"+address {
				return items["transaction_hash"].(string)
			}
		}
	}
	cost := time.Since(start)
	fmt.Printf("cost2=[%s],", cost)

	return ""
}

func GetHeightIndex(txHash string) string {
	start := time.Now()

	data := getapi.GetData("https://api.blockchair.com/ethereum/dashboards/transaction/" + txHash)
	log.Printf("data %s ", data)
	if data == "error" {
		return data + "  " + txHash
	}
	//dataTmp := gojsonq.New().FromString(data).From("data."+txHash+".transaction").Select("block_id","index")
	blockIdTmp := gojsonq.New().FromString(data).Find("data." + txHash + ".transaction.block_id")
	indexTmp := gojsonq.New().FromString(data).Find("data." + txHash + ".transaction.index")
	blockId := GetRLPCode(int64(blockIdTmp.(float64))) //strconv.ParseInt(blockIdTmp.(string), 16, 64)
	index := GetRLPCode(int64(indexTmp.(float64)))     //strconv.ParseInt(indexTmp.(string), 16, 64)
	fmt.Println("blockId:" + strconv.Itoa(int(blockIdTmp.(float64))) + " indexTmp:" + strconv.Itoa(int(indexTmp.(float64))))
	return blockId + index

	cost := time.Since(start)
	fmt.Printf("cost2=[%s],", cost)

	return ""
}
