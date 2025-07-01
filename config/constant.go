package config

import "fmt"

// captcha prefix key in redis
const CAPTCHA_PREFIX = "user:login:captcha:uid"

const LOGIN_USER_TOKEN_PREFIX = "user:login:user:token"

const LOGIN_USER_UID_PREFIX = "user:login:user:uid"

const API_V1_PREFIX = "/boxdb/api/v1"

// skip check token
var SKIP_AUTH_API = []string{fmt.Sprintf("%s/user/login/username", API_V1_PREFIX), fmt.Sprintf("%s/user/captcha/generate", API_V1_PREFIX)}

const BOXDB_TOKEN_KEY = "boxdb-token"
