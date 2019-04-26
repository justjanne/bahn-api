package bahn

type CoachTypeInfo struct {
	RawType string `json:"raw_type"yaml:"raw_type"`

	DieselLocomotive         bool `json:"diesel_locomotive,omitempty"yaml:"diesel_locomotive,omitempty"`
	DieselElectricLocomotive bool `json:"diesel_electric_locomotive,omitempty"yaml:"diesel_electric_locomotive,omitempty"`
	ElectricLocomotive       bool `json:"electric_locomotive,omitempty"yaml:"electric_locomotive,omitempty"`

	FirstClass              bool `json:"first_class,omitempty"yaml:"first_class,omitempty"`
	SecondClass             bool `json:"second_class,omitempty"yaml:"second_class,omitempty"`
	DoubleDeck              bool `json:"double_deck,omitempty"yaml:"double_deck,omitempty"`
	Sleeping                bool `json:"sleeping,omitempty"yaml:"sleeping,omitempty"`
	Restaurant              bool `json:"restaurant,omitempty"yaml:"restaurant,omitempty"`
	Bistro                  bool `json:"bistro,omitempty"yaml:"bistro,omitempty"`
	CarTransport            bool `json:"car_transport,omitempty"yaml:"car_transport,omitempty"`
	Saloon                  bool `json:"saloon,omitempty"yaml:"saloon,omitempty"`
	Accessible              bool `json:"accessible,omitempty"yaml:"accessible,omitempty"`
	Couchette               bool `json:"couchette,omitempty"yaml:"couchette,omitempty"`
	AirConditioning         bool `json:"air_conditioning,omitempty"yaml:"air_conditioning,omitempty"`
	InterRegio              bool `json:"interregio,omitempty"yaml:"interregio,omitempty"`
	InterCity               bool `json:"intercity,omitempty"yaml:"intercity,omitempty"`
	Bicycle                 bool `json:"bicycle,omitempty"yaml:"bicycle,omitempty"`
	Compartments            bool `json:"compartments,omitempty"yaml:"compartments,omitempty"`
	OpenCoach               bool `json:"open_coach,omitempty"yaml:"open_coach,omitempty"`
	ControlCar              bool `json:"control_car,omitempty"yaml:"control_car,omitempty"`
	ServicePoint            bool `json:"service_point,omitempty"yaml:"service_point,omitempty"`
	ReducedCompartmentCount bool `json:"former_first_class,omitempty"yaml:"former_first_class,omitempty"`
}
