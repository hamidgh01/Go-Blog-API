ALTER TABLE follows_m2m ADD CONSTRAINT cant_follow_yourself
CHECK (followed_by != followed);
