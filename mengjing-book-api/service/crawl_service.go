package service

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"mengjing-book/dao/dto"
	"mengjing-book/dao/entity"
	"mengjing-book/dao/model"
	"mengjing-book/global"
	"mengjing-book/result"
	"time"
)

func CrawlSaveService(c *gin.Context) {
	bookSourceInfo := dto.BookSourceInfoDto{}
	err := c.Bind(&bookSourceInfo)
	fmt.Printf("bookinfo:%v\n", bookSourceInfo)
	if err != nil {
		result.Err(c, "序列化数据失败")
		return
	}
	// 获取当前时间
	nowTime := time.Now().Unix()
	// 保存书源 info 信息
	node, _ := snowflake.NewNode(1)
	bookSourceInfo.BookSourceInfo.BookSourceID = node.Generate().Int64()
	bookSourceInfo.BookSourceInfo.LastUpdateTime = nowTime
	if global.DB.Save(&bookSourceInfo.BookSourceInfo).Error != nil {
		result.Err(c, "保存基础信息失败")
		return
	}
	// 保存书源 rule info 信息
	bookSourceInfo.BookSourceRuleBookInfo.BookSourceID = bookSourceInfo.BookSourceInfo.BookSourceID
	bookSourceInfo.BookSourceRuleBookInfo.RuleBookInfoID = node.Generate().Int64()
	if global.DB.Save(&bookSourceInfo.BookSourceRuleBookInfo).Error != nil {
		result.Err(c, "保存基础信息失败")
		return
	}
	// 保存书源 rule content info 信息
	bookSourceInfo.BookSourceRuleContent.BookSourceID = bookSourceInfo.BookSourceInfo.BookSourceID
	bookSourceInfo.BookSourceRuleContent.RuleContentID = node.Generate().Int64()
	if global.DB.Save(&bookSourceInfo.BookSourceRuleContent).Error != nil {
		result.Err(c, "保存基础信息失败")
		return
	}
	// 保存书源 rule search info 信息
	bookSourceInfo.BookSourceRuleSearch.BookSourceID = bookSourceInfo.BookSourceInfo.BookSourceID
	bookSourceInfo.BookSourceRuleSearch.RuleSearchID = node.Generate().Int64()
	if global.DB.Save(&bookSourceInfo.BookSourceRuleSearch).Error != nil {
		result.Err(c, "保存基础信息失败")
		return
	}
	// 保存书源 rule toc info 信息
	bookSourceInfo.BookSourceRuleToc.BookSourceID = bookSourceInfo.BookSourceInfo.BookSourceID
	bookSourceInfo.BookSourceRuleToc.RuleTocID = node.Generate().Int64()
	if global.DB.Save(&bookSourceInfo.BookSourceRuleToc).Error != nil {
		result.Err(c, "保存基础信息失败")
		return
	}
	result.Success(c, "保存成功")
}

func CrawlUpdateService(c *gin.Context) {
}

func CrawlGetService(c *gin.Context) {
}

func CrawlListService(c *gin.Context) {
	infoModel := model.BookSourceInfoModel{}
	err := c.Bind(&infoModel)
	if err != nil {
		result.Err(c, "序列化数据失败")
		return
	}
	var bookSourceInfo []entity.BookSourceInfo
	db := global.DB
	if infoModel.BookSourceName != "" {
		db = db.Where("book_source_name like ?", "%"+infoModel.BookSourceName+"%")
	}
	if infoModel.BookSourceURL != "" {
		db = db.Where("book_source_url like ?", "%"+infoModel.BookSourceURL+"%")
	}
	if db.
		Offset((infoModel.PageNum-1)*infoModel.PageSize).
		Limit(infoModel.PageSize).
		Find(&bookSourceInfo).Error != nil {
		result.Err(c, "查询失败")
		return
	}
	// 查询总数
	var total int64
	if db.Model(&entity.BookSourceInfo{}).Count(&total).Error != nil {
		result.Err(c, "查询失败")
		return
	}
	result.Ok(c, gin.H{
		"list":  bookSourceInfo,
		"total": total,
	}, "查询成功")
}

func CrawlDeleteService(c *gin.Context) {
	// 获取参数id
	id := c.Param("id")
	// 删除书源 rule info 信息
	if global.DB.Where("book_source_id = ?", id).Delete(&entity.BookSourceRuleBookInfo{}).Error != nil {
		result.Err(c, "删除失败")
		return
	}
	// 删除书源 rule content info 信息
	if global.DB.Where("book_source_id = ?", id).Delete(&entity.BookSourceRuleContent{}).Error != nil {
		result.Err(c, "删除失败")
		return
	}
	// 删除书源 rule search info 信息
	if global.DB.Where("book_source_id = ?", id).Delete(&entity.BookSourceRuleSearch{}).Error != nil {
		result.Err(c, "删除失败")
		return
	}

	// 删除书源 rule toc info 信息
	if global.DB.Where("book_source_id = ?", id).Delete(&entity.BookSourceRuleToc{}).Error != nil {
		result.Err(c, "删除失败")
		return
	}

	// 删除书源 info 信息
	if global.DB.Where("book_source_id = ?", id).Delete(&entity.BookSourceInfo{}).Error != nil {
		result.Err(c, "删除失败")
		return
	}

	result.Success(c, "删除成功")
}
