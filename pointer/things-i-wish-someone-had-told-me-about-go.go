package main

func test(expr bool) {
	if !expr {
		panic("failed")
	}
}

func main() {

	// Everything passed as parameter will be copied in Golang
	u1 := &User{Name: "Leto"}
	test(u1.Name == "Leto")

	Modify1(u1, "Paul")
	test(u1.Name == "Leto")

	u2 := User{Name: "Leto"}
	Modify2(u2, "Duncan")
	test(u2.Name == "Leto")

	u3 := &User{Name: "Leto"}
	Modify3(u3, "Paul")
	test(u3.Name == "Paul")

	/**
	 *  func(ServeHTTP res http.ResponseWriter, req *http.Request) { ... }
	 */

	// Technically, we don't know whether the value being passed is
	// a copy of pointer / a copy of value
	// but it's probably the former

	StoreCache(&CachedUser{Name: "Leto", id: "leto"})
	StoreCache(Application{id: "letoApp"})
}

type User struct {
	Name string
}

func Modify1(u *User, name string) {
	u = &User{Name: name}
}

func Modify2(u User, name string) {
	u.Name = name
}

func Modify3(u *User, name string) {
	u.Name = name
}

type CacheItem interface {
	GetId() string
}

func StoreCache(item CacheItem) {
	// do something iwth item.GetId()
}

type Application struct {
	id string
}

type CachedUser struct {
	id   string
	Name string
}

func (a Application) GetId() string {
	return a.id
}

func (u *CachedUser) GetId() string {
	return u.id
}
