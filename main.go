package main

import (
	"DataCompress/bzip2"
	"DataCompress/ethtx"
	"DataCompress/flate"
	"DataCompress/gzip"
	"DataCompress/huffman"
	"DataCompress/key"
	"DataCompress/lzw"
	"DataCompress/rwfile"
	"DataCompress/zlib"
	"bufio"
	"bytes"
	"fmt"
	bzip2w "github.com/larzconwell/bzip2"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main(){
	//a := big.NewInt(111111112222222233333333);//1111111122222222333333334444444455555555666666667777777788888888
	//fmt.Println(unsafe.Sizeof(a))
	go putget()
	fmt.Println("completed!")
	select { }
}

func putget(){
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		data, err0 := stdReader.ReadString('\n')
		if err0 != nil {
			fmt.Printf("putValue error0: %s\n", err0)
			continue
		}
		cmd, v1, v2,v3,v4 := getDatas(data)//parseCmd(flag.Args())
		log.Printf("Values %s %s %s %s %s", cmd, v1, v2,v3,v4)

		switch cmd {
		case "readfile":
			rwfile.ReadFile(v1)

		case "writefile":
			rwfile.WriteFile2(v1,v2)

		case "test":
			tmpInt, selfLength := rwfile.CompactSize(v1)
			log.Printf("tmpInt %s selfLength %s ", tmpInt,selfLength)

		case "get1":
			rwfile.GetTxSimple("0437cd7f8525ceed2324359c2d0ba26006d92d856a9c20fa0241106ee5a597c9")

		case "opt":
			rwfile.Opt(v1)
		case "unopt":
			rwfile.UnOpt(v1)

		case "huffman":
			huffman.Encode(v1) //代码需要优化

		case "bzip0":
			bzip0() //压缩有问题
		case "bzip1":
			bzip1()

		case "flate0":
			//m,_:=strconv.Atoi(v1)
			flate0(v1)
		case "flate1":
			flate1(v1)

		case "gzip0":
			gzip0()
		case "gzip1":
			gzip1()

		case "lzw0":
			lzw0()
		case "lzw1":
			lzw1()

		case "zlib0":
			zlib0()
		case "zlib1":
			zlib1()

		case "txhash":
			Txhash(v1)

		case "encode":
			//pub_hash_2 := []byte("bbde7b2e3f41502c8be452b269ac45941a8cfb6f")
			//address := key.B58checkencode(0x00, pub_hash_2)
			address,_ := key.Encode("00", v1)
			fmt.Println("B58check encode address ", address)

		case "decode":
			decoded, _ := key.Decode(v1)
			fmt.Println("B58check decode pub hash ", decoded)

		case "ethopt":
			//ethtx.OptTest(v1)
			ethtx.Opt(v1)

		default:
			log.Printf("Command %s unrecognized", cmd)
		}
	}
}

func getDatas(data string) (string, string, string, string, string){
	datas := strings.Fields(data)
	if len(datas) == 5 {
		return datas[0], datas[1], datas[2], datas[3], datas[4]
	} else if len(datas) == 4 {
		return datas[0], datas[1], datas[2], datas[3],""
	} else if len(datas) == 3 {
		return datas[0], datas[1], datas[2],"",""
	} else if len(datas) == 2 {
		return datas[0], datas[1],"","",""
	} else if len(datas) == 1 {
		return datas[0], "","","",""
	} else {
		return "","","","",""
	}
}

var fileSource ="demo.txt"
var fileDes ="demo.compress"
var fileUn ="demo.uncompress"

func bzip0() {
	var buf bytes.Buffer
	var out bytes.Buffer
	expected := []byte("banana")

	writer := bzip2w.NewWriter(&buf)
	_, err := writer.Write(expected)
	if err == nil {
		err = writer.Close()
	}
	if err != nil {
		log.Fatalln(err)
	}

	reader := bzip2.NewReader(&buf)
	_, err = io.Copy(&out, reader)
	if err != nil {
		log.Fatalln(err)
	}

	if out.String() != string(expected) {
		log.Fatalln("Output is incorrect. Got", out.String(), "wanted",
			string(expected))
	}
	return

	newfile, err := os.Create(fileDes)
	if err != nil {
		log.Fatalln(err)
	}
	defer newfile.Close()

	file, err := os.Open(fileSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	write := bzip2w.NewWriter(newfile)
	defer write.Close()

	_, err = io.Copy(write, file)
	if err != nil {
		return
	}

	if err := write.Close(); err != nil {
		return
	}
	return
}

func bzip1() {
	file, err := os.Open(fileDes)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create(fileUn)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	zr := bzip2.NewReader(file)

	_, err = io.Copy(newfile, zr)
	if err != nil {
		panic(err)
	}

	return
}

func flate0(fileSource string) {
	fileDes := fileSource[0:len(fileSource)-4] + "_2.txt"
	newfile, err := os.Create(fileDes)
	if err != nil {
		log.Fatalln(err)
	}
	defer newfile.Close()

	file, err := os.Open(fileSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	start := time.Now()
	flateWrite, err := flate.NewWriter(newfile, flate.BestCompression)  //9
	//flateWrite, err := flate.NewWriter(newfile, typ)
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWrite.Close()

	_, err = io.Copy(flateWrite, file)
	if err != nil {
		return
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	flateWrite.Flush()
	if err := flateWrite.Close(); err != nil {
		return
	}
	return
}

func flate1(fileSource string) {
	fileDes := fileSource[0:len(fileSource)-4] + "_3.txt"

	file, err := os.Open(fileSource)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create(fileDes)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	start := time.Now()
	zr := flate.NewReader(file)
	defer zr.Close()

	_, err = io.Copy(newfile, zr)
	if err != nil {
		panic(err)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	if err := zr.Close(); err != nil {
		panic(err)
	}
	return
}

func gzip0() {
	newfile, err := os.Create(fileDes)
	if err != nil {
		log.Fatalln(err)
	}
	defer newfile.Close()

	file, err := os.Open(fileSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	start := time.Now()
	zw := gzip.NewWriter(newfile)

	//filestat, err := file.Stat()
	//if err != nil {
	//	return
	//}
	//zw.Name = filestat.Name()
	//zw.ModTime = filestat.ModTime()

	_, err = io.Copy(zw, file)
	if err != nil {
		return
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	zw.Flush()
	if err := zw.Close(); err != nil {
		return
	}
	return

	fw, err := os.Create("demo.gzip")   // 创建gzip包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	// 实例化新的gzip.Writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 获取要打包的文件信息
	fr, err := os.Open("demo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 获取文件头信息
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建gzip.Header
	gw.Header.Name = fi.Name()

	// 读取文件数据
	buf := make([]byte, fi.Size())
	fmt.Println("File Size:", fi.Size())
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// 写入数据到zip包
	_, err = gw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
}

func gzip1() {
	file, err := os.Open(fileDes)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create(fileUn)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	start := time.Now()
	zr, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}

	//filestat, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//zr.Name = filestat.Name()
	//zr.ModTime = filestat.ModTime()
	_, err = io.Copy(newfile, zr)
	if err != nil {
		panic(err)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	if err := zr.Close(); err != nil {
		panic(err)
	}
	return

	// 打开gzip文件
	fr, err := os.Open("demo.gzip")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 创建gzip.Reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Fatalln(err)
	}
	defer gr.Close()

	// 读取文件内容
	buf := make([]byte, 1024 * 1024 * 3)// 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩gzip文件中有多个文件时，强烈建议直接和tar组合使用）
	n, err := gr.Read(buf)
	fmt.Println("File Size:", n)
	// 将包中的文件数据写入
	fw, err := os.Create(gr.Header.Name)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fw.Write(buf[:n])
	if err != nil {
		log.Fatalln(err)
	}
}

func lzw0() {
	newfile, err := os.Create(fileDes)
	if err != nil {
		log.Fatalln(err)
	}
	defer newfile.Close()

	file, err := os.Open(fileSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	start := time.Now()
	write := lzw.NewWriter(newfile, lzw.LSB,8)
	defer write.Close()

	_, err = io.Copy(write, file)
	if err != nil {
		return
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	if err := write.Close(); err != nil {
		return
	}
	return
}

func lzw1() {
	file, err := os.Open(fileDes)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create(fileUn)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	start := time.Now()
	zr := lzw.NewReader(file, lzw.LSB,8)

	_, err = io.Copy(newfile, zr)
	if err != nil {
		panic(err)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	if err := zr.Close(); err != nil {
		panic(err)
	}
	return
}

func zlib0() {
	newfile, err := os.Create(fileDes)
	if err != nil {
		log.Fatalln(err)
	}
	defer newfile.Close()

	file, err := os.Open(fileSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	start := time.Now()
	write := zlib.NewWriter(newfile)
	defer write.Close()

	_, err = io.Copy(write, file)
	if err != nil {
		return
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	if err := write.Close(); err != nil {
		return
	}
	return
}

func zlib1() {
	file, err := os.Open(fileDes)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create(fileUn)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	start := time.Now()
	zr,_ := zlib.NewReader(file)

	_, err = io.Copy(newfile, zr)
	if err != nil {
		panic(err)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

	if err := zr.Close(); err != nil {
		panic(err)
	}
	return
}

func Txhash(file string) {
	fileTmp, err := os.Open(file)
	defer func() { fileTmp.Close() }()
	if err != nil && os.IsNotExist(err) {
		log.Printf("Not Find File!")
		return
	}


	strData := rwfile.ReadFile(file)


	count:=strings.Count(strData,"error  ")
	for i:=0; i<count;i++  {
		index := strings.Index(strData, "error  ")

		if index > -1 {
			indexStart := index + 7
			hashStr := strData[indexStart : indexStart+64]

			newStr := rwfile.GetTxSimple(hashStr)

			strData = strings.Replace(strData, "error  "+hashStr, newStr, -1)
		}
	}

	//fileNew := file[0:len(file)-4] + "_9.txt"
	rwfile.WriteFile2(file,strData)
}