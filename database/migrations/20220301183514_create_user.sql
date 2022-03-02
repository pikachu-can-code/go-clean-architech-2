-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO users (id,created_at,updated_at,first_name,last_name,email,password,phone) VALUES
(1,NOW(),NOW(),'Phi','Khanh','khanh@gmail.com','pass','phone')
;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DELETE FROM users where id = 1;

