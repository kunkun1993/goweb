package access

import (
	"iissy.com/src/models"
	"iissy.com/src/utils"
	"math"
)

func ClanList(clan_id int, page int, size int) (*models.Clans, error) {
	var result models.Clans
	result.Items = []*models.Clan{}

	rows,err := db.Query("select clan_id,clan_name,description,create_time,delete_time from clans order by clan_id desc limit ?, ?", (page-1)*size, size)
	utils.CheckErr(err)

	for rows.Next() {
		item := models.Clan{}
		err = rows.Scan(&item.Clan_id, &item.Clan_name, &item.Description, &item.Create_time, &item.Delete_time)
		utils.CheckErr(err)

		result.Items = append(result.Items, &item)
	}

	total := 0
	err = db.QueryRow("select count(*) from clans").Scan(&total)
	utils.CheckErr(err)

	pageCount := int(math.Ceil(float64(total) / float64(size)))
	result.PageArgs = models.PageArgs{PageNumber: page, TotalCount: total, PageSize: size, PageCount: pageCount}
	defer rows.Close()

	return &result, nil
}


func PostClan(clan models.Clan) (bool,error) {
	var result int64
	if clan.Clan_id > 0 {
		res, err := db.Exec("update clans set clan_name=?,Description=? where id=?", clan.Clan_name, clan.Description, clan.Clan_id)
		utils.CheckErr(err)
		result, err = res.RowsAffected()
		utils.CheckErr(err)
	} else {
		res, err := db.Exec("update clans set clan_name=?,Description=? where id=?", clan.Clan_name, clan.Description, clan.Clan_id)
		utils.CheckErr(err)
		result, err = res.RowsAffected()
		utils.CheckErr(err)
	}
	return result > 0, nil
}
