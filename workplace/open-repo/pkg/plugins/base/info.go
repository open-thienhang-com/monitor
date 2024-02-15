package base

import "time"

type Info struct {
	Title            string    `json:"title" yaml:"title" ini:"title"`
	Description      string    `json:"description" yaml:"description" ini:"description"`
	OldVersion       string    `json:"old_version" yaml:"old_version" ini:"old_version"`
	Version          string    `json:"version" yaml:"version" ini:"version"`
	Author           string    `json:"author" yaml:"author" ini:"author"`
	Banners          []string  `json:"banners" yaml:"banners" ini:"banners"`
	Url              string    `json:"url" yaml:"url" ini:"url"`
	Cover            string    `json:"cover" yaml:"cover" ini:"cover"`
	MiniCover        string    `json:"mini_cover" yaml:"mini_cover" ini:"mini_cover"`
	Website          string    `json:"website" yaml:"website" ini:"website"`
	Agreement        string    `json:"agreement" yaml:"agreement" ini:"agreement"`
	CreateDate       time.Time `json:"create_date" yaml:"create_date" ini:"create_date"`
	UpdateDate       time.Time `json:"update_date" yaml:"update_date" ini:"update_date"`
	ModulePath       string    `json:"module_path" yaml:"module_path" ini:"module_path"`
	Name             string    `json:"name" yaml:"name" ini:"name"`
	Uuid             string    `json:"uuid" yaml:"uuid" ini:"uuid"`
	Downloaded       bool      `json:"downloaded" yaml:"downloaded" ini:"downloaded"`
	ExtraDownloadUrl string    `json:"extra_download_url" yaml:"extra_download_url" ini:"extra_download_url"`
	Price            []string  `json:"price" yaml:"price" ini:"price"`
	GoodUUIDs        []string  `json:"good_uuids" yaml:"good_uuids" ini:"good_uuids"`
	GoodNum          int64     `json:"good_num" yaml:"good_num" ini:"good_num"`
	CommentNum       int64     `json:"comment_num" yaml:"comment_num" ini:"comment_num"`
	Order            int64     `json:"order" yaml:"order" ini:"order"`
	Features         string    `json:"features" yaml:"features" ini:"features"`
	Questions        []string  `json:"questions" yaml:"questions" ini:"questions"`
	HasBought        bool      `json:"has_bought" yaml:"has_bought" ini:"has_bought"`
	CanUpdate        bool      `json:"can_update" yaml:"can_update" ini:"can_update"`
	Legal            bool      `json:"legal" yaml:"legal" ini:"legal"`
	SkipInstallation bool      `json:"skip_installation" yaml:"skip_installation" ini:"skip_installation"`
}

func (i Info) IsFree() bool {
	return len(i.Price) == 0
}
