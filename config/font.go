package config

type FontConfiguration struct {
	DosisFontPath  string `env:"DOSIS_FONT_PATH, default=./etc/Dosis.ttf"`
	RobotoSlabPath string `env:"ROBOTO_SLAB_FONT_PATH, default=./etc/RobotoSlab.ttf"`
}
