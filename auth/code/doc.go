// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package code implements Atlas user authentication using OAuth 2.0.

Both flows share `Config`, `Token`, and common options. `NewConfig` and
`NewConfigWithOptions` create a `Config`. `SetClientID`, `SetScopes`,
`SetUserAgent`, and `SetWithRaw` configure both flows. `SetAuthURL` configures
the device endpoint. `SetAuthServerURL` configures the Authorization Server issuer.

# Device Authorization Grant (Legacy)

The legacy device flow remains available for existing consumers. Its `Config`
methods are `RequestCode`, `GetToken`, `PollToken`, `RefreshToken`, `RevokeToken`,
and `RegistrationConfig`. It uses `Config.NewRequest` and `Config.Do` for HTTP
requests. `IsTimeoutErr` reports device-flow expiration errors.

# Authorization Code + PKCE

The Authorization Code flow uses `Config.AuthServerURL`, configured with
`SetAuthServerURL`. Its `Config` methods are `DiscoverAuthServer`,
`AuthorizationURL`, `ExchangeCode`, `RefreshAccessToken`, and
`RevokeAuthServerToken`. `GenerateState` and `GenerateVerifier` create per-login
state and PKCE values. `StartCallbackServer` validates state and, when discovery
metadata advertises `AuthorizationResponseIssParameterSupported`, requires the
callback `iss` to match the discovered issuer. The returned `CallbackServer`
provides `RedirectURI`, `Port`, `WaitForCallback`, and `Close` methods.

Token-returning methods in both flows return `*Token` values.
`Token.SetAuthHeader` adds the access token as a Bearer header. See
`examples/auth_code` for using an exchanged token with the Atlas Admin API.
Flow methods accept contexts for cancellation and deadlines.
*/
package code
