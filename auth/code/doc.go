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
Package code provides a way to follow a Device Authorization Grant
https://datatracker.ietf.org/doc/html/rfc8628 and the OAuth 2.0 Authorization
Code Grant https://datatracker.ietf.org/doc/html/rfc6749#section-4.1 with PKCE
https://datatracker.ietf.org/doc/html/rfc7636.

Usage

	import "go.mongodb.org/atlas-sdk/v20250312026/auth/code"

Construct a new client Config, then use the various methods to complete a flow.
For example:

	config, err := code.NewConfigWithOptions(nil, code.SetClientID("my-client-id"), code.SetScopes([]string{"openid"}))
	if err != nil {
		panic(err)
	}
	deviceCode, _, err := config.RequestCode(context.TODO())
	if err != nil {
		panic(err)
	}
	token, _, err := config.PollToken(context.TODO(), deviceCode)
	if err != nil {
		panic(err)
	}
	fmt.Println(token.AccessToken)

NOTE: Using the https://godoc.org/context package, one can easily
pass cancellation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then context.Background()
can be used as a starting point.

Authorization Code + PKCE

	The authorization code flow is separate from device authorization. The caller
	provides a redirect URI, receives the authorization response, and validates
	state before exchanging the code:

	config, err := code.NewConfigWithOptions(nil,
	    code.SetClientID("my-client-id"),
	    code.SetScopes([]string{"atlas"}),
	)
	if err != nil {
	    panic(err)
	}
	metadata, _, err := config.DiscoverAuthServer(context.Background())
	if err != nil {
	    panic(err)
	}
	pkce, err := code.GeneratePKCE()
	if err != nil {
	    panic(err)
	}
	state, err := code.GenerateState()
	if err != nil {
	    panic(err)
	}
	callback, err := code.StartCallbackServer(state)
	if err != nil {
	    panic(err)
	}
	defer callback.Close()
	redirectURI := callback.RedirectURI()
	authorizationURL, err := config.AuthorizationURL(
	    metadata, redirectURI, state, pkce,
	)
	if err != nil {
	    panic(err)
	}
	_ = authorizationURL // open this URL in a browser
	authorizationCode, err := callback.WaitForCallback(context.Background())
	if err != nil {
	    panic(err)
	}
	token, _, err := config.ExchangeCode(
	    context.Background(), metadata, authorizationCode, redirectURI, pkce.CodeVerifier,
	)
	if err != nil {
	    panic(err)
	}
	fmt.Println(token.AccessToken)
*/
package code
