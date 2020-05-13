package endpoints

import (
	"github.com/go-pg/pg"
)

type PostgreConfig struct {
	User     string
	Password string
	Port     string
	Host     string
}

func NewPostgreBot(config PostgreConfig) MoviesDb {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database:"movies",
	})
	return &postgreStore{db: db}
}

type postgreStore struct {
	db *pg.DB
}

func (p postgreStore) UpdateUser(id int, user *UserMovies) (*UserMovies, error) {
	user.ID = id

	err := p.db.Update(user)
	return user, err
}
func (p postgreStore) GetUser(id int) (*UserMovies, error) {
	user := &UserMovies{ID: id}
	err := p.db.Select(user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p postgreStore) CreateUser(user *UserMovies) (*UserMovies, error) {
	res := p.db.Insert(user)
	return user, res
}



func (p postgreStore) GetMovie(id int) (*Movie, error) {
	movie := &Movie{ID: id}
	err := p.db.Select(movie)

	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (p postgreStore) CreateMovie(movie *Movie) (*Movie, error) {
	res := p.db.Insert(movie)
	return movie, res
}

func (p postgreStore) DeleteMovie(id int) error {
	panic("implement me")
}

