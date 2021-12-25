package services

import (
	"notificationService/configuration"
	"notificationService/sender"
	"notificationService/transports"
	"reflect"
)

type Transporter interface {
	GetTransport(config *configuration.Config) []sender.ISender
}

type Transport struct {
	*Logger
}

func (t *Transport) GetTransport(config *configuration.Config) []sender.ISender {
	var tr []sender.ISender

	funcs := map[string]interface{}{
		"transports.NewTgBot": transports.NewTgBot,
	}

	val := reflect.ValueOf(config.Transports)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			f := val.Field(i).FieldByName("Active")
			if f.Bool() == true {
				if _, ok := funcs[val.Field(i).FieldByName("Callback").String()]; ok {
					initFunc := reflect.ValueOf(funcs[val.Field(i).FieldByName("Callback").String()])
					in := make([]reflect.Value, 1)
					in[0] = reflect.ValueOf(config)
					result := initFunc.Call(in)
					err, ok := result[1].Interface().(error)
					if ok {
						t.logger.Errorf("cannot create transport from %s, got err %v", val.Field(i).Type().Name(), err)
						continue
					}
					tr = append(tr, result[0].Interface().(sender.ISender))
					t.logger.Infof("created transport from %s", val.Field(i).Type().Name())
				}
			}
		}
	}

	return tr
}

func NewTransport(logger *Logger) *Transport {
	return &Transport{
		logger,
	}
}
