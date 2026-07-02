CREATE INDEX IF NOT EXISTS idx_post_likes_m2m_user_id ON post_likes_m2m (user_id);
-- this index will help with queries that filter or join by user_id (fetching the posts a user liked)
