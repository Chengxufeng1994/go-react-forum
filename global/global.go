package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GRF_DB *gorm.DB
	GRF_VP *viper.Viper
)
