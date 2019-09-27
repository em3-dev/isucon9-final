package isutrain

import (
	"errors"
	"time"
)

var (
	ErrNoSeats = errors.New("座席がありません")
)

// TODO: Statusのconst

// ReservationStatus は予約状態です
type ReservationStatus string

const (
	// Pending は予約が確定していない状態です
	Pending ReservationStatus = "pending"
	// Ok は予約が確定した状態です
	Ok ReservationStatus = "ok"
)

type ReserveResponse struct {
	ReservationID int  `json:"reservation_id"`
	IsOk          bool `json:"is_ok"`
}

type ReservationRequest struct {
	// Train構造体
	TrainClass string `json:"train_class"`
	TrainName  string `json:"train_name"`
	// TrainSeat構造体
	SeatClass string     `json:"seat_class"`
	Seats     TrainSeats `json:"seats"`
	// それ以外
	//// 区間
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	// 日付
	Date   time.Time `json:"date"`
	CarNum int       `json:"car_number"`
	Child  int       `json:"child"`
	Adult  int       `json:"adult"`
	// 座席位置(通路、真ん中、窓側)
	Type string `json:"type"`
}

func NewReservationRequest(trains Train, seats TrainSeats) (*ReservationRequest, error) {
	req := &ReservationRequest{}

	// Train構造体
	req.TrainClass = trains.Class
	req.TrainName = trains.Name

	// TrainSeat構造体
	if len(seats) == 0 {
		return nil, ErrNoSeats
	}
	req.SeatClass = seats[0].Class
	req.Seats = seats

	return req, nil
}

// SeatReservation は座席予約です
type ReservationResponse struct {
	ReservationID int        `json:"reservation_id"`
	UserID        *int       `json:"user_id"`
	Date          *time.Time `json:"date"`
	TrainClass    string     `json:"train_class"`
	TrainName     string     `json:"train_name"`
	Departure     string     `json:"departure"`
	Arrival       string     `json:"arrival"`
	PaymentStatus string     `json:"payment_status"`
	Status        string     `json:"status"`
	PaymentID     string     `json:"payment_id,omitempty"`
	Adult         int        `json:"adult"`
	Child         int        `json:"child"`
	Amount        int64      `json:"amount"`
}


type CommitReservationRequest struct {
	ReservationID int    `json:"reservation_id"`
	CardToken     string `json:"card_token"`
}

type ShowReservationResponse struct {
	ReservationID int `json:"reservation_id"`
}
