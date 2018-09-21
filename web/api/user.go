package main

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
)

type User struct {
	Id, Name string
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{user-id}").To(u.findUser))
	ws.Route(ws.POST("").To(u.updateUser))
	ws.Route(ws.PUT("/{user-id}").To(u.createUser))
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))

	container.Add(ws)
}

// GET http://localhost:8083/users/1
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr, ok := u.users[id]
	if !ok {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// POST http://localhost:8083/users
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func (u UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(usr)
	if err == nil {
		u.users[usr.Id] = *usr
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// PUT http://localhost:8083/users/1
// <User><Id>1</Id><Name>Melissa</Name></User>
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := User{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = usr
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8083/users/1
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
}

func main() {
	container := restful.NewContainer()
	container.Router(restful.CurlyRouter{})

	u := &UserResource{
		users: make(map[string]User),
	}
	u.Register(container)

	log.Printf("start linstening on localhost:8083")
	serve := &http.Server{Addr: ":8083", Handler: container}
	log.Fatal(serve.ListenAndServe())
}
