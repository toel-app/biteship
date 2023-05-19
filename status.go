package biteship

type Status string

const (
	StatusPlaced          Status = "placed"
	StatusScheduled       Status = "scheduled"
	StatusConfirmed       Status = "confirmed"
	StatusAllocated       Status = "allocated"
	StatusPickingUp       Status = "picking_up"
	StatusPicked          Status = "picked"
	StatusDroppingOff     Status = "dropping_off"
	StatusDelivered       Status = "delivered"
	StatusRejected        Status = "rejected"
	StatusCancelled       Status = "cancelled"
	StatusOnHold          Status = "on_hold"
	StatusCourierNotFound Status = "courier_not_found"
)
