package parser

import (
	"regexp"
	"strconv"

	"github.com/Felyne/crawler/distributed/config"
	"github.com/Felyne/crawler/distributed/engine"
	"github.com/Felyne/crawler/distributed/model"
)

//预先编译
var (
	//nameRe       = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	ageRe        = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
	hokouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xinzuoRe     = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
	guessRe      = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
	idUrlRe      = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func parseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	//profile.Name = extractString(contents, nameRe)
	profile.Gender = extractString(contents, genderRe)
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

//序列化传输过去，反序列化后用于NewProfileParser()
func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(username string) *ProfileParser {
	return &ProfileParser{userName: username}
}
