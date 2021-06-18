package main

type Film struct {
	Title         string     `json:"title,omitempty"`
	EpisodeID     int        `json:"episode_id,omitempty"`
	OpeningCrawl  string     `json:"opening_crawl,omitempty"`
	Director      string     `json:"director,omitempty"`
	Producer      string     `json:"producer,omitempty"`
	Character     []Person   `json:"characters,omitempty"`
	Planet        []Planet   `json:"planets,omitempty"`
	Starship      []Starship `json:"starships,omitempty"`
	Vehicle       []Vehicle  `json:"vehicles,omitempty"`
	Species       []Specie   `json:"species,omitempty"`
}

type Person struct {
	Name         string   `json:"name,omitempty"`
	Height       string   `json:"height,omitempty"`
	Mass         string   `json:"mass,omitempty"`
	HairColor    string   `json:"hair_color,omitempty"`
	SkinColor    string   `json:"skin_color,omitempty"`
	EyeColor     string   `json:"eye_color,omitempty"`
	BirthYear    string   `json:"birth_year,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	Homeworld    string   `json:"homeworld,omitempty"`
}

type Planet struct {
	Name           string   `json:"name,omitempty"`
	RotationPeriod string   `json:"rotation_period,omitempty"`
	OrbitalPeriod  string   `json:"orbital_period,omitempty"`
	Diameter       string   `json:"diameter,omitempty"`
	Climate        string   `json:"climate,omitempty"`
	Gravity        string   `json:"gravity,omitempty"`
	Terrain        string   `json:"terrain,omitempty"`
	SurfaceWater   string   `json:"surface_water,omitempty"`
	Population     string   `json:"population,omitempty"`
}

type Specie struct {
	Name            string   `json:"name,omitempty"`
	Classification  string   `json:"classification,omitempty"`
	Designation     string   `json:"designation,omitempty"`
	AverageHeight   string   `json:"average_height,omitempty"`
	SkinColors      string   `json:"skin_colors,omitempty"`
	HairColors      string   `json:"hair_colors,omitempty"`
	EyeColors       string   `json:"eye_colors,omitempty"`
	AverageLifespan string   `json:"average_lifespan,omitempty"`
	Homeworld       string   `json:"homeworld,omitempty"`
	Language        string   `json:"language,omitempty"`
}

type Starship struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	HyperdriveRating     string   `json:"hyperdrive_rating,omitempty"`
	MGLT                 string   `json:"MGLT,omitempty"`
	StarshipClass        string   `json:"starship_class,omitempty"`
}

type Vehicle struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	VehicleClass         string   `json:"vehicle_class,omitempty"`
}
