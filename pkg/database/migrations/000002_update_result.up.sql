ALTER TABLE results ADD type varchar(255) NOT NULL;
ALTER TABLE results ADD form_id int not null;
ALTER TABLE results ADD user_id int not null;
ALTER TABLE results ADD FOREIGN KEY (form_id) REFERENCES forms(id) ON DELETE CASCADE;
ALTER TABLE results ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
