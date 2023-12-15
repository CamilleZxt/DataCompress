package getapi

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetData2(urlStr string) (string){
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, _ := http.NewRequest("GET", urlStr, nil)

	reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding","gzip, deflate, br")
	reqest.Header.Set("Accept-Language","	zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	reqest.Header.Set("Cache-Control","max-age=0")
	reqest.Header.Set("Connection","keep-alive")
	reqest.Header.Set("User-Agent","	Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:75.0) Gecko/20100101 Firefox/75.0")
	reqest.Header.Set("Host","sochain.com")
	reqest.Header.Set("Upgrade-Insecure-Requests","1")

	response,err := client.Do(reqest)
	if err != nil {
		log.Printf("GetData Error: %s  %s", err, urlStr)
		return "error"
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body,err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("GetData Error: %s  %s", err, urlStr)
			return "error"
		}
		return string(body)
	}else{
		return "error"
	}
}

func GetData3(urlStr string) (string){
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response,err := client.Get(urlStr)

	if err != nil {
		log.Printf("GetData Error: %s  %s", err, urlStr)
		return "error"
	}
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("GetData Error: %s  %s", err, urlStr)
		return "error"
	}
	//fmt.Println(string(body))

	if response.StatusCode == 200 {
		return string(body)
	}else{
		return "error"
	}
}

func GetData(urlStr string) (string){
	response,err := http.Get(urlStr)

	if err != nil {
		log.Printf("GetData Error: %s  %s", err, urlStr)
		return "error"
	}
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("GetData Error: %s  %s", err, urlStr)
		return "error"
	}
	//fmt.Println(string(body))

	if response.StatusCode == 200 {
		return string(body)
	} else if response.StatusCode == 429 {
		return "error"
	} else{
		return "error"
	}

	//client := &http.Client{}
	//reqest, _ := http.NewRequest("GET", urlStr, nil)
	//
	//reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	//reqest.Header.Set("Cache-Control","max-age=0")
	//reqest.Header.Set("Connection","keep-alive")
	//reqest.Header.Set("User-Agent","chrome 100")
	//
	//response,_ := client.Do(reqest)
	//if response.StatusCode == 200 {
	//	body, _ := ioutil.ReadAll(response.Body)
	//	bodystr := string(body);
	//	fmt.Println(bodystr)
	//}

	//  reqest, _ = http.NewRequest("POST","http:/127.0.0.1/", bytes.NewBufferString(data.Encode()))
	//    respet1,_ := http.NewRequest("POST","http://127.0.0.1/",url.Values{"key":"Value"})
	//    reqest1.Header.Set("User-Agent","chrome 100")
	//    client.Do(reqest1)
}

func PostData(urlStr string){
	http.PostForm(urlStr,
		url.Values{"name": {"ruifengyun"}, "blog": {"xiaorui.cc"},
			"aihao":{"python golang"},"content":{"nima,fuck "}})
}


func httpRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Content-Length", "1604")
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")
	request.Header.Add("TE","Trailers")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	return client.Do(request)
}
// 根据指定的URL进行数据抓取
func Fetch(url string)(string,error)  {
	resp, err := httpRequest(url)
	if err != nil{
		return "error",err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "error", fmt.Errorf("wrong status code: %d of %s", resp.StatusCode, url)
	}
	data,_ := ioutil.ReadAll(resp.Body)
	return string(data),nil

}