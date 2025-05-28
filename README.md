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

// token: JYqsNem8DP-Cd3VPkQG3DVQegWxQr8lXW2Qi8BC6U6lBk-8-Ek836sOnkH5cIbmm2aUdYB3GaHdsocTZ2-rHiDZoDTphg47NuXh3cv2T3kRx1oOt3j3a0EDU52JsRbHmt3dCoN-83u26Qv05V3g_U2G4I87V7_ic79ix4iS9MxSGiuEEuKQyP3UJ3U2zIhfyCMVxIeMDxsLAABT92ImCHyBALFf0Rw_isdlGP8IZcgF5fzAzreptLaYDTcPvW02sA2qJes_qG5Z9oS2MaSA9vYXaIjk8_3eYErJLzG4fn-xCZfWaNynNCkQTuW19SyCsj-X7LlqwIJUXWcLU4X5crXvRS2egEFi5U5EbBC2FIEpIEaQXnlDIpqCMD-fJxYl6
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
