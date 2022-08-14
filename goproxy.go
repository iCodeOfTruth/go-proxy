package go_proxy

import (
	"errors"
	"fmt"
	golist "github.com/iCodeOfTruth/go-list"
	"strings"
)

type ProxyManager struct {
	listManager *golist.ListManager
}

func NewProxyManager(proxyList []interface{}) *ProxyManager {
	return &ProxyManager{listManager: golist.NewListManager(proxyList)}
}

func (p *ProxyManager) GetRandomProxy() (string, error) {
	return ParseProxy(p.listManager.Random().(string))
}

func (p *ProxyManager) GetNextProxy() (string, error) {
	return ParseProxy(p.listManager.Next().(string))
}

func ParseProxy(proxy string) (string, error) {
	if !strings.Contains(proxy, ":") {
		return "", errors.New("invalid_format")
	}

	parts := strings.Split(proxy, ":")
	if len(parts) >= 4 {
		return fmt.Sprintf("%s:%s@%s:%s", parts[2], parts[3], parts[0], parts[1]), nil
	}

	return proxy, nil
}
