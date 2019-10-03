package lunarsolar

import (
	"github.com/sirupsen/logrus"
	"testing"
)

/**
*公历转农历
 */
//func SolarToLunar(solar Solar) *Lunar
func TestSolarToLunar(t *testing.T) {
	lunar := SolarToLunar(Solar{
		SolarDay:   3,
		SolarMonth: 10,
		SolarYear:  2019,
	})
	logrus.Infof("lunar : %#v", lunar)
}

/**
*农历转公历
 */
//func  LunarToSolar(lunar Lunar) *Solar
func TestLunarToSolar(t *testing.T) {
	lunar := SolarToLunar(Solar{
		SolarDay:   3,
		SolarMonth: 10,
		SolarYear:  2019,
	})
	solar := LunarToSolar(*lunar)
	logrus.Infof("solar : %#v", solar)
}
