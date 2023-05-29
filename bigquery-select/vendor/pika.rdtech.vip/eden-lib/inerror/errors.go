package inerror

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

// Service Service
type Service string

const (
	// AppleTree banking
	AppleTree Service = "APPLETREE"
	// Aristotle 結算
	Aristotle Service = "ARISTOTLE"
	// Raphael 參數
	Raphael Service = "RAPHAEL"
	// Sage 收單
	Sage Service = "SAGE"
	// Lucifer gcc
	Lucifer Service = "LUCIFER"
	// Archimedes 統計
	Archimedes Service = "ARCHIMEDES"
	// Ianus Ianus
	Ianus Service = "Ianus"
	// Dove 我是信鴿我驕傲
	Dove Service = "DOVE"
	// Dolly 桃莉羊
	Dolly Service = "DOLLY"
	// Snow 追號機
	Snow Service = "SNOW"
	// Umai 遊戲庫
	Umai Service = "UMAI"
	// LifeBook 操作記錄
	LifeBook Service = "LIFEBOOK"
	// Shit 全部都在這
	Shit Service = "SHIT"
	// God 賽果來源
	God Service = "GOD"
	// Isaiah 壓力測試
	Isaiah Service = "ISAIAH"
	// Copycat 外部抄單api
	Copycat Service = "COPYCAT"
	// Robot 內部抄單機器人
	Robot Service = "ROBOT"
	// Hades 稽核機
	Hades Service = "HADES"
	// Najash 外部賽果機
	Najash Service = "NAJASH"
	// Iris 權限系統
	Iris Service = "IRIS"
	//Monitor 監控系統
	Monitor Service = "MONITOR"
	// CashRobot 現金抄單機
	CashRobot Service = "CASH_ROBOT"
	// Ladon 長龍統計
	Ladon Service = "LADON"
	// UserRobot 會員抄寫機
	UserRobot Service = "USERROBOT"
	// MeltonCalf 派彩監控阻擋系統
	MeltonCalf Service = "MELTONCALF"
	// Artemis 客端最近下注API
	Artemis Service = "ARTEMIS"
	// Trident 緊急停止出款相關
	Trident Service = "TRIDENT"
	// ParamsMachine 參數機
	ParamsMachine Service = "ParamsMachine"
	// Gardener apple-tree cash, user 抄單
	Gardener Service = "GARDENER"
	// Sales 推薦系統
	Sales Service = "SALES"
	// Zion Tag系統
	Zion Service = "ZION"
	// Medusa 廣宣系統
	Medusa Service = "MEDUSA"
	// Walle 廣宣系統
	Walle Service = "WALL-E"
	// Ahab 假平台商
	Ahab Service = "AHAB"
	// Coeus XBB用IP驗證服務
	Coeus Service = "COEUS"
	// 抄現金紀錄 包網
	CpCashTo Service = "CP-CASH-TO"
	// 賽果中心
	Scan Service = "SCAN"
	// Beetles 聊天室後台系統
	Beetles Service = "BEETLES"
	// Apis 聊天室
	Apis Service = "APIS"
	// Cicada 聊天室佇列
	Cicada Service = "CICADA"
	// Phoenix Phoenix
	Phoenix Service = "phoenix"
	// Sparta 唯一額度
	Sparta Service = "Sparta"
	// tang-proxy
	TangProxy Service = "tang-proxy"
	// tang-third
	TangThird Service = "tang-third"
	// TangDragon 包網金流中介處理
	TangDragon Service = "tang-dragon"
	// TangTiger 包網金流交易紀錄抄單機"
	TangTiger Service = "tang-tiger"
	// TangAsgard 包網操作記錄
	TangAsgard Service = "tang-asgard"
	// Accountant 包網注單查詢、退傭活動
	Accountant Service = "accountant"
	// Student 抄單機
	Student Service = "student"
	// Iron 包網廣宣 活動
	Iron Service = "iron"
	// Ives 包網權限系統
	Ives Service = "ives"
	// 不知道什麼系統
	Soda Service = "SODA"
	// Report 包網報表
	Report Service = "report"
	// Barachiel 包網用ip驗證服務
	Barachiel Service = "BARACHIEL"
	// Rd5Router 包網廳主端proxy
	Rd5Router Service = "rd5-router"
	// Dobby 包網批次人工優惠處理
	Dobby Service = "DOBBY"
	// Money XBB報表
	Money Service = "MONEY"
	// HandyMan 包網文章、排序系統
	HandyMan Service = "HANDYMAN"
	// Solo Proxy機
	Solo Service = "SOLO"
	// Nasa 火箭
	Nasa Service = "NASA"
	// Gabriel graphql系統
	Gabriel Service = "gabriel"

	// Nike open-match client
	Nike Service = "NIKE"
	// Hermes open-match director
	Hermes Service = "HERMES"
	// Themis open-match mmf
	Themis Service = "THEMIS"

	Aurora Service = "AURORA"

	// Wolf 電子
	Wolf Service = "WOLF"

	// Valhalla xbb操作記錄
	Valhalla Service = "VALHALLA"
	// Mars mattermost robot
	Mars Service = "MARS"

	// 對戰奪寶
	OdinOp Service = "ODIN-OP"

	// 百家
	Bacc Service = "BACC"

	// 對戰百人
	OdinHp Service = "ODIN-HP"

	// api站上雲轉單
	Divert Service = "DIVERT"

	// 直播彩票
	Voice Service = "VOICE"

	// 越南奪寶Bot proxy用
	Calbee Service = "CALBEE"

	// 模擬球機
	Shakespeare Service = "SHAKESPEARE"

	// Michael 換洗分
	Michael Service = "MICHAEL"

	// BB購彩車
	Meiji Service = "MEIJI"

	// 火箭風控
	Astronaut2 Service = "Astronaut2"

	// 虛擬體育賽事
	Suntory Service = "SUNTORY"

	// 試玩檢查服務
	Phantasos Service = "Phantasos"

	// 機器人氣氛產生服務
	Lazy Service = "LAZY"

	// 限制下注服務
	Prison Service = "PRISON"

	// 標籤系統
	Metatron Service = "METATRON"

	// 新型態遊戲風控
	Luna Service = "LUNA"

	// Unknown Unknown
	Unknown Service = "UNKNOWN"
)

// Code
const (
	// xbb 專案 請用 125 開頭
	AppleTreeCode     = 12501
	AristotleCode     = 12502
	RaphaelCode       = 12503
	SageCode          = 12504
	LuciferCode       = 12505
	ArchimedesCode    = 12506
	IanusCode         = 12507
	DoveCode          = 12508
	DollyCode         = 12509 //已作廢
	SnowCode          = 12510
	UmaiCode          = 12511
	LifeBookCode      = 12512
	ShitCode          = 12513
	GodCode           = 12514
	ParamsMachineCode = 12515
	// satan 12516
	HadesCode       = 12517
	CopycatCode     = 12518
	IrisCode        = 12519
	NajashCode      = 12520
	MonitorCode     = 12521
	CashRobotCode   = 12522
	LadonCode       = 12523
	UserRobotCode   = 12524
	MeltonCalfCode  = 12525
	ArtemisCode     = 12526
	TridentCode     = 12527
	RobotCode       = 12528
	IsaiahCode      = 12529
	GardenerCode    = 12530
	SalesCode       = 12531
	ZionCode        = 12532
	PhoenixCode     = 12533
	MedusaCode      = 12534
	WalleCode       = 12535
	AhabCode        = 12536
	ScanCode        = 12537
	SpartaCode      = 12538
	CoeusCode       = 12539
	MoneyCode       = 12540
	SoloCode        = 12541
	ValhallaCode    = 12542
	NasaCode        = 12543
	NikeCode        = 12544
	HermesCode      = 12545
	ThemisCode      = 12546
	MarsCode        = 12547
	WolfCode        = 12548
	GabrielCode     = 12549
	OdinOpCode      = 12550
	BaccCode        = 12551
	OdinHpCode      = 12552
	DivertCode      = 12553
	VoiceCode       = 12554
	CalbeeCode      = 12555
	ShakespeareCode = 12556
	MichaelCode     = 12557
	MeijiCode       = 12558
	Astronaut2Code  = 12559
	SuntoryCode     = 12560
	PhantasosCode   = 12561
	LazyCode        = 12562
	PrisonCode      = 12563
	AuroraCode      = 12564
	MetatronCode    = 12565
	LunaCode        = 12566

	// tang 專案 請用 126 開頭
	CpCashToCode   = 12601
	TangProxyCode  = 12602
	TangThirdCode  = 12603
	TangDragonCode = 12604
	TangTigerCode  = 12605
	TangAsgardCode = 12606
	AccountantCode = 12607
	StudentCode    = 12608
	IronCode       = 12609
	IvesCode       = 12610
	SodaCode       = 12611
	BarachielCode  = 12612
	ReportCode     = 12613
	Rd5RouterCode  = 12614
	DobbyCode      = 12615
	HandyManCode   = 12616
	// chatroom 專案 請用 127開頭
	ApisCode    = 12700
	BeetlesCode = 12701
	CicadaCode  = 12702
)

//var reError = regexp.MustCompile(`\[Error (\d+)\]\t([A-Z]{1,}):([\w\s]{1,}),Extra:(\S+.)\t(?:Origin:(\[\w+.*\].\w+.*|\w+.*))?`)
var defaultError = Error{
	Msg:       "inerror parse error",
	Code:      129999999,
	ExtraInfo: make(map[string]interface{}),
	Service:   Unknown,
	OriginErr: "",
}
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Errors Errors
type Errors struct {
	s Service
}

// Wrap Wrap
func (e *Errors) Wrap(code int, msg string, extrainfo map[string]interface{}, err error) (er Error) {
	return Wrap(code, msg, e.s, extrainfo, err)
}

// Parse Parse
func (e *Errors) Parse(s interface{}) (er Error) {
	return Parse(s)
}

// New New
func New(s Service) (er *Errors) {
	return &Errors{
		s: s,
	}
}

// Error Error
type Error struct {
	Msg       string                 `json:"msg"`
	Code      int                    `json:"code"`
	ExtraInfo map[string]interface{} `json:"extrainfo"`
	Time      int64                  `json:"time"`
	Service   Service                `json:"service"`
	OriginErr string                 `json:"origin_err"`
}

// SwitchSysCode SwitchSysCode
func SwitchSysCode(s Service) (code int) {
	switch s {
	case AppleTree:
		code = AppleTreeCode
	case Aristotle:
		code = AristotleCode
	case Raphael:
		code = RaphaelCode
	case Sage:
		code = SageCode
	case Lucifer:
		code = LuciferCode
	case Archimedes:
		code = ArchimedesCode
	case Ianus:
		code = IanusCode
	case Dove:
		code = DoveCode
	case Snow:
		code = SnowCode
	case Umai:
		code = UmaiCode
	case LifeBook:
		code = LifeBookCode
	case Shit:
		code = ShitCode
	case God:
		code = GodCode
	case Isaiah:
		code = IsaiahCode
	case Copycat:
		code = CopycatCode
	case Robot:
		code = RobotCode
	case Hades:
		code = HadesCode
	case Najash:
		code = NajashCode
	case Iris:
		code = IrisCode
	case Monitor:
		code = MonitorCode
	case CashRobot:
		code = CashRobotCode
	case Ladon:
		code = LadonCode
	case UserRobot:
		code = UserRobotCode
	case MeltonCalf:
		code = MeltonCalfCode
	case Artemis:
		code = ArtemisCode
	case Trident:
		code = TridentCode
	case ParamsMachine:
		code = ParamsMachineCode
	case Gardener:
		code = GardenerCode
	case Sales:
		code = SalesCode
	case Zion:
		code = ZionCode
	case Medusa:
		code = MedusaCode
	case Walle:
		code = WalleCode
	case Ahab:
		code = AhabCode
	case CpCashTo:
		code = CpCashToCode
	case Scan:
		code = ScanCode
	case Beetles:
		code = BeetlesCode
	case Apis:
		code = ApisCode
	case Cicada:
		code = CicadaCode
	case Sparta:
		code = SpartaCode
	case TangProxy:
		code = TangProxyCode
	case TangThird:
		code = TangThirdCode
	case TangDragon:
		code = TangDragonCode
	case TangTiger:
		code = TangTigerCode
	case TangAsgard:
		code = TangAsgardCode
	case Phoenix:
		code = PhoenixCode
	case Accountant:
		code = AccountantCode
	case Student:
		code = StudentCode
	case Iron:
		code = IronCode
	case Ives:
		code = IvesCode
	case Soda:
		code = SodaCode
	case Barachiel:
		code = BarachielCode
	case Coeus:
		code = CoeusCode
	case Report:
		code = ReportCode
	case Rd5Router:
		code = Rd5RouterCode
	case Dobby:
		code = DobbyCode
	case Money:
		code = MoneyCode
	case HandyMan:
		code = HandyManCode
	case Solo:
		code = SoloCode
	case Valhalla:
		code = ValhallaCode
	case Nasa:
		code = NasaCode
	case Nike:
		code = NikeCode
	case Hermes:
		code = HermesCode
	case Themis:
		code = ThemisCode
	case Mars:
		code = MarsCode
	case Wolf:
		code = WolfCode
	case Gabriel:
		code = GabrielCode
	case OdinOp:
		code = OdinOpCode
	case Bacc:
		code = BaccCode
	case OdinHp:
		code = OdinHpCode
	case Divert:
		code = DivertCode
	case Voice:
		code = VoiceCode
	case Calbee:
		code = CalbeeCode
	case Shakespeare:
		code = ShakespeareCode
	case Michael:
		code = MichaelCode
	case Meiji:
		code = MeijiCode
	case Astronaut2:
		code = Astronaut2Code
	case Suntory:
		code = SuntoryCode
	case Phantasos:
		code = PhantasosCode
	case Lazy:
		code = LazyCode
	case Prison:
		code = PrisonCode
	case Aurora:
		code = AuroraCode
	case Metatron:
		code = MetatronCode
	case Luna:
		code = LunaCode
	default:
		code = 12999
	}
	return
}

func (e Error) Error() (s string) {
	var b []byte
	b, _ = json.Marshal(e)
	return string(b)
}

// Wrap Wrap
func Wrap(code int, msg string, service Service, extrainfo map[string]interface{}, err error) (er Error) {
	c := SwitchSysCode(service)
	code = (c * 10000) + code
	var errs string
	if err != nil {
		errs = err.Error()
	} else {
		errs = ""
	}
	er = Error{
		Code:      code,
		Msg:       msg,
		ExtraInfo: extrainfo,
		Service:   service,
		Time:      time.Now().Unix(),
		OriginErr: errs,
	}
	return
}

// Parse Parse
func Parse(s interface{}) (e Error) {
	var ss string
	switch v := s.(type) {
	case string:
		ss = v
	case error:
		ss = grpc.ErrorDesc(v)
	default:
		ss = "can't parse"
	}
	e = Error{}
	err := json.Unmarshal([]byte(ss), &e)
	if err != nil {
		e = defaultError
	}
	return

}
