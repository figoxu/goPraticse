package main
import (
	"bufio"
	"io"
	"bytes"
	"strings"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

func main(){
	pem:="-----BEGIN CERTIFICATE-----\nfriendlyName: Apple Development IOS Push Services: com.td.demo\nlocalKeyId: 2d86d047afccd69901d4d41db2c015a167ff15cb\n\nMIIFejCCBGKgAwIBAgIIJjF8domciiQwDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQKDApBcHBsZSBJbmMuMSwwKgYDVQQLDCNBcHBsZSBXb3Js\nZHdpZGUgRGV2ZWxvcGVyIFJlbGF0aW9uczFEMEIGA1UEAww7QXBwbGUgV29ybGR3\naWRlIERldmVsb3BlciBSZWxhdGlvbnMgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkw\nHhcNMTUwOTEwMDU1MjA2WhcNMTYwOTA5MDU1MjA2WjB6MRswGQYKCZImiZPyLGQB\nAQwLY29tLnRkLmRlbW8xOTA3BgNVBAMMMEFwcGxlIERldmVsb3BtZW50IElPUyBQ\ndXNoIFNlcnZpY2VzOiBjb20udGQuZGVtbzETMBEGA1UECwwKRVRaWDhWNk44MzEL\nMAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCcHnXS\nHFMP2bqR7l+tKKdDF3K68rTCG8HMrHFHotiqb3dUj1b/lqnykKu2s5ywadQoLAmX\nfr6m/pXVhtIAxqJkxaa8Qm8ZbktWhDxoa2aGxlawg+Gys33PD6op3wj55nE8UQJI\nUzGOXbNcR9wWLYQzKGWmuFN11MZfII3LHIYXFUg8O5s2m0Evd4bQ9GjkyhYVU9pV\nJKnBysa6TkTOnNPZVbZ6iF7Cb/8hnOwWcy4HY498k2Hur7mRKH8rAl3uNMDxs5N6\nZQVkxnCwPHVXpEB/qQuCPiXGIV9vj5MDM6tKB4rBnmktvwXcCwxxxZtiEk3nlFAR\ndtC5FxGx6FD2hk9jAgMBAAGjggHlMIIB4TAdBgNVHQ4EFgQULYbQR6/M1pkB1NQd\nssAVoWf/FcswCQYDVR0TBAIwADAfBgNVHSMEGDAWgBSIJxcJqbYYYIvs67r2R1nF\nUlSjtzCCAQ8GA1UdIASCAQYwggECMIH/BgkqhkiG92NkBQEwgfEwgcMGCCsGAQUF\nBwICMIG2DIGzUmVsaWFuY2Ugb24gdGhpcyBjZXJ0aWZpY2F0ZSBieSBhbnkgcGFy\ndHkgYXNzdW1lcyBhY2NlcHRhbmNlIG9mIHRoZSB0aGVuIGFwcGxpY2FibGUgc3Rh\nbmRhcmQgdGVybXMgYW5kIGNvbmRpdGlvbnMgb2YgdXNlLCBjZXJ0aWZpY2F0ZSBw\nb2xpY3kgYW5kIGNlcnRpZmljYXRpb24gcHJhY3RpY2Ugc3RhdGVtZW50cy4wKQYI\nKwYBBQUHAgEWHWh0dHA6Ly93d3cuYXBwbGUuY29tL2FwcGxlY2EvME0GA1UdHwRG\nMEQwQqBAoD6GPGh0dHA6Ly9kZXZlbG9wZXIuYXBwbGUuY29tL2NlcnRpZmljYXRp\nb25hdXRob3JpdHkvd3dkcmNhLmNybDALBgNVHQ8EBAMCB4AwEwYDVR0lBAwwCgYI\nKwYBBQUHAwIwEAYKKoZIhvdjZAYDAQQCBQAwDQYJKoZIhvcNAQEFBQADggEBAJz/\nvCNpH7rfn3IeDc6JNHA3aLjvEf+E3Pm3jkVZSf6JPc0oUsySfmVN0Zs48JSXv+Gm\nFybpn6j3eFI6w3cJ0Gi4NX/71q7/WufaT7nctN6ebwjloIz8T7lYljz6HXZIMR3A\nM3Lv+oNAQoZFasGicjibDfpbjNYRi1YjFViHWQDya9oM9AOFqWO/dBrkfVIGxvz/\nx5dB5LHgeDWACGvVjfAMkvbAM3TOndv42fV/hsgodcJy4lw59/nQVD3SYJHaQG2Z\n11vAdbPI/ZkALB/sojb/ERWrCwMI4NuFVJHVk8yCVX3P12rTJbhOm974suR7GJg9\nJWsILPcgGVLddQ1vlxc=\n-----END CERTIFICATE-----\n-----BEGIN PRIVATE KEY-----\nfriendlyName: liweiqiang\nlocalKeyId: 2d86d047afccd69901d4d41db2c015a167ff15cb\n\nMIIEowIBAAKCAQEAnB510hxTD9m6ke5frSinQxdyuvK0whvBzKxxR6LYqm93VI9W\n/5ap8pCrtrOcsGnUKCwJl36+pv6V1YbSAMaiZMWmvEJvGW5LVoQ8aGtmhsZWsIPh\nsrN9zw+qKd8I+eZxPFECSFMxjl2zXEfcFi2EMyhlprhTddTGXyCNyxyGFxVIPDub\nNptBL3eG0PRo5MoWFVPaVSSpwcrGuk5EzpzT2VW2eohewm//IZzsFnMuB2OPfJNh\n7q+5kSh/KwJd7jTA8bOTemUFZMZwsDx1V6RAf6kLgj4lxiFfb4+TAzOrSgeKwZ5p\nLb8F3AsMccWbYhJN55RQEXbQuRcRsehQ9oZPYwIDAQABAoIBABSn1xDgnIDJXuz1\n7AS+DztKO+zONepEv+RzoF1JB3+tVc2DVZBlf8DPMVjv9LhGmEJkaTR5GYKAxpRE\nzT47Cbtph0D7TTXBKQieYFu0chE85rVeDZuWMfwWZ344uzhNLAg9855cG39pmFSK\nUW1Bwm7+VkyzsJ8zcredWAValccx3cCI7hcObS++axUWWdO0iSUg/L0tSg/O/eiq\nOnBCDx1tQR1+1YYZtVNwyVJzB+z+dvzGLtuT1KnoFc1JmqsOORFaXdi/x0MF9SCI\n+GP1UVNnriA4vc68dqqYdEY6haB3Qs2nfBrjp7oS04v5co5KfZnDfEG9/td3/dcJ\nr2uFUCkCgYEAybGsQ3GZQ+BuBnh3FaBDu4d+ycCqpbsre1+fFIIiYx5hKrF0zGU4\nFL7nTlGRrrz3zk2pDtFvAJ2Q6SRAeScc+7GS5aiHKjQyWzHj8OrutRGmFeeIuWl3\nSyScAqm0PZnz8GPD/Zf/y9ECpRn1X55CKCc5bazNSS4HLIbY6gGWQycCgYEAxidm\nMqieyF4iqSHqvhy+5WJmJeu34+5De+a9n2leBwz2ELViuWV6a30DAUHJAtFr1/9+\nx9upDVaIkSxq6wv5kGpkjE6V9zau8jK7kk2KboEzqGca7Malfb/rtuMsPCW7WUT0\nysIg0Fsz7s27Mmnh7IbmPFDQmgLnWZvazpQyR2UCgYEAkrRGPWRCe8mhndk/nR5O\nRff/M9aNCTFEJl/eAPfK9Veii4A6GXXCPezBqAjSs0vF32xDfIC/ga4aOkphDv5x\nKW33EE7tybffrM90IcdwS4oDyUj4/QwGdkSxqYowjOIPOhoG88z8hQ6JIvfnr14Y\ndsZZEjovs3vkQfHvkv4Ggt8CgYBvzGrsyhZMFlCX/HrlKHLMGOhpHMVz5EpO7bCu\n8FQMxlNOuggpcgfP/YhkfSlcXavrAkwVlumajOgggF8Snn7/7Acu4mfaQxARtm39\n7aHdFyh+Ky4VyT415MKpPSMIDIHXrAAxJbY92DdE1O55UyQUv0fHYBFR07GT7UtW\nomZ7wQKBgDVciiBMcfO+4m0n4RyE7XzNfZxVMcFSI8VXF7EG7LsfIX6KzhC3YgsQ\nsLeJ2k4EZL7O6xBj4r/JS4g6zu6m836YhdvX2F25AUW7VHmvCIXg8mHTOsqGZbhL\nczE1sH2q4LuVDTFu4jV3V5iU2QLAAQyasNom7pW07fIuIKtmRuSV\n-----END PRIVATE KEY-----\n"
	//fmt.Println(pem)
//	fmt.Println("----------------")

	log.Println("------")
	log.Println(readPemBundle(pem))


	v := "   How are you     "
	v=strings.TrimSpace(v)
	log.Println("--",v,"---")
	log.Println(v)
}


func readPemBundle(pem string) string {
	if pem == "" {
		return ""
	}

	meta := readPemMeta(strings.NewReader(pem))
	if meta == nil {
		return ""
	}

	return meta.UID
}

type PemMeta struct {
	UID string
	CN  string
}
func readPemMeta(file io.Reader) *PemMeta {
	scanner := bufio.NewScanner(file)
	buf := bytes.NewBufferString("")
	for scanner.Scan() {
		line := scanner.Text()
		buf.WriteString(line)
		buf.WriteString("\n")
		if !strings.Contains(line, "UID=") || !strings.Contains(line, "CN=") {
			continue
		}

		meta := &PemMeta{}
		for _, s := range strings.Split(line, "/") {
			tmp := strings.Split(s, "=")
			if len(tmp) != 2 {
				continue
			}

			if tmp[0] == "UID" {
				meta.UID = tmp[1]
			}
			if tmp[0] == "CN" {
				meta.CN = tmp[1]
			}
		}

		return meta

	}
	log.Println("002")
	//the way of pem can not found result,using p12 way
	v := buf.String()
	cer, err := readPem(v)
	if err != nil {
		return nil
	}
	cname := cer.Subject.CommonName
	meta := &PemMeta{}
	meta.CN = cname
	meta.UID = strings.Split(cname, ":")[1]


	log.Println(cname)
	log.Println("----")
	return meta
}


func readPem(p string) (*x509.Certificate, error) {
	block, e := pem.Decode([]byte(p))
	if block != nil {
		return x509.ParseCertificate(block.Bytes)
	}
	if e != nil {
		return x509.ParseCertificate(e)
	}
	return nil, errors.New("x509 format error")
}