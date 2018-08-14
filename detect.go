package botdetect

import (
	"net/http"
	"regexp"
)

// BotDetect need for checking request is from bot or not
type BotDetect struct {
	userAgent string
	rules     []string
}

// NewBotDetect function create instance of BotDetect
func NewBotDetect(r *http.Request, rules []string) *BotDetect {
	if rules == nil {
		rules = makeRules()
	}

	d := &BotDetect{
		userAgent: r.UserAgent(),
		rules:     rules,
	}

	return d
}

func makeRules() []string {
	return []string{
		"Googlebot|Mediapartners-Google|AdsBot-Google|Google Keyword Suggestion|Googlebot-Mobile|AdsBot-Google-Mobile|APIs-Google",
		"YandexAccessibilityBot|YandexBot|YandexMobileBot|YandexImages|yandex.com/bots",
		"YahooSeeker/M1A1-R2D2|facebookexternalhit|Facebot|bingbot|ia_archiver|AhrefsBot|Ezooms|GSLFbot|WBSearchBot|Twitterbot|TweetmemeBot|Twikle|PaperLiBot|Wotbox|UnwindFetchor|Exabot|MJ12bot|TurnitinBot|Pingdom",
	}
}

// IsBot method detect is bot or not
func (d *BotDetect) IsBot() bool {
	for _, ruleValue := range d.rules {
		if d.match(ruleValue) {
			return true
		}
	}

	return false
}

func (d *BotDetect) match(ruleValue string) bool {
	ruleValue = `(?is)` + ruleValue
	var re *regexp.Regexp
	re = regexp.MustCompile(ruleValue)
	return re.MatchString(d.userAgent)
}
