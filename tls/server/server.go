package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	//	"io"
	"log"
	"net"
	"os"
)

var encrypt *bool = flag.Bool("encrypt", false, "Encrypt underlying connection")
var address *string = flag.String("address", ":5858", "Address to listen on")
var serverName *string = flag.String("servername", "Harry", "Name of server")
var keyWriterFile *string = flag.String("keywriter", "", "Path to TLS key file. For use with Wireshark.")

var certificate string = `
-----BEGIN CERTIFICATE-----
MIIFlzCCA3+gAwIBAgIUI/8Te0MzVwMWBDTiKtwBrer6p+YwDQYJKoZIhvcNAQEL
BQAwWjELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDETMBEGA1UEAwwKVGVzdFNlcnZlcjAg
Fw0yMDAxMjYwMzE0NDRaGA8yMTIwMDEwMjAzMTQ0NFowWjELMAkGA1UEBhMCQVUx
EzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMg
UHR5IEx0ZDETMBEGA1UEAwwKVGVzdFNlcnZlcjCCAiIwDQYJKoZIhvcNAQEBBQAD
ggIPADCCAgoCggIBALrL/SjUrDsbttRAfgY79EykAbjfrCO2PU2JtSnYt3tReNNF
YbBK+WxOESF3HNDLNdsKBK8SepiFw3QXgksfnbIKllUcRJmZ1ZP2ljkO4NzC9GhI
GLRV2uXv+k5VC5Rdp0X1fqZj9Rxv0KUT8MRSSXLdzorLYRD4lHgPwtXtGRyM1nDE
rf5QaQW6reyWbAuystnOkd0QvF2+RYyIZZlUd6q23Rgy3HSJGd1HQQwtEPITTC/A
0UWptfpI0rXLijthqeX37ZDZnoFzITA1SGB7gSpN2sjprqO/imlMfR8Kt/IgR6og
Svi7LWd5M3HyFjCI3srW9Ae7fGMFrTRCQ4P/p6MgxKNAcGSCsDcvnsPodyQjWope
oFTnDIAXSNGvOGx7/ILlIIQWYdLhJiWtCo7vsZVzW56DcGwX0LcxHuvjoruEQvaW
HxUB6xjo0XhJUatcvVwRaonZUcQ00O1U4iXp427837DctWnSd2POFxwnm6mpZUOm
iRkRgbfgzr+qQblUqTA3wv2fcb0IRcMjvptRROxiFPNvRvX9gkg3WZgKCZUaqDQy
XQRqECDekRzUw4Bw+KCqedQC31Cbib6cLflsYD6VKejt6Av6UoVu8sxagHFKubOa
tWnxzhbxRrT/LvHIIjrxbOoEhAeyet/T06vMJSgh/lH2icsXOv06u5X6B7dRAgMB
AAGjUzBRMB0GA1UdDgQWBBRk+vCuGffDOhkW0D38F1pnETbbmzAfBgNVHSMEGDAW
gBRk+vCuGffDOhkW0D38F1pnETbbmzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
DQEBCwUAA4ICAQB+PZfeYu+VA/v83HiSGXrMIAILArY0v2YRfXp6sHt71FZLgbm2
SHEeWM/upqSwrdiqBBdnCVoEHWbzqCWkT1ZZXfTs1DYAKa2gD5XLLbBi+2757Gk1
ha+4s0LhZBOk7i58J9dGmHbDOakwSbwpwLDXZoomb6A92ze+Wop71Zws1mF23tnl
e3ARtDd7rG8kON6JPxy6mUqmhXBgzYyLnIc4lBURJ3nRa0YgOqcfX/VxXrjwB8av
iVzyHntWC9QB9DdC5d4Gip1rgf7JsBoTw5d0t8P9tKYaKdygRgdAysDbys8737NP
67yt9lyqKlULB3R/ipqlruR/Ilns8d5caPhqbj7kp5eFhziL7gxQI/mDoUwE2+c2
8W80LOy/kVM6JniMg0O2bU1ZiwunKYHe2Yt5sEVXGnCt0cgEfiDr6owzU3bZxFAX
QyZGxzHsoceOKwHThGNRk5pcpAPTdXsDa8HAmhDOWpG38fBdvb7sWOT/1sgjU2Bq
VRl1FNypcdgBTLbLVixYdvfPm/UPEr+Sda8tHw1T6T60k+J8jB22MoztrcU43z5A
khp0guwbwcMEoLnLpPkWL+14jzMRVvnWGppdyG9qLtVtMjS1R1/NO0hKhZtdh7RW
qb3f8jipi/0NywcpySr1lJpIzkkQJTmLNmXx/TcMDGdtmwe/VckQDn+pVw==
-----END CERTIFICATE-----
`

var key string = `
-----BEGIN RSA PRIVATE KEY-----
MIIJKQIBAAKCAgEAusv9KNSsOxu21EB+Bjv0TKQBuN+sI7Y9TYm1Kdi3e1F400Vh
sEr5bE4RIXcc0Ms12woErxJ6mIXDdBeCSx+dsgqWVRxEmZnVk/aWOQ7g3ML0aEgY
tFXa5e/6TlULlF2nRfV+pmP1HG/QpRPwxFJJct3OisthEPiUeA/C1e0ZHIzWcMSt
/lBpBbqt7JZsC7Ky2c6R3RC8Xb5FjIhlmVR3qrbdGDLcdIkZ3UdBDC0Q8hNML8DR
Ram1+kjStcuKO2Gp5fftkNmegXMhMDVIYHuBKk3ayOmuo7+KaUx9Hwq38iBHqiBK
+LstZ3kzcfIWMIjeytb0B7t8YwWtNEJDg/+noyDEo0BwZIKwNy+ew+h3JCNail6g
VOcMgBdI0a84bHv8guUghBZh0uEmJa0Kju+xlXNbnoNwbBfQtzEe6+Oiu4RC9pYf
FQHrGOjReElRq1y9XBFqidlRxDTQ7VTiJenjbvzfsNy1adJ3Y84XHCebqallQ6aJ
GRGBt+DOv6pBuVSpMDfC/Z9xvQhFwyO+m1FE7GIU829G9f2CSDdZmAoJlRqoNDJd
BGoQIN6RHNTDgHD4oKp51ALfUJuJvpwt+WxgPpUp6O3oC/pShW7yzFqAcUq5s5q1
afHOFvFGtP8u8cgiOvFs6gSEB7J639PTq8wlKCH+UfaJyxc6/Tq7lfoHt1ECAwEA
AQKCAgA4oPbSlgbQtIostpB+G1bolR/giA6LlQfopcLCOO0G+aADjJkc9N3As7oF
xzJ5KeHd0Z+x44w/CO5EF4xscd1AyGziyHsThct129+W1KfexUuLAbBbm1uNMb9U
V/v2sp2vW8lVcCAyysMaCH2JRtj1dcDT2uxGVNXNwBl0+TgPq7Km5TlKS95K2lk3
zDJdKYOG/FTEC/rPAIvdGsySO0U7/8vvNhwXSCkLKDrUEiV8/dSu9Wl8EQMWKVWX
v2jcr7LuAMeIjjC7VcfN8ZlLhqWVyZx2JEHJapvvaDzWWE2Dnb2cGcYgub3+DIRI
GvgXJqdEHv5ECdafvmzP1vbY4XxUNwR/qBei+cUwoN16d74AKVE7kvKuwnDy2IMO
ldZ08Jd3lW2nK36reFC+0Lp4U0sxtGovqyQ2plxyvqXWdHi5dbXbcIZiE1SRLlzd
9/dFbhrY/e8l0cKP2EKstNs4NC0LmxM5Bek9sRJDVww1dHhiX8H2TZobFFYkAU9l
wtp7wUFgSpkKBjirJ+8MlZRRgLEzZSSRjFyjF04jj4tLhCTAZXh6AquJeN9hL0ek
NFLs5SDsTLrhnFhaWHKihLnLz+9pxUKQ6lA/N0d5BrUGZU2VlUWY8cI5+cF66E4i
8lmB6Bs8HF6cG5TOjr+Mln579StqMikghhqJOHg37rRgrpmLAQKCAQEA6pBl52vc
inzKteWC0QylrQcaz6DPdE52g+FETn99Y86ryPr0sx1VMKhm0bGfDVk5pryzGSbY
NeCgMFLa18XvEs4X4MjsuMmtx1kDgr4H20E7pCT6SujzKlSHYINCMQN27NdEpWoM
PBpiutDE4Erx4pCdxvxlXLLJlhqJdRObG2RRb7cFceZey7JwZ9ra3Zg87SlEGnhw
QUpzNTiXp1IOUrU0cYuuS5DSSDHmNYjuZZvL/7KtyvZDcwliXee6drwz6SvZzOco
SW7ry1swxbGH3DN/mTgVXJkOFhq7nx+crtihe8W0E7MuaccK0M8Yo9iGlKQqPP/f
UZq4+C/LGQWBmQKCAQEAy94U4Y6EIy7oqSo9KJ/W3MUSAsehNsDbBCZdlqvo+TK7
/QAiXZMNMLho0TqeK5cEwMSYKSRk5XRGsH10INmUBuFY4R9I/dkMq+GK0vte/b6g
B7xdMJn3F8A2cMohhtAzVQ9uUIUYfmXP8Mzasn7CFTgK2JYgAyI3IAlRB/ntAZdQ
h+nktIuBRrNSrJ9ER1CvX/7tHa6gR3a5xf8MdAfsFJQEZx4uUykNxGL7gYBXFOz4
XPZ8CvEo+sXMWfKJq6ObfbR7tnR5a9eWXSlzwySjMn2RQBHskzRAwRQ5R13F83E7
JWvu8Q/NpEMfTm7e39eKPFb6BL963DvPRxPdzq7meQKCAQBChXUA7ov7Em0CrPYu
hyGtMmieHYL3/xCJUidnA6zx0zjQpKsk6NqyE/Ak4/Sxem5pJPa92VBT50JGshiy
PMYSVTRcYV8RANExycK/H1lnCtb5NCtvdyUPCi3iZxcsg0kE3f/v0WVq3ijFxlMv
MNHsaQr11bqBUYrt8NSuyUKhwA+AWS1IIgccSZyrN1v+oCAXOi7AOwvK2GxX2ZbY
suKw/gbDdNOXRpj8NHqPEChb1JVEDM5Q0wpJ441sCD3PUox+QhtgiuXX/YcNgu0Q
A8r93fT/5PHZ8uYVyrsO444x3+ncCjOJqrUs9m/QzAq23L8+BFieOAqDQBfY+uTk
UbVxAoIBAQCoPL4YoWaULkrcBzpvQvCqQYsqdhmpOJ/FHfAPvhBFTcPq9mhltkul
UBlXyLrsl/TZK6OyGBGXdUw8q1rhHQzWXLLfHNU8fxjA4yCQGdb7KYugtqZkzDoo
BHwoufXO7hPedxx/IEblUBm4yyUTNh3uKtBwifsi+uJo8qdHIM2giYFwl+kfwRxO
/v8T618KRyBi5NpCq3AjaWvHZEGfo0YEeV3kxvhNskxlK5YH+aRjZWdUOCiHUxqR
UBfiho9r8FkJ9J4/JIFrKH8ypFmeyaZPrWXnbKNBm1Zwv8LBDTalPmUj7Z1Cm9sF
WksEi+Qq7xFp2BD9SBMYmdbk/hUGUjlpAoIBAQCYKw/FrYWaT9ai+FVyMMSqHaYS
UaA4eHiQOaLHwMIzP5C51dotW3NzqosyrsZ4YX3kXjqCl5HF5kAGRpZQfX1QXOqh
Ysx0RPzmCEo8z+P6tzm6UfSlFRC31GrvKl7WikE/DWVubMG/UUg4FsN4G6CTjCz9
3/vSjE94onPnebJjdt8Fc59QTC9Jtpo/yOtntz/VFT1umvi/2PsMDsi8woCFamkD
ZQ/DDkTQxwWxT0R5FXvH1jvKBUJyrTyD/rJgah4GuC3Kl2SW2e/HqdIcM+/KjZ3H
EgcdPCN/uamIooJUSsBofgR3D+3JED7fpfrWMDfnN7MMnTTkz6NQtoVVRtz9
-----END RSA PRIVATE KEY-----
`

func main() {
	fmt.Println("Server")

	flag.Parse()

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		panic(err)
	}

	if *encrypt {

		cert, err := tls.X509KeyPair([]byte(certificate), []byte(key))
		if err != nil {
			panic(err)
		}

		keyWriter := os.Stdout

		if *keyWriterFile != "" {
			keyWriter, err = os.Create(*keyWriterFile)
			if err != nil {
				panic(err)
			}
		}

		config := &tls.Config{
			ServerName:   *serverName,
			Certificates: []tls.Certificate{cert},
			KeyLogWriter: keyWriter,
		}
		listener = tls.NewListener(listener, config)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection.", err)
			continue
		}
		defer conn.Close()

		n, err := conn.Write([]byte("Hello from Server"))
		if err != nil {
			log.Println("Error writing bytes to connection.", err)
			continue
		}

		fmt.Println("Wrote", n, "bytes")
	}
}
