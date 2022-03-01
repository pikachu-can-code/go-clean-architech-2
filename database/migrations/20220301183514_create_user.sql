-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO user (id,created_at,updated_at,first_name,last_name,email,password,phone) VALUES
(1,NOW(),NOW(),'Phi','Khanh','khanh@gmail.com','pass','phone')
;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM user where id = 1;
-- +goose StatementEnd

