package client

import (
	"fmt"
	. "github.com/eradszewski/dos-template-minon/cmd/domain"
	loggerI "github.com/eradszewski/dos-template-minon/cmd/logger"
	"github.com/fatih/color"
	"github.com/guptarohit/asciigraph"
	"github.com/rodaine/table"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var logger = loggerI.NewLogger()
var tracingList []TracingTransport

type TracingTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

func newTransport() *TracingTransport {

	tr := &TracingTransport{
		dialer: &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	}
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	t.Proxy = http.ProxyFromEnvironment
	t.Dial = tr.dial
	t.TLSHandshakeTimeout = 10 * time.Second
	tr.rtp = t
	return tr
}

func (tr *TracingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

func (tr *TracingTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

func (tr *TracingTransport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

func (tr *TracingTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

func (tr *TracingTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}

// OLD
type requestMetric struct {
	RequestCount          int `json:"requestCount"`
	AverageTimePerRequest int `json:"averageTimePerRequest"`
}

func request(config *Config) {

	var url = config.Client.Dos.Url
	var show bool
	show = false
	//
	//flag.BoolVar(&show, "show", false, "Display the response content")
	//flag.Parse()
	//
	//url := flag.Args()[0]
	//fmt.Println("URL:", url)

	tp := newTransport()
	client := &http.Client{Transport: tp}
	//config.Client.Dos.RequestCount = 5
	logger.Info(config.Client.Dos.RequestCount)
	for config.Client.Dos.RequestCount > 0 {
		if config.Client.Dos.Method == "GET" {
			resp, err := client.Get(url)
			if err != nil {
				log.Fatalf("get error: %s: %s", err, url)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(resp.Body)
			output := ioutil.Discard
			if show {
				output = os.Stdout
			}
			io.Copy(output, resp.Body)
		}
		if config.Client.Dos.Method == "POST" {
			resp, err := client.Get(url)
			if err != nil {
				log.Fatalf("get error: %s: %s", err, url)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(resp.Body)
			output := ioutil.Discard
			if show {
				output = os.Stdout
			}
			io.Copy(output, resp.Body)
		}
		logger.PrintTracingTransport(tp.Duration(), tp.ConnDuration(), tp.ReqDuration())
		config.Client.Dos.RequestCount--
		tracingList = append(tracingList, *tp)

	}
	calcTracingResult(tracingList)
	//log.Println("Duration:", tp.Duration())
	//log.Println("Request duration:", tp.ReqDuration())
	//log.Println("Connection duration:", tp.ConnDuration())

	//
	//_, err := http.Get(url)
	//if err != nil {
	//	fmt.Print(err.Error())
	//	os.Exit(1)
	//}

	//responseData, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//logger.PrettyHttpRequestClient(response)
}

func calcTracingResult(tracingList []TracingTransport) {
	//data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	data := []float64{}
	lenTraceList := int64(len(tracingList))
	for _, e := range tracingList {
		//e.Duration().Minutes()
		data = append(data, float64(e.Duration().Milliseconds()))
		//data = append(data, e.Duration().Seconds())
	}

	//data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println("Duration for all Requests :", lenTraceList, " in Milliseconds")
	fmt.Println(graph)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Duration Sum (MillSec)", "Trace Counts", "Duration Avarage(MillSec)")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var durationSum int64
	for _, trace := range tracingList {
		durationSum = trace.Duration().Milliseconds() + durationSum
	}

	tbl.AddRow(durationSum, lenTraceList, durationSum/lenTraceList)

	tbl.Print()

}

func Run(config *Config) {

	logger.Info("Client Config enabled: ", config.Client)
	request(config)

}
