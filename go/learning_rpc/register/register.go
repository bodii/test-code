package register

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

// LearningRegistery is a simple register conter, provide following functions.
type LearningRegistery struct {
	timeout time.Duration
	mu      sync.Mutex
	servers map[string]*ServerItem
}

// ServerItem ...
type ServerItem struct {
	Addr  string
	start time.Time
}

const (
	defaultPath    = "/_learning_rpc/register"
	defaultTimeout = time.Minute * 5
)

// New create a register instance with timeout setting
func New(timeout time.Duration) *LearningRegistery {
	return &LearningRegistery{
		servers: make(map[string]*ServerItem),
		timeout: timeout,
	}
}

// DefaultLearningRegister ...
var DefaultLearningRegister = New(defaultTimeout)

func (r *LearningRegistery) putServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	s := r.servers[addr]
	if s == nil {
		r.servers[addr] = &ServerItem{
			Addr:  addr,
			start: time.Now(),
		}
	} else {
		s.start = time.Now()
	}
}

func (r *LearningRegistery) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()

	var alive []string
	for addr, s := range r.servers {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, addr)
		} else {
			delete(r.servers, addr)
		}
	}
	sort.Strings(alive)
	return alive
}

// Runs at /_learning_rpc_/registry
func (r *LearningRegistery) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set("X-Learning-rpc-Servers", strings.Join(r.aliveServers(), ","))
	case "POST":
		addr := req.Header.Get("X-Learning-rpc-Server")
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(addr)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleHTTP registers an HTTP handler for LearningRegister messages on registeryPath
func (r *LearningRegistery) HandleHTTP(registeryPath string) {
	http.Handle(registeryPath, r)
	log.Println("rpc registry path:", registeryPath)
}

// HandleHTTP ...
func HandleHTTP() {
	DefaultLearningRegister.HandleHTTP(defaultPath)
}

// Heartbeat send a heartbeat message every once in a while
func Heartbeat(registery, addr string, duration time.Duration) {
	if duration == 0 {
		duration = defaultTimeout - time.Duration(1)*time.Minute
	}

	var err error
	err = sendHeartbeat(registery, addr)
	go func() {
		t := time.NewTicker(duration)
		for err == nil {
			<-t.C
			err = sendHeartbeat(registery, addr)
		}
	}()
}

func sendHeartbeat(registery, addr string) error {
	log.Println(addr, "send heart beat to registry", registery)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", registery, nil)
	req.Header.Set("X-Learning-rpc-Server", addr)
	if _, err := httpClient.Do(req); err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}
	return nil
}
