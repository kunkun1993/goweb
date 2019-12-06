package access

import (
	"fmt"
	"iissy.com/src/models"
	"iissy.com/src/utils"
	"os"
)

// Index is yes.
func Index() (*models.Course, error) {
	list, err := db.Query("select ID,Subject,Picture,Description from Article order by AddDate desc limit ?", 30)
	utils.CheckErr(err)
	course := models.Course{}
	course.ArticleItems = []*models.Article{}
	for list.Next() {
		var article models.Article
		err = list.Scan(&article.ID, &article.Subject, &article.Picture, &article.Description)
		utils.CheckErr(err)

		course.ArticleItems = append(course.ArticleItems, &article)
	}

	list.Close()
	fmt.Println(fmt.Sprintf("%+v",course))
	os.Exit(3)
	return &course, nil
}
