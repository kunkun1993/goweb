package access

import (
	"fmt"
	"html/template"
	"math"

	"iissy.com/src/models"
	"iissy.com/src/utils"
)

// UserArticle is yes.
func UserArticle(userid int, page int, size int) (*models.Articles, error) {
	var result models.Articles
	result.Items = []*models.Article{}

	rows, err := db.Query("select ID,Subject,Visited,PostType,AddDate,Origin from Article where UserId = ? order by AddDate desc limit ?, ?", userid, (page-1)*size, size)
	utils.CheckErr(err)

	for rows.Next() {
		item := models.Article{}
		err = rows.Scan(&item.ID, &item.Subject, &item.Visited, &item.PostType, &item.AddDate, &item.Origin)
		utils.CheckErr(err)

		result.Items = append(result.Items, &item)
	}

	total := 0
	err = db.QueryRow("select count(*) from Article where UserId = ?", userid).Scan(&total)
	utils.MustErr(err)
	fmt.Println(1234)
	pageCount := int(math.Ceil(float64(total) / float64(size)))
	result.PageArgs = models.PageArgs{PageNumber: page, TotalCount: total, PageSize: size, PageCount: pageCount}
	rows.Close()
	return &result, nil
}

// Detail is for article
func Detail(id string) (*models.Article, error) {
	var result models.Article

	body := ""
	err := db.QueryRow("select Subject,AddDate,Body,Origin from Article where Id = ?", id).Scan(&result.Subject, &result.AddDate, &body, &result.Origin)
	utils.CheckErr(err)
	result.Body = template.HTML(body)

	list, err := db.Query("select ID,Subject from Article where AddDate < ? order by AddDate desc limit 10", result.AddDate)
	utils.CheckErr(err)
	result.List = []*models.Article{}
	for list.Next() {
		var simple models.Article
		err = list.Scan(&simple.ID, &simple.Subject)
		utils.CheckErr(err)

		result.List = append(result.List, &simple)
	}

	db.Exec("update Article set Visited = Visited + 1 where Id = ?", id)

	list.Close()
	return &result, nil
}

// Post is yes
func Post(article models.Article) (bool, error) {
	res, err := db.Exec("insert Article set ID=?,Subject=?,Picture=?,Description=?,Body=?,UserID=?,NickName=?,PostType=?,Origin=?", article.ID, article.Subject, article.Picture, article.Description, string(article.Body), article.UserID, article.NickName, article.PostType, article.Origin)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}

// GetArticle is for article
func GetArticle(id string) (*models.Article, error) {
	var result models.Article

	body := ""
	err := db.QueryRow("select Id,PostType,Subject,Picture,Body,Origin,Description from Article where Id = ?", id).Scan(&result.ID, &result.PostType, &result.Subject, &result.Picture, &body, &result.Origin, &result.Description)
	utils.CheckErr(err)
	result.Body = template.HTML(body)
	return &result, nil
}

// Update is yes
func Update(article models.Article) (bool, error) {
	res, err := db.Exec("update Article set Subject=?,Picture=?,Description=?,Body=?,PostType=?,Origin=? where id=? and UserId=?", article.Subject, article.Picture, article.Description, string(article.Body), article.PostType, article.Origin, article.ID, article.UserID)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}

// DelArticle is yes
func DelArticle(uid int, id string) (bool, error) {
	res, err := db.Exec("delete from Article where Id=? and UserId=?", id, uid)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}
