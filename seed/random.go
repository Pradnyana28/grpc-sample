package seed

import (
	"math/rand"
	"time"

	"github.com/Pradnyana28/uploads/uploadpb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() uploadpb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return uploadpb.Keyboard_QWERTY;
	case 2:
		return uploadpb.Keyboard_QWERTZ;
	default:
		return uploadpb.Keyboard_AZERTY;
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750M",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max - min + 1)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32() * (max-min)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64() * (max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1060-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomScreenResolution() *uploadpb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &uploadpb.Screen_Resolution{
		Height: uint32(height),
		Width: uint32(width),
	}
	return resolution
}

func randomScreenPanel() uploadpb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return uploadpb.Screen_IPS
	}
	return uploadpb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Asus", "Dell")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet(
			"MacBook Pro 13 Inch + Touchbar",
			"MacBook Air 13 Inch",
			"MacBook",
		)
	case "Asus":
		return randomStringFromSet(
			"Republic Of Gamer",
			"ThinkPad",
		)
	default:
		return randomStringFromSet(
			"Latitude",
			"Vostro",
			"XPS",
		)
	}
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}