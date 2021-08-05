package bmak_tools

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Bmak struct {
	UserAgent  string
	ScreenInfo struct {
		AvailWidth  int `json:"availWidth"`
		AvailHeight int `json:"availHeight"`
		Width       int `json:"width"`
		Height      int `json:"height"`
		OuterWidth  int `json:"outerWidth"`
		InnerWidth  int `json:"innerWidth"`
		InnerHeight int `json:"innerHeight"`
	}
	Navigator struct {
		ProductSub    int    `json:"productSub"`
		Language      string `json:"language"`
		Product       string `json:"product"`
		PluginsLength int    `json:"pluginsLength"`
	}
	CookieValue string
	Site        string
	FormInfo    string
	Fpcf        struct {
		Fpvalstr        string `json:"fpValstr"`
		Fpvalcalculated bool   `json:"fpValCalculated"`
		Rval            string `json:"rVal"`
		Rcfp            string `json:"rCFP"`
		Cache           struct {
		} `json:"cache"`
		Td      int      `json:"td"`
		Plugins []string `json:"PLUGINS"`
	}
	Ver               float64       `json:"ver"`
	KeCntLmt          int           `json:"ke_cnt_lmt"`
	MmeCntLmt         int           `json:"mme_cnt_lmt"`
	MduceCntLmt       int           `json:"mduce_cnt_lmt"`
	PmeCntLmt         int           `json:"pme_cnt_lmt"`
	PduceCntLmt       int           `json:"pduce_cnt_lmt"`
	TmeCntLmt         int           `json:"tme_cnt_lmt"`
	TduceCntLmt       int           `json:"tduce_cnt_lmt"`
	DoeCntLmt         int           `json:"doe_cnt_lmt"`
	DmeCntLmt         int           `json:"dme_cnt_lmt"`
	VcCntLmt          int           `json:"vc_cnt_lmt"`
	DoaThrottle       int           `json:"doa_throttle"`
	DmaThrottle       int           `json:"dma_throttle"`
	SessionID         string        `json:"session_id"`
	JsPost            bool          `json:"js_post"`
	Loc               string        `json:"loc"`
	CfURL             string        `json:"cf_url"`
	ParamsURL         string        `json:"params_url"`
	Auth              string        `json:"auth"`
	APIPublicKey      string        `json:"api_public_key"`
	AjLmtDoact        int           `json:"aj_lmt_doact"`
	AjLmtDmact        int           `json:"aj_lmt_dmact"`
	AjLmtTact         int           `json:"aj_lmt_tact"`
	CeJsPost          int           `json:"ce_js_post"`
	InitTime          int           `json:"init_time"`
	Informinfo        string        `json:"informinfo"`
	Prevfid           int           `json:"prevfid"`
	Fidcnt            int           `json:"fidcnt"`
	SensorData        string        `json:"sensor_data"`
	Ins               string        `json:"ins"`
	Cns               string        `json:"cns"`
	Engetloc          int           `json:"enGetLoc"`
	Enreaddocurl      int           `json:"enReadDocUrl"`
	Disfpcalontimeout int           `json:"disFpCalOnTimeout"`
	Xagg              int           `json:"xagg"`
	Pen               int           `json:"pen"`
	Brow              string        `json:"brow"`
	Browver           string        `json:"browver"`
	Psub              string        `json:"psub"`
	Lang              string        `json:"lang"`
	Prod              string        `json:"prod"`
	Plen              int           `json:"plen"`
	DoadmaEn          int           `json:"doadma_en"`
	Sdfn              []interface{} `json:"sdfn"`
	D2                int           `json:"d2"`
	D3                int           `json:"d3"`
	Thr               int           `json:"thr"`
	Cs                string        `json:"cs"`
	Hn                string        `json:"hn"`
	Z1                int           `json:"z1"`
	O9                int           `json:"o9"`
	Vc                string        `json:"vc"`
	Y1                int           `json:"y1"`
	Ta                int           `json:"ta"`
	Tst               int           `json:"tst"`
	TTst              int64         `json:"t_tst"`
	Ckie              string        `json:"ckie"`
	NCk               string        `json:"n_ck"`
	Ckurl             int           `json:"ckurl"`
	Bm                bool          `json:"bm"`
	Mr                string        `json:"mr"`
	Altfonts          bool          `json:"altFonts"`
	Rst               bool          `json:"rst"`
	Runfonts          bool          `json:"runFonts"`
	Fsp               bool          `json:"fsp"`
	Firstload         bool          `json:"firstLoad"`
	Pstate            bool          `json:"pstate"`
	MnMcLmt           int           `json:"mn_mc_lmt"`
	MnState           int           `json:"mn_state"`
	MnMcIndx          int           `json:"mn_mc_indx"`
	MnSen             int           `json:"mn_sen"`
	MnTout            int           `json:"mn_tout"`
	MnStout           int           `json:"mn_stout"`
	MnCt              int           `json:"mn_ct"`
	MnCc              string        `json:"mn_cc"`
	MnCd              int           `json:"mn_cd"`
	Mnrts             int
	MnLc              []string          `json:"mn_lc"`
	MnLd              []int             `json:"mn_ld"`
	MnLcl             int               `json:"mn_lcl"`
	MnAl              []string          `json:"mn_al"`
	MnIl              []int             `json:"mn_il"`
	MnTcl             []int             `json:"mn_tcl"`
	MnR               map[string]string `json:"mn_r"`
	MnRt              int               `json:"mn_rt"`
	MnWt              int               `json:"mn_wt"`
	MnAbck            string            `json:"mn_abck"`
	MnPsn             string            `json:"mn_psn"`
	MnTs              string            `json:"mn_ts"`
	MnLg              []string          `json:"mn_lg"`
	Loap              int               `json:"loap"`
	Dcs               int               `json:"dcs"`
	Listfunctions     struct {
	} `json:"listFunctions"`
	StartTs     int64  `json:"start_ts"`
	Kact        string `json:"kact"`
	KeCnt       int    `json:"ke_cnt"`
	KeVel       int    `json:"ke_vel"`
	Mact        string `json:"mact"`
	MmeCnt      int    `json:"mme_cnt"`
	MduceCnt    int    `json:"mduce_cnt"`
	MeVel       int    `json:"me_vel"`
	Pact        string `json:"pact"`
	PmeCnt      int    `json:"pme_cnt"`
	PduceCnt    int    `json:"pduce_cnt"`
	PeVel       int    `json:"pe_vel"`
	Tact        string `json:"tact"`
	TmeCnt      int    `json:"tme_cnt"`
	TduceCnt    int    `json:"tduce_cnt"`
	TeVel       int    `json:"te_vel"`
	Doact       string `json:"doact"`
	DoeCnt      int    `json:"doe_cnt"`
	DoeVel      int    `json:"doe_vel"`
	Dmact       string `json:"dmact"`
	DmeCnt      int    `json:"dme_cnt"`
	DmeVel      int    `json:"dme_vel"`
	Vcact       string `json:"vcact"`
	VcCnt       int    `json:"vc_cnt"`
	AjIndx      int    `json:"aj_indx"`
	AjSs        int    `json:"aj_ss"`
	AjType      int    `json:"aj_type"`
	AjIndxDoact int    `json:"aj_indx_doact"`
	AjIndxDmact int    `json:"aj_indx_dmact"`
	AjIndxTact  int    `json:"aj_indx_tact"`
	MeCnt       int    `json:"me_cnt"`
	PeCnt       int    `json:"pe_cnt"`
	TeCnt       int    `json:"te_cnt"`
	NavPerm     string `json:"nav_perm"`
	Brv         int    `json:"brv"`
	Hbcalc      bool   `json:"hbCalc"`
	Fmh         string `json:"fmh"`
	Fmz         int    `json:"fmz"`
	SSH         string `json:"ssh"`
	Wv          string `json:"wv"`
	Wr          string `json:"wr"`
	Weh         string `json:"weh"`
	Wl          int    `json:"wl"`
	Wen         int    `json:"wen"`
	Den         int    `json:"den"`
}

//GetCfDate, its basically an almost mirror function of akamai's get_cf_date() func
func (bm *Bmak) GetCfDate() int64 {
	return time.Now().UTC().UnixNano() / 1e6
}

//Jrs, returns 2 float64 variables, it does some weird shit, i think it has to do with time parsing not sure
func (bm *Bmak) Jrs(t int64) (float64, float64) {
	var a float64
	var e string
	var n int
	var m bool
	var o []int64
	a = math.Floor(1e5*rand.Float64() + 1e4)
	e = fmt.Sprintf("%f", float64(t)*a)
	n = 0
	m = len(e) >= 18
	for len(o) < 6 {
		intN, _ := strconv.ParseInt(e[n:n+2], 10, 32)
		o = append(o, intN)
		if m {
			n += 3
		} else {
			n += 2
		}
	}
	return a, bm.Cal_dis(o)
}

//Mn_s, completely diff than the akamai's version but what it does is it basically hashes the param and returns the hash chunk
func (bm *Bmak) Mn_s(t string) []int {
	bytesdt := []byte(t)
	hasher := sha256.New()
	hasher.Write(bytesdt)
	return convertBytetoInt(hasher.Sum(nil))
}

func (bm *Bmak) Bdm(t []int, a int) int {
	e := 0
	for n := 0; n < len(t); n++ {
		e = e<<8 | t[n]>>0
		e %= a
	}
	return e
}

func (bm *Bmak) Mn_pr() string {
	return fmt.Sprintf("%s;%s;%s;%s;", strings.Join(bm.MnAl, ","), strings.Join(convertIntToString(bm.MnTcl), ","), strings.Join(convertIntToString(bm.MnIl), ","), strings.Join(bm.MnLg, ","))
}

//Cal_dis does some weird math things that i dont rly understand but i was lucky enough to just basically copy and paste it
func (bm *Bmak) Cal_dis(t []int64) float64 {
	var a = t[0] - t[1]
	var e = t[2] - t[3]
	var n = t[4] - t[5]
	var o = math.Sqrt(float64(a*a + e*e + n*n))
	return math.Floor(o)
}

//Pi just takes a float and returns the int part as an int type var
func (bm *Bmak) Pi(t float64) int {
	return int(t)
}

//Ab does some weird ascii code shit, dont rly understand it but it returns an int, i think it just sums up all the ascii
//codes that are less than 128
func (bm *Bmak) Ab(t *string) int {
	if t == nil {
		return -1
	}
	var a = 0
	for e := 0; e < len(*t); e++ {
		n := []rune(*t)[e]
		if n < 128 {
			a += int(n)
		}
	}
	return a
}

//Od, does some weird ascii character shit that returns a result, its huge twice inside the generateSensor (bpd) function
func (bm *Bmak) Od(a, t string) string {
	var e []string
	n := len(t)
	if n > 0 {
		for o, v := range a {
			r := string(v)
			i := t[o%n]

			vInt := int(v)
			m := bm.Rir(vInt, 47, 57, int(i))
			if m != vInt {
				r = string(rune(m))
			}
			e = append(e, r)
		}
		if len(e) > 0 {
			return strings.Join(e, "")
		}
	}
	return a
}

//Rir takes in 4 diff params, if the a param is between 47 and 57 it does some calculations, at the end of the day it returns
//always an int
func (bm *Bmak) Rir(a, b, c, d int) int {
	if a > 47 && a <= 57 {
		a += d % (c - b)
		if a > c {
			a = a - c + b
		}
	}
	return a
}

//Cc returns a function that returns an int64, weird right? It gets even weirder, that func's return value is determined by
//the t parameter of the parent funcion
func (bm *Bmak) Cc(t int) func(t, a int) int64 {
	var a = t % 4
	if a == 2 {
		a = 3
	}
	var e = 42 + a
	n := func(t, a int) int64 {
		return 0
	}
	if e == 42 {
		n = func(t, a int) int64 {
			return int64(t * a)
		}
	} else if e == 43 {
		n = func(t, a int) int64 {
			return int64(t + a)
		}
	} else {
		n = func(t, a int) int64 {
			return int64(t - a)
		}
	}

	return n
}

func (bm *Bmak) getBrowser() {
	if bm.Navigator.ProductSub != 0 {
		bm.Psub = strconv.Itoa(bm.Navigator.ProductSub)
	}

	if bm.Navigator.Language != "" {
		bm.Lang = bm.Navigator.Language
	}

	if bm.Navigator.Product != "" {
		bm.Prod = bm.Navigator.Product
	}

	if bm.Navigator.PluginsLength != 0 {
		bm.Plen = bm.Navigator.PluginsLength
	}
}

func (bm *Bmak) bc() {
	t := 1
	a := 1
	n := 0
	e := 0
	o := 1
	m := 1
	r := 1
	i := 0
	c := 0
	if strings.Contains(bm.UserAgent, "Chrome") {
		c = 1
	}
	b := 1
	d := 0
	s := 1
	k := 0
	if bm.ScreenInfo.InnerWidth != 0 {
		k = 1
	}
	l := 0
	if bm.ScreenInfo.OuterWidth != 0 {
		l = 1
	}

	bm.Xagg = t + (a << 1) + (e << 2) + (n << 3) + (o << 4) + (m << 5) + (r << 6) + (i << 7) + (k << 8) + (l << 9) + (c << 10) + (b << 11) + (d << 12) + (s << 13)
}

func (bm *Bmak) bmisc() {
	bm.Pen = 0
	bm.Wen = 0
	bm.Den = 0
}

func (bm *Bmak) bd() string {
	a := 0
	e := 0
	n := 1
	o := 0
	m := 1
	r := 0
	if strings.Contains(strings.ToLower(bm.UserAgent), "opera") {
		r = 1
	}
	i := 1
	c := 0
	b := 0
	d := 0
	bm.D2 = bm.Pi(float64(bm.Z1 / 23))
	s := 0
	k := 0
	if strings.Contains(bm.UserAgent, "Chrome") {
		s = 1
		k = 1
	}
	l := 0
	u := 1
	return ",cpen:" + fmt.Sprint(a) + ",i1:" + fmt.Sprint(e) + ",dm:" + fmt.Sprint(n) + ",cwen:" + fmt.Sprint(o) + ",non:" + fmt.Sprint(m) + ",opc:" + fmt.Sprint(r) + ",fc:" + fmt.Sprint(i) + ",sc:" + fmt.Sprint(c) + ",wrc:" + fmt.Sprint(b) + ",isc:" + fmt.Sprint(d) + ",vib:" + fmt.Sprint(s) + ",bat:" + fmt.Sprint(k) + ",x11:" + fmt.Sprint(l) + ",x12:" + fmt.Sprint(u)
}

func (bm *Bmak) Gd() string {
	t := bm.UserAgent
	a := bm.Ab(&t)
	e := bm.StartTs / 2
	n := -1
	o := -1
	m := -1
	r := -1
	i := -1
	c := -1
	b := -1

	if bm.ScreenInfo.AvailWidth != 0 {
		n = bm.ScreenInfo.AvailWidth
	}

	if bm.ScreenInfo.AvailHeight != 0 {
		o = bm.ScreenInfo.AvailHeight
	}

	if bm.ScreenInfo.Width != 0 {
		m = bm.ScreenInfo.Width
	}

	if bm.ScreenInfo.Height != 0 {
		r = bm.ScreenInfo.Height
	}

	if bm.ScreenInfo.InnerHeight != 0 {
		i = bm.ScreenInfo.InnerHeight
	}

	if bm.ScreenInfo.InnerWidth != 0 {
		c = bm.ScreenInfo.InnerWidth
	}

	if bm.ScreenInfo.OuterWidth != 0 {
		b = bm.ScreenInfo.OuterWidth
	}

	bm.Z1 = bm.Pi(float64(bm.StartTs / int64((bm.Y1 * bm.Y1))))
	_rand := rand.New(rand.NewSource(99))
	d := _rand.Float64()
	s := bm.Pi(1e3 * d / 2)
	k := fmt.Sprint(d)
	k = k[:11] + strconv.Itoa(s)
	bm.Brv = 0
	bm.getBrowser()
	bm.bc()
	bm.bmisc()

	return t + ",uaend," + fmt.Sprint(bm.Xagg) + "," + fmt.Sprint(bm.Psub) + "," + bm.Lang + "," + bm.Prod + "," + fmt.Sprint(bm.Plen) + "," + fmt.Sprint(bm.Pen) + "," + fmt.Sprint(bm.Wen) + "," + fmt.Sprint(bm.Den) + "," + fmt.Sprint(bm.Z1) + "," + fmt.Sprint(bm.D3) + "," + fmt.Sprint(n) + "," + fmt.Sprint(o) + "," + fmt.Sprint(m) + "," + fmt.Sprint(r) + "," + fmt.Sprint(c) + "," + fmt.Sprint(i) + "," + fmt.Sprint(b) + "," + bm.bd() + "," + fmt.Sprint(a) + "," + k + "," + fmt.Sprint(e) + "," + fmt.Sprint(bm.Brv) + ",loc:" + bm.Loc
}
