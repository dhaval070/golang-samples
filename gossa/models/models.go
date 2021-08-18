package models

type Location struct {
    Id int `json:"id"`
    Location string `json:"location"`
    Status int `json:"status"`
    CopyMethod string `json:"copy_method"`
    LockedByLeague string `json:"locked_by_league"`
    LockedByEventId int `json:"locked_by_event_id"`
    UseRclone int `json:"use_rclone"`
}
