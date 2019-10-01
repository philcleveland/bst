package main

import "fmt"

//VehicleCondition describes the condition of the vehicle
type VehicleCondition int

const (
	New VehicleCondition = iota + 1
	LikeNew
	Excellent
	Good
	Fair
	Salvage
)

//Drive describes the wheels drive of the vehicle
type Drive int

const (
	FrontWheelDrive Drive = iota + 1
	RearWheelDrive
	FourWheelDrive
	AllWheelDrive
)

//Fuel describes the fuel type of the vehicle
type Fuel int

const (
	Gas Fuel = iota + 1
	Diesal
	Hybrid
	Electric
	other
)

//VehicleSize describes the size of the vehicle
type VehicleSize int

const (
	Compact VehicleSize = iota + 1
	FullSize
	MidSize
	SubCompact
)

//TitleStatus describes the title status of the vehicle
type TitleStatus int

const (
	Clean TitleStatus = iota + 1
	Rebuilt
	PartsOnly
	Lien
	Missing
)

type VehicleType int

const (
	Bus VehicleType = iota + 1
	Convertible
	Coupe
	Hatchback
	MiniVan
	Offroad
	Pickup
	Sedan
	SUV
	Van
	Other
)

//VehiclePost is the read model for a vehicle
type VehiclePost struct {
	VIN      string `json:"type"`
	Odometer uint32 `json:"Odometer"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Year     uint16 `json:"year"`
	Post     `json:"post"`
}

func newVehiclePost(post Post, year uint16, miles uint32, make, model, vin string) *VehiclePost {
	return &VehiclePost{
		Make:     make,
		Odometer: miles,
		Model:    model,
		Year:     year,
		VIN:      vin,
		Post:     post,
	}
}

func (p *VehiclePost) details() string {
	return p.Post.details() +
		fmt.Sprintln(p.Make) +
		fmt.Sprintln(p.Model) +
		fmt.Sprintln(p.Year) +
		fmt.Sprintln(p.Odometer) +
		fmt.Sprintln(p.VIN)
}
