package models

import (
	"time"
)

type WarningEvents struct {
	IdHolocron                    int32     `json:"id_holocron"`
	IdTcloudWatch                 string    `json:"id_tcloud_watch"`
	EventDetailData               string    `json:"event_detail_data" gorm:"serializer:json"`
	EventData                     *Job      `json:"event_data" gorm:"serializer:json"`
	ChamadoEncerradoIsUnprotected bool      `json:"chamado_encerrado_is_unprotected"`
	Status                        string    `json:"status"`
	CreatedAt                     time.Time `json:"created_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
}
