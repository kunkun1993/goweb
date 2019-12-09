package models

type Clan struct {
	Clan_id      int
	Clan_name    string
	Description  string
	Create_time  string
	Delete_time  string

}

type Clans struct {
	PageArgs PageArgs
	Items    []*Clan
}