package config

type ValidationConfig struct {
	Type             string
	MinutesToExpires int
	TokenLength      int
	CodeSize         int
}

var (
	UserForgotPasswordValidationConfig = &ValidationConfig{
		Type:             "user-forgot-password-validation",
		MinutesToExpires: 10, // 10 Min
		CodeSize:         5,
	}
)
