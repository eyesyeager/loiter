package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 15:54
 */

var (
	AppLogger       *zap.SugaredLogger
	BackstageLogger *zap.SugaredLogger
	GatewayLogger   *zap.SugaredLogger
	MDB             *gorm.DB
)
