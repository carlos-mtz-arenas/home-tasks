package services

import (
  "errors"
  "hometasks/internal/entities"
)

func AuthDevice(key string, secret string) (*entities.DeviceEntity, error) {
  if key == "hello" && secret == "secret" {
    return &entities.DeviceEntity {
      Uuid: "some-uuid",
      Name: "casa device",
      Key: "hello",
      Secret: "secret",
    }, nil
  }

  return nil, errors.New("invalid credentials")
}
