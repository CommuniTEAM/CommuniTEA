// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Business struct {
	ID              pgtype.UUID `json:"id"`
	Name            string      `json:"name"`
	StreetAddress   string      `json:"street_address"`
	City            pgtype.UUID `json:"city"`
	State           string      `json:"state"`
	Zipcode         string      `json:"zipcode"`
	BusinessOwnerID pgtype.UUID `json:"business_owner_id"`
}

type BusinessFollower struct {
	ID         pgtype.UUID `json:"id"`
	UserID     pgtype.UUID `json:"user_id"`
	BusinessID pgtype.UUID `json:"business_id"`
}

type BusinessOfferedTea struct {
	ID         pgtype.UUID `json:"id"`
	BusinessID pgtype.UUID `json:"business_id"`
	TeaID      pgtype.UUID `json:"tea_id"`
}

type BusinessReview struct {
	ID       pgtype.UUID `json:"id"`
	Business pgtype.UUID `json:"business"`
	Author   pgtype.UUID `json:"author"`
	Rating   int16       `json:"rating"`
	Comment  pgtype.Text `json:"comment"`
	Date     pgtype.Date `json:"date"`
}

type Event struct {
	ID             pgtype.UUID `json:"id"`
	Name           string      `json:"name"`
	Host           pgtype.UUID `json:"host"`
	LocationName   pgtype.Text `json:"location_name"`
	StreetAddress  string      `json:"street_address"`
	City           pgtype.UUID `json:"city"`
	State          string      `json:"state"`
	Zipcode        string      `json:"zipcode"`
	Date           pgtype.Date `json:"date"`
	StartTime      pgtype.Time `json:"start_time"`
	EndTime        pgtype.Time `json:"end_time"`
	Description    string      `json:"description"`
	HeadlinerOne   pgtype.Text `json:"headliner_one"`
	HeadlinerTwo   pgtype.Text `json:"headliner_two"`
	HighlightOne   pgtype.Text `json:"highlight_one"`
	HighlightTwo   pgtype.Text `json:"highlight_two"`
	HighlightThree pgtype.Text `json:"highlight_three"`
	Rsvps          bool        `json:"rsvps"`
	Capacity       pgtype.Int4 `json:"capacity"`
}

type EventCategory struct {
	Name string `json:"name"`
}

type EventCategoryTag struct {
	ID       pgtype.UUID `json:"id"`
	EventID  pgtype.UUID `json:"event_id"`
	Category string      `json:"category"`
}

type EventCohost struct {
	ID          pgtype.UUID `json:"id"`
	EventID     pgtype.UUID `json:"event_id"`
	UserID      pgtype.UUID `json:"user_id"`
	Permissions string      `json:"permissions"`
}

type EventCohostPermission struct {
	Name string `json:"name"`
}

type EventRsvp struct {
	ID    pgtype.UUID `json:"id"`
	Event pgtype.UUID `json:"event"`
	User  pgtype.UUID `json:"user"`
}

type EventWatcher struct {
	ID      pgtype.UUID `json:"id"`
	EventID pgtype.UUID `json:"event_id"`
	UserID  pgtype.UUID `json:"user_id"`
}

type LocationsCity struct {
	ID    pgtype.UUID `json:"id"`
	Name  string      `json:"name"`
	State string      `json:"state"`
}

type LocationsState struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Tea struct {
	ID          pgtype.UUID   `json:"id"`
	Name        string        `json:"name"`
	ImgUrl      pgtype.Text   `json:"img_url"`
	Description string        `json:"description"`
	BrewTime    pgtype.Text   `json:"brew_time"`
	BrewTemp    pgtype.Float8 `json:"brew_temp"`
	Published   bool          `json:"published"`
}

type TeaAromatic struct {
	Name string `json:"name"`
}

type TeaAromaticTag struct {
	ID    pgtype.UUID `json:"id"`
	Name  string      `json:"name"`
	TeaID pgtype.UUID `json:"tea_id"`
}

type TeaFlavorProfile struct {
	Name string `json:"name"`
}

type TeaFlavorProfileTag struct {
	ID    pgtype.UUID `json:"id"`
	Name  string      `json:"name"`
	TeaID pgtype.UUID `json:"tea_id"`
}

type TeaOrigin struct {
	Name string `json:"name"`
}

type TeaOriginTag struct {
	ID    pgtype.UUID `json:"id"`
	Name  pgtype.Text `json:"name"`
	TeaID pgtype.UUID `json:"tea_id"`
}

type User struct {
	ID        pgtype.UUID `json:"id"`
	Role      string      `json:"role"`
	Username  string      `json:"username"`
	FirstName pgtype.Text `json:"first_name"`
	LastName  pgtype.Text `json:"last_name"`
	Email     pgtype.Text `json:"email"`
	Password  string      `json:"password"`
	Location  pgtype.UUID `json:"location"`
}

type UserFavoriteTea struct {
	ID     pgtype.UUID `json:"id"`
	UserID pgtype.UUID `json:"user_id"`
	TeaID  pgtype.UUID `json:"tea_id"`
}

type UserRole struct {
	Name string `json:"name"`
}
