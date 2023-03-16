package main

import "fmt"

/*
	Шаблон Строитель

	Позволяет вынести производство объекта за пределы класса, что бы отдельно
	контролировать способы создания класса не порождая большое количество
	разнообразных конструкторов.
*/

// Структура, которую мы будем создавать с помощью шаблона Строитель
type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

// Интерфейс для создания пользователя
type UserBuilderI interface {
	SetFirstName(firstName string) UserBuilderI
	SetLastName(lastName string) UserBuilderI
	SetEmail(email string) UserBuilderI
	SetAge(age int) UserBuilderI
	Build() User
}

// Реализация интерфейса UserBuilder
type UserBuilder struct {
	user User
}

func NewUserBuilder() UserBuilderI {
	return &UserBuilder{}
}

func (b *UserBuilder) SetFirstName(firstName string) UserBuilderI {
	b.user.FirstName = firstName
	return b
}

func (b *UserBuilder) SetLastName(lastName string) UserBuilderI {
	b.user.LastName = lastName
	return b
}

func (b *UserBuilder) SetEmail(email string) UserBuilderI {
	b.user.Email = email
	return b
}

func (b *UserBuilder) SetAge(age int) UserBuilderI {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

// Интерфейс для Директора
type UserDirector interface {
	CreateUser(firstName, lastName, email string, age int) User
}

// Реализация интерфейса UserDirector
type userDirector struct {
	builder UserBuilderI
}

func NewUserDirector(builder UserBuilderI) UserDirector {
	return &userDirector{builder: builder}
}

func (d *userDirector) CreateUser(firstName, lastName, email string, age int) User {
	return d.builder.
		SetFirstName(firstName).
		SetLastName(lastName).
		SetEmail(email).
		SetAge(age).
		Build()
}

// Пример использования шаблона Строитель
func main() {
	userBuilder := NewUserBuilder()
	user := userBuilder.SetFirstName("Егор").
		SetLastName("Могилевич").
		SetEmail("egormog@mail.com").
		SetAge(23).
		Build()

	fmt.Printf("%+v\n", user)

	userDirector := NewUserDirector(userBuilder)
	user = userDirector.CreateUser("Иван", "Батраков", "egormog@mail.com", 30)

	fmt.Printf("%+v\n", user)
}
