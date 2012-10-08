package session

import(
    "net/http"
    "crypto/rand"
    "time"
    "sync"
    "net/url"
    "math/big"
    "errors"
    // "fmt"
)

// cookie name to save sessionid
var CookieName = "GOSESSION"

// rate to start GC
var gcRate int64

// session struct
type Session struct{
    // sessionid
    id string

    // sessionValues which is map format
    values map[string]string
    
    // session create time, and used by GC
    time time.Time
}

// global session
var Sessions map[string]*Session
var rwm sync.RWMutex

func init() {
    Sessions = make(map[string]*Session)
    gcRate = 10
    // TODO : update gcRate from config
}

func SessionStart(r *http.Request, w http.ResponseWriter) (*Session, error){
    // check session-cookie
    cookie, err := r.Cookie(CookieName)
    if err != nil {
        if err == http.ErrNoCookie {
            return newSession(w)
        }
        return nil, err
    }

    sessionid := cookie.Value
    if session, ok := Sessions[sessionid]; ok == true {
        return session, nil
    }
    return newSession(w)
}

func newSession(w http.ResponseWriter) (*Session, error) {
    id, err := newSessionId()
    if err != nil {
        return nil, err
    }
    id = url.QueryEscape(id)
    Sessions[id] = &Session{id, map[string]string{}, time.Now()}
    
    // save to cookie
    c := &http.Cookie{Name: CookieName, Value: id, Path: "/"}
    http.SetCookie(w, c)

    // start Gc
    r, err := rand.Int(rand.Reader, big.NewInt(gcRate))
    if err == nil && r.Cmp(big.NewInt(1)) == 0{
        Gc()
    }

    return Sessions[id], nil
}

func (this *Session)Get (k string) string {
    if v, ok := this.values[k]; ok == true {
        return v
    }
    return ""
}

func (this *Session)Set (k, v string) {
    this.values[k] = v
}

func newSessionId() (string, error){
    // 32byte id
    k := make([]byte, 32)
    if _, err := rand.Read(k); err != nil {
        return "", errors.New("GetSessionId error")
    }
	return string(k), nil
}

func Gc() {
    for id, session := range Sessions {
        // 60 * 60 * 1
        if time.Now().Sub(session.time).Seconds() > 3600 {
            rwm.Lock()
            Sessions[id] = nil
            rwm.Unlock()
        }
    }
}