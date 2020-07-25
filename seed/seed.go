package seed

import (
	"github.com/Pradnyana28/uploads/uploadpb"
	"github.com/golang/protobuf/ptypes"
)

// NewKeyboard ...
func NewKeyboard() *uploadpb.Keyboard {
	keyboard := &uploadpb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU ...
func NewCPU() *uploadpb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &uploadpb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz: minGhz,
		MaxGhz: maxGhz,
	}

	return cpu
}

// NewGPU ...
func NewGPU() *uploadpb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &uploadpb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit: uploadpb.Memory_GIGABYTE,
	}

	gpu := &uploadpb.GPU{
		Brand: brand,
		Name: name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

// NewRAM ...
func NewRAM() *uploadpb.Memory {
	ram := &uploadpb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit: uploadpb.Memory_MEGABYTE,
	}

	return ram
}

// NewSSD ...
func NewSSD() *uploadpb.Storage {
	ssd := &uploadpb.Storage{
		Driver: uploadpb.Storage_SSD,
		Memory: &uploadpb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit: uploadpb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD ...
func NewHDD() *uploadpb.Storage {
	hdd := &uploadpb.Storage{
		Driver: uploadpb.Storage_HDD,
		Memory: &uploadpb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit: uploadpb.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen ...
func NewScreen() *uploadpb.Screen {
	screen := &uploadpb.Screen{
		SizeInch: randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel: randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop ...
func NewLaptop() *uploadpb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &uploadpb.Laptop{
		Id: randomID(),
		Brand: brand,
		Name: name,
		Cpu: NewCPU(),
		Ram: NewRAM(),
		Gpus: []*uploadpb.GPU{NewGPU(), NewGPU()},
		Storages: []*uploadpb.Storage{NewSSD(), NewHDD()},
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &uploadpb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd: randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt: ptypes.TimestampNow(),
	}

	return laptop
}