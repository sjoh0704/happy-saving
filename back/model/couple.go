package model

import (
	"fmt"
	"time"
)

type Couple struct {
	ID         int64     `json:"id"`
	SenderId   int64     `json:"send_id"` //SenderId와 Sender가 같은 포멧으로 있어야 함
	Sender     *User     `json:"sender" pg:"rel:has-one"`
	ReceiverId int64     `json:"recv_id"`
	Receiver   *User     `json:"receiver" pg:"rel:has-one"`
	Phase      Phase     `json:"phase"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Posts      []*Post   `json:"posts" pg:"rel:has-many"`
}

func (c Couple) String() string {
	return fmt.Sprintf("Couple id: %d, phase: %s, Send Id: %d, Recv Id: %d", c.ID, c.Phase, c.SenderId, c.ReceiverId)
}

func (c *Couple) SetMID(mid int64) {
	c.SenderId = mid
}

func (c *Couple) SetFmID(fmid int64) {
	c.ReceiverId = fmid
}

func (c *Couple) UpdateTime() {
	c.UpdatedAt = time.Now()
}

func (c *Couple) GetMID() int64 {
	return c.SenderId
}

func (c *Couple) GetFmID() int64 {
	return c.ReceiverId
}
