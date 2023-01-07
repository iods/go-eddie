package parseutil

// ParseFlag wraps the pflag.FlagSet struct
type ParseFlag interface {
	GetBool(name string) (bool, error)
	GetString(name string) (string, error)
	GetStringArray(name string) ([]string, error)
	GetInt(name string) (int, error)
	Set(name, value string) error
}
