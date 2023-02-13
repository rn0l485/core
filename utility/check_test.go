package utility

import (
	"testing"
)


func TestIsAlpha(t *testing.T) {
	strList := []string{"", "asdzxcQWEASD", "ASDZXCasqwe", "123548", "#$#@qwesa", "AWDSAFDA"}
	for _, v := range strList {
		if !IsAlpha(v) {
			t.Error("fail: "+v)
		}
	}
	t.Log("success")
}

func TestIsDigit(t *testing.T) {
	strList := []string{"", "asdzxcQWEASD", "ASDZXCasqwe", "123548", "#$#@qwesa", "AWDSAFDA"}
	for _, v := range strList {
		if !IsDigit(v) {
			t.Error("fail: "+v)
		}
	}
	t.Log("success")	
}

func TestIsCapital(t *testing.T) {
	strList := []string{"", "asdzxcQWEASD", "ASDZXCasqwe", "123548", "#$#@qwesa", "AWDSAFDA"}
	for _, v := range strList {
		if !IsCapital(v) {
			t.Error("fail: "+v)
		}
	}
	t.Log("success")		
}

func TestSwap(t *testing.T) {
	strList := []string{"", "asdzxcQWEASD", "ASDZXCasqwe", "123548", "#$#@qwesa", "AWDSAFDA"}
	Swap(strList, 0, 1)
	if strList[1] != "" {
		t.Error("fail")
	} else {
		t.Log("success")	
	}
}