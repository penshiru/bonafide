package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/lestrrat-go/jwx/jwa"
	jwt "github.com/lestrrat-go/jwx/jwt"
)

//Auth Verifies if the token is valid. If error is nil, is valid
func Auth(config *Config, v io.Reader) error {
	token, err := jwt.Parse(v, jwt.WithVerify(jwa.HS256, config.secret))
	if err != nil {
		return errors.New("Could not parse JWT token")
	}
	buf, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		fmt.Printf("failed to generate JSON: %s\n", err)
		return err
	}

	fmt.Printf("%s\n", buf)

	return nil
}
