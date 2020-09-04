package tanyibot-api-go

import(
	// "crypto/hmac"
	"crypto/sha256"
	"sort"
	"bytes"
	"encoding/hex"
	"fmt"
)

type KV struct {
	V string
	K string
}

type KVS []KV
// implement sort.Interface
func (p KVS) Len() int           { return len(p) }
func (p KVS) Less(i, j int) bool { return p[i].K < p[j].K }
func (p KVS) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sign(pairs map[string]string, tenantSign string, needSort bool) string {	
	kvs := make(KVS, 0, len(pairs))
	for k, v := range pairs {
		kvs = append(kvs, KV{k, v})
	}

	if needSort {
		sort.Sort(kvs)
	}
	
	return SignArr(kvs, tenantSign)
}

func SignArr(kvs KVS, tenantSign string) string {
	// h := hmac.New(sha256.New)
	h := sha256.New()

	var buffer bytes.Buffer
	if len(kvs) > 0 {
		buffer.WriteString(kvs[0].K)
		buffer.WriteString("=")
		buffer.WriteString(kvs[0].V)
		for _, kv := range kvs[1:] {
			buffer.WriteString("&")
			buffer.WriteString(kv.K)
			buffer.WriteString("=")
			buffer.WriteString(kv.V)
		}
	}

	h.Write([]byte(tenantSign))
	h.Write([]byte(buffer.String()))
	signature := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("params: %s, signature: %s\n", buffer.String(), signature)
	return signature
}
