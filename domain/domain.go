package domain

import (
	"net"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var searcher *xdb.Searcher

// 加载地址库
func Load(dbpath string) (err error) {
	buffer, err := xdb.LoadContentFromFile(dbpath)
	if err != nil {
		return
	}

	if searcher, err = xdb.NewWithBuffer(buffer); err != nil {
		return
	}

	return
}

// 关闭地址库
func Close() {
	if searcher != nil {
		searcher.Close()
		searcher = nil
	}
}

// 通过网络地址查找ip区域
func Search(addr string) (region string, err error) {
	host := addr
	if strings.Contains(host, ":") {
		if host, _, err = net.SplitHostPort(host); err != nil {
			return
		}
	}

	// 域名转ip地址
	ipArray, err := net.LookupIP(host)
	if err != nil || len(ipArray[0]) <= 0 {
		return
	}

	// 判断是否内网
	ip := ipArray[0].String()
	localRegion := "0|0|0|内网IP|内网IP"

	ip4 := ipArray[0].To4()
	if ip4 == nil {
		region = localRegion
		return
	}

	region, err = SearchIP(ip)
	return
}

// 通过ip查询地域
func SearchIP(ip string) (region string, err error) {
	region, err = searcher.SearchByStr(ip)
	return
}

// 判断ip是否在国内
func HostInPRC(host string) bool {
	region, err := Search(host)
	if err != nil {
		return false
	}

	v := strings.Split(region, "|")

	if v[0] == "0" || v[0] == "中国" {
		return true
	}

	return false
}
