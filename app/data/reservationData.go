package data

import (
	_ "github.com/lib/pq"
	"github.com/speix/aws-lambda-go-cron-case/app"
	"github.com/speix/aws-lambda-go-cron-case/app/config"
)

type ReservationData struct {
	*config.Database
}

func (s ReservationData) GetAll() (app.Reservations, error) {

	reservations := make(app.Reservations, 0)

	sql := `select
				r.lang, 
				r.phone_code || r.phone as phone, 
				r.template
			from reservation r 
				inner join store s on s.store_id = r.store_id
			where 
				(s.integrations->'promo'->>'active')::boolean
				and r.promo_enabled and r.valid
				and to_char(current_timestamp, 'yyyy-mm-dd HH24:MI') = to_char(r.reservation_date, 'yyyy-mm-dd HH24:MI')`

	err := s.DB.Select(&reservations, sql)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}
