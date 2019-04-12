# Auth

- `user.go` : create/update user
- `token.go` : create custom token with claims
- `claim.go` : set claim
- `revoke.go` : revoke refresh token, check id token, revoke user not id token
- `session.go` : session cookie (5 minutes to 14 days)

## revoke previous user
 App   | Firebase | server 
:-----:|:--------:|:------:
acc/pwd | --- | ---
 -> | acc/pwd  | ---
 -- | id/token | ---
id/token | <- | ---
 -> | --- | session cookie
session cookie | --- | <-
api with cookie | --- | ---
 -> | --- | verify cookie

# Firestore

- `quote.go` : set/get firestore
