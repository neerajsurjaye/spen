package model

type Config struct {
	Circles []CircleConfig `json:"circles"`
	Walls   []WallConfig   `json:"walls"`
	Ui      UIConfig       `json:"ui"`
}

type CircleConfig struct {
	Transform CircleTransformConfig `json:"transforms"`
	Color     []float32             `json:"color"`
}

type CircleTransformConfig struct {
	PosX  float32 `json:"posX"`
	PosY  float32 `json:"posY"`
	Scale float32 `json:"scale"`
}

type WallConfig struct {
	Transform WallTransformConfig `json:"transforms"`
	Color     []float32           `json:"color"`
}

type WallTransformConfig struct {
	PosX   float32 `json:"posX"`
	PosY   float32 `json:"posY"`
	ScaleX float32 `json:"scaleX"`
	ScaleY float32 `json:"scaleY"`
}

type UIConfig struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}
