package models

type PrefectureId int

// 都道府県
const (
	Hokkaido PrefectureId = iota + 1
	Aomori
	Iwate
	Miyagi
	Akita
	Yamagata
	Fukushima
	Ibaraki
	Tochigi
	Gunma
	Saitama
	Chiba
	Tokyo
	Kanagawa
	Niigata
	Toyama
	Ishikawa
	Fukui
	Yamanashi
	Nagano
	Gifu
	Shizuoka
	Aichi
	Mie
	Shiga
	Kyoto
	Osaka
	Hyogo
	Nara
	Wakayama
	Tottori
	Shimane
	Okayama
	Hiroshima
	Yamaguchi
	Tokushima
	Kagawa
	Ehime
	Kochi
	Fukuoka
	Saga
	Nagasaki
	Kumamoto
	Oita
	Miyazaki
	Kagoshima
	Okinawa
)

func (p PrefectureId) ToString() string {
	switch p {
	case Hokkaido:
		return "北海道"
	case Aomori:
		return "青森県"
	case Iwate:
		return "岩手県"
	case Miyagi:
		return "宮城県"
	case Akita:
		return "秋田県"
	case Yamagata:
		return "山形県"
	case Fukushima:
		return "福島県"
	case Ibaraki:
		return "茨城県"
	case Tochigi:
		return "栃木県"
	case Gunma:
		return "群馬県"
	case Saitama:
		return "埼玉県"
	case Chiba:
		return "千葉県"
	case Tokyo:
		return "東京都"
	case Kanagawa:
		return "神奈川県"
	case Niigata:
		return "新潟県"
	case Toyama:
		return "富山県"
	case Ishikawa:
		return "石川県"
	case Fukui:
		return "福井県"
	case Yamanashi:
		return "山梨県"
	case Nagano:
		return "長野県"
	case Gifu:
		return "岐阜県"
	case Shizuoka:
		return "静岡県"
	case Aichi:
		return "愛知県"
	case Mie:
		return "三重県"
	case Shiga:
		return "滋賀県"
	case Kyoto:
		return "京都府"
	case Osaka:
		return "大阪府"
	case Hyogo:
		return "兵庫県"
	case Nara:
		return "奈良県"
	case Wakayama:
		return "和歌山県"
	case Tottori:
		return "鳥取県"
	case Shimane:
		return "島根県"
	case Okayama:
		return "岡山県"
	case Hiroshima:
		return "広島県"
	case Yamaguchi:
		return "山口県"
	case Tokushima:
		return "徳島県"
	case Kagawa:
		return "香川県"
	case Ehime:
		return "愛媛県"
	case Kochi:
		return "高知県"
	case Fukuoka:
		return "福岡県"
	case Saga:
		return "佐賀県"
	case Nagasaki:
		return "長崎県"
	case Kumamoto:
		return "熊本県"
	case Oita:
		return "大分県"
	case Miyazaki:
		return "宮崎県"
	case Kagoshima:
		return "鹿児島県"
	case Okinawa:
		return "沖縄県"
	default:
		return "未定義"
	}
}
