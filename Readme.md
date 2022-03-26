ENV Example Client:

````
CLIENT_ENABLED=true;CLIENT_DOS_URL=https://google.de;CLIENT_DOS_METHOD=GET;CLIENT_DOS_REQUEST_COUNT=5
````


YAML Example Client:

````
client:
  enabled: true
  dos:
    url: "https://google.de"
    method: "GET"
    requestCount: 5
````


Example Result:

````
GoVersion: go1.18
GOOS: darwin
GOARCH: amd64
NumCPU: 10
GOPATH: .....
GOROOT: /usr/local/go
Compiler: gc
ENV: /Users/..../go
Now: Saturday, 26 Mar 2022INFO[03-26-03-2022 18:59:57.3378+01:00] Client Config enabled: {true {https://google.de GET 20}   }                                       
INFO[03-26-03-2022 18:59:57.5165+01:00] >------> ✔ Request statistic ( Duration: 98.866125ms ConnectionDuration : 0s RequestDuration : 98.866125ms) ✔ >------> 
INFO[03-26-03-2022 18:59:57.6043+01:00] >------> ✔ Request statistic ( Duration: 55.876792ms ConnectionDuration : 0s RequestDuration : 55.876792ms) ✔ >------> 
INFO[03-26-03-2022 18:59:57.6846+01:00] >------> ✔ Request statistic ( Duration: 54.211334ms ConnectionDuration : 0s RequestDuration : 54.211334ms) ✔ >------> 
INFO[03-26-03-2022 18:59:57.7753+01:00] >------> ✔ Request statistic ( Duration: 64.620084ms ConnectionDuration : 0s RequestDuration : 64.620084ms) ✔ >------> 
INFO[03-26-03-2022 18:59:57.866+01:00] >------> ✔ Request statistic ( Duration: 55.535792ms ConnectionDuration : 0s RequestDuration : 55.535792ms) ✔ >------> 
INFO[03-26-03-2022 18:59:57.943+01:00] >------> ✔ Request statistic ( Duration: 54.5605ms ConnectionDuration : 0s RequestDuration : 54.5605ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.0238+01:00] >------> ✔ Request statistic ( Duration: 54.356ms ConnectionDuration : 0s RequestDuration : 54.356ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.1123+01:00] >------> ✔ Request statistic ( Duration: 64.484708ms ConnectionDuration : 0s RequestDuration : 64.484708ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.1943+01:00] >------> ✔ Request statistic ( Duration: 53.412959ms ConnectionDuration : 0s RequestDuration : 53.412959ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.2792+01:00] >------> ✔ Request statistic ( Duration: 57.996916ms ConnectionDuration : 0s RequestDuration : 57.996916ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.374+01:00] >------> ✔ Request statistic ( Duration: 68.476167ms ConnectionDuration : 0s RequestDuration : 68.476167ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.4638+01:00] >------> ✔ Request statistic ( Duration: 62.537042ms ConnectionDuration : 0s RequestDuration : 62.537042ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.5464+01:00] >------> ✔ Request statistic ( Duration: 56.259917ms ConnectionDuration : 0s RequestDuration : 56.259917ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.6335+01:00] >------> ✔ Request statistic ( Duration: 58.7255ms ConnectionDuration : 0s RequestDuration : 58.7255ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.7147+01:00] >------> ✔ Request statistic ( Duration: 56.678916ms ConnectionDuration : 0s RequestDuration : 56.678916ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.8037+01:00] >------> ✔ Request statistic ( Duration: 63.224458ms ConnectionDuration : 0s RequestDuration : 63.224458ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.8786+01:00] >------> ✔ Request statistic ( Duration: 53.321167ms ConnectionDuration : 0s RequestDuration : 53.321167ms) ✔ >------> 
INFO[03-26-03-2022 18:59:58.9482+01:00] >------> ✔ Request statistic ( Duration: 47.76625ms ConnectionDuration : 0s RequestDuration : 47.76625ms) ✔ >------> 
INFO[03-26-03-2022 18:59:59.0545+01:00] >------> ✔ Request statistic ( Duration: 69.147167ms ConnectionDuration : 0s RequestDuration : 69.147167ms) ✔ >------> 
INFO[03-26-03-2022 18:59:59.146+01:00] >------> ✔ Request statistic ( Duration: 49.45525ms ConnectionDuration : 0s RequestDuration : 49.45525ms) ✔ >------> 
Duration for all Requests : 20  in Milliseconds
 98.00 ┼╮
 97.00 ┤│
 96.00 ┤│
 95.00 ┤│
 94.00 ┤│
 93.00 ┤│
 92.00 ┤│
 91.00 ┤│
 90.00 ┤│
 89.00 ┤│
 88.00 ┤│
 87.00 ┤│
 86.00 ┤│
 85.00 ┤│
 84.00 ┤│
 83.00 ┤│
 82.00 ┤│
 81.00 ┤│
 80.00 ┤│
 79.00 ┤│
 78.00 ┤│
 77.00 ┤│
 76.00 ┤│
 75.00 ┤│
 74.00 ┤│
 73.00 ┤│
 72.00 ┤│
 71.00 ┤│
 70.00 ┤│
 69.00 ┤│                ╭╮
 68.00 ┤│        ╭╮      ││
 67.00 ┤│        ││      ││
 66.00 ┤│        ││      ││
 65.00 ┤│        ││      ││
 64.00 ┤│ ╭╮  ╭╮ ││      ││
 63.00 ┤│ ││  ││ ││   ╭╮ ││
 62.00 ┤│ ││  ││ │╰╮  ││ ││
 61.00 ┤│ ││  ││ │ │  ││ ││
 60.00 ┤│ ││  ││ │ │  ││ ││
 59.00 ┤│ ││  ││ │ │  ││ ││
 58.00 ┤│ ││  ││ │ │╭╮││ ││
 57.00 ┤│ ││  ││╭╯ │││││ ││
 56.00 ┤│ ││  │││  ╰╯╰╯│ ││
 55.00 ┤╰╮│╰╮ │││      │ ││
 54.00 ┤ ╰╯ ╰─╯││      │ ││
 53.00 ┤       ╰╯      ╰╮││
 52.00 ┤                │││
 51.00 ┤                │││
 50.00 ┤                │││
 49.00 ┤                ││╰
 48.00 ┤                ││
 47.00 ┤                ╰╯
Duration Sum (MillSec)  Trace Counts  Duration Avarage(MillSec)  
1189                    20            59                         


````