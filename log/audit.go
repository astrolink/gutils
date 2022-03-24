package log

import (
	gqueue "github.com/astrolink/gutils/queue"
	"log"
	"net/http"
)

type AuditItem struct {
	Cod              string `json:"cod"`
	Id               string `json:"id"`
	Token            string `json:"token"`
	Login            string `json:"login"`
	ServiceName      string `json:"serviceName"`
	File             string `json:"file"`
	Ip               string `json:"ip"`
	UserAgent        string `json:"userAgent"`
	Action           string `json:"action"`
	Referer          string `json:"referer"`
	Route            string `json:"route"`
	StateBefore      string `json:"stateBefore"`
	StateAfter       string `json:"stateAfter"`
	CustomParameters string `json:"customParameters"`
}

func NewAuditItem() *AuditItem {
	return &AuditItem{}
}

func (a *AuditItem) SetCod(cod string) *AuditItem {
	a.Cod = cod
	return a
}

func (a *AuditItem) SetId(id string) *AuditItem {
	a.Id = id
	return a
}

func (a *AuditItem) SetToken(token string) *AuditItem {
	a.Token = token
	return a
}

func (a *AuditItem) SetLogin(login string) *AuditItem {
	a.Login = login
	return a
}

type AuditLogger struct {
	Config  gqueue.Config
	Request http.Request
}

func NewAuditLogger(c gqueue.Config) *AuditLogger {
	return &AuditLogger{Config: c}
}

func (a *AuditLogger) SetRequest(r *http.Request) {
	a.Request = *r
}

func (a *AuditLogger) buildAuditItem() AuditItem {
	built := *NewAuditItem()
	built.File = getCallerFullInfo()
	built.ServiceName = getBinName()
	built.Ip = GetRequestRealIp(&a.Request)
	built.UserAgent = GetUserAgent(&a.Request)
	return built
}

func (a *AuditLogger) SendAuditLog(item AuditItem) error {
	var queue *gqueue.RabbitMQ
	var err error

	if queue, err = gqueue.NewRabbitMQ(a.Config); err != nil {
		log.Println(err)
		return err
	}

	defer queue.Close()

	auditItem := a.buildAuditItem()
	auditItem.Cod = item.Cod
	auditItem.Id = item.Id
	auditItem.Token = item.Token
	auditItem.Login = item.Login
	auditItem.Action = item.Action
	auditItem.Referer = item.Referer
	auditItem.Route = item.Route

	if item.Route == "" {
		item.Route = GetCurrentRoute(&a.Request)
	}

	auditItem.StateBefore = item.StateBefore
	auditItem.StateAfter = item.StateAfter
	auditItem.CustomParameters = item.CustomParameters

	if err = queue.Publish(auditItem); err != nil {
		log.Println(err)
	}

	return err
}
