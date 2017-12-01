package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"strings"
	"sync"
	"testing"
)

type StringSet map[string]struct{}

func assert(t *testing.T, a interface{}, b interface{}, msg string) {
	if a == b {
		return
	}
	if (a == nil || reflect.ValueOf(a).IsNil()) && (b == nil || reflect.ValueOf(b).IsNil()) {
		return
	}
	t.Errorf("%s %+v != %+v", msg, a, b)
}

func replace(tmpl string, mapping map[string]string) string {
	for k, v := range mapping {
		tmpl = strings.Replace(tmpl, fmt.Sprintf("{%s}", k), v, -1)
	}
	return tmpl
}

func substr(str, begin, end string) string {
	bpos := strings.Index(str, begin)
	if bpos == -1 {
		return ""
	}
	sub := str[bpos+len(begin):]
	e := strings.Index(sub, end)
	if e == -1 {
		return ""
	}
	return sub[:e]
}

func loadRedisSet(client *redis.Client, key string, set *StringSet, mutex *sync.Mutex) error {
	list, err := client.SMembers(key).Result()
	log.Infof("Redis SMembers %s result, len: %d, err: %v", key, len(list), err)
	if err != nil {
		log.Errorf("Redis SMembers failed,key: %s, len: %d, err: %v", key, len(list), err)
		return err
	}
	ninjaSet := make(StringSet)
	for _, v := range list {
		ninjaSet[v] = struct{}{}
	}
	if mutex != nil {
		mutex.Lock()
	}
	*set = ninjaSet
	if mutex != nil {
		mutex.Unlock()
	}
	return nil
}
