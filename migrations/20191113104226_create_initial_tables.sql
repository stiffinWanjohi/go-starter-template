-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

create table daily_cost(
    id serial PRIMARY KEY,
    user_id integer,
    target_date date,
    amount decimal(12, 4),
    created_at timestamp with time zone not null default clock_timestamp(),
    updated_at timestamp with time zone
);

create index idx_daily_cost_user_id
    ON daily_cost(user_id);

create index idx_daily_cost_target_date
    ON daily_cost(target_date);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

drop index if exists idx_daily_cost_target_date;

drop index if exists idx_daily_cost_user_id;

drop table if exists daily_cost;
