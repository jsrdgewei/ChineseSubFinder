package subhd

import (
	"fmt"
	commonValue "github.com/allanpk716/ChineseSubFinder/internal/common"
	"github.com/allanpk716/ChineseSubFinder/internal/logic/series_helper"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/settings"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/something_static"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/unit_test_helper"
	"path/filepath"
	"testing"
)

// 无需关注这个测试用例，这个方案暂时弃用
func TestSupplier_GetSubListFromFile(t *testing.T) {

	//movie1 := "X:\\电影\\The Devil All the Time (2020)\\The Devil All the Time (2020) WEBDL-1080p.mkv"
	//movie1 := "X:\\电影\\Luca (2021)\\Luca (2021) WEBDL-1080p.mkv"
	//movie1 := "X:\\电影\\The Boss Baby Family Business (2021)\\The Boss Baby Family Business (2021) WEBDL-1080p.mkv"
	//movie1 := "X:\\电影\\Oslo (2021)\\Oslo (2021) WEBDL-1080p.mkv"
	//movie1 := "X:\\电影\\Spiral From the Book of Saw (2021)\\Spiral From the Book of Saw (2021) WEBDL-1080p.mkv"
	//movie1 := "X:\\电影\\消失爱人 (2016)\\消失爱人 (2016) 720p AAC.rmvb"
	//movie1 := "X:\\电影\\机动战士Z高达：星之继承者 (2005)\\机动战士Z高达：星之继承者 (2005) 1080p TrueHD.mkv"
	getCode()
	rootDir := unit_test_helper.GetTestDataResourceRootPath([]string{"sub_spplier"}, 5, true)
	movie1 := filepath.Join(rootDir, "zimuku", "movies", "消失爱人 (2016)", "消失爱人 (2016) 720p AAC.rmvb")

	subhd := NewSupplier(*settings.NewSettings())
	outList, err := subhd.getSubListFromFile4Movie(movie1)
	if err != nil {
		t.Error(err)
	}
	println(outList)

	if len(outList) == 0 {
		println("now sub found")
	}

	for i, sublist := range outList {
		println(i, sublist.Name, sublist.Ext, sublist.Language.String(), sublist.Score, sublist.FileUrl, len(sublist.Data))
	}

	alive, _ := subhd.CheckAlive()
	if alive == false {
		t.Fatal("CheckAlive == false")
	}
}

// 无需关注这个测试用例，这个方案暂时弃用
func TestSupplier_GetSubListFromFile4Series(t *testing.T) {

	//ser := "X:\\连续剧\\The Bad Batch"	// tt12708542
	//ser := "X:\\连续剧\\瑞克和莫蒂 (2013)" //
	//ser := "X:\\连续剧\\杀死伊芙 (2018)"	// tt7016936
	//ser := "X:\\连续剧\\Money.Heist"
	//ser := "X:\\连续剧\\黑钱胜地 (2017)"
	getCode()
	rootDir := unit_test_helper.GetTestDataResourceRootPath([]string{"sub_spplier"}, 5, true)
	ser := filepath.Join(rootDir, "zimuku", "series", "黄石 (2018)")
	// 读取本地的视频和字幕信息
	seriesInfo, err := series_helper.ReadSeriesInfoFromDir(ser, nil, false)
	if err != nil {
		t.Fatal(err)
	}
	s := NewSupplier(*settings.NewSettings())
	outList, err := s.GetSubListFromFile4Series(seriesInfo)
	if err != nil {
		t.Fatal(err)
	}
	println(outList)
	for i, sublist := range outList {
		println(i, sublist.Name, sublist.Ext, sublist.Language.String(), sublist.Score, len(sublist.Data))
	}
}

// 无需关注这个测试用例，这个方案暂时弃用
func TestSupplier_getSubListFromKeyword4Movie(t *testing.T) {

	//imdbID := "tt12708542" // 星球大战：残次品
	//imdbID := "tt7016936" // 杀死伊芙
	imdbID := "tt2990738" // 恐怖直播
	//imdbID := "tt3032476" 	// 风骚律师
	//imdbID := "tt6468322" 	// 纸钞屋
	//imdbID := "tt15299712" // 云南虫谷
	//imdbID := "tt3626476" // Vacation Friends (2021)
	getCode()
	subhd := NewSupplier(*settings.NewSettings())
	subInfos, err := subhd.getSubListFromKeyword4Movie(imdbID)
	if err != nil {
		t.Fatal(err)
	}
	for i, sublist := range subInfos {
		println(i, sublist.Name, sublist.Ext, sublist.Language.String(), sublist.Score, len(sublist.Data))
	}
}

func getCode() {
	updateTimeString, code, err := something_static.GetCodeFromWeb()
	if err != nil {
		commonValue.SubhdCode = ""
	} else {
		commonValue.SubhdCode = code
	}
	fmt.Println("UpdateTime", updateTimeString)
}
