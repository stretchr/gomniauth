# CSS3 Social Sign-in Buttons

CSS3 Social Sign-in Buttons with icons. Small and large sizes.

## Installation

* [Bower](http://bower.io/): bower install --save necolas/css3-social-signin-buttons
* Download:
  [zip](http://github.com/necolas/css3-social-signin-buttons/zipball/master) or
  [tar](http://github.com/necolas/css3-social-signin-buttons/tarball/master)
  formats.
* Git: git clone git://github.com/necolas/css3-social-signin-buttons.git

## Buttons

### Default

To create the default sign-in button, add a class of `btn-auth` and
`btn-[service]` (where `[service]` is one of the supported social sign-in
services) to any appropriate element (most likely an anchor).

```html
<a class="btn-auth btn-[service]" href="#">
    Sign in with <b>[service]</b>
</a>
```

### Large

To create large buttons include an additional class of `large`.

```html
<a class="btn-auth btn-[service] large" href="#">
    Sign in with <b>[service]</b>
</a>
```

## Supported services

* Facebook
* GitHub
* Google
* OpenID
* Twitter
* Windows Live ID
* Yahoo!

## Browser support

* Google Chrome
* Firefox 3.5+
* Safari 4+
* IE 8+
* Opera

**Note:** Some CSS3 enhancements are not supported in older versions of Opera
and IE. The use of icons is not supported in IE 6 or IE 7.

## License

Public domain: [http://unlicense.org](http://unlicense.org)
