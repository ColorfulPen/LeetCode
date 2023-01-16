// package main

// import (
// 	"fmt"
// 	"math/rand"

// 	"time"
// )

// func main() {
// 	maxNum := 100
// 	rand.Seed(time.Now().UnixNano())
// 	secretNumber := rand.Intn(maxNum)

// 	fmt.Println("Please input your guess")

// 	for {
// 		var guess int
// 		_,err:=fmt.Scanf("%d",&guess)
// 		if err!=nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		fmt.Println("You guess is", guess)
// 		if guess > secretNumber {
// 			fmt.Println("Your guess is bigger than the secret number. Please try again")
// 		} else if guess < secretNumber {
// 			fmt.Println("Your guess is smaller than the secret number. Please try again")
// 		} else {
// 			fmt.Println("Correct, you Legend!")
// 			break
// 		}
// 	}
// }

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

type YoudaoDictResponse struct {
	WebTrans struct {
		WebTranslation []struct {
			Same string `json:"@same,omitempty"`
			Key string `json:"key"`
			KeySpeech string `json:"key-speech"`
			Trans []struct {
				Summary struct {
					Line []string `json:"line"`
				} `json:"summary"`
				Value string `json:"value"`
				Support int `json:"support"`
				URL string `json:"url"`
				Cls struct {
					Cl []string `json:"cl"`
				} `json:"cls,omitempty"`
			} `json:"trans"`
		} `json:"web-translation"`
	} `json:"web_trans"`
	OxfordAdvanceHTML struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvanceHtml"`
	Ee struct {
		Source struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"source"`
		Word struct {
			Trs []struct {
				Pos string `json:"pos"`
				Tr []struct {
					Tran string `json:"tran"`
					SimilarWords []string `json:"similar-words,omitempty"`
				} `json:"tr"`
			} `json:"trs"`
			Speech string `json:"speech"`
			ReturnPhrase string `json:"return-phrase"`
		} `json:"word"`
	} `json:"ee"`
	BlngSentsPart struct {
		SentenceCount int `json:"sentence-count"`
		SentencePair []struct {
			Sentence string `json:"sentence"`
			SentenceEng string `json:"sentence-eng"`
			SentenceTranslation string `json:"sentence-translation"`
			SpeechSize string `json:"speech-size"`
			AlignedWords struct {
				Src struct {
					Chars []struct {
						S string `json:"@s"`
						E string `json:"@e"`
						Aligns struct {
							Sc []struct {
								ID string `json:"@id"`
							} `json:"sc"`
							Tc []struct {
								ID string `json:"@id"`
							} `json:"tc"`
						} `json:"aligns"`
						ID string `json:"@id"`
					} `json:"chars"`
				} `json:"src"`
				Tran struct {
					Chars []struct {
						S string `json:"@s"`
						E string `json:"@e"`
						Aligns struct {
							Sc []struct {
								ID string `json:"@id"`
							} `json:"sc"`
							Tc []struct {
								ID string `json:"@id"`
							} `json:"tc"`
						} `json:"aligns"`
						ID string `json:"@id"`
					} `json:"chars"`
				} `json:"tran"`
			} `json:"aligned-words"`
			Source string `json:"source"`
			URL string `json:"url"`
			SentenceSpeech string `json:"sentence-speech"`
		} `json:"sentence-pair"`
		More string `json:"more"`
		TrsClassify []struct {
			Proportion string `json:"proportion"`
			Tr string `json:"tr"`
		} `json:"trs-classify"`
	} `json:"blng_sents_part"`
	Individual struct {
		Trs []struct {
			Pos string `json:"pos"`
			Tran string `json:"tran"`
		} `json:"trs"`
		Idiomatic []struct {
			Colloc struct {
				En string `json:"en"`
				Zh string `json:"zh"`
			} `json:"colloc"`
		} `json:"idiomatic"`
		Level string `json:"level"`
		ExamInfo struct {
			Year int `json:"year"`
			QuestionTypeInfo []struct {
				Time int `json:"time"`
				Type string `json:"type"`
			} `json:"questionTypeInfo"`
			RecommendationRate int `json:"recommendationRate"`
			Frequency int `json:"frequency"`
		} `json:"examInfo"`
		ReturnPhrase string `json:"return-phrase"`
		PastExamSents []struct {
			En string `json:"en"`
			Source string `json:"source"`
			Zh string `json:"zh"`
		} `json:"pastExamSents"`
	} `json:"individual"`
	CollinsPrimary struct {
		Words struct {
			Indexforms []string `json:"indexforms"`
			Word string `json:"word"`
		} `json:"words"`
		Gramcat []struct {
			Audiourl string `json:"audiourl"`
			Pronunciation string `json:"pronunciation"`
			Senses []struct {
				Examples []struct {
					Sense struct {
						Lang string `json:"lang"`
						Word string `json:"word"`
					} `json:"sense"`
					Example string `json:"example"`
				} `json:"examples"`
				Labelgrammar string `json:"labelgrammar"`
				Definition string `json:"definition"`
				Lang string `json:"lang"`
				Word string `json:"word"`
			} `json:"senses"`
			Partofspeech string `json:"partofspeech"`
			Audio string `json:"audio"`
			Forms []interface{} `json:"forms"`
		} `json:"gramcat"`
	} `json:"collins_primary"`
	VideoSents struct {
		SentsData []struct {
			VideoCover string `json:"video_cover"`
			Contributor string `json:"contributor"`
			SubtitleSrt string `json:"subtitle_srt"`
			ID int `json:"id"`
			Video string `json:"video"`
		} `json:"sents_data"`
		WordInfo struct {
			ReturnPhrase string `json:"return-phrase"`
			Sense []string `json:"sense"`
		} `json:"word_info"`
	} `json:"video_sents"`
	AuthSentsPart struct {
		SentenceCount int `json:"sentence-count"`
		More string `json:"more"`
		Sent []struct {
			Score float64 `json:"score"`
			Speech string `json:"speech"`
			SpeechSize string `json:"speech-size"`
			Source string `json:"source"`
			URL string `json:"url"`
			Foreign string `json:"foreign"`
		} `json:"sent"`
	} `json:"auth_sents_part"`
	Simple struct {
		Query string `json:"query"`
		Word []struct {
			Usphone string `json:"usphone"`
			Ukphone string `json:"ukphone"`
			Ukspeech string `json:"ukspeech"`
			ReturnPhrase string `json:"return-phrase"`
			Usspeech string `json:"usspeech"`
		} `json:"word"`
	} `json:"simple"`
	ExpandEc struct {
		ReturnPhrase string `json:"return-phrase"`
		Source struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"source"`
		Word []struct {
			TransList []struct {
				Content struct {
					DetailPos string `json:"detailPos"`
					Sents []struct {
						SentOrig string `json:"sentOrig"`
						SourceType string `json:"sourceType"`
						SentSpeech string `json:"sentSpeech"`
						SentTrans string `json:"sentTrans"`
						Source string `json:"source"`
					} `json:"sents"`
				} `json:"content,omitempty"`
				Trans string `json:"trans"`
				Contenta struct {
					DetailPos string `json:"detailPos"`
					ExamType []struct {
						En string `json:"en"`
						Zh string `json:"zh"`
					} `json:"examType"`
					Sents []struct {
						SentOrig string `json:"sentOrig"`
						SourceType string `json:"sourceType"`
						SentSpeech string `json:"sentSpeech"`
						SentTrans string `json:"sentTrans"`
						Source string `json:"source"`
						Type string `json:"type,omitempty"`
					} `json:"sents"`
				} `json:"content,omitempty"`
				Contentb struct {
					DetailPos string `json:"detailPos"`
				} `json:"content,omitempty"`
				Contentc struct {
					DetailPos string `json:"detailPos"`
					ExamType []struct {
						Zh string `json:"zh"`
						En string `json:"en,omitempty"`
					} `json:"examType"`
					Sents []struct {
						SentOrig string `json:"sentOrig"`
						SourceType string `json:"sourceType"`
						SentSpeech string `json:"sentSpeech"`
						SentTrans string `json:"sentTrans"`
						Source string `json:"source"`
						Type string `json:"type,omitempty"`
					} `json:"sents"`
				} `json:"content,omitempty"`
			} `json:"transList"`
			Pos string `json:"pos"`
		} `json:"word"`
	} `json:"expand_ec"`
	Etym struct {
		Etyms struct {
			Zh []struct {
				Source string `json:"source"`
				Word string `json:"word"`
				Value string `json:"value"`
				URL string `json:"url"`
				Desc string `json:"desc"`
			} `json:"zh"`
		} `json:"etyms"`
		Word string `json:"word"`
	} `json:"etym"`
	Oxford struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxford"`
	Special struct {
		Summary struct {
			Sources struct {
				Source struct {
					Site string `json:"site"`
					URL string `json:"url"`
				} `json:"source"`
			} `json:"sources"`
			Text string `json:"text"`
		} `json:"summary"`
		CoAdd string `json:"co-add"`
		Total string `json:"total"`
		Entries []struct {
			Entry struct {
				Major string `json:"major"`
				Trs []struct {
					Tr struct {
						Nat string `json:"nat"`
						ChnSent string `json:"chnSent"`
						Cite string `json:"cite"`
						DocTitle string `json:"docTitle"`
						EngSent string `json:"engSent"`
						URL string `json:"url"`
					} `json:"tr"`
				} `json:"trs"`
				Num int `json:"num"`
			} `json:"entry"`
		} `json:"entries"`
	} `json:"special"`
	Senior struct {
		EncryptedData string `json:"encryptedData"`
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
	} `json:"senior"`
	Syno struct {
		Synos []struct {
			Pos string `json:"pos"`
			Ws []string `json:"ws"`
			Tran string `json:"tran"`
		} `json:"synos"`
		Word string `json:"word"`
	} `json:"syno"`
	Input string `json:"input"`
	Collins struct {
		CollinsEntries []struct {
			Entries struct {
				Entry []struct {
					TranEntry []struct {
						PosEntry struct {
							Pos string `json:"pos"`
							PosTips string `json:"pos_tips"`
						} `json:"pos_entry"`
						ExamSents struct {
							Sent []struct {
								ChnSent string `json:"chn_sent"`
								EngSent string `json:"eng_sent"`
							} `json:"sent"`
						} `json:"exam_sents"`
						Tran string `json:"tran"`
					} `json:"tran_entry"`
				} `json:"entry"`
			} `json:"entries"`
			Phonetic string `json:"phonetic"`
			BasicEntries struct {
				BasicEntry []struct {
					Cet string `json:"cet"`
					Headword string `json:"headword"`
				} `json:"basic_entry"`
			} `json:"basic_entries"`
			Headword string `json:"headword"`
			Star string `json:"star"`
		} `json:"collins_entries"`
	} `json:"collins"`
	Baike struct {
		Summarys []struct {
			Summary string `json:"summary"`
			Key string `json:"key"`
		} `json:"summarys"`
		Source struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"source"`
	} `json:"baike"`
	Meta struct {
		Input string `json:"input"`
		GuessLanguage string `json:"guessLanguage"`
		IsHasSimpleDict string `json:"isHasSimpleDict"`
		Le string `json:"le"`
		Lang string `json:"lang"`
		Dicts []string `json:"dicts"`
	} `json:"meta"`
	Webster struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"webster"`
	Le string `json:"le"`
	Lang string `json:"lang"`
	Ec struct {
		WebTrans []string `json:"web_trans"`
		Special []struct {
			Nat string `json:"nat"`
			Major string `json:"major"`
		} `json:"special"`
		ExamType []string `json:"exam_type"`
		Source struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"source"`
		Word struct {
			Usphone string `json:"usphone"`
			Ukphone string `json:"ukphone"`
			Ukspeech string `json:"ukspeech"`
			Trs []struct {
				Pos string `json:"pos"`
				Tran string `json:"tran"`
			} `json:"trs"`
			ReturnPhrase string `json:"return-phrase"`
			Usspeech string `json:"usspeech"`
		} `json:"word"`
	} `json:"ec"`
	OxfordAdvance struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvance"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse YoudaoDictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Ec.Word.Ukphone, "US:", dictResponse.Ec.Word.Usphone)
	// for _, item := range dictResponse.Dictionary.Explanations {
	// 	fmt.Println(item)
	// }
}

func queryYoudao(word string){
	client := &http.Client{}
	querys:="q="+word+"&le=en&t=2&client=web&sign=454545e97f9956f826c4573861e65207&keyfrom=webdict"
	var data = strings.NewReader(querys)
	req, err := http.NewRequest("POST", "https://dict.youdao.com/jsonapi_s?doctype=json&jsonversion=4", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID=857364429@10.112.57.88; OUTFOX_SEARCH_USER_ID_NCOO=1428061955.5303416; __yadk_uid=YMABMtPd2xQfCw6YfCRTyxoYt4nhGlog; advertiseCookie=advertiseValue; ___rl__test__cookies=1673790933150")
	req.Header.Set("Origin", "https://dict.youdao.com")
	req.Header.Set("Referer", "https://dict.youdao.com/result?word=china&lang=en")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	queryYoudao(word)
}