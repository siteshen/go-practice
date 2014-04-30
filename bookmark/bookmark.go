package bookmark

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo/bson"

	"github.com/PuerkitoBio/purell"
	"github.com/gorilla/mux"
	"upper.io/db"
	_ "upper.io/db/mongo"
)

var sess db.Database

func Router() (r *mux.Router) {
	r = mux.NewRouter()

	return
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Md5sum(in string) string {
	hash := md5.New()
	io.WriteString(hash, in)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

type Content struct {
	Id      bson.ObjectId `bson:"_id"`
	Url     string
	Md5     string
	Header  http.Header
	Content string
}

func Save(u string) (content *Content, err error) {
	url, err := purell.NormalizeURLString(u, purell.FlagsSafe|purell.FlagAddTrailingSlash)
	md5sum := Md5sum(url)

	coll, err := sess.Collection("urls")
	// checkError(err)

	err = coll.Find(db.Cond{"md5": md5sum}).One(&content)
	if err == db.ErrNoMoreRows {
		// not found, create one
		content = &Content{Url: url, Md5: md5sum}
	} else {
		checkError(err)
	}

	log.Printf("fetching %s...", url)
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	content.Header = resp.Header
	content.Content = string(bytes)

	log.Printf("saving %s...", url)
	if content.Id == "" {
		content.Id = bson.NewObjectIdWithTime(time.Now())
	}
	_, err = coll.Append(content)

	return
}

var settings = struct {
	db.Settings
	Driver string
}{}

func init() {}

func Main() {
	flag.StringVar(&settings.Driver, "driver", "mongo", "connection driver")
	flag.StringVar(&settings.Host, "host", "", "database host")
	flag.IntVar(&settings.Port, "port", 0, "database host")
	flag.StringVar(&settings.Database, "database", "dbapp", "database name")
	flag.StringVar(&settings.User, "username", "", "database username")
	flag.StringVar(&settings.Password, "password", "", "database password")
	flag.StringVar(&settings.Charset, "charset", "", "database charset")

	sess, _ = db.Open(settings.Driver, settings.Settings)

	flag.Parse()

	_, err := Save("http://www.google.com")
	checkError(err)
}
