package object

import (
	"github.com/kr/pretty"
	"testing"
)

func Test_Attribute_Set_AND_Get(t *testing.T) {
	mapTest := NewAttribute(&HashMap{
		"gun": "model",
	})

	mapTest.SetAttribute("weapon.bullet", 100)
	mapTest.SetAttribute("weapon.shield.strength", "strong")

	bulletCount := mapTest.GetAttribute("weapon.bullet", 0)
	if bulletCount != 100 {
		t.Error("get bullet error")
		pretty.Log(bulletCount)
	}

	shieldStrength := mapTest.Get("weapon.shield.strength", "")
	if shieldStrength != "strong" {
		t.Error("get shield error")
		pretty.Log(shieldStrength)
	}

}
