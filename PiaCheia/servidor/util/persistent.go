package util

import (
	"time"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Account struct {
	AccountId int
	AccountName string
	Passhash uint64
	IsAdmin bool
}
const SchemaAccounts = `
CREATE TABLE IF NOT EXISTS accounts (
	accountId INTEGER NOT NULL PRIMARY KEY,
	accountName TEXT NOT NULL,
	passhash INT NOT NULL,
	isAdmin BOOLEAN NOT NULL DEFAULT false,
	UNIQUE (accountName)
);`
const INSERT_ACCOUNT = `
INSERT INTO accounts
	(accountName, passhash, idAdmin)
 VALUES
 	(?, ?, 0);
`

func newAccount(name, password string) *Account {
	r := Unpack(db.EXEC(INSERT_ACCOUNT), name, Hash(password)
	return nil
}

const INSERT_ADMIN_ACCOUNT = `
INSERT INTO accounts
	(accountName, passhash, idAdmin)
 VALUES
 	(?, ?, 1);
`


const INSERT_ARTICLE = `
INSERT INTO articles
	(path, lastEditor, lastEdit)
VALUES
	(?, ?, DATETIME("now"));
`
func (a *Account) NewArticle(filepath string) (int64, error) {
	r := Unpack(db.Exec(INSERT_ARTICLE, filepath, a.AccountId))
	return r.LastInsertId()
}

const ADD_ACCOUNT = `
INSERT INTO accounts
	(accountName, passhash)
VALUES
	(?, ?);
`

type Article struct {
	ArticleId int
	Path string
	LastEditor int
	LastEdit time.Time
	Comments []int // not in sql
}
const SchemaArticle = `
CREATE TABLE IF NOT EXISTS articles (
	articleId INTEGER NOT NULL PRIMARY KEY,
	path TEXT NOT NULL,
	lastEditor INTERGER NOT NULL,
	lastEdit DATE NOT NULL,
	UNIQUE (path),
	FOREIGN KEY(lastEditor) REFERENCES accounts(accountId)
);`

// figure out how to check if lastEditor is valid account
// CHECK ( SELECT MAX(rowid) FROM accounts ) >= lastEditor

func (art Article) GetComments() (comments []*Comment) {
	comments = make([]*Comment, len(art.Comments))
	for i, cmntId := range art.Comments {
		comments[i] = &Comments[cmntId]
	}
	return comments
}

const INSERT_HEAD_COMMENT = `
INSERT INTO comments
	(articleId, posterId, text)
VALUES
	(?, ?, ?);
`
func (art *Article) AddComent(posterId int, text string) {
	Panic(db.Exec(INSERT_HEAD_COMMENT,
		art.ArticleId, posterId, text,
	))
}


const UPDATE_ARTICLE_PATH = `
UPDATE articles SET
	path=?
WHERE
	( articleId IS ? );
`
func (art *Article) UpdatePath(newPath string) error {
	/*
	 if (Exists(newPath)) {
		return fmt.Errorf("%s is already used", newPath)
  	}
 	*/
	Unpack(db.Exec(UPDATE_ARTICLE_PATH), newPath, art.ArticleId)
}

type Comment struct {
	CommentId int
	ParentCommentId int
	ArticleId int
	PosterId int
	Text string
	Children []int // not in sql
}
const SchemaComment = `
CREATE TABLE IF NOT EXISTS comments (
	commentId INTEGER NOT NULL PRIMARY KEY,
	parentCommentId INTEGER NOT NULL DEFAULT 0,
	articleId INTEGER NOT NULL,
	posterId INTEGER NOT NULL,
	text TEXT NOT NULL,
	FOREIGN KEY(posterId) REFERENCES accounts(accountId),
	FOREIGN KEY(articleId) REFERENCES articles(articleId),
	FOREIGN KEY(parentCommentId) REFERENCES comments(commentId),
	CHECK ( parentCommentId IS NULL OR parentCommentId < commentId )
);`
// figure out how to check if posterId is valid account
// CHECK ( MAX(rowid) FROM accounts ) >= posterId
// figure out how to check if parentComment's articleId = articleId


func (cmt Comment) GetParent() *Comment {
	if cmt.ParentCommentId == 0 {return nil}
	return &Comments[cmt.ParentCommentId]
}
func (cmt Comment) GetChildren() (comments []*Comment) {
	comments = make([]*Comment, len(cmt.Children))
	for i, cmntId := range cmt.Children {
		comments[i] = &Comments[cmntId]
	}
	return comments
}
const INSERT_COMMENT = `
INSERT INTO comments
	(parentCommentId, articleId, posterId, text)
VALUES
	(?, ?, ?, ?);
`
func (cmt *Comment) AddComent(posterId int, text string) {
	Panic(db.Exec(INSERT_COMMENT,
		cmt.CommentId, cmt.ArticleId,
		posterId, text,
	))
}


const GET_ACCOUNTS = `
SELECT * FROM accounts;
`

const GET_ARTICLES = `
SELECT * FROM articles;
`

const GET_COMMENTS = `
SELECT * FROM comments;
`

// TODO: _could_ COUNT Articles, Accounts, Comments
// : to preallocate arrays
var Accounts = []Account{Account{0, "Super Admin", 0, true}}
var Articles = []Article{Article{ Path: "index.html" }}
var Comments = []Comment{Comment{ ArticleId: -1 }}

func InitSQL(SQLFILE string) {
	db = Unpack(sql.Open("sqlite3", SQLFILE))
	Unpack(db.Exec(SchemaAccounts))
	Unpack(db.Exec(SchemaArticle))
	Unpack(db.Exec(SchemaComment))


	accountRows := Unpack(db.Query(GET_ACCOUNTS))
	for accountRows.Next() {
		var acc Account
		Panic(accountRows.Scan(
			&acc.AccountId,
			&acc.AccountName,
			&acc.Passhash,
			&acc.IsAdmin,
		))
		Accounts = append(Accounts, acc)
	}

	articleRows := Unpack(db.Query(GET_ARTICLES))
	for articleRows.Next() {
		var art Article
		Panic(articleRows.Scan(
			&art.ArticleId,
			&art.Path,
			&art.LastEditor,
			&art.LastEdit,
		))
		Articles = append(Articles, art)
	}

	commentRows := Unpack(db.Query(GET_COMMENTS))
	for commentRows.Next() {
		var cmt Comment
		Panic(commentRows.Scan(
			&cmt.CommentId,
			&cmt.ParentCommentId,
			&cmt.ArticleId,
			&cmt.PosterId,
			&cmt.Text,
		))
		prt := cmt.GetParent()
		if (prt != nil) {
			prt.Children = append(prt.Children, cmt.CommentId)
		}
		Articles[cmt.ArticleId].Comments = append(Articles[cmt.ArticleId].Comments, cmt.CommentId)
		Comments = append(Comments, cmt)
	}
}

