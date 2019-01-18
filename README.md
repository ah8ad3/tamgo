# tamgo
lightweight golang bundle
- app structure in building
- gorm orm for postgresql can use as many dbms 
- every app has its own db and routes and models
- .env support
- session support
- simple cookie based authentication in progress
- auth app created in separate package to manage all auth staff
- static files served

# TODO
- create a bundle to fast create and deploy golang projects in backend
- find some good environment library management like package.json
- create bundle structure like patty
- add travis.yml
- add test coverage and test library
- ...


# NOTE
- after create app you should first pass DB to models in server package, then create subroutes in settings/routes  and pass it to route in app this is
how you can work offline in app, simple app common has the base pattern for app
- i will create this section easier in future



i create this repo in my free time so feel free to help me
contact `ah8ad3@gmail.com`


to run this
`go install` and then run `./tamgo` binary file in your'e gopath bin file