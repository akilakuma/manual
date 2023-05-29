```
   _____                     ______      __
  / ___/____  ____ _      __/ __/ /___ _/ /_____
  \__ \/ __ \/ __ \ | /| / / /_/ / __ `/ //_/ _ \
 ___/ / / / / /_/ / |/ |/ / __/ / /_/ / ,< /  __/
/____/_/ /_/\____/|__/|__/_/ /_/\__,_/_/|_|\___/

http://pika.rdtech.vip/genesis-lib/snowflake
```

## Core

* Reference:
  - [bwmarrin/snowflake](https://github.com/bwmarrin/snowflake/blob/master/snowflake.go)
  - [sony/sonyflake](https://github.com/sony/sonyflake)

* ID Bits

```
+--------------------------------------------------------------------------------+
| 39 Bit Time in units of 10 msec | 8 Bit Sequence number |  16 Bit Machine ID |
+--------------------------------------------------------------------------------+
```

## Method
* Options
  - SetStartTime(t time.Time) Option
  - SetMachineID(callback func() (uint16, error)) Option
* Node
  - NewNode(opts ...Option) *node
  - Generate() ID
* ID
  - Int64() int64
  - String() string
  - Bytes() []byte

## Usage
```go
package main

import "pika.rdtech.vip/genesis-lib/snowflake"

func main() {
    node := snowflake.NewNode()
    num1 := node.Generate()

    fmt.Println(num1.Int64())
    fmt.Println(num1.String())
}
```

## Options

### K8s deployment
```yaml
template:
  spec:
    containers:
      - name: {{ .Chart.Name }}
        image: "{{ .Values.global.hub }}/{{ .Values.image.repository }}:{{ .Values.image.tag }}"
        imagePullPolicy: {{ .Values.global.pullPolicy }}
        env:
          - name: MY_POD_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
```

### Example
```go
package main

import (
  "os"
  "net"
  "pika.rdtech.vip/genesis-lib/snowflake"
)

func main() {
    //set generator start time
    startTime := snowflake.SetStartTime(time.Date(1983, 1, 1, 0, 0, 0, 0, time.UTC))

    //set machine unique ID with callback func
    myIP := os.Getenv("MY_POD_IP")
    mcID := snowflake.SetMachineID(lower16BitPrivateIP(myIP))

    //apply setting while create new node
    node := snowflake.NewNode(startTime, mcID)
    num1 := node.Generate()

    fmt.Println(num1.String())
}

//Your machine ID (This case created by ip mask 2&3)
func lower16BitPrivateIP(ipString string) func() (uint16, error) {
	ip := net.ParseIP(ipString).To4()
	switch {
	case ip == nil:
		panic("parse ip err")
	case !isPrivateIPv4(ip):
		panic("ip is not private")
	}

	return func() (uint16, error) {
		return uint16(ip[2])<<8 + uint16(ip[3]), nil
	}
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

```
