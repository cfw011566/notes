# Auth

- `user.go` : create/update user
- `token.go` : create custom token with claims
- `claim.go` : set claim
- `revoke.go` : revoke refresh token, check id token

## revoke previous user
 App   | Firebase | server 
:-----:|:--------:|:------:
acc/pwd | --- | ---
 >> | acc/pwd  | ---
 -- | id/token | ---
id/token | << | ---
>> | --- | id/token and revoke previous token
api with token | --- | ---
>> | --- | verify token and check revoked

# Firestore

- `quote.go` : set/get firestore
