package app

type Reservation struct {
	Template string `db:"template" json:"-"`
	Language string `db:"lang" json:"lang"`
	Phone    string `db:"phone" json:"-"`
}

type Reservations []Reservation

type ReservationService interface {
	GetAll() (Reservations, error)
}
