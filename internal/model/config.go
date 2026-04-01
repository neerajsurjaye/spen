package model

type Config struct {
	Circles []CircleConfig `json:"circles"`
	Walls   []WallConfig   `json:"walls"`
	Ui      UIConfig       `json:"ui"`
}

type CircleConfig struct {
	Transform CircleTransformConfig `json:"transforms"`
	Color     []float32             `json:"color"`
	Body      *BodyConfig           `json:"body"`
}

type CircleTransformConfig struct {
	PosX   float32 `json:"posX"`
	PosY   float32 `json:"posY"`
	Radius float32 `json:"radius"`
}

type WallConfig struct {
	Transform WallTransformConfig `json:"transforms"`
	Color     []float32           `json:"color"`
}

type WallTransformConfig struct {
	PosX     float32 `json:"posX"`
	PosY     float32 `json:"posY"`
	ScaleX   float32 `json:"scaleX"`
	ScaleY   float32 `json:"scaleY"`
	Rotation float32 `json:"rotation"`
}

type UIConfig struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}

type BodyConfig struct {
	VelX       float32 `json:"velX"`
	VelY       float32 `json:"velY"`
	Mass       float32 `json:"mass"`
	Restititon float32 `json:"restitution"`
}
