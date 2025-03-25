CREATE TABLE story_views(
    id SERIAL PRIMARY KEY,
    story_id INT NOT NULL,
    viewer_id UUID NOT NULL,
    viewed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_story_views_stories FOREIGN KEY(story_id) REFERENCES stories(id) ON DELETE CASCADE,
    CONSTRAINT fk_story_views_users FOREIGN KEY(viewer_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE(story_id, viewer_id)
);
