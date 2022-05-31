package models

import "time"

type Community struct {
	ID   int64  `json:"id,omitempty" db:"community_id"`
	Name string `json:"name,omitempty" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id,omitempty" db:"community_id"`
	Name         string    `json:"name,omitempty" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
