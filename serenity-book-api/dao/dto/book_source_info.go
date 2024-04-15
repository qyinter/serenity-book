package dto

import (
	"serenity-book/dao/entity"
)

type BookSourceInfoDto struct {
	BookSourceInfo         entity.BookSourceInfo         `json:"bookSourceInfo"`
	BookSourceRuleBookInfo entity.BookSourceRuleBookInfo `json:"bookSourceRuleBookInfo"`
	BookSourceRuleContent  entity.BookSourceRuleContent  `json:"bookSourceRuleContent"`
	BookSourceRuleSearch   entity.BookSourceRuleSearch   `json:"bookSourceRuleSearch"`
	BookSourceRuleToc      entity.BookSourceRuleToc      `json:"bookSourceRuleToc"`
}
