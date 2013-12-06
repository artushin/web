package oauth2

import (
	"code.google.com/p/goauth2/oauth"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var GoogleOauthCfg = &oauth.Config{
	ClientId:     "151342988894-4kfbh1b7lf81pabbkmtgkcla2n6hgtr2.apps.googleusercontent.com",
	ClientSecret: "-dD6oR2i1ku-scSMYa3S1kij",
	RedirectURL:  "http://www.ribrak.com:8080/oauth2callback",
	Scope:        "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/drive",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
}

const profileInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"

type Profile struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Link       string `json:"link"`
	Picture    string `json:"picture"`
	Gender     string `json:"gender"`
	Locale     string `json:"local"`
}

func Transport(code string) (t *oauth.Transport, err error) {
	// set up the transport
	t = &oauth.Transport{Config: GoogleOauthCfg}

	// Exchange the received code for a token
	tok, _ := t.Exchange(code)

	{
		tokenCache := oauth.CacheFile("./request.token")

		err = tokenCache.PutToken(tok)
		if err != nil {
			return
		}
	}

	// Skip TLS Verify
	t.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return t, err
}

func GetAccount(t *oauth.Transport) (profile *Profile, err error) {

	req, err := t.Client().Get(profileInfoURL)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(req.Body)
	err = json.Unmarshal([]byte(body), &profile)
	return
}
