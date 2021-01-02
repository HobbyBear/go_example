package msgbus
//
//import (
//	"fmt"
//	"testing"
//
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
//var Logger *zap.Logger
//
//func init() {
//	config := zap.Config{
//		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
//		Development: false,
//		Sampling: &zap.SamplingConfig{
//			Initial:    100,
//			Thereafter: 100,
//		},
//		Encoding:         "json",
//		EncoderConfig:    zap.NewProductionEncoderConfig(),
//		OutputPaths:      []string{"stderr"},
//		ErrorOutputPaths: []string{"stderr"},
//	}
//	Logger, _ = config.Build(zap.Hooks(func(entry zapcore.Entry) error {
//		fmt.Println(entry)
//		return nil
//	}))
//
//
//}
//
//func Test01(t *testing.T) {
//
//	config := zap.Config{
//		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
//		Development: false,
//		Sampling: &zap.SamplingConfig{
//			Initial:    100,
//			Thereafter: 100,
//		},
//		Encoding:         "json",
//		EncoderConfig:    zap.NewProductionEncoderConfig(),
//		OutputPaths:      []string{"stderr"},
//		ErrorOutputPaths: []string{"stderr"},
//	}
//	Logger, _ = config.Build(zap.Hooks(func(entry zapcore.Entry) error {
//		fmt.Println(entry)
//		return nil
//	}))
//
//	Logger.Info("test", zap.String("name", "xch"))
//
//}
//
//func Test02(t *testing.T)  {
//	//var logger *zap.Logger
//
//}