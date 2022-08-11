package couple

import (
	"fmt"
	"time"
	"github.com/sjoh0704/happysaving/user"
)

type Couple struct{
	ID          int64 `json:"id`
	MaleID 		int64	`json:"male_id"`
	Male 		*user.User `pg:"rel:has-one"` 
	FemaleID 	int64	`json:"female_id"`
	Female 		*user.User `pg:"rel:has-one"` 
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

func (c Couple) String() string{
	return fmt.Sprintf("Couple id: %d, male Id: %d, female Id: %d", c.ID, c.MaleID, c.FemaleID)
}

func (c *Couple) SetMID(mid int64){
	c.MaleID = mid
}

func (c *Couple) SetFmID(fmid int64){
	c.FemaleID = fmid
}

func (c *Couple) UpdateTime(){
	c.UpdatedAt = time.Now()
}

func (c *Couple) GetMID() int64{
	return c.MaleID
}

func (c *Couple) GetFmID() int64{
	return c.FemaleID
}
