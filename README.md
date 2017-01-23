# Gomniauth

Authentication framework for Go applications.

* EDITOR NOTE: It is recommended that you transition to the [Goth package](https://github.com/markbates/goth).

---

## Features

  * OAuth2 today
  * Supports other protocols too
  * [Multiple providers](https://github.com/stretchr/gomniauth/tree/master/providers) - Supports Google, GitHub and Facebook and [more](https://github.com/stretchr/gomniauth/tree/master/providers)
  * Easily extensible
  * [Example web app](https://github.com/stretchr/gomniauth/tree/master/example) to copy
  * Works beautifully with [Goweb](https://github.com/stretchr/goweb)
  * Fully [TDD](http://en.wikipedia.org/wiki/Test-driven_development)

## Documentation

  * Jump right into the [API Documentation](http://godoc.org/github.com/stretchr/gomniauth)

## Get started

Install Gomniauth by doing:

```shell
$ go get github.com/stretchr/gomniauth
```

Check out the [example web app code](https://github.com/stretchr/gomniauth/tree/master/example) to see how to use Gomniauth using [Goweb](https://github.com/stretchr/goweb).

## Contributing

  * If you add a provider that others could also make use of, please send us a Pull Request and we'll add it to the repo.

## Implementing Gomniauth

### Set up Gomniauth

First and only once for your application, you need to setup the security key and providers.  The security key is used when hashing any data that gets transmitted to ensure its integrity.

You are free to use the [signature package's RandomKey function](http://godoc.org/github.com/stretchr/signature#RandomKey) to generate a unique code every time your application starts.

```go
gomniauth.SetSecurityKey(signature.RandomKey(64))
```

A provider represents an authentication service that will be available to your users.  Usually, you'll have to add some configuration of your own, such as your application `key` and `secret` (provided by the auth service), and the `callback` into your app where users will be sent following successful (or not) authentication.

```go
gomniauth.WithProviders(
  github.New("key", "secret", "callback"),
  google.New("key", "secret", "callback"),
)
```    

#### What kind of callback?

The callback should be an absolute URL to your application and should include the provider name in some way.

For example, in the [example web app](https://github.com/stretchr/gomniauth/tree/master/example) we used the following format for callbacks:

    http://mydomain.com/auth/{provider}/callback

### Are they logged in?

When a user tries to access a protected resource, or if you want to make third party authenticated API calls, you need a mechanism to decide whether a user is logged in or not.  For web apps, cookies usually work, if you're building an API, then you should consider some kind of auth token.

### Decide how to log in

If they are not logged in, you need to provide them with a list of providers from which they can choose.

You can access a list of the providers you are supporting by calling the `gomniauth.Providers()` function.

### Redirecting them to the login page

Once a provider has been chosen, you must redirect them to be authenticated.  You can do this by using the `gomniauth.Provider` function, that will return a [Provider](http://godoc.org/github.com/stretchr/gomniauth/common#Provider) by name.

So if the user chooses to login using Github, you would do:

```go
provider, err := gomniauth.Provider("github")
```

Once you have your provider, you can get the URL to redirect users to by calling:

```go
authUrl, err := provider.GetBeginAuthURL(state, options)
```

You should then redirect the user to the `authUrl`.

#### State and options

The `state` parameter is a `State` object that contains information that will be hashed and passed (via the third party) through to your callback (see below).  Usually, this object contains the URL to redirect to once authentication has completed, but you can store whatever you like in here.

The `options` parameter is an `objx.Map` containing additional query-string parameters that will be sent to the authentication service.  For example, in OAuth2 implementations, you can specify a `scope` parameter to get additional access to other services you might need.

### Handling the callback

Once the third party authentication service has finished processing the request, they will send the user back to your callback URL.

Remember, you specified the callback URL when you setup your providers.

For example, the user might hit:

    http://yourdomain.com/auth/github/callback?code=abc123

You don't need to worry about the detail of the parameters passed back, because Gomniauth takes care of those for you.  You just need to pass them into the `CompleteAuth` method of your provider:

```go
provider, err := gomniauth.Provider("github")
creds, err := provider.CompleteAuth(queryParams)
```

NOTE: It's unlikely you'll hard-code the provider name, we have done it here to make it obvious what's going on.

The provider will then do the work in the background to complete the authentication and return the `creds` for that user.  The `creds` are then used to make authenticated requests to the third party (in this case Github) on behalf of the user.

### Getting the user information

If you then want some information about the user who just authenticated, you can call the `GetUser` method on the provider (passing in the `creds` from the `CompleteAuth` method.)

The [User](https://github.com/stretchr/gomniauth/blob/master/common/user.go) you get back will give you access to the common user data you will need (like name, email, avatar URL etc) and also an `objx.Map` of `Data()` that contains everything else.

### Caching in

Once you had the credentials for a user for a given provider, you should cache them in your own datastore.  This will mean that if the cookie hasn't expired, or if the client has stored the auth token, they can continue to use the service without having to log in again.
