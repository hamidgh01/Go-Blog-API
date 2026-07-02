package postgres_repository

import "database/sql"

var (
	// create
	createDraftPostQuery = `
		INSERT INTO posts (title, content, isPrivate, userID) VALUES ($1, $2, $3, $4)
		RETURNING id, title, content, status, isPrivate, userID, createdAt, modifiedAt, firstPublishedAt
	`

	createPublishedPostQuery = `
		INSERT INTO posts (title, content, status, isPrivate, firstPublishedAt, userID)
		VALUES ($1, $2, 'published', $3, CURRENT_TIMESTAMP, $4)
		RETURNING id, title, content, status, isPrivate, userID, createdAt, modifiedAt, firstPublishedAt
	`

	// update
	updatePostQuery = `
		UPDATE posts
		SET title = $1, content = $2, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $3 AND (status = 'draft' OR status = 'published')
	`
	updatePostStatusQuery = `
		UPDATE posts
		SET status = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2
	`
	publishDraftPostQuery = `
		UPDATE posts
		SET status = 'published', modifiedAt = CURRENT_TIMESTAMP, firstPublishedAt = CURRENT_TIMESTAMP
		WHERE id = $1 AND status = 'draft'
	`
	updatePostPrivacyQuery = `
		UPDATE posts
		SET isPrivate = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2 AND (status = 'draft' OR status = 'published')
	`

	// delete
	deletePostQuery = "DELETE FROM posts WHERE id = $1"

	// read
	getPostByIDQuery = `
		SELECT
		p.id, p.title, p.content, p.status, p.isPrivate, p.userID, p.createdAt, p.modifiedAt, p.firstPublishedAt,
		u.id, u.username
		FROM posts as p
		JOIN users as u
		ON p.userID = u.id
		WHERE p.id = $1 AND status = 'published'
	`

	getDraftByIDQuery = `` // add later (and implement other layers) access: admin_or_owner

	getPostOwnerIDQuery = "SELECT userID FROM posts WHERE id = $1"

	// get other sources that has FK to `Post`
	countPostCommentsQuery = "SELECT COUNT(id) FROM comments WHERE postParentID = $1 AND status = 'published'"
	countPostLikesQuery    = "SELECT COUNT(user_id) FROM post_likes_m2m WHERE post_id = $1"
	countPostTagsQuery     = "SELECT COUNT(tag_id) FROM posts_tags_m2m WHERE post_id = $1"

	countListsThatSavedThisPostQuery = "SELECT COUNT(list_id) FROM saved_posts_m2m WHERE post_id = $1"

	getPostCommentsQuery = `
		SELECT
		c.id, c.content, c.status, c.postParentID, c.userID, c.createdAt, c.modifiedAt,
		u.id, u.username
		FROM comments as c
		JOIN users as u ON c.userID = u.id
		WHERE c.postParentID = $1 AND c.status = 'published'
		ORDER BY c.createdAt DESC
		LIMIT %d OFFSET %d
	`
	getPostLikesQuery = `
		SELECT u.id, u.username
		FROM users as u
		WHERE u.id IN (
			SELECT pl.user_id FROM post_likes_m2m as pl
			WHERE pl.post_id = $1
			ORDER BY pl.liked_at DESC
			LIMIT %d OFFSET %d
		) AND enabled = true
	`
	getPostTagsQuery = `
		SELECT id, name FROM tags
		WHERE id IN (
			SELECT pt.tag_id FROM posts_tags_m2m as pt
			WHERE pt.post_id = $1
			LIMIT %d OFFSET %d
		)
	`
	getListsThatSavedThisPostQuery = `
		SELECT
		l.id, l.title, l.isPrivate, l.userID, l.createdAt, l.modifiedAt,
		u.id, u.username
		FROM lists as l
		JOIN users as u ON l.userID = u.id
		WHERE l.id IN (
			SELECT sp.list_id FROM saved_posts_m2m as sp
			WHERE sp.post_id = $1
			ORDER BY sp.saved_at DESC
			LIMIT %d OFFSET %d
		) AND l.isPrivate = false
	`
)

var (
	// create
	createDraftPostStmt     *sql.Stmt
	createPublishedPostStmt *sql.Stmt

	// update
	updatePostStmt        *sql.Stmt
	updatePostStatusStmt  *sql.Stmt
	publishDraftPostStmt  *sql.Stmt
	updatePostPrivacyStmt *sql.Stmt

	// delete
	deletePostStmt *sql.Stmt

	// read
	getPostByIDStmt *sql.Stmt

	getPostOwnerIDStmt *sql.Stmt
)

func prepareAllPostStatements(db *sql.DB) {
	// create
	createDraftPostStmt = prepareStatement(db, "createDraftPost", createDraftPostQuery)
	createPublishedPostStmt = prepareStatement(db, "createPublishedPost", createPublishedPostQuery)

	// update
	updatePostStmt = prepareStatement(db, "updatePost", updatePostQuery)
	updatePostStatusStmt = prepareStatement(db, "updatePostStatus", updatePostStatusQuery)
	publishDraftPostStmt = prepareStatement(db, "publishDraftPost", publishDraftPostQuery)
	updatePostPrivacyStmt = prepareStatement(db, "updatePostPrivacy", updatePostPrivacyQuery)

	// delete
	deletePostStmt = prepareStatement(db, "deletePost", deletePostQuery)

	// read
	getPostByIDStmt = prepareStatement(db, "getPostByID", getPostByIDQuery)

	getPostOwnerIDStmt = prepareStatement(db, "getPostOwnerID", getPostOwnerIDQuery)
}

func closeAllPostStatements() {
	createDraftPostStmt.Close()
	createPublishedPostStmt.Close()
	updatePostStmt.Close()
	updatePostStatusStmt.Close()
	publishDraftPostStmt.Close()
	updatePostPrivacyStmt.Close()
	deletePostStmt.Close()
	getPostByIDStmt.Close()
	getPostOwnerIDStmt.Close()
}
