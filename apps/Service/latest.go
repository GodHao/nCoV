package Service

import (
	"encoding/json"
	"fmt"
	"log"
	"nCoV-API/lib/conf"
	"nCoV-API/lib/util"
	"time"
)

type LatestData struct {
	Count struct {
		Confirmed int `json:"confirmed"`
		Suspected int `json:"suspected"`
		Cure      int `json:"cure"`
		Death     int `json:"death"`
	} `json:"count"`
	Info struct {
		// 疫情名称
		Name string `json:"name"`
		// 易感人群
		VulnerableGroup string `json:"vulnerable_group"`
		// 潜伏期
		IncubationPeriod string `json:"incubation_period"`
		// 数据源
		DataSource string `json:"data_source"`
		// 传染源
		InfectionSource string `json:"infection_source"`
		// 传播路径
		Transmission string `json:"transmission"`
	} `json:"info"`
	// 数据更新时间
	UpdateTime int64 `json:"update_time"`
}

type NcovApiRes struct {
	Results []struct {
		AbroadRemark   string `json:"abroadRemark"`
		Confirmed      int    `json:"confirmed"`
		ConfirmedCount int    `json:"confirmedCount"`
		CountRemark    string `json:"countRemark"`
		Cured          int    `json:"cured"`
		CuredCount     int    `json:"curedCount"`
		DailyPic       string `json:"dailyPic"`
		DeadCount      int    `json:"deadCount"`
		Death          int    `json:"death"`
		GeneralRemark  string `json:"generalRemark"`
		InfectSource   string `json:"infectSource"`
		PassWay        string `json:"passWay"`
		Remark1        string `json:"remark1"`
		Remark2        string `json:"remark2"`
		Remark3        string `json:"remark3"`
		Remark4        string `json:"remark4"`
		Remark5        string `json:"remark5"`
		Summary        string `json:"summary"`
		Suspect        int    `json:"suspect"`
		SuspectedCount int    `json:"suspectedCount"`
		UpdateTime     int64  `json:"updateTime"`
		Virus          string `json:"virus"`
	} `json:"results"`
	Success bool `json:"success"`
}

type TxApiRes struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Newslist []struct {
		News []struct {
			ID           int    `json:"id"`
			PubDate      int64  `json:"pubDate"`
			PubDateStr   string `json:"pubDateStr"`
			Title        string `json:"title"`
			Summary      string `json:"summary"`
			InfoSource   string `json:"infoSource"`
			SourceURL    string `json:"sourceUrl"`
			ProvinceID   string `json:"provinceId"`
			ProvinceName string `json:"provinceName,omitempty"`
			CreateTime   int64  `json:"createTime"`
			ModifyTime   int64  `json:"modifyTime"`
		} `json:"news"`
		Case []struct {
			ID                int    `json:"id"`
			CreateTime        int64  `json:"createTime"`
			ModifyTime        int64  `json:"modifyTime"`
			Tags              string `json:"tags"`
			CountryType       int    `json:"countryType"`
			ProvinceID        string `json:"provinceId"`
			ProvinceName      string `json:"provinceName"`
			ProvinceShortName string `json:"provinceShortName"`
			CityName          string `json:"cityName"`
			ConfirmedCount    int    `json:"confirmedCount"`
			SuspectedCount    int    `json:"suspectedCount"`
			CuredCount        int    `json:"curedCount"`
			DeadCount         int    `json:"deadCount"`
			Comment           string `json:"comment"`
			Sort              int    `json:"sort"`
			Operator          string `json:"operator"`
		} `json:"case"`
		Desc struct {
			ID             int    `json:"id"`
			CreateTime     int64  `json:"createTime"`
			ModifyTime     int64  `json:"modifyTime"`
			InfectSource   string `json:"infectSource"`
			PassWay        string `json:"passWay"`
			ImgURL         string `json:"imgUrl"`
			DailyPic       string `json:"dailyPic"`
			Summary        string `json:"summary"`
			Deleted        bool   `json:"deleted"`
			CountRemark    string `json:"countRemark"`
			ConfirmedCount int    `json:"confirmedCount"`
			SuspectedCount int    `json:"suspectedCount"`
			CuredCount     int    `json:"curedCount"`
			DeadCount      int    `json:"deadCount"`
			Virus          string `json:"virus"`
			Remark1        string `json:"remark1"`
			Remark2        string `json:"remark2"`
			Remark3        string `json:"remark3"`
			Remark4        string `json:"remark4"`
			Remark5        string `json:"remark5"`
			GeneralRemark  string `json:"generalRemark"`
			AbroadRemark   string `json:"abroadRemark"`
		} `json:"desc"`
	} `json:"newslist"`
}

var Latest = make(map[string]LatestData)
var Original = make(map[string]interface{})

// 获取最新数据
func GetLatestData() LatestData {
	_, ok := Latest["latest"]
	if ok == false {
		RequestLatestData()
	}
	return Latest["latest"]
}

// 获取原版输出
func GetOriginalLatestData() interface{} {
	_, ok := Original["latest"]
	if ok == false {
		RequestLatestData()
	}
	return Original["latest"]
}

func GetTxApiData() LatestData {
	_, ok := Latest["txApi"]

	if ok == false {
		RequestTxApiData()
	}
	log.Println(Latest["txApi"])
	return Latest["txApi"]
}

func GetOriginalTxApiData() interface{} {
	_, ok := Original["txApi"]
	if ok == false {
		RequestTxApiData()
	}
	return Original["txApi"]
}

// 请求最新数据
func RequestLatestData() error {
	url := fmt.Sprintf(conf.Conf.String("api::ncovapi"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	var resp NcovApiRes
	json.Unmarshal(ret, &resp)
	if err != nil {
		return err
	}
	if resp.Success == false {
		return fmt.Errorf("接口请求失败")
	}
	var latestData LatestData
	//make()
	latestData.Count.Confirmed = resp.Results[0].ConfirmedCount
	latestData.Count.Cure = resp.Results[0].CuredCount
	latestData.Count.Suspected = resp.Results[0].SuspectedCount
	latestData.Count.Death = resp.Results[0].DeadCount
	latestData.Info.DataSource = resp.Results[0].GeneralRemark
	latestData.Info.IncubationPeriod = resp.Results[0].Remark2
	latestData.Info.VulnerableGroup = resp.Results[0].Remark1
	latestData.Info.InfectionSource = resp.Results[0].InfectSource
	latestData.Info.Name = resp.Results[0].Virus
	latestData.UpdateTime = resp.Results[0].UpdateTime
	Latest["latest"] = latestData
	Original["latest"] = resp

	return nil
}

func RequestTxApiData() error {
	url := fmt.Sprintf(conf.Conf.String("api::txApi"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	var resp TxApiRes
	json.Unmarshal(ret, &resp)
	if err != nil {
		log.Println(err)
		return err
	}
	if resp.Code != 200 {
		return nil
	}
	var latestData LatestData
	//make()
	latestData.Count.Confirmed = resp.Newslist[0].Desc.ConfirmedCount
	latestData.Count.Cure = resp.Newslist[0].Desc.CuredCount
	latestData.Count.Suspected = resp.Newslist[0].Desc.SuspectedCount
	latestData.Count.Death = resp.Newslist[0].Desc.DeadCount
	latestData.Info.DataSource = resp.Newslist[0].Desc.GeneralRemark
	latestData.Info.IncubationPeriod = resp.Newslist[0].Desc.Remark2
	latestData.Info.VulnerableGroup = resp.Newslist[0].Desc.Remark1
	latestData.Info.InfectionSource = resp.Newslist[0].Desc.InfectSource
	latestData.Info.Name = resp.Newslist[0].Desc.Virus
	latestData.UpdateTime = resp.Newslist[0].Desc.ModifyTime

	Latest["txApi"] = latestData
	Original["txApi"] = resp

	return nil
}

//刷新缓存
func CrontabFunc(d time.Duration, hander func() error) {
	for {
		log.Println("crontab func runing")
		hander()
		time.Sleep(d)
	}
}

func Crond() {
	go CrontabFunc(time.Second * 30 ,RequestLatestData)
	go CrontabFunc(time.Second * 30 ,RequestTxApiData)
	go CrontabFunc(time.Second * 300 ,RequestTogetherData)
}
