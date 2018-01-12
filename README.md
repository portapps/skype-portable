<p align="center"><a href="https://github.com/portapps/skype-portable" target="_blank"><img width="100" src="https://github.com/portapps/skype-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/skype-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/skype-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/skype-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/skype-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/skype-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/skype-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/skype-portable"><img src="https://goreportcard.com/badge/github.com/portapps/skype-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/skype-portable"><img src="https://img.shields.io/codacy/grade/07946201a8a74eab9c6021a26f32fb4e.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://saythanks.io/to/crazymax"><img src="https://img.shields.io/badge/thank-crazymax-426aa5.svg?style=flat-square" alt="Say Thanks"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[Skype](https://www.skype.com) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are different kinds of artifacts :

* `skype-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of Skype as a setup. **Recommended way**!
* `skype-portable-{ia32,x64}-x.x.x-x.7z` : Full portable release of Skype as a 7z archive.
* `skype-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `skype-portable.exe`)
* `Skype-{ia32,x64}-x.x.x.x-setup.exe` : The original setup from the [official website](https://www.skype.com/fr/get-skype/).

For a **fresh installation**, install `skype-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `skype-portable.exe`.

If **you have already installed Skype from the original setup**, do the same thing as a fresh installation and move :

* `%APPDATA%\Microsoft\Skype for Desktop\*` to `data`.

Run `skype-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Skype from your computer.

**For an upgrade**, simply download and install the [latest setup](https://github.com/portapps/skype-portable/releases/latest).

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

[![Donate Paypal](https://raw.githubusercontent.com/portapps/portapps/master/res/paypal.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
