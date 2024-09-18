package setting

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type App struct {
	AccessTokenSecret   string
	RefreshTokenSecret  string
	AccessTokenExpireIn int

	RefreshTokenExpireIn int

	PageSize  int
	PrefixUrl string
	Env       string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Sslmode  string
}

var DatabaseSetting = &Database{}

type RedisCache struct {
	Host     string
	Port     int
	Password string
}

var RedisCacheSetting = &RedisCache{}

// Setup initializes the configuration instance
func Setup() {
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mapEnvToStruct("app", AppSetting, envs)
	mapEnvToStruct("server", ServerSetting, envs)
	mapEnvToStruct("database", DatabaseSetting, envs)
	mapEnvToStruct("redis", RedisCacheSetting, envs)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapEnvToStruct(prefix string, v interface{}, envs map[string]string) {
	for _, field := range getFields(v) {
		envVar := envs[prefix+"_"+toSnakeCase(field)]
		if envVar != "" {
			setField(v, field, envVar)
		}
	}
}

func getFields(v interface{}) []string {
	var fields []string
	valueType := reflect.TypeOf(v)
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		// Ignore unexported fields
		if field.PkgPath != "" {
			continue
		}
		fields = append(fields, field.Name)
	}
	return fields
}

func setField(v interface{}, fieldName string, value string) error {
	valueType := reflect.TypeOf(v)
	if valueType.Kind() != reflect.Ptr {
		return fmt.Errorf("setField: v must be a pointer to a struct")
	}

	valueElem := reflect.ValueOf(v).Elem()
	fieldValue := valueElem.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return fmt.Errorf("setField: field %s not found in struct", fieldName)
	}

	fieldType := fieldValue.Type()
	var parsedValue reflect.Value

	// Convert string value to the appropriate type
	switch fieldType.Kind() {
	case reflect.String:
		parsedValue = reflect.ValueOf(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		parsedValue = reflect.ValueOf(intValue).Convert(fieldType)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		parsedValue = reflect.ValueOf(uintValue).Convert(fieldType)
	// Add cases for other types as needed
	default:
		return fmt.Errorf("setField: unsupported field type %s", fieldType.Kind())
	}

	fieldValue.Set(parsedValue)
	return nil
}

func toSnakeCase(s string) string {
	var snake string
	for i, r := range s {
		if 'A' <= r && r <= 'Z' {
			if i > 0 {
				snake += "_"
			}
			snake += string(r + 32) // convert to lowercase
		} else {
			snake += string(r)
		}
	}
	return snake
}
