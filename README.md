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

// token: D97xRIS9a47POjX6R8pkjwyssFM2IF6IG2q_zSmRxINGtSmOweiy7NifAxi3b4Efk1TSh7yP6_dfQ6Exc_Twe1VhUv5B8gUs2KWA3eoNsV3nKMhdjj4gNYMwTgoXR1zbvFdnlgwoMJdJaLr93Fr-u1MZD6rKn31ZmswjWzyN76awGIiKORAL8x0Uc1NkJXDce7xsDEFgYHpWO1yj_8Z5QQ
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
```
