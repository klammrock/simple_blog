package blog_service

import (
	"fmt"
	"net/http"
	"simple_blog/db"

	"github.com/sirupsen/logrus"
)

// [url]/api/v1/posts
// GET posts/id
// { id, html, }
//
// POST posts/id
// PUT posts/id
// DELETE posts/id

// comments
// GET for post, POST add to post

// categories

////
// html
// [url]/ ?page=
// [url]/about
// [url]/posts/[post_id]
// [url]/categories/[category_id]
// [url]/admin
// [url]/admin/posts
// [url]/admin/logs

type AdminConfig struct {
	Password string    `json:"password"`
	Jwt      jwtConfig `json:"jwt"`
}

type jwtConfig struct {
	Secret    string `json:"secret"`
	ExpiresIn int    `json:"expires_in"`
}

func Do_work() {
	fmt.Println("hello blog")
}

type BlogService struct {
	isDev       bool
	servicePort int
	webPath     string
	publicPath  string
	db          db.DBer
}

func New(isDev bool, servicePort int, webPath, publicPath string, db db.DBer) (*BlogService, error) {
	if err := checkConfig(servicePort); err != nil {
		return nil, err
	}
	return &BlogService{isDev, servicePort, webPath, publicPath, db}, nil
}

func checkConfig(servicePort int) error {
	if servicePort <= 0 {
		return fmt.Errorf("ServicePort <= 0")
	}

	return nil
}

func (service *BlogService) Start() error {
	// http.HandleFunc("/api/v1/events", logPanics(service.serveEvents))
	// http.HandleFunc("/api/v1/users/login", logPanics(service.serveLogin))
	// http.HandleFunc("/api/v1/logs", logPanics(service.serveLogs))

	// Mandatory root-based resources
	serveSingle("/", service.webPath+"index.html")
	// TODO: DEV/PROD in config
	// for debug
	// if service.isDev {
	// 	serveSingle("/test", service.webPath+"test/index.html")
	// }
	serveSingle("/sitemap.xml", service.webPath+"sitemap.xml")
	serveSingle("/favicon.ico", service.webPath+"favicon.ico")
	serveSingle("/robots.txt", service.webPath+"robots.txt")

	// Normal resources
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(service.publicPath))))

	// return http.ListenAndServeTLS(fmt.Sprintf(":%d", service.servicePort),
	// 	service.sslConfig.Crt, service.sslConfig.Key, nil)
	return http.ListenAndServe(fmt.Sprintf(":%d", service.servicePort), nil)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		logRequest(logWithTagSERVICE(), r, fmt.Sprintf("ServeFile: %s", filename))
		http.ServeFile(w, r, filename)
	})
}

func logRequest(l *logrus.Entry, r *http.Request, handler string) {
	// ip, _, err := net.SplitHostPort(r.RemoteAddr)
	// if err != nil {
	//   ip = fmt.Sprintf("%s parse error: %v", r.RemoteAddr, err)
	// }
	// // Accessed via non-anonymous proxy r.Header.Get("X-Forwarded-For")
	// // https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request-golang

	// // TODO: with fields?
	// //log.Printf("%s by %s %v from %s\n", message, r.Method, r.URL, ip)
	// logrus.Infof("%s by %s %v from %s", message, r.Method, r.URL, ip)

	getRequestLogFields(l, r).Debug("Handler: ", handler)
}

func getRequestLogFields(l *logrus.Entry, r *http.Request) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"RemoteAddr": r.RemoteAddr,
		"Method":     r.Method,
		"URL":        fmt.Sprintf("%v", r.URL),
	})
}

func logWithTagSERVICE() *logrus.Entry {
	return logrus.WithField("Tag", "SERVICE")
}
