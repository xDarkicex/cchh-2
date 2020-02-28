package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/url"
	"regexp"
	"time"

	"github.com/alecthomas/template"
	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
	e "github.com/scorredoira/email"
)

type Res struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Time    string      `json:"time"`
	Payload interface{} `json:"payload"`
}

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func PostEmail(c echo.Context) error {
	var res RecaptchaResponse
	//decoder := json.NewDecoder(c.Request().Body)
	//encoder := json.NewEncoder(c.Response())

	//err = decoder.Decode(&res)
	err := c.Bind(&res)
	if err != nil {
		c.Response().Header().Set("content-type", "application/json; charset=UTF-8")
		return c.JSON(http.StatusOK, &Res{
			Success: false,
			Message: "Malformed Request: " + errors.Cause(err).Error(),
			Time:    time.Now().Local().Format(time.Stamp),
			Payload: ""})
	}
	senderEmail := c.FormValue("contact_email")
	add := c.FormValue("contact_address")
	phone := c.FormValue("contact_phone")
	Sender := c.FormValue("contact_name")
	google_chaptcha := c.FormValue("g-recaptcha-response")

	google_struct := validate_google_rechaptcha(google_chaptcha, c)
	if !google_struct.Success {
		c.Response().Header().Set("content-type", "application/json; charset=UTF-8")
		err = c.JSON(http.StatusOK, &Res{
			Success: false,
			Message: "Malformed Request: " + errors.Cause(err).Error(),
			Time:    time.Now().Local().Format(time.Stamp),
			Payload: ""})
		err = c.Redirect(http.StatusFound, "/contact")
	}

	if !validate_email(senderEmail) {
		// Start of html email sending
		subject := fmt.Sprintf("New Contact Request from %s", Sender)
		t, err := template.ParseFiles("./app/views/emails/contact.gotmpl")
		if err != nil {
			c.Response().Header().Set("content-type", "application/json; charset=UTF-8")

			return c.JSON(http.StatusOK, &Res{
				Success: false,
				Message: "Malformed Request: " + errors.Cause(err).Error(),
				Time:    time.Now().Local().Format(time.Stamp),
				Payload: ""})
		}
		buf := new(bytes.Buffer)
		if err = t.Execute(buf, map[string]interface{}{
			"Phone":   phone,
			"Address": add,
			"Sender":  Sender,
			"Email":   senderEmail,
			"Content": c.FormValue("contact_body"),
		}); err != nil {
			c.Response().Header().Set("content-type", "application/json; charset=UTF-8")

			return c.JSON(http.StatusOK, &Res{
				Success: false,
				Message: "Malformed Request: " + errors.Cause(err).Error(),
				Time:    time.Now().Local().Format(time.Stamp),
				Payload: ""})
		}
		Body := buf.String()
		msg := Body
		m := e.NewHTMLMessage(subject, msg)
		m.Subject = subject
		m.BodyContentType = "text/html"
		m.From = mail.Address{Name: "Compassionate Care Mobile Health", Address: "admin@cchha.com"}
		m.To = []string{"xDarkicex@gmail.com"}
		auth := smtp.PlainAuth("", "admin@cchha.com", "Vh2@cchha#G0!", "smtp.gmail.com")
		SMTP := "smtp.gmail.com:587"
		if err := e.Send(SMTP, auth, m); err != nil {
			c.Response().Header().Set("content-type", "application/json; charset=UTF-8")
			return c.JSON(http.StatusOK, &Res{
				Success: false,
				Message: "Malformed Request: " + errors.Cause(err).Error(),
				Time:    time.Now().Local().Format(time.Stamp),
				Payload: ""})
		}
	}
	return c.NoContent(http.StatusNotModified)
}

func validate_email(email string) bool {
	regex, err := regexp.Compile(`\S+@\S+`)
	if err != nil {
		fmt.Println(err)
	}
	if !regex.MatchString(email) {
		return false
	}
	return true
}

func validate_google_rechaptcha(chaptcha string, c echo.Context) (r RecaptchaResponse) {
	var encoder = json.NewEncoder(c.Response())
	var google_check = url.Values{
		"secret":   {"6LchUqEUAAAAALM_u_okQofqiw7Htdcp96jJGn1p"},
		"response": {chaptcha},
	}
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", google_check)
	if err != nil {
		c.Response().Header().Set("content-type", "application/json; charset=UTF-8")
		err := c.JSON(http.StatusOK, &Res{
			Success: false,
			Message: "Malformed Request: " + errors.Cause(err).Error(),
			Time:    time.Now().Local().Format(time.Stamp),
			Payload: ""})
		if err != nil {
			fmt.Println(err)
		}
		return RecaptchaResponse{}
	}
	defer c.Request().Body.Close()
	google_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Response().Header().Set("content-type", "application/json; charset=UTF-8")
		_ = encoder.Encode(&Res{
			Success: false,
			Message: "Malformed Request: " + errors.Cause(err).Error(),
			Time:    time.Now().Local().Format(time.Stamp),
			Payload: ""})
		return RecaptchaResponse{}
	}
	err = json.Unmarshal(google_body, &r)
	if err != nil {
		c.Response().Header().Set("content-type", "application/json; charset=UTF-8")
		_ = encoder.Encode(&Res{
			Success: false,
			Message: "Malformed Request: " + errors.Cause(err).Error(),
			Time:    time.Now().Local().Format(time.Stamp),
			Payload: ""})
		return RecaptchaResponse{}
	}
	return r
}

func GenerateCookie(status string, success bool, c echo.Context) *http.Cookie {
	encoder := json.NewEncoder(c.Response())
	type data struct {
		Status  string
		Success bool
	}
	cookie_value := data{
		Status:  status,
		Success: success,
	}

	d, err := json.Marshal(cookie_value)
	if err != nil {
		_ = encoder.Encode(&Res{
			Success: false,
			Message: "Internal Server Error: " + errors.Cause(err).Error(),
			Time:    time.Now().Local().Format(time.Stamp),
			Payload: ""})
	}
	cookie := &http.Cookie{
		Name:     "contact-limit",
		Value:    base64.StdEncoding.EncodeToString(d),
		Path:     "http://www.compassionatecaremobileclinic.org/contact.html",
		Domain:   "compassionatecaremobileclinic.org",
		Expires:  time.Now().Add(time.Minute * 1),
		Secure:   false,
		HttpOnly: false,
	}
	return cookie
}
