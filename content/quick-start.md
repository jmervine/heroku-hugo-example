+++
date        = "2015-12-12T13:39:29-08:00"
title       = "Quick Start"
description = "A basic example of a Hugo generated site running on Heroku using a simple custom Go static file server."

+++

A basic example of a [Hugo] generated site running on [Heroku] using a simple custom [Go] static file server.

#### Requirements

- [Go]
- [Hugo]
- [Godep]
- [Heroku] Account

[Go]: https://golang.org/
[Hugo]: https://gohugo.io/
[Godep]: https://github.com/tools/godep
[Heroku]: https://www.heroku.com/

#### Using

Fork and clone this repo.

```
heroku create
git push heroku master
```

Making and testing content updates...

```
# install deps
go get -u -v github.com/spf13/hugo
go get -u -v github.com/tools/godep

#
make build server
```

Saving and pushing content updates...

```
git add .
git commit -m 'content updates'
git push heroku master
```

