package data

//获取其他网站的资源
// 是否可以开仓

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

var ISOK bool

// 返回股票信息
func GetStockJson(stock_code string) {
	var stock_birth string
	if strings.HasPrefix(stock_code, "300") || strings.HasPrefix(stock_code, "002") || strings.HasPrefix(stock_code, "000") {
		stock_birth = "sz"
	} else {
		stock_birth = "sh"
	}
	stock_url := fmt.Sprintf("http://web.juhe.cn:8080/finance/stock/hs?gid=%v%v&type=&key=3dad52c26c6cb338ba75dced2c865488", stock_birth, stock_code)
	r, err := http.Get(stock_url)
	fmt.Printf("r: %#v\n", r)
	if err != nil {
		panic(err)
	}
	var result string
	for {
		buf := make([]byte, 4096)
		n, err2 := r.Body.Read(buf)
		defer r.Body.Close()
		if n == 0 {
			break
		}
		if err2 != io.EOF && err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		}
		result += string(buf[0:n])
	}
	// fmt.Printf("result: %v\n", result)
	store_filename := fmt.Sprintf("templates/html/%v.json", stock_code)
	f, _ := os.OpenFile(store_filename, os.O_CREATE|os.O_RDWR, 0777)
	defer f.Close()
	n, err2 := f.WriteString(result)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Println("保存成功!")
		fmt.Printf("一共%v个字符\n", n)
	}
}

func GetData() {
	// 获取是否可以开仓
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
	c.OnHTML(".board-infos dl:nth-child(6)", func(e *colly.HTMLElement) {
		incre := e.DOM.Find("dd").Text()
		if incre[1:4] < "1.5" {
			ISOK = true
			fmt.Println("今天环境适合开仓")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Cookie", "spversion=20130314; __utmz=156575163.1663516834.2.2.utmcsr=stock.10jqka.com.cn|utmccn=(referral)|utmcmd=referral|utmcct=/company.shtml; Hm_lvt_78c58f01938e4d85eaf619eae71b4ed1=1663516756,1663581512,1663598080,1663637253; __utma=156575163.1716118374.1663301989.1663598084.1663637260.4; __utmc=156575163; historystock=000428|*|001238|*|000620|*|688035|*|000002; log=; Hm_lpvt_78c58f01938e4d85eaf619eae71b4ed1=1663644455; v=AxkN3FL_nqFSAkI1l4OBH3F5KA72pg1Y95ox7DvOlcC_Qjdwg_YdKIfqQbDI")
		fmt.Printf("r.URL: %v\n", r.URL)
	})
	c.Visit("http://q.10jqka.com.cn/thshy/detail/code/883987/")

	//获取问财数据
	c2 := colly.NewCollector()
	c2.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
	c2.OnHTML("", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Html())
	})
	c2.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Cookie", "other_uid=Ths_iwencai_Xuangu_hgsrhbs9urjq18r5c1309dkan3dxdagr; ta_random_userid=1ozlmimb9g; cid=7134ae1675e327338ab7d9d7dbc356d81663306959; cid=7134ae1675e327338ab7d9d7dbc356d81663306959; ComputerID=7134ae1675e327338ab7d9d7dbc356d81663306959; WafStatus=0; PHPSESSID=e9d46eb5f98dc3741736325766c3e3c2; v=A_rvgZJMTVwV18EwAm5SEtbsSysZq36E8C_yKQTzpg1Y95SV7DvOlcC_QjXX")
		fmt.Printf("r.URL: %v\n", r.URL)
	})
	c2.Visit("http://www.iwencai.com/customized/chart/get-robot-data")

}
