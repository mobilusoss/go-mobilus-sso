# go-mobilus-sso

[![Build Status](https://cloud.drone.io/api/badges/mobilusoss/go-mobilus-sso/status.svg)](https://cloud.drone.io/mobilusoss/go-mobilus-sso)
[![codecov](https://codecov.io/gh/mobilusoss/go-mobilus-sso/branch/master/graph/badge.svg)](https://codecov.io/gh/mobilusoss/go-mobilus-sso)

mobiSeries SSOを行うライブラリーです

## Marshal

```golang
sso := go_mobilus_sso.New(secret)
token, err := sso.Marshal(go_mobilus_sso.User{
    Name:                  "テストマン",
    PermitLevel:           0,
    Token:                 "thisistesttoken",
    DomainID:              "adm",
    PlusID:                "testman@example.com",
    UserID:                "testman@example.com",
    TenantID:              "mobilus",
    PermissionDisplayName: "管理者",
})
if err != nil {
    panic(err)
}

// token: D97xRIS9a47POjX6R8pkjwyssFM2IF6IG2q_zSmRxIMsZoFJGzOCMzeK1Rriwr5e2LBp6g-qA-59HSyMwhfm9ACjoE16XQF0tU6U0IziX-QrE2u1vVVDh05ecUPZ-KSvMQIn6nXr1mSQBCbDfATosl6jL9wUgMjCsbu0qABI9e0SPmb5R2J_L9vDuxp-weCpQXKKEk2N2gFn70GsbxEXOybJ1qbhocSH6m4ZJid3ZGwyStIp-kyMvd6z5P0DfYJyy0gkJ4xggqHUXk-W9mE4_qHzONRSPojKqqIBo1TgLIk
```

## Unmarshal

```golang
sso := go_mobilus_sso.New(secret)
user, err := sso.Unmarshal(cookie)
if err != nil {
    panic(err)
}

// user.Name
// user.PermitLevel
// user.Token
// user.DomainID
// user.PlusID
// user.UserID
// user.TenantID
// user.PermissionDisplayName
```
