package list_query

import (
	"fmt"
	"gorm.io/gorm"
	"im-server/common/models"
)

type Option struct {
	PageInfo models.PageInfo
	Where    *gorm.DB
	Likes    []string
	Preload  []string
}

func ListQuery[T any](db *gorm.DB, model T, option Option) (list []T, count int64, err error) {
	query := db.Where(model) //查询哪个表

	//模糊匹配
	if option.PageInfo.Key != "" && len(option.Likes) > 0 {
		likeQuery := db.Where("")
		for index, column := range option.Likes {
			//where name like '%fff%' or abstract like '%fff%'
			//查询用户所有的属性，如果任何一个属性包含key就符合条件
			if index == 0 {
				likeQuery = likeQuery.Where(column+" LIKE ?", "%"+option.PageInfo.Key+"%")
			} else {
				likeQuery = likeQuery.Or(column+" LIKE ?", "%"+option.PageInfo.Key+"%")
			}
		}
		query = query.Where(likeQuery)
	}

	//预加载
	if len(option.Preload) > 0 {
		for _, s := range option.Preload {
			fmt.Printf("add preload: %s\n", s)
			query = query.Preload(s)
		}
	}

	//分页查询
	option.PageInfo.Page = 1
	option.PageInfo.Limit = 10

	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit

	err = query.Limit(option.PageInfo.Limit).Offset(offset).Find(&list).Error
	if err != nil {
		fmt.Printf("查询列表失败: %v", err)
		return
	}
	//查询总数
	err = query.Count(&count).Error
	if err != nil {
		fmt.Printf("查询列表count: %v", err)
		return
	}
	return
}
