package namespace

import "strconv"

type Namespace int

const (
	Main Namespace = iota
	Talk
	User
	UserTalk
	Project
	ProjectTalk
	File
	FileTalk
	MediaWiki
	MediaWikiTalk
	Template
	TemplateTalk
	Help
	HelpTalk
	Category
	CategoryTalk
	Media   Namespace = -2
	Special Namespace = -1
)

func (ns Namespace) String() string {
	return strconv.Itoa(int(ns))
}
